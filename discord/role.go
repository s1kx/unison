package discord

import "github.com/s1kx/unison/twitter/snowflake"

type Role struct {
	ID          snowflake.ID `json:"id"`
	Name        string       `json:"name"`
	Managed     bool         `json:"managed"`
	Mentionable bool         `json:"mentionable"`
	Hoist       bool         `json:"hoist"`
	Color       int          `json:"color"`
	Position    int          `json:"position"`
	Permissions uint64       `json:"permissions"`
}

func NewRole() *Role {
	return &Role{}
}
