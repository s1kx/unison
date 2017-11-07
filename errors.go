package unison

import "fmt"

// Duplicate command error
type DuplicateCommandError struct {
	Existing *Command
	New      *Command
	Name     string
}

func (e DuplicateCommandError) Error() string {
	return fmt.Sprintf("commands: name/alias '%s' already exists", e.Name)
}

// Duplicate event hook error
type DuplicateEventHookError struct {
	Existing *EventHook
	New      *EventHook
	Name     string
}

func (e DuplicateEventHookError) Error() string {
	return fmt.Sprintf("event hooks: name '%s' already exists", e.Name)
}

// Duplicate event hook error
type DuplicateServiceError struct {
	Existing *Service
	New      *Service
	Name     string
}

func (e DuplicateServiceError) Error() string {
	return fmt.Sprintf("service: name '%s' already exists", e.Name)
}

// Duplicate command prefix error
type DuplicateCommandPrefixError struct {
	Prefix string
}

func (e DuplicateCommandPrefixError) Error() string {
	return fmt.Sprintf("command prefix: '%s' already exists", e.Prefix)
}

// Too short command prefix error
type TooShortCommandPrefixError struct {
	Prefix string
}

func (e TooShortCommandPrefixError) Error() string {
	return fmt.Sprintf("command prefix: '%s' is too short, must me minimum 1 char", e.Prefix)
}
