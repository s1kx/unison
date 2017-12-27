package unison

import (
	"strings"

	"github.com/Sirupsen/logrus"
	"github.com/bwmarrin/discordgo"
)

func handleMessageCreate(ctx *Context, m *discordgo.MessageCreate) {
	var err error
	content := strings.TrimSpace(m.Content)

	// start [SwissCheeze]
	// Check if message is a command and should be handled
	// Don't communicate with other bots
	if m.Author.Bot {
		return
	}

	// Don't communicate with myself
	if m.Author.ID == ctx.Bot.User.ID {
		return
	}

	// Only handle commands with the right prefix, when not in a PM channel
	// This also removes potential command prefixes just in case
	var legitCommandPrefix bool
	var request string
	for _, prefix := range ctx.Bot.CommandPrefix {
		legitCommandPrefix, request = identifiesAsCommand(content, prefix)

		if legitCommandPrefix {
			break
		}
	}
	channel, err := ctx.Discord.Channel(m.ChannelID)
	if err != nil || (channel.Type != 1 && !legitCommandPrefix) {
		return
	}
	// end [SwissCheeze]

	// Find command, if exists
	for name, cmd := range ctx.Bot.commandMap {
		if !strings.HasPrefix(request, name) {
			continue
		}

		// remove the command prefix
		//request = removePrefix(request, name) // depratecated after subcommands was added
		// TODO: call func recursively for sub commands

		// commands that works on guild related matters should not be runnable from a PM(!)
		// TODO: write guild requirement check

		// verify that user has permission to invoke this command
		memberPermissions, err := ctx.Bot.Discord.UserChannelPermissions(m.Author.ID, m.ChannelID)
		if err != nil || !cmd.invokableWithPermissions(DiscordPermissionFlags(memberPermissions)) {
			logrus.Info("[unison] User " + m.Author.String() + " tried to invoke unaccessable command: " + cmd.Name)
			break //command was found but permission was denied, so just stop looking for another command
		}

		logrus.Info("[unison] Invoking command " + cmd.Name)
		go cmd.invoke(ctx, m.Message, request)

		// command was found, stop looping
		break
	}
}

// Check if a message content string is a valid command by it's prefix "!" or bot mention
func identifiesAsCommand(content, prefix string) (status bool, updatedContent string) {
	if strings.HasPrefix(content, prefix) {
		// remove duplicate command prefixes
		result := removePrefix(content, prefix)
		for {
			if !strings.HasPrefix(result, prefix) {
				break
			} else {
				result = removePrefix(result, prefix)
			}
		}

		// make sure there is content after removing the command prefix
		if len(result) > 0 {
			return true, result
		}
		return false, result
	}

	return false, content
}

// Removes a substring from the string and cleans up leading & trailing spaces.
func removePrefix(str, remove string) string {
	result := strings.TrimPrefix(str, remove)
	result = strings.TrimSpace(result)

	if str[0] == ' ' {
		return result[1:len(result)]
	}
	return result
}
