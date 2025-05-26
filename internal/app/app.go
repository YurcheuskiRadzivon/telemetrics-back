package app

import (
	"context"
	"database/sql"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/YurcheuskiRadzivon/telemetrics-back/config"
	"github.com/YurcheuskiRadzivon/telemetrics-back/internal/adapters/inbound/http"
	"github.com/YurcheuskiRadzivon/telemetrics-back/internal/adapters/outbound/repositories"
	"github.com/YurcheuskiRadzivon/telemetrics-back/internal/core/service"
	"github.com/YurcheuskiRadzivon/telemetrics-back/internal/infrastructure/database/queries"
	sm "github.com/YurcheuskiRadzivon/telemetrics-back/internal/infrastructure/session-manager"
	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/lib/pq"
	"github.com/pressly/goose"
	"github.com/redis/go-redis/v9"

	"github.com/YurcheuskiRadzivon/telemetrics-back/pkg/generator"
	"github.com/YurcheuskiRadzivon/telemetrics-back/pkg/httpserver"
	jwts "github.com/YurcheuskiRadzivon/telemetrics-back/pkg/jwtservice"
	"github.com/YurcheuskiRadzivon/telemetrics-back/pkg/logger"
)

func Run(cfg *config.Config) {
	if err := migrate(cfg.PG.URL); err != nil {
		log.Fatal("migrate: ", err)
	}

	conn, err := pgxpool.New(context.Background(), cfg.PG.URL)
	if err != nil {
		log.Fatal("connection: ", err)
	}

	q := queries.New(conn)

	_ = q

	userRepo := repositories.NewUserRepository(q, conn)
	viewOptRepo := repositories.NewViewOptRepository(q, conn)

	userService := service.NewUserService(userRepo)
	viewOptService := service.NewviewOptService(viewOptRepo)

	lgr := logger.NewLogger()

	gnrt := &generator.Generator{}

	jwts := jwts.New(cfg.JWT.SECRET_KEY)

	sessionManager := sm.New()

	Opt, err := redis.ParseURL(cfg.REDIS.URL)
	if err != nil {
		lgr.ErrorLogger.Fatalf("parse Redis options error: %w", err)
	}

	redisClient := redis.NewClient(Opt)

	sessionRepository := repositories.NewSessionRepository(gnrt, redisClient)

	httpServer := httpserver.New(httpserver.Port(cfg.HTTP.Port), httpserver.Prefork(cfg.HTTP.UsePreforkMode))

	http.NewRouter(
		httpServer.App,
		gnrt, sessionManager,
		cfg, lgr, sessionRepository,
		jwts, userService,
		viewOptService,
	)

	httpServer.Start()

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		lgr.InfoLogger.Printf("app - Run - signal: %s", s.String())
	case err := <-httpServer.Notify():
		lgr.ErrorLogger.Printf("app - Run - httpServer.Notify: %v", err)
	}

	err = httpServer.Shutdown()
	if err != nil {
		lgr.ErrorLogger.Printf("app - Run - httpServer.Shutdown: %v", err)
	}

}

func migrate(url string) error {
	db, err := sql.Open("postgres", url)
	if err != nil {
		return err
	}
	if err := goose.Up(db, "sql/migrations"); err != nil {
		return err
	}
	if err := db.Close(); err != nil {
		return err
	}
	return nil
}
