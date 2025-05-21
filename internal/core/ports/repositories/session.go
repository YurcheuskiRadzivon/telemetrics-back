package repositories

import (
	"context"

	"github.com/YurcheuskiRadzivon/telemetrics-back/internal/core/entity"
)

type SessionRepository interface {
	SaveSession(ctx context.Context, sessionID string, session entity.Session) error
	GetSession(ctx context.Context, sessionID string) (entity.Session, error)
	DeleteSession(ctx context.Context, sessionId string)
}
