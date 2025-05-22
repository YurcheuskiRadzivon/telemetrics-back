package repositories

import (
	"context"
)

type SessionRepository interface {
	SaveSession(ctx context.Context, session []byte) (string, error)
	GetSession(ctx context.Context, sessionID string) ([]byte, error)
	DeleteSession(ctx context.Context, sessionID string) error
}
