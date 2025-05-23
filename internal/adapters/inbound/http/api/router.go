package api

import (
	"github.com/YurcheuskiRadzivon/telemetrics-back/config"
	"github.com/YurcheuskiRadzivon/telemetrics-back/internal/adapters/inbound/http/api/auth"
	"github.com/YurcheuskiRadzivon/telemetrics-back/internal/adapters/inbound/http/api/manage"
	sm "github.com/YurcheuskiRadzivon/telemetrics-back/internal/infrastructure/session-manager"
	"github.com/YurcheuskiRadzivon/telemetrics-back/pkg/generator"
	"github.com/YurcheuskiRadzivon/telemetrics-back/pkg/logger"
	"github.com/gofiber/fiber/v2"
)

func NewApiRoutes(api fiber.Router, gnrt *generator.Generator, sm *sm.SessionManager, cfg *config.Config, lgr *logger.Logger) {
	authGroup := api.Group("/auth")
	{
		auth.NewAuthRoutes(authGroup, gnrt, sm, cfg, lgr)
	}

	manageGroup := api.Group("/manage")
	{
		manage.NewManageRoutes(manageGroup)
	}
}
