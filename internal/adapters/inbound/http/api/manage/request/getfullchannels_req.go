package request

import "github.com/YurcheuskiRadzivon/telemetrics-back/internal/core/entity"

type GetFullChannelsReq struct {
	Token    string               `json:"token"`
	Channels []entity.ChannelInfo `json:"channels"`
}
