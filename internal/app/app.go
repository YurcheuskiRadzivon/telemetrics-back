package app

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/YurcheuskiRadzivon/telemetrics-back/config"
	"github.com/YurcheuskiRadzivon/telemetrics-back/internal/adapters/inbound/http"
	sm "github.com/YurcheuskiRadzivon/telemetrics-back/internal/infrastructure/session-manager"

	"github.com/YurcheuskiRadzivon/telemetrics-back/pkg/generator"
	"github.com/YurcheuskiRadzivon/telemetrics-back/pkg/httpserver"
	"github.com/YurcheuskiRadzivon/telemetrics-back/pkg/logger"
)

func Run(cfg *config.Config) {
	lgr := logger.NewLogger()

	gnrt := &generator.Generator{}

	sm := sm.New()

	httpServer := httpserver.New(httpserver.Port(cfg.HTTP.Port), httpserver.Prefork(cfg.HTTP.UsePreforkMode))
	http.NewRouter(httpServer.App, gnrt, sm, cfg, lgr)

	httpServer.Start()

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		lgr.InfoLogger.Printf("app - Run - signal: %s", s.String())
	case err := <-httpServer.Notify():
		lgr.ErrorLogger.Printf("app - Run - httpServer.Notify: %v", err)
	}

	err := httpServer.Shutdown()
	if err != nil {
		lgr.ErrorLogger.Printf("app - Run - httpServer.Shutdown: %v", err)
	}

}
