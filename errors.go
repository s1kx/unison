package unison

import "fmt"

// DuplicateCommandError command error
type DuplicateCommandError struct {
	Existing *Command
	New      *Command
	Name     string
}

func (e DuplicateCommandError) Error() string {
	return fmt.Sprintf("commands: name/alias '%s' already exists", e.Name)
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

// DuplicateServiceError service error
type DuplicateServiceError struct {
	Existing *Service
	New      *Service
	Name     string
}

func (e DuplicateServiceError) Error() string {
	return fmt.Sprintf("service: name '%s' already exists", e.Name)
}

// TooShortCommandPrefixError short command prefix error
type TooShortCommandPrefixError struct {
	Prefix string
}

func (e TooShortCommandPrefixError) Error() string {
	return fmt.Sprintf("command prefix: '%s' is too short, must me minimum 1 char", e.Prefix)
}
