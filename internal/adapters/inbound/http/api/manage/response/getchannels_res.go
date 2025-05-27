package response

import "github.com/YurcheuskiRadzivon/telemetrics-back/internal/core/entity"

type GetChannelsRes struct {
	Channels []entity.ChannelInfo `json:"channels"`
}
