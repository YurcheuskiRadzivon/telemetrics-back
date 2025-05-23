package session_manager

import (
	"sync"
)

type AuthSession struct {
	SessionID    string
	Phone        string
	CodeChan     chan string
	PasswordChan chan string
}

type SessionManager struct {
	sessions map[string]*AuthSession
	mu       sync.Mutex
}

func New() *SessionManager {
	return &SessionManager{
		sessions: make(map[string]*AuthSession),
	}
}

func (sm *SessionManager) CreateSession(phone, sessionID string) *AuthSession {
	codeChan := make(chan string, 1)
	passwordChan := make(chan string, 1)

	session := &AuthSession{
		SessionID:    sessionID,
		Phone:        phone,
		CodeChan:     codeChan,
		PasswordChan: passwordChan,
	}

	sm.mu.Lock()
	sm.sessions[sessionID] = session
	sm.mu.Unlock()

	return session
}

func (sm *SessionManager) GetSession(sessionID string) (*AuthSession, bool) {
	sm.mu.Lock()
	defer sm.mu.Unlock()

	session, ok := sm.sessions[sessionID]

	return session, ok
}

func (sm *SessionManager) DeleteSession(sessionID string) {
	sm.mu.Lock()
	defer sm.mu.Unlock()

	delete(sm.sessions, sessionID)
}
