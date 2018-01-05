package unison

// Config contains the definition of bot behavior.
// It is used while creating/setting up the actual bot.
type Config struct {
	Token         string
	CommandPrefix []string
	BotState      uint8

	// DisableMentionTrigger when true, bot commands won't execute when using mention as prefix
	DisableMentionTrigger bool

	Commands   []*Command
	EventHooks []*EventHook
	Services   []*Service
}
