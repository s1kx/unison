package unison

import "github.com/s1kx/unison/state"

// Config defines individual behaviors of the bot.
// This type is meant to be populated by either reading from a file or from
// the command line.
// Use the `Bot` type to define programmatic bot behaviors.
type Config struct {
	Token string

	CommandPrefixes []string
	BotState        state.Type

	// Put an environment prefix on all environment variables
	EnvironmentPrefix string

	// Unison generates a boltd key/value database with each guild as a bucket by default
	// disable this behavior to avoid file generation and extra logic on joining guilds
	DisableBoltDatabase bool

	// DisableMentionTrigger when true, bot commands won't execute when using mention as prefix
	DisableMentionTrigger bool
}
