package discord

type Webhook struct {
	ID        uint64 `json:"id"`
	GuildID   uint64 `json:"guild_id"`
	ChannelID uint64 `json:"channel_id"`
	User      *User  `json:"user"`
	Name      string `json:"name"`
	Avatar    string `json:"avatar"`
	Token     string `json:"token"`
}
