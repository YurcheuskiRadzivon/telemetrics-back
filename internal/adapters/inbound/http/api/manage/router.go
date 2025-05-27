package manage

import (
	"github.com/YurcheuskiRadzivon/telemetrics-back/config"
	"github.com/YurcheuskiRadzivon/telemetrics-back/internal/adapters/outbound/repositories"
	"github.com/YurcheuskiRadzivon/telemetrics-back/internal/core/service"
	sm "github.com/YurcheuskiRadzivon/telemetrics-back/internal/infrastructure/session-manager"
	"github.com/YurcheuskiRadzivon/telemetrics-back/pkg/generator"
	jwts "github.com/YurcheuskiRadzivon/telemetrics-back/pkg/jwtservice"
	"github.com/YurcheuskiRadzivon/telemetrics-back/pkg/logger"
	"github.com/gofiber/fiber/v2"
)

func NewManageRoutes(
	manage fiber.Router,
	gnrt *generator.Generator,
	sm *sm.SessionManager,
	cfg *config.Config,
	lgr *logger.Logger,
	sr *repositories.SessionRepository,
	jwts *jwts.JWTService,
	us *service.UserService,
	vs *service.ViewOptService,
) {
	manageController := &Manage{
		gnrt: gnrt,
		sm:   sm,
		cfg:  cfg,
		lgr:  lgr,
		sr:   sr,
		jwts: jwts,
		us:   us,
		vs:   vs,
	}

	{
		manage.Use(manageController.AuthMiddleware)
		manage.Get("/userinfo", manageController.GetUserInfo)
		manage.Get("/viewparams", manageController.GetViewOptParams)
		manage.Put("/viewparams", manageController.UpdateViewOptParams)
		manage.Get("/channels", manageController.GetChannels)
		manage.Get("/metrics", manageController.GetFullChannels)
	}
}
