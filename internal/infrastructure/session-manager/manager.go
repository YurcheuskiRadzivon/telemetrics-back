package session_manager

import (
	"sync"
)

type AuthData struct {
	SessionID string
	UserID    string
}

type ManageSession struct {
	ManageSessionID string
	Phone           string
	CodeChan        chan string
	PasswordChan    chan string
	ErrorChan       chan error
	AuthDataChan    chan AuthData
}

type SessionManager struct {
	manageSessions map[string]*ManageSession
	mu             sync.Mutex
}

func New() *SessionManager {
	return &SessionManager{
		manageSessions: make(map[string]*ManageSession),
	}
}

func (sm *SessionManager) CreateSession(phone, manageSessionID string) *ManageSession {
	codeChan := make(chan string, 1)
	passwordChan := make(chan string, 1)
	errorChan := make(chan error, 1)
	AuthDataChan := make(chan AuthData, 1)

	manageSession := &ManageSession{
		ManageSessionID: manageSessionID,
		Phone:           phone,
		CodeChan:        codeChan,
		PasswordChan:    passwordChan,
		ErrorChan:       errorChan,
		AuthDataChan:    AuthDataChan,
	}

	sm.mu.Lock()
	sm.manageSessions[manageSessionID] = manageSession
	sm.mu.Unlock()

	return manageSession
}

func (sm *SessionManager) GetSession(manageSessionID string) (*ManageSession, bool) {
	sm.mu.Lock()
	defer sm.mu.Unlock()

	manageSession, ok := sm.manageSessions[manageSessionID]

	return manageSession, ok
}

func (sm *SessionManager) DeleteSession(manageSessionID string) {
	sm.mu.Lock()
	defer sm.mu.Unlock()

	delete(sm.manageSessions, manageSessionID)
}
