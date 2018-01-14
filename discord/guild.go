package discord

import (
	"time"

	"github.com/s1kx/unison/twitter/snowflake"
)

type Guild struct {
	ID                          snowflake.ID   `json:"id"`
	Name                        string         `json:"name"`
	Icon                        string         `json:"icon"`
	Region                      string         `json:"region"`
	AfkChannelID                snowflake.ID   `json:"afk_channel_id"`
	EmbedChannelID              snowflake.ID   `json:"embed_channel_id"`
	OwnerID                     snowflake.ID   `json:"owner_id"`
	JoinedAt                    time.Time      `json:"joined_at"`
	Splash                      string         `json:"splash"`
	AfkTimeout                  uint           `json:"afk_timeout"`
	MemberCount                 uint           `json:"member_count"`
	VerificationLevel           uint           `json:"verification_level"`
	EmbedEnabled                bool           `json:"embed_enabled"`
	Large                       bool           `json:"large"` // ??
	DefaultMessageNotifications int            `json:"default_message_notifications"`
	Roles                       []*Role        `json:"roles"`
	Emojis                      []*Emoji       `json:"emojis"`
	Members                     []*GuildMember `json:"members"`
	Presences                   []*Presence    `json:"presences"`
	Channels                    []*Channel     `json:"channels"`
	VoiceStates                 []*VoiceState  `json:"voice_states"`
	Unavailable                 bool           `json:"unavailable"`
}
