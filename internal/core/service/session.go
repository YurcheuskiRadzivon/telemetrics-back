package service

import (
	"context"

	"github.com/YurcheuskiRadzivon/telemetrics-back/internal/core/entity"
	ports "github.com/YurcheuskiRadzivon/telemetrics-back/internal/core/ports/repositories"
	json "github.com/goccy/go-json"
)

type SessionService struct {
	sessionRepo ports.SessionRepository
}

func NewSessionService(sessionRepo ports.SessionRepository) *SessionService {
	return &SessionService{
		sessionRepo: sessionRepo,
	}
}

func (ss *SessionService) SaveSession(ctx context.Context, session entity.Session) (string, error) {
	sessionData, err := json.Marshal(session)
	if err != nil {
		return "", err
	}

	return ss.sessionRepo.SaveSession(ctx, sessionData)
}

func (ss *SessionService) GetSession(ctx context.Context, sessionID string) (entity.Session, error) {
	sessionData, err := ss.sessionRepo.GetSession(ctx, sessionID)
	if err != nil {
		return entity.Session{}, err
	}

	if err = ss.sessionRepo.DeleteSession(ctx, sessionID); err != nil {
		return entity.Session{}, err
	}

	var session entity.Session
	err = json.Unmarshal(sessionData, &session)
	if err != nil {
		return entity.Session{}, err
	}

	return session, nil
}

func (ss *SessionService) DeleteSession(ctx context.Context, sessionID string) error {
	return ss.sessionRepo.DeleteSession(ctx, sessionID)
}
