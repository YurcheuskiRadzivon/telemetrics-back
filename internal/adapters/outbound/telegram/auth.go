package telegram

import (
	"context"
	"time"

	sm "github.com/YurcheuskiRadzivon/telemetrics-back/internal/infrastructure/session-manager"
	"github.com/YurcheuskiRadzivon/telemetrics-back/pkg/ctxutil"
	"github.com/gotd/td/telegram/auth"
)

const (
	authProcessionTimeOut = 2 * time.Minute
)

func (tg *TelegramClient) AuthProcession(flow auth.Flow, manageSession *sm.ManageSession) {
	go func() {
		ctx, cancel := context.WithTimeout(context.Background(), authProcessionTimeOut)
		defer cancel()

		sessionID := tg.gnrt.NewSessionID()

		ctxWithSessionID := ctxutil.WithSessionID(ctx, sessionID)

		tg.sr.StoreSession(ctxWithSessionID, []byte{})

		if err := tg.Client.Run(ctxWithSessionID, func(ctx context.Context) error {
			if err := tg.Client.Auth().IfNecessary(ctxWithSessionID, flow); err != nil {
				return err
			}

			self, err := tg.Client.Self(ctx)
			if err != nil {
				return err
			}

			tg.lgr.InfoLogger.Printf("✅ Auth successful for %s (ID: %d)\n", self.Phone, self.ID)

			return nil
		}); err != nil {
			manageSession.ErrorChan <- err
			tg.lgr.ErrorLogger.Printf("❌ Auth failed for %s: %v\n", manageSession.Phone, err)
		} else {
			authData := sm.AuthData{
				SessionID: sessionID,
				UserID:    "",
			}
			manageSession.AuthDataChan <- authData
			tg.lgr.InfoLogger.Printf("✅ Auth completed for %s\n", manageSession.Phone)
		}
	}()

}
