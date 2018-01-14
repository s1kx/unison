package discord

type Webhook struct {
	ID        Snowflake `json:"id"`
	GuildID   Snowflake `json:"guild_id"`
	ChannelID Snowflake `json:"channel_id"`
	User      *User     `json:"user"`
	Name      string    `json:"name"`
	Avatar    string    `json:"avatar"`
	Token     string    `json:"token"`
}
