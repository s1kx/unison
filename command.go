package unison

import "github.com/bwmarrin/discordgo"

const DefaultCommandPrefix = "!"

// CommandActionFunc is the action to execute when a command is called.
type CommandActionFunc func(ctx *Context, m *discordgo.Message, content string) error

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

	// Command permissions
	Permission CommandPermission
}

// Checks if given user has permission to use this command.
func (cmd Command) deniedUserAccess(author *discordgo.User) bool {
	id := author.ID

	// verify that user hasn't been banned from using command
	for _, v := range cmd.Permission.BannedUserIDs {
		if id == v || "*" == v {
			return true
		}
	}

	// if not match is found:
	return false
}

// Checks if given user has permission to use this command.
func (cmd Command) invokableByUser(author *discordgo.User) bool {
	id := author.ID

	// check if this user has a unique access
	for _, v := range cmd.Permission.AllowedUserIDs {
		if id == v || "*" == v {
			return true
		}
	}

	// if not match is found:
	return false
}

// Checks if given user has permission to use this command.
func (cmd Command) invokableByMember(member *discordgo.Member) bool {

	// check if user has access
	if cmd.deniedUserAccess(member.User) {
		return false
	}

	// check if the user role has access.
	// Sort every role to speed up this process
	// check the first char
	for _, ur := range member.Roles {
		for _, ar := range cmd.Permission.AllowedRoles {
			urc := ur[0]
			arc := ar[0]

			if urc != arc {
				continue
			} else if ur == ar || "*" == ar {
				return true
			} else if urc < arc {
				// since the roles are sorted, the first char in ar (accepted roles) should never
				// be higher than the ur (user roles) first char.
				break
			}
		}
	}

	// he might not have the role, but what if he has special permissions
	if cmd.invokableByUser(member.User) {
		return true
	}

	// if not match is found:
	return false
}
