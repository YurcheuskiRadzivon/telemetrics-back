package telegram

import (
	"context"
	"time"

	"github.com/YurcheuskiRadzivon/telemetrics-back/internal/core/entity"
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

		var username, phone string

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

			username, phone = self.Username, self.Phone

			return nil
		}); err != nil {
			manageSession.ErrorChan <- err
			tg.lgr.ErrorLogger.Printf("❌ Auth failed for %s: %v\n", manageSession.Phone, err)
		} else {
			userID := tg.gnrt.NewUserID()
			if err := tg.us.CreateUser(ctx, &entity.User{
				UserID:      userID,
				Username:    username,
				PhoneNumber: phone,
			}); err != nil {
				manageSession.ErrorChan <- err
			}

			if err := tg.vs.CreateViewOptions(ctx, &entity.ViewOptions{
				UserID:            userID,
				ChannelCount:      0,
				Tittle:            false,
				About:             false,
				ChannelID:         false,
				ChannelDate:       false,
				ParticipantsCount: false,
				Photo:             false,
				MessageCount:      0,
				MessageID:         false,
				Views:             false,
				PostDate:          false,
				ReactionsCount:    false,
				Reactions:         false,
			}); err != nil {
				manageSession.ErrorChan <- err
			}

			authData := sm.AuthData{
				SessionID: sessionID,
				UserID:    userID,
			}
			manageSession.AuthDataChan <- authData
			tg.lgr.InfoLogger.Printf("✅ Auth completed for %s\n", manageSession.Phone)
		}
	}()

}
