package http

import (
	"github.com/YurcheuskiRadzivon/telemetrics-back/config"
	"github.com/YurcheuskiRadzivon/telemetrics-back/internal/adapters/inbound/http/api"
	"github.com/YurcheuskiRadzivon/telemetrics-back/internal/adapters/inbound/http/middleware"
	"github.com/YurcheuskiRadzivon/telemetrics-back/pkg/logger"
	"github.com/gofiber/fiber/v2"
)

func NewRouter(app *fiber.App, cfg *config.Config, lgr *logger.Logger) {
	app.Use(middleware.Logger(lgr))
	app.Use(middleware.Recovery(lgr))

	apiGroup := app.Group("")
	{
		api.NewApiRoutes(apiGroup, lgr)
	}

}
