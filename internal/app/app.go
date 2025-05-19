package app

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/YurcheuskiRadzivon/telemetrics-back/config"
	"github.com/YurcheuskiRadzivon/telemetrics-back/pkg/httpserver"
)

func Run(cfg *config.Config) {
	fmt.Println(cfg.App.Name)

	httpServer := httpserver.New(httpserver.Port(cfg.HTTP.Port), httpserver.Prefork(cfg.HTTP.UsePreforkMode))
	httpServer.Start()

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		log.Printf("app -Run - signal: %s", s.String())
	case err := <-httpServer.Notify():
		log.Printf("app -Run - httpServer.Notify: %w", err)
	}

	err := httpServer.Shutdown()
	if err != nil {
		log.Printf("app -Run - httpServer.Shutdown: %w", err)
	}

}
