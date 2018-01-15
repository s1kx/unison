package unison

import (
	"errors"
)

var (
	ErrDatabaseDisabled = errors.New("bolt(key-value) database is disabled")
)

type Bot struct {
	Commands   []*Command
	EventHooks []*EventHook
	Services   []*Service
}

var defaultBot = Bot{
// Commands, EventHooks, Services are automatically initialized as empty arrays.
}
