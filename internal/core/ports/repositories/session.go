package repositories

import (
	"context"
)

type SessionRepository interface {
	LoadSession(ctx context.Context) ([]byte, error)
	StoreSession(ctx context.Context, session []byte) error
	DeleteSession(ctx context.Context) error
}
