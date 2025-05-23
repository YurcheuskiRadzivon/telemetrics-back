package auth

import (
	"github.com/YurcheuskiRadzivon/telemetrics-back/config"
	sm "github.com/YurcheuskiRadzivon/telemetrics-back/internal/infrastructure/session-manager"
	"github.com/YurcheuskiRadzivon/telemetrics-back/pkg/generator"
	"github.com/YurcheuskiRadzivon/telemetrics-back/pkg/logger"
)

type Auth struct {
	gnrt *generator.Generator
	sm   *sm.SessionManager
	cfg  *config.Config
	lgr  *logger.Logger
}
