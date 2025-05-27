package entity

type Channel struct {
	Tittle            string  `json:"tittle"`
	About             string  `json:"about"`
	ChannelID         int     `json:"channel_id"`
	ChannelDate       string  `json:"channel_date"`
	ParticipantsCount int     `json:"participants_count"`
	Photo             []byte  `json:"photo"`
	MessageCount      int     `json:"message_count"`
	MessageID         int     `json:"message_id"`
	Views             int     `json:"views"`
	PostDate          string  `json:"post_date"`
	ReactionsCount    int     `json:"reactions_count"`
	Reactions         []Emoji `json:"reactions"`
	Messages          []Post  `json:"messages"`
}
