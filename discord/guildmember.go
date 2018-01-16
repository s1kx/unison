package discord

import (
	"time"

	"github.com/s1kx/unison/discord/snowflake"
)

type GuildMember struct {
	GuildID  snowflake.ID   `json:"guild_id"`
	JoinedAt time.Time      `json:"joined_at"`
	Nick     string         `json:"nick"`
	Deaf     bool           `json:"deaf"`
	Mute     bool           `json:"mute"`
	User     *User          `json:"user"`
	Roles    []snowflake.ID `json:"roles"`
}
