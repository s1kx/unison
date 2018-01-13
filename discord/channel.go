package discord

import (
	"fmt"
)

type Channel struct {
	ID                   uint64                 `json:"id"`
	GuildID              uint64                 `json:"guild_id"`
	Name                 string                 `json:"name"`
	Topic                string                 `json:"topic"`
	Type                 uint8                  `json:"type"`
	LastMessageID        uint64                 `json:"last_message_id"`
	NSFW                 bool                   `json:"nsfw"`
	Position             uint                   `json:"position"`
	Bitrate              int                    `json:"bitrate"`
	Recipients           []*User                `json:"recipient"`
	Messages             []*Message             `json:"-"`
	PermissionOverwrites []*PermissionOverwrite `json:"permission_overwrites"`
}

func New() *Channel {
	return &Channel{}
}

func (c *Channel) Mention() string {
	return fmt.Sprintf("<#%d>", c.ID)
}

type PermissionOverwrite struct {
	ID    uint64 `json:"id"`    // role or user id
	Type  string `json:"type"`  // either `role` or `member`
	Deny  int    `json:"deny"`  // permission bit set
	Allow int    `json:"allow"` // permission bit set
}
