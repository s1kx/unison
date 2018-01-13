package unison

import "github.com/s1kx/unison/state"

// Config contains the definition of bot behavior.
// It is used while creating/setting up the actual bot.
type Config struct {
	Token         string
	CommandPrefix []string
	BotState      state.Type

	// Put an environment prefix on all environment variables
	EnvironmentPrefix string

	// Unison generates a boltd key/value database with each guild as a bucket by default
	// disable this behavior to avoid file generation and extra logic on joining guilds
	DisableBoltDatabase bool

	// DisableMentionTrigger when true, bot commands won't execute when using mention as prefix
	DisableMentionTrigger bool

	Commands   []*Command
	EventHooks []*EventHook
	Services   []*Service
}
