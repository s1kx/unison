package discord

type Emoji struct {
	ID            uint64   `json:"id"`
	Name          string   `json:"name"`
	User          *User    `json:"user"` // the user who created the emoji
	Roles         []uint64 `json:"roles"`
	RequireColons bool     `json:"require_colons"`
	Managed       bool     `json:"managed"`
}
