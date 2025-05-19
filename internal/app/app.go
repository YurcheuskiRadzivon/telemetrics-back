package app

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/YurcheuskiRadzivon/telemetrics-back/config"
	"github.com/YurcheuskiRadzivon/telemetrics-back/pkg/httpserver"
	"github.com/YurcheuskiRadzivon/telemetrics-back/pkg/logger"
)

func Run(cfg *config.Config) {
	lgr := logger.NewLogger()
	lgr.InfoLogger.Println(cfg.App.Name)

	httpServer := httpserver.New(httpserver.Port(cfg.HTTP.Port), httpserver.Prefork(cfg.HTTP.UsePreforkMode))
	httpServer.Start()

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		lgr.InfoLogger.Printf("app -Run - signal: %s", s.String())
	case err := <-httpServer.Notify():
		lgr.ErrorLogger.Printf("app -Run - httpServer.Notify: %w", err)
	}

	err := httpServer.Shutdown()
	if err != nil {
		lgr.ErrorLogger.Printf("app -Run - httpServer.Shutdown: %w", err)
	}

}
