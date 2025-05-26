package auth

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

func NewAuthRoutes(
	auth fiber.Router,
	gnrt *generator.Generator,
	sm *sm.SessionManager,
	cfg *config.Config,
	lgr *logger.Logger,
	sr *repositories.SessionRepository,
	jwts *jwts.JWTService,
	us *service.UserService,
	vs *service.ViewOptService,
) {
	authController := &Auth{
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
		auth.Post("/start", authController.StartAuth)
		auth.Post("/code", authController.SubmitCode)
		auth.Post("/password", authController.SubmitPassword)
	}
}
