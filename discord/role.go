package discord

type Role struct {
	ID          Snowflake `json:"id"`
	Name        string    `json:"name"`
	Managed     bool      `json:"managed"`
	Mentionable bool      `json:"mentionable"`
	Hoist       bool      `json:"hoist"`
	Color       int       `json:"color"`
	Position    int       `json:"position"`
	Permissions uint64    `json:"permissions"`
}

func NewRole() *Role {
	return &Role{}
}
