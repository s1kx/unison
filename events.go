package unison

import (
	"fmt"

	"github.com/s1kx/unison/discord/events"
)

// EventHandlerFunc handles a discord event and returns whether it handles the
// event type and if an error occured.
// self is true if event was fired by bot. eg bot sent a message to someone.
type EventHandlerFunc func(ctx *Context, ev *events.DiscordEvent, self bool) error

// EventHook interface for anything that is supposed to react on a event, besides commands.
type EventHook struct {
	// Name of the hook
	Name string

	// Description of what the hook does
	Usage string

	// Events that the hook should react to
	Events []events.EventType

	// Check if this hook is deactivated
	Deactivated bool

	// Command behavior
	OnEvent EventHandlerFunc
}

// DuplicateEventHookError event hook error
type DuplicateEventHookError struct {
	Existing *EventHook
	New      *EventHook
	Name     string
}

func (e DuplicateEventHookError) Error() string {
	return fmt.Sprintf("event hooks: name '%s' already exists", e.Name)
}
