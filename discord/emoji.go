package discord

import "github.com/s1kx/unison/discord/snowflake"

type Emoji struct {
	ID            snowflake.ID   `json:"id"`
	Name          string         `json:"name"`
	User          *User          `json:"user"` // the user who created the emoji
	Roles         []snowflake.ID `json:"roles"`
	RequireColons bool           `json:"require_colons"`
	Managed       bool           `json:"managed"`
}
