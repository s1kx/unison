package unison

import "github.com/s1kx/unison/state"

// Config defines individual behaviors of the bot.
// This type is meant to be populated by either reading from a file or from
// the command line.
// Use the `Bot` type to define programmatic bot behaviors.
type Config struct {
	// Name is the canonical name for the application/bot.
	// This is used for default values of environment variables, storage paths
	// and logging.
	Name string

	// Token is the Discord Token for this bot
	Token string

	CommandPrefixes []string

	// BotState is not sure why it's here
	BotState state.Type

	// Specify a custom environment prefix (optional).
	// If not given, the Name is capitalized and separated by underscores.
	// For instance, if Name is 'AmazingBot', environment variables would be taken
	// in the format of `AMAZING_BOT_TOKEN`.
	EnvPrefix *string

	// Unison generates a boltd key/value database with each guild as a bucket by default
	// disable this behavior to avoid file generation and extra logic on joining guilds
	DisableBoltDatabase bool

	// DisableMentionTrigger when true, bot commands won't execute when using mention as prefix
	DisableMentionTrigger bool
}
