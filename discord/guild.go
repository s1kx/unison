package discord

import (
	"time"

	"gopkg.in/bwmarrin/Discordgo.v0"
)

type Guild struct {
	ID                          uint64         `json:"id"`
	Name                        string         `json:"name"`
	Icon                        string         `json:"icon"`
	Region                      string         `json:"region"`
	AfkChannelID                uint64         `json:"afk_channel_id"`
	EmbedChannelID              uint64         `json:"embed_channel_id"`
	OwnerID                     uint64         `json:"owner_id"`
	JoinedAt                    time.Time      `json:"joined_at"`
	Splash                      string         `json:"splash"`
	AfkTimeout                  int            `json:"afk_timeout"`
	MemberCount                 uint           `json:"member_count"`
	VerificationLevel           int            `json:"verification_level"`
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

func NewGuildFromDiscordgo(guild *discordgo.Guild) *Guild {
	return &Guild{
		ID: discordIDStringToUint64(guild.ID),
	}
}
