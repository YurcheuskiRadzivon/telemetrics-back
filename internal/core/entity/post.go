package entity

type Post struct {
	ChannelName string
	ChannelID   int
	Emojis      []Emoji
	Views       int
	Date        string
}
