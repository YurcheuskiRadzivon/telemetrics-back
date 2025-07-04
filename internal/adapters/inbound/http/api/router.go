package api

import (
	"github.com/YurcheuskiRadzivon/telemetrics-back/config"
	"github.com/YurcheuskiRadzivon/telemetrics-back/internal/adapters/inbound/http/api/auth"
	"github.com/YurcheuskiRadzivon/telemetrics-back/internal/adapters/inbound/http/api/manage"
	"github.com/YurcheuskiRadzivon/telemetrics-back/internal/adapters/outbound/repositories"
	"github.com/YurcheuskiRadzivon/telemetrics-back/internal/core/service"
	sm "github.com/YurcheuskiRadzivon/telemetrics-back/internal/infrastructure/session-manager"
	"github.com/YurcheuskiRadzivon/telemetrics-back/pkg/generator"
	jwts "github.com/YurcheuskiRadzivon/telemetrics-back/pkg/jwtservice"
	"github.com/YurcheuskiRadzivon/telemetrics-back/pkg/logger"
	"github.com/gofiber/fiber/v2"
)

func NewApiRoutes(
	api fiber.Router,
	gnrt *generator.Generator,
	sm *sm.SessionManager,
	cfg *config.Config,
	lgr *logger.Logger,
	sr *repositories.SessionRepository,
	jwts *jwts.JWTService,
	us *service.UserService,
	vs *service.ViewOptService,
) {
	authGroup := api.Group("/auth")
	{
		auth.NewAuthRoutes(
			authGroup,
			gnrt,
			sm,
			cfg,
			lgr,
			sr,
			jwts,
			us,
			vs,
		)
	}

	manageGroup := api.Group("/manage")
	{
		manage.NewManageRoutes(
			manageGroup,
			gnrt,
			sm,
			cfg,
			lgr,
			sr,
			jwts,
			us,
			vs,
		)
	}
}
