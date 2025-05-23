package telegram

import (
	"github.com/gotd/td/telegram"
)

type Client struct {
	Tgc *telegram.Client
}

func New(sessionStoragePath string, appID int, appHash string) (*Client, error) {
	c := &Client{}

	opt := telegram.Options{
		SessionStorage: &telegram.FileSessionStorage{
			Path: sessionStoragePath,
		},
		Device: telegram.DeviceConfig{
			DeviceModel: "TeleMD",
		},
	}

	c.Tgc = telegram.NewClient(appID, appHash, opt)

	return c, nil
}
