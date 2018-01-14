package discord

type Emoji struct {
	ID            Snowflake   `json:"id"`
	Name          string      `json:"name"`
	User          *User       `json:"user"` // the user who created the emoji
	Roles         []Snowflake `json:"roles"`
	RequireColons bool        `json:"require_colons"`
	Managed       bool        `json:"managed"`
}
