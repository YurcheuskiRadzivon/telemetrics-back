package repositories

import (
	"context"
	"errors"
	"time"

	"github.com/YurcheuskiRadzivon/telemetrics-back/pkg/ctxutil"
	"github.com/YurcheuskiRadzivon/telemetrics-back/pkg/generator"
	"github.com/redis/go-redis/v9"
)

const (
	ErrNoSessionID = "NO_SESSSION_ID"

	_sessionTTL = 120 * time.Hour
)

type SessionRepository struct {
	gnrt   *generator.Generator
	client *redis.Client
}

func NewSessionRepository(gnrt *generator.Generator, client *redis.Client) *SessionRepository {
	return &SessionRepository{
		gnrt:   gnrt,
		client: client,
	}
}

func (sr *SessionRepository) StoreSession(ctx context.Context, session []byte) error {
	sessionID, ok := ctxutil.GetSessionID(ctx)
	if !ok {
		return errors.New(ErrNoSessionID)
	}

	err := sr.client.Set(ctx, sessionID, session, _sessionTTL).Err()
	if err != nil {
		return err
	}
	return nil

}

func (sr *SessionRepository) LoadSession(ctx context.Context) ([]byte, error) {
	sessionID, ok := ctxutil.GetSessionID(ctx)
	if !ok {
		return nil, errors.New(ErrNoSessionID)
	}

	var result []byte
	err := sr.client.Get(ctx, sessionID).Scan(&result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (sr *SessionRepository) DeleteSession(ctx context.Context) error {
	sessionID, ok := ctxutil.GetSessionID(ctx)
	if !ok {
		return errors.New(ErrNoSessionID)
	}

	err := sr.client.Del(ctx, sessionID).Err()
	if err != nil {
		return err
	}
	return nil

}
