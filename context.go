package unison

// Context is a type that is passed to every handler
// in a bot application.
// It can be used to refer back to main components.
type Context struct {
	Bot *Bot
}

func NewContext(bot *Bot) *Context {
	return &Context{Bot: bot}
}
