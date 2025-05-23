package auth

import (
	"context"
	"net/http"
	"path/filepath"
	"time"

	"github.com/YurcheuskiRadzivon/telemetrics-back/internal/adapters/inbound/http/api/auth/request"
	"github.com/YurcheuskiRadzivon/telemetrics-back/internal/adapters/inbound/http/api/auth/response"
	"github.com/YurcheuskiRadzivon/telemetrics-back/internal/adapters/outbound/telegram"
	customauth "github.com/YurcheuskiRadzivon/telemetrics-back/internal/infrastructure/telegram/custom-auth"
	"github.com/gofiber/fiber/v2"
	"github.com/gotd/td/telegram/auth"
)

const (
	sessionFormat = ".session"
	authTimeOut   = 2 * time.Minute
)

func (a *Auth) StartAuth(ctx *fiber.Ctx) error {
	var req request.StartBody
	if err := ctx.BodyParser(&req); err != nil {
		return errorResponse(ctx, http.StatusBadRequest, response.StatusInvalidRequest)
	}

	sessionID := a.gnrt.NewSessionID()

	sessionPath := filepath.Join("sessions", sessionID+sessionFormat)

	tg, err := telegram.New(sessionPath, a.cfg.TG.APP_ID, a.cfg.TG.APP_HASH)
	if err != nil {
		return errorResponse(ctx, http.StatusBadRequest, response.StatusInvalidRequest)
	}

	session := a.sm.CreateSession(req.PhoneNumber, sessionID)

	flow := auth.NewFlow(
		customauth.CustomAuthenticator{Session: session},
		auth.SendCodeOptions{},
	)

	res := response.StartResponse{
		Status:    response.StatusCodeRequested,
		SessionID: sessionID,
	}

	go func() {
		ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
		defer cancel()
		if err := tg.Tgc.Run(ctx, func(ctx context.Context) error {
			if err := tg.Tgc.Auth().IfNecessary(ctx, flow); err != nil {
				return err
			}

			self, err := tg.Tgc.Self(ctx)
			if err != nil {
				return err
			}

			a.lgr.InfoLogger.Printf("✅ Auth successful for %s (ID: %d)\n", self.Phone, self.ID)

			return nil
		}); err != nil {
			a.lgr.ErrorLogger.Printf("❌ Auth failed for %s: %v\n", req.PhoneNumber, err)
		} else {
			a.lgr.InfoLogger.Printf("✅ Auth completed for %s\n", req.PhoneNumber)
		}

		a.sm.DeleteSession(sessionID)
	}()

	return ctx.Status(http.StatusOK).JSON(res)
}

func (a *Auth) SubmitCode(ctx *fiber.Ctx) error {
	var req request.CodeBody
	if err := ctx.BodyParser(&req); err != nil {
		return errorResponse(ctx, http.StatusBadRequest, response.StatusInvalidRequest)
	}

	session, ok := a.sm.GetSession(req.SessionID)
	if !ok {
		return errorResponse(ctx, http.StatusBadRequest, response.StatusInvalidRequest)
	}

	session.CodeChan <- req.Code

	res := response.CodeRespose{
		Status: response.StatusCodeRequested,
	}

	return ctx.Status(http.StatusOK).JSON(res)
}

func (a *Auth) SubmitPassword(ctx *fiber.Ctx) error {
	var req request.PasswordBody
	if err := ctx.BodyParser(&req); err != nil {
		return errorResponse(ctx, http.StatusBadRequest, response.StatusInvalidRequest)
	}

	session, ok := a.sm.GetSession(req.SessionID)
	if !ok {
		return errorResponse(ctx, http.StatusBadRequest, response.StatusInvalidRequest)
	}

	session.PasswordChan <- req.Password

	res := response.CodeRespose{
		Status: response.StatusCodeRequested,
	}

	return ctx.Status(http.StatusOK).JSON(res)
}
