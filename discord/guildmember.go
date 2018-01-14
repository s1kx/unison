package discord

import "time"

type GuildMember struct {
	GuildID  Snowflake   `json:"guild_id"`
	JoinedAt time.Time   `json:"joined_at"`
	Nick     string      `json:"nick"`
	Deaf     bool        `json:"deaf"`
	Mute     bool        `json:"mute"`
	User     *User       `json:"user"`
	Roles    []Snowflake `json:"roles"`
}
