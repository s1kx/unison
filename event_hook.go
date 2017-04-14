package unison

import "github.com/bwmarrin/discordgo"

// EventHandlerFunc handles a discord event and returns whether it handles the
// event type and if an error occured.
type EventHandlerFunc func(ctx *Context, ds *discordgo.Session, event interface{}) (handled bool, err error)

// Hook interface for anything that is supposed to react on a event, besides commands.
type EventHook struct {
	// Name of the hook
	Name string

	// Description of what the hook does
	Description string

	// Events that the hook should react to
	// Events []string

	// Check if this hook is deactivated
	Deactivated bool

	// Command behavior
	OnEvent EventHandlerFunc
}
