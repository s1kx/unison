package discord

type VoiceState struct {
	UserID    Snowflake `json:"user_id"`
	SessionID Snowflake `json:"session_id"`
	ChannelID Snowflake `json:"channel_id"`
	GuildID   Snowflake `json:"guild_id"`
	Suppress  bool      `json:"suppress"`
	SelfMute  bool      `json:"self_mute"`
	SelfDeaf  bool      `json:"self_deaf"`
	Mute      bool      `json:"mute"`
	Deaf      bool      `json:"deaf"`
}
