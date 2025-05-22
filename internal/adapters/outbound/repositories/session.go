package repositories

import (
	"context"
	"time"

	"github.com/YurcheuskiRadzivon/telemetrics-back/pkg/generator"
	"github.com/redis/go-redis/v9"
)

type SessionRepository struct {
	gnrt   *generator.Generator
	client *redis.Client
	ttl    time.Duration
}

func NewSessionRepository(gnrt *generator.Generator, client *redis.Client, ttl time.Duration) *SessionRepository {
	return &SessionRepository{
		gnrt:   gnrt,
		client: client,
		ttl:    ttl,
	}
}

func (sr *SessionRepository) SaveSession(ctx context.Context, session []byte) (string, error) {
	sessionID := sr.gnrt.NewSessionID()
	err := sr.client.Set(ctx, sessionID, session, sr.ttl).Err()
	if err != nil {
		return "", err
	}
	return sessionID, nil

}

func (sr *SessionRepository) GetSession(ctx context.Context, sessionID string) ([]byte, error) {
	var result []byte
	err := sr.client.Get(ctx, sessionID).Scan(&result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (sr *SessionRepository) DeleteSession(ctx context.Context, sessionID string) error {
	err := sr.client.Del(ctx, sessionID).Err()
	if err != nil {
		return err
	}
	return nil

}
