package api

import (
	"github.com/YurcheuskiRadzivon/telemetrics-back/internal/adapters/inbound/http/api/auth"
	"github.com/YurcheuskiRadzivon/telemetrics-back/internal/adapters/inbound/http/api/manage"
	"github.com/YurcheuskiRadzivon/telemetrics-back/pkg/logger"
	"github.com/gofiber/fiber/v2"
)

func NewApiRoutes(api fiber.Router, lgr *logger.Logger) {
	authGroup := api.Group("/auth")
	{
		auth.NewAuthRoutes(authGroup, lgr)
	}

	manageGroup := api.Group("/manage")
	{
		manage.NewManageRoutes(manageGroup, lgr)
	}
}
