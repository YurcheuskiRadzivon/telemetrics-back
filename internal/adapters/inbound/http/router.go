package http

import (
	"github.com/YurcheuskiRadzivon/telemetrics-back/config"
	"github.com/YurcheuskiRadzivon/telemetrics-back/internal/adapters/inbound/http/api"
	"github.com/YurcheuskiRadzivon/telemetrics-back/internal/adapters/inbound/http/middleware"
	sm "github.com/YurcheuskiRadzivon/telemetrics-back/internal/infrastructure/session-manager"
	"github.com/YurcheuskiRadzivon/telemetrics-back/pkg/generator"
	"github.com/YurcheuskiRadzivon/telemetrics-back/pkg/logger"
	"github.com/gofiber/fiber/v2"
)

func NewRouter(app *fiber.App, gnrt *generator.Generator, sm *sm.SessionManager, cfg *config.Config, lgr *logger.Logger) {
	app.Use(middleware.Logger(lgr))
	app.Use(middleware.Recovery(lgr))

	apiGroup := app.Group("")
	{
		api.NewApiRoutes(apiGroup, gnrt, sm, cfg, lgr)
	}

}
