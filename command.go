package unison

import "github.com/bwmarrin/discordgo"

const DefaultCommandPrefix = "!"

// CommandActionFunc is the action to execute when a command is called.
type CommandActionFunc func(ctx *Context, ds *discordgo.Session, m *discordgo.Message, content string) error

// Command is the interface that every bot command must implement.
type Command struct {
	// Name of the command
	Name string

	// Aliases for the command
	Aliases []string

	// Flags supported in this command
	// Flags []*FlagSet

	// Description of what the command does
	Description string

	// Check if this command is deactivated
	Deactivated bool

	// Command behavior
	Action CommandActionFunc
}
