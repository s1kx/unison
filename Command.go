package unison

import (
	"errors"
	"fmt"

	"github.com/Sirupsen/logrus"
	arg "github.com/alexflint/go-arg"
	"github.com/andersfylling/unison/constant"
	"github.com/bwmarrin/discordgo"
)

// CommandAction command logic to be executed
type CommandAction func(ctx *Context, msg *discordgo.Message, request string) error

// Command struct holds all the command details
type Command struct {
	// Name is the command title used to trigger the CommandAction
	Name string

	// Aliases are alternative command names. Usually shorter versions of the main.
	//Aliases []string. Should be stored on a per server basis, not on a per command basis.

	// Usage is a description of the command goal
	Usage string

	// Flags are optionals and bools that can be used in triggering the command
	Flags interface{} // used to generate a *arg.Parser

	// Action is the func run to execute the command
	Action CommandAction

	// Sub commands
	SubCommands []*Command

	// Deactivated true if the command should be ignored and viewed as "dead"
	Deactivated bool

	// Set the minimum required permissions for this command
	//	This level is inherited into each subcommand and must be overwritten if else is desired
	//	https://discordapp.com/developers/docs/topics/permissions
	Permissions uint32

	// Private
	//

	// go-arg parser for user input
	flagParser *arg.Parser
}

// buildCommand configurates the command from the public fields before it gets stored.
func (cmd *Command) buildCommand() *Command {
	errArr := []error{}

	// Build the command in layers..

	// Build a go-arg.Parser to handle the flags
	errArr = append(errArr, cmd.createParser(cmd.Flags))

	// make sure the depth of sub commands is acceptable
	errArr = append(errArr, cmd.insistSubCommandDepth(constant.SubCommandDepthLimit))

	// Not sure if this matters, but make sure the User can write in the same channel that the bot
	// gets triggered from. Some one that cannot send a message should never be able to trigger a command.
	if cmd.Permissions == 0 {
		cmd.Permissions = 0x00000800
	}

	// check for issues
	for _, err := range errArr {
		if err != nil {
			logrus.Fatal(err)
		}
	}

	// completed and usable command instance
	return cmd
}

// createParser creates a go-arg parser that can be used to parse user input for flags/optionals
func (cmd *Command) createParser(dests ...interface{}) error {
	p, err := arg.NewParser(arg.Config{}, dests)
	if err != nil {
		return err
	}

	cmd.flagParser = p

	return nil
}

// insistSubCommandDepth Make sure there aren't infinite depth of sub commands
func (cmd *Command) insistSubCommandDepth(depth int) error {

	if depth <= 0 && len(cmd.SubCommands) > 0 {
		errMsg := fmt.Sprintf("Too many recursive sub commands. Max depth is %d", constant.SubCommandDepthLimit)
		return errors.New(errMsg)
	}

	for _, c := range cmd.SubCommands {
		err := c.insistSubCommandDepth(depth - 1)
		if err != nil {
			return err
		}
	}

	return nil
}

// invokableWithPermissions checks if the permissions given has the minimum access level
func (cmd *Command) invokableWithPermissions(permissions uint32) bool {
	return (cmd.Permissions & permissions) == permissions
}
