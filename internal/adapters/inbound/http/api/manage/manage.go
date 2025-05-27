package manage

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/YurcheuskiRadzivon/telemetrics-back/internal/adapters/inbound/http/api/manage/request"
	"github.com/YurcheuskiRadzivon/telemetrics-back/internal/adapters/inbound/http/api/manage/response"
	"github.com/YurcheuskiRadzivon/telemetrics-back/internal/adapters/outbound/telegram"
	"github.com/YurcheuskiRadzivon/telemetrics-back/internal/core/entity"
	"github.com/YurcheuskiRadzivon/telemetrics-back/pkg/ctxutil"
	"github.com/gofiber/fiber/v2"
	"github.com/gotd/td/tg"
)

func (m *Manage) GetUserInfo(ctx *fiber.Ctx) error {
	var req request.GetUserInfoBody
	if err := ctx.BodyParser(&req); err != nil {
		return errorResponse(ctx, http.StatusBadRequest, "INVALID_JSON_BODY")
	}

	userID, err := m.jwts.GetUserID(req.Token)
	if err != nil {
		return errorResponse(ctx, http.StatusBadRequest, "INVALID_TOKEN")
	}

	userInfo, err := m.us.GetUser(ctx.Context(), userID)
	if err != nil {
		return errorResponse(ctx, http.StatusBadRequest, "NO_USER_WITH_THIS_TOKEN")
	}

	return ctx.Status(http.StatusOK).JSON(response.GetUserInfoRes{
		UserID:      userInfo.UserID,
		Username:    userInfo.Username,
		PhoneNumber: userInfo.PhoneNumber,
	})
}

func (m *Manage) GetViewOptParams(ctx *fiber.Ctx) error {
	var req request.GetViewOptBody
	if err := ctx.BodyParser(&req); err != nil {
		return errorResponse(ctx, http.StatusBadRequest, "INVALID_JSON_BODY")
	}

	userID, err := m.jwts.GetUserID(req.Token)
	if err != nil {
		return errorResponse(ctx, http.StatusBadRequest, "INVALID_TOKEN")
	}

	viewOpt, err := m.vs.GetViewOptions(ctx.Context(), userID)
	if err != nil {
		return errorResponse(ctx, http.StatusBadRequest, "NO_PARAMS_WITH_THIS_TOKEN")
	}

	return ctx.Status(http.StatusOK).JSON(response.GetViewOptRes{
		ChannelCount:      viewOpt.ChannelCount,
		Tittle:            viewOpt.Tittle,
		About:             viewOpt.About,
		ChannelID:         viewOpt.ChannelID,
		ChannelDate:       viewOpt.ChannelDate,
		ParticipantsCount: viewOpt.ParticipantsCount,
		Photo:             viewOpt.Photo,
		MessageCount:      viewOpt.MessageCount,
		MessageID:         viewOpt.MessageID,
		Views:             viewOpt.Views,
		PostDate:          viewOpt.PostDate,
		ReactionsCount:    viewOpt.ReactionsCount,
		Reactions:         viewOpt.Reactions,
	})
}

func (m *Manage) UpdateViewOptParams(ctx *fiber.Ctx) error {
	var req request.UpdateViewOptParams
	if err := ctx.BodyParser(&req); err != nil {
		return errorResponse(ctx, http.StatusBadRequest, "INVALID_JSON_BODY")
	}

	userID, err := m.jwts.GetUserID(req.Token)
	if err != nil {
		return errorResponse(ctx, http.StatusBadRequest, "INVALID_TOKEN")
	}

	if err = m.vs.UpdateViewOptions(ctx.Context(), entity.ViewOptions{
		UserID:            userID,
		ChannelCount:      req.ChannelCount,
		Tittle:            req.Tittle,
		About:             req.About,
		ChannelID:         req.ChannelID,
		ChannelDate:       req.ChannelDate,
		ParticipantsCount: req.ParticipantsCount,
		Photo:             req.Photo,
		MessageCount:      req.MessageCount,
		MessageID:         req.MessageID,
		Views:             req.Views,
		PostDate:          req.PostDate,
		ReactionsCount:    req.ReactionsCount,
		Reactions:         req.Reactions,
	}); err != nil {
		return errorResponse(ctx, http.StatusBadRequest, "NO_PARAMS_WITH_THIS_TOKEN")
	}

	return ctx.Status(http.StatusOK).JSON(response.UpdateViewOptRes{
		Status: "UPDATE_SUCCESFULLY",
	})
}

func (m *Manage) GetChannels(ctx *fiber.Ctx) error {
	var req request.GetChannelsReq
	if err := ctx.BodyParser(&req); err != nil {
		return errorResponse(ctx, http.StatusBadRequest, "INVALID_JSON_BODY")
	}

	sessionID, err := m.jwts.GetSessionID(req.Token)
	if err != nil {
		return errorResponse(ctx, http.StatusBadRequest, "INVALID_TOKEN")
	}

	ctxWithSessionID := ctxutil.WithSessionID(ctx.Context(), sessionID)
	channels := make([]entity.ChannelInfo, 0)

	tgc := telegram.New(m.cfg.TG.APP_ID, m.cfg.TG.APP_HASH, m.lgr, m.gnrt, m.sr, m.us, m.vs)
	err = tgc.Client.Run(ctxWithSessionID, func(ctx context.Context) error {
		s, err := tgc.Client.Auth().Status(ctx)
		if err != nil {
			return fmt.Errorf("error getting auth status: %w", err)
		}
		if !s.Authorized {
			return fmt.Errorf("error authorizing user: %w", err)
		}
		res, err := tgc.Client.API().MessagesGetDialogs(ctx, &tg.MessagesGetDialogsRequest{
			OffsetDate: 0,
			OffsetID:   0,
			OffsetPeer: &tg.InputPeerEmpty{},
			Limit:      70,
			Hash:       0,
		})

		if err != nil {
			return fmt.Errorf("failed to get dialogs: %w", err)
		}

		dialogs := res.(*tg.MessagesDialogsSlice)
		for _, chat := range dialogs.Chats {
			if channel, ok := chat.(*tg.Channel); ok {
				channels = append(channels, entity.ChannelInfo{
					Name:      channel.Title,
					ChannelID: int(channel.ID),
				})
			}
		}
		return nil

	})

	res := response.GetChannelsRes{
		Channels: channels,
	}

	return ctx.Status(http.StatusOK).JSON(res)
}

func (m *Manage) GetFullChannels(ctx *fiber.Ctx) error {
	var req request.GetFullChannelsReq
	if err := ctx.BodyParser(&req); err != nil {
		return errorResponse(ctx, http.StatusBadRequest, "INVALID_JSON_BODY")
	}
	requirechannels := req.Channels
	sessionID, err := m.jwts.GetSessionID(req.Token)
	if err != nil {
		return errorResponse(ctx, http.StatusBadRequest, "INVALID_TOKEN")
	}

	ctxWithSessionID := ctxutil.WithSessionID(ctx.Context(), sessionID)
	channels := make([]entity.Channel, 0)

	tgc := telegram.New(m.cfg.TG.APP_ID, m.cfg.TG.APP_HASH, m.lgr, m.gnrt, m.sr, m.us, m.vs)
	err = tgc.Client.Run(ctxWithSessionID, func(ctx context.Context) error {
		s, err := tgc.Client.Auth().Status(ctx)
		if err != nil {
			return fmt.Errorf("error getting auth status: %w", err)
		}
		if !s.Authorized {
			return fmt.Errorf("error authorizing user: %w", err)
		}

		res, err := tgc.Client.API().MessagesGetDialogs(ctx, &tg.MessagesGetDialogsRequest{
			OffsetDate: 0,
			OffsetID:   0,
			OffsetPeer: &tg.InputPeerEmpty{},
			Limit:      70,
			Hash:       0,
		})

		if err != nil {
			return fmt.Errorf("failed to get dialogs: %w", err)
		}

		switch dialogs := res.(type) {
		case *tg.MessagesDialogsSlice:
			for _, chat := range dialogs.Chats {
				if channel, ok := chat.(*tg.Channel); ok {
					if checkChannel(int(channel.ID), requirechannels) == true {
						var parseChannel entity.Channel
						date := time.Unix(int64(channel.Date), 0)
						parseChannel.Tittle = channel.Title
						parseChannel.ChannelID = int(channel.ID)
						parseChannel.ChannelDate = date.Format("2006-01-02 15:04:05")
						parseChannel.ParticipantsCount = channel.ParticipantsCount

						if photo, ok := channel.Photo.(*tg.ChatPhoto); ok {
							parseChannel.Photo = photo.StrippedThumb
						}

						inputChannel := &tg.InputChannel{
							ChannelID:  channel.ID,
							AccessHash: channel.AccessHash,
						}

						fullChannel, err := tgc.Client.API().ChannelsGetFullChannel(ctx, inputChannel)
						if err != nil {
							continue
						}
						switch full := fullChannel.FullChat.(type) {
						case *tg.ChannelFull:
							parseChannel.MessageCount = full.ReadInboxMaxID
							parseChannel.About = full.About
						default:

						}
						history, err := tgc.Client.API().MessagesGetHistory(ctx, &tg.MessagesGetHistoryRequest{
							Peer: &tg.InputPeerChannel{
								ChannelID:  channel.ID,
								AccessHash: channel.AccessHash,
							},
							Limit: 10,
						})

						var messages []tg.MessageClass

						switch h := history.(type) {
						case *tg.MessagesChannelMessages:
							messages = h.Messages
						case *tg.MessagesMessages:
							messages = h.Messages
						case *tg.MessagesMessagesSlice:
							messages = h.Messages
						default:
							return fmt.Errorf("неизвестный тип ответа: %T", history)
						}
						msgs := make([]entity.Post, 0, 100)
						for _, msg := range messages {
							var post entity.Post
							message, ok := msg.(*tg.Message)
							if !ok {
								continue
							}
							post.ChannelID = message.ID
							date := time.Unix(int64(message.Date), 0)
							post.Date = date.Format("2006-01-02 15:04:05")
							post.Views = message.Views
							em := make([]entity.Emoji, 0)
							if len(message.Reactions.Results) > 0 {
								c := 0
								for _, reaction := range message.Reactions.Results {
									em = append(em, entity.Emoji{fmt.Sprintf("%s", reaction.Reaction), reaction.Count})
									c += reaction.Count
								}
								post.EmojisCount = c
							}
							post.Emojis = em
							msgs = append(msgs, post)
						}
						parseChannel.Messages = msgs
						channels = append(channels, parseChannel)
					}
				}
			}
		case *tg.MessagesDialogsNotModified:
			fmt.Println("Список диалогов не изменился.")
		default:
			fmt.Println("Неизвестный тип ответа.")
		}

		return nil

	})

	return ctx.Status(http.StatusOK).JSON(channels)
}

func checkChannel(channelID int, requirechannels []entity.ChannelInfo) bool {
	for _, val := range requirechannels {
		if val.ChannelID == channelID {
			return true
		}
	}
	return false
}
