package unison

import "fmt"

type DuplicateCommandError struct {
	Existing *Command
	New      *Command
	Name     string
}

func (e DuplicateCommandError) Error() string {
	return fmt.Sprintf("commands: name/alias '%s' already exists", e.Name)
}

type DuplicateEventHookError struct {
	Existing *EventHook
	New      *EventHook
	Name     string
}

func (e DuplicateEventHookError) Error() string {
	return fmt.Sprintf("event hooks: name '%s' already exists", e.Name)
}
