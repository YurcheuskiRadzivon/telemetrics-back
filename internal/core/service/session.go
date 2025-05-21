package service

import (
	ports "github.com/YurcheuskiRadzivon/telemetrics-back/internal/core/ports/repositories"
)

type SessionService struct {
	sessionRepo ports.SessionRepository
}
