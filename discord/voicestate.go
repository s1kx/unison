package discord

type VoiceState struct {
	UserID    uint64 `json:"user_id"`
	SessionID uint64 `json:"session_id"`
	ChannelID uint64 `json:"channel_id"`
	GuildID   uint64 `json:"guild_id"`
	Suppress  bool   `json:"suppress"`
	SelfMute  bool   `json:"self_mute"`
	SelfDeaf  bool   `json:"self_deaf"`
	Mute      bool   `json:"mute"`
	Deaf      bool   `json:"deaf"`
}
