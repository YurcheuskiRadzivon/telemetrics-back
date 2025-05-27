package request

type UpdateViewOptParams struct {
	Token             string `json:"token"`
	ChannelCount      int    `json:"channel_count"`
	Tittle            bool   `json:"tittle"`
	About             bool   `json:"about"`
	ChannelID         bool   `json:"channel_id"`
	ChannelDate       bool   `json:"channel_date"`
	ParticipantsCount bool   `json:"participants_count"`
	Photo             bool   `json:"photo"`
	MessageCount      int    `json:"message_count"`
	MessageID         bool   `json:"message_id"`
	Views             bool   `json:"views"`
	PostDate          bool   `json:"post_date"`
	ReactionsCount    bool   `json:"reactions_count"`
	Reactions         bool   `json:"reactions"`
}
