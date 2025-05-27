package entity

type Post struct {
	ChannelName string
	ChannelID   int
	Emojis      []Emoji
	EmojisCount int
	Views       int
	Date        string
}
