package discord

import "github.com/s1kx/unison/discord/snowflake"

type Webhook struct {
	ID        snowflake.ID `json:"id"`
	GuildID   snowflake.ID `json:"guild_id"`
	ChannelID snowflake.ID `json:"channel_id"`
	User      *User        `json:"user"`
	Name      string       `json:"name"`
	Avatar    string       `json:"avatar"`
	Token     string       `json:"token"`
}
