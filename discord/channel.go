package discord

import (
	"fmt"

	"github.com/s1kx/unison/twitter/snowflake"
)

type Channel struct {
	ID                   snowflake.ID           `json:"id"`
	GuildID              snowflake.ID           `json:"guild_id"`
	Name                 string                 `json:"name"`
	Topic                string                 `json:"topic"`
	Type                 uint                   `json:"type"`
	LastMessageID        snowflake.ID           `json:"last_message_id"`
	NSFW                 bool                   `json:"nsfw"`
	Position             uint                   `json:"position"`
	Bitrate              int                    `json:"bitrate"`
	Recipients           []*User                `json:"recipient"`
	Messages             []*Message             `json:"-"`
	PermissionOverwrites []*PermissionOverwrite `json:"permission_overwrites"`
}

func NewChannel() *Channel {
	return &Channel{}
}

func (c *Channel) Mention() string {
	return fmt.Sprintf("<#%d>", c.ID)
}

type PermissionOverwrite struct {
	ID    snowflake.ID `json:"id"`    // role or user id
	Type  string       `json:"type"`  // either `role` or `member`
	Deny  int          `json:"deny"`  // permission bit set
	Allow int          `json:"allow"` // permission bit set
}
