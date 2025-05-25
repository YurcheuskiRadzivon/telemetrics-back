package custom_authentificator

import (
	"context"

	sm "github.com/YurcheuskiRadzivon/telemetrics-back/internal/infrastructure/session-manager"
	"github.com/gotd/td/telegram/auth"
	"github.com/gotd/td/tg"
)

type CustomAuthenticator struct {
	ManageSession *sm.ManageSession
}

func (a CustomAuthenticator) Phone(ctx context.Context) (string, error) {
	return a.ManageSession.Phone, nil
}

func (a CustomAuthenticator) Password(ctx context.Context) (string, error) {
	select {
	case pass := <-a.ManageSession.PasswordChan:
		return pass, nil
	case <-ctx.Done():
		return "", ctx.Err()
	}
}

func (a CustomAuthenticator) Code(ctx context.Context, sentCode *tg.AuthSentCode) (string, error) {
	select {
	case code := <-a.ManageSession.CodeChan:
		return code, nil
	case <-ctx.Done():
		return "", ctx.Err()
	}
}

func (a CustomAuthenticator) AcceptTermsOfService(ctx context.Context, tos tg.HelpTermsOfService) error {
	return nil
}

func (a CustomAuthenticator) SignUp(ctx context.Context) (auth.UserInfo, error) {
	return auth.UserInfo{}, nil
}
