package constant

const (
	// SubCommandDepthLimit decides the max sub command depth allowed
	SubCommandDepthLimit = 0 // 0: disable sub-command feature

	// EnvUnisonDiscordToken environment string to collect discord bot token.
	EnvUnisonDiscordToken = "UNISON_DISCORD_TOKEN"

	// EnvUnisonCommandPrefix The command prefix to trigger commands. Defaults to mention. @botname
	EnvUnisonCommandPrefix = "UNISON_COMMAND_PREFIX"

	// EnvUnisonState the default bot state. Defaults to 0. "normal"
	EnvUnisonState = "UNISON_STATE"

	// DiscordGoBotTokenPrefix discordgo requires this token prefix
	DiscordGoBotTokenPrefix = "Bot "

	// package state
	//

	StateKey = "state"

	DefaultDatabaseFile = "unisonStates.db"
)
