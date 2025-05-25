package auth

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/YurcheuskiRadzivon/telemetrics-back/internal/adapters/inbound/http/api/auth/request"
	"github.com/YurcheuskiRadzivon/telemetrics-back/internal/adapters/inbound/http/api/auth/response"
	"github.com/YurcheuskiRadzivon/telemetrics-back/internal/adapters/outbound/telegram"
	customauth "github.com/YurcheuskiRadzivon/telemetrics-back/internal/infrastructure/telegram/custom-auth"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"github.com/gotd/td/telegram/auth"
	"github.com/gotd/td/tgerr"
)

const (
	sessionFormat          = ".session"
	authTimeOut            = 3 * time.Minute
	passwordNeedTimeOut    = 10 * time.Second
	ErrAuthFlowInvalidPass = "callback: auth flow: sign in with password: invalid password"
	sessionIDHeader        = "session_id"
	userIDHeader           = "user_id"
)

func (a *Auth) StartAuth(ctx *fiber.Ctx) error {
	var req request.StartBody
	if err := ctx.BodyParser(&req); err != nil {
		return errorResponse(ctx, http.StatusBadRequest, response.ErrInvalidRequest)
	}

	manageSessionID := a.gnrt.NewSessionID()

	tg := telegram.New(a.cfg.TG.APP_ID, a.cfg.TG.APP_HASH, a.lgr, a.gnrt, a.sr)

	manageSession := a.sm.CreateSession(req.PhoneNumber, manageSessionID)

	flow := auth.NewFlow(
		customauth.CustomAuthenticator{ManageSession: manageSession},
		auth.SendCodeOptions{},
	)

	res := response.StartResponse{
		Status:          response.StatusCodeRequested,
		ManageSessionID: manageSessionID,
	}

	tg.AuthProcession(flow, manageSession)

	return ctx.Status(http.StatusOK).JSON(res)
}

func (a *Auth) SubmitCode(ctx *fiber.Ctx) error {
	var req request.CodeBody
	if err := ctx.BodyParser(&req); err != nil {
		return errorResponse(ctx, http.StatusBadRequest, response.ErrInvalidRequest)
	}

	session, ok := a.sm.GetSession(req.ManageSessionID)
	if !ok {
		return errorResponse(ctx, http.StatusBadRequest, response.ErrInvalidRequest)
	}

	session.CodeChan <- req.Code
	ctxPassord, cancel := context.WithTimeout(ctx.Context(), passwordNeedTimeOut)
	defer cancel()

	select {
	case authData := <-session.AuthDataChan:
		payload := jwt.MapClaims{
			sessionIDHeader: authData.SessionID,
			userIDHeader:    authData.UserID,
		}

		token, err := a.jwts.CreateToken(payload)
		if err != nil {
			return errorResponse(ctx, http.StatusBadRequest, response.ErrJWT)
		}

		res := response.CodeRespose{
			Token:  token,
			Status: response.StatusSuccessfully,
		}

		return ctx.Status(http.StatusCreated).JSON(res)

	case err := <-session.ErrorChan:
		var tgErr *tgerr.Error
		if errors.As(err, &tgErr) {
			a.lgr.InfoLogger.Println(tgErr.Message)
			switch tgErr.Message {
			case response.StatusAuthRestart:
				res := response.CodeRespose{
					Token:  "",
					Status: response.StatusAuthRestart,
				}
				return ctx.Status(http.StatusOK).JSON(res)
			case response.StatusCodeEmpty:
				res := response.CodeRespose{
					Token:  "",
					Status: response.StatusCodeEmpty,
				}
				return ctx.Status(http.StatusOK).JSON(res)
			case response.StatusCodeExpired:
				res := response.CodeRespose{
					Token:  "",
					Status: response.StatusCodeExpired,
				}
				return ctx.Status(http.StatusOK).JSON(res)
			case response.ErrCodeInvalid, response.ErrSignInFailed:
				return errorResponse(ctx, http.StatusBadRequest, response.ErrCodeInvalid)
			default:
				return errorResponse(ctx, http.StatusBadRequest, response.ErrUnknown)
			}
		} else {
			return errorResponse(ctx, http.StatusBadRequest, response.ErrInvalidRequest)
		}

	case <-ctxPassord.Done():
		res := response.CodeRespose{
			Token:  "",
			Status: response.Status2FANeeded,
		}

		return ctx.Status(http.StatusOK).JSON(res)
	}
}

func (a *Auth) SubmitPassword(ctx *fiber.Ctx) error {
	var req request.PasswordBody
	if err := ctx.BodyParser(&req); err != nil {
		return errorResponse(ctx, http.StatusBadRequest, response.ErrInvalidRequest)
	}

	session, ok := a.sm.GetSession(req.ManageSessionID)
	if !ok {
		return errorResponse(ctx, http.StatusBadRequest, response.ErrInvalidRequest)
	}

	session.PasswordChan <- req.Password

	select {
	case authData := <-session.AuthDataChan:
		payload := jwt.MapClaims{
			sessionIDHeader: authData.SessionID,
			userIDHeader:    authData.UserID,
		}

		token, err := a.jwts.CreateToken(payload)
		if err != nil {
			return errorResponse(ctx, http.StatusBadRequest, response.ErrJWT)
		}

		res := response.CodeRespose{
			Token:  token,
			Status: response.StatusSuccessfully,
		}

		return ctx.Status(http.StatusCreated).JSON(res)

	case err := <-session.ErrorChan:
		switch err.Error() {
		case ErrAuthFlowInvalidPass:
			return errorResponse(ctx, http.StatusBadRequest, response.ErrPassHashInvalid)
		default:
			return errorResponse(ctx, http.StatusBadRequest, response.ErrUnknown)
		}
	}
}
