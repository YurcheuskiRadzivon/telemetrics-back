package app

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/YurcheuskiRadzivon/telemetrics-back/config"
	"github.com/YurcheuskiRadzivon/telemetrics-back/internal/adapters/inbound/http"
	"github.com/redis/go-redis/v9"

	"github.com/YurcheuskiRadzivon/telemetrics-back/pkg/generator"
	"github.com/YurcheuskiRadzivon/telemetrics-back/pkg/httpserver"
	"github.com/YurcheuskiRadzivon/telemetrics-back/pkg/logger"
)

func Run(cfg *config.Config) {
	lgr := logger.NewLogger()
	lgr.InfoLogger.Println(cfg.App.Name)

	gnrtr := generator.Generator{}
	lgr.DebugLogger.Println(gnrtr.NewSessionID())

	Opt, err := redis.ParseURL(cfg.REDIS.URL)
	if err != nil {
		lgr.ErrorLogger.Fatalf("parse Redis options error: %w", err)
	}

	redisClient := redis.NewClient(Opt)
	info, err := redisClient.Info(context.Background()).Result()
	if err != nil {
		lgr.ErrorLogger.Fatalf("get Redis info error: %w", err)
	}
	fmt.Println(info)

	httpServer := httpserver.New(httpserver.Port(cfg.HTTP.Port), httpserver.Prefork(cfg.HTTP.UsePreforkMode))
	http.NewRouter(httpServer.App, cfg, lgr)

	httpServer.Start()

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		lgr.InfoLogger.Printf("app - Run - signal: %s", s.String())
	case err := <-httpServer.Notify():
		lgr.ErrorLogger.Printf("app - Run - httpServer.Notify: %w", err)
	}

	err = httpServer.Shutdown()
	if err != nil {
		lgr.ErrorLogger.Printf("app - Run - httpServer.Shutdown: %w", err)
	}

}
