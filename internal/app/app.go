package app

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/YurcheuskiRadzivon/telemetrics-back/config"
	"github.com/YurcheuskiRadzivon/telemetrics-back/internal/adapters/inbound/http"
	"github.com/YurcheuskiRadzivon/telemetrics-back/internal/adapters/outbound/repositories"
	sm "github.com/YurcheuskiRadzivon/telemetrics-back/internal/infrastructure/session-manager"
	"github.com/redis/go-redis/v9"

	"github.com/YurcheuskiRadzivon/telemetrics-back/pkg/generator"
	"github.com/YurcheuskiRadzivon/telemetrics-back/pkg/httpserver"
	jwts "github.com/YurcheuskiRadzivon/telemetrics-back/pkg/jwtservice"
	"github.com/YurcheuskiRadzivon/telemetrics-back/pkg/logger"
)

func Run(cfg *config.Config) {
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
	http.NewRouter(httpServer.App, gnrt, sessionManager, cfg, lgr, sessionRepository, jwts)

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
