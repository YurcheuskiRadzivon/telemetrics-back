package auth

import (
	"github.com/YurcheuskiRadzivon/telemetrics-back/config"
	sm "github.com/YurcheuskiRadzivon/telemetrics-back/internal/infrastructure/session-manager"
	"github.com/YurcheuskiRadzivon/telemetrics-back/pkg/generator"
	"github.com/YurcheuskiRadzivon/telemetrics-back/pkg/logger"
	"github.com/gofiber/fiber/v2"
)

func NewAuthRoutes(auth fiber.Router, gnrt *generator.Generator, sm *sm.SessionManager, cfg *config.Config, lgr *logger.Logger) {
	authController := &Auth{gnrt: gnrt, sm: sm, cfg: cfg, lgr: lgr}

	{
		auth.Post("/start", authController.StartAuth)
		auth.Post("/code", authController.SubmitCode)
		auth.Post("/password", authController.SubmitPassword)
	}
}
