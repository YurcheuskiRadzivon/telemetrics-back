package telegram

import (
	"github.com/YurcheuskiRadzivon/telemetrics-back/internal/adapters/outbound/repositories"
	"github.com/YurcheuskiRadzivon/telemetrics-back/pkg/generator"
	"github.com/YurcheuskiRadzivon/telemetrics-back/pkg/logger"
	"github.com/gotd/td/telegram"
)

type TelegramClient struct {
	Client *telegram.Client
	lgr    *logger.Logger
	gnrt   *generator.Generator
	sr     *repositories.SessionRepository
}

func New(appID int, appHash string, lgr *logger.Logger, gnrt *generator.Generator, sr *repositories.SessionRepository) *TelegramClient {
	c := &TelegramClient{
		lgr:  lgr,
		gnrt: gnrt,
		sr:   sr,
	}

	opt := telegram.Options{
		SessionStorage: sr,
		Device: telegram.DeviceConfig{
			DeviceModel: "TeleMD",
		},
	}

	c.Client = telegram.NewClient(appID, appHash, opt)

	return c
}
