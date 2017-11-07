package unison

import (
	"strings"

	"github.com/Sirupsen/logrus"
	"github.com/bwmarrin/discordgo"
)

func handleMessageCreate(ctx *Context, m *discordgo.MessageCreate) {
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
	legitCommandPrefix, request := identifiesAsCommand(content, ctx)
	channel, channelErr := ctx.Discord.Channel(m.ChannelID)
	if channelErr != nil || (channel.Type != 1 && !legitCommandPrefix) {
		return
	}
	// end [SwissCheeze]

	// Find command, if exists
	for name, cmd := range ctx.Bot.commandMap {
		if !strings.HasPrefix(request, name) {
			continue
		}

		request = removePrefix(request, name)

		// commands that works on guild related matters should not be runnable from a PM(!)
		// TODO: write guild requirement check

		// verify that user has permission to invoke this command
		member, memberErr := ctx.Discord.GuildMember(channel.GuildID, m.Author.ID)
		if memberErr != nil || !cmd.invokableByMember(member) {
			break //command was found but permission was denied, so just stop looking for another command
		}

		// Invoke command
		err := cmd.Action(ctx, m.Message, request)
		if err != nil {
			logrus.Errorf("Command [%s]: %s", name, err)
		}

		// command was found, stop looping
		break
	}
}

// Check if a message content string is a valid command by it's prefix "!" or bot mention
func identifiesAsCommand(content string, ctx *Context) (status bool, updatedContent string) {
	failure := false
	success := true

	// For every prefix entry set by botSettings, go through until a match
	for _, prefix := range ctx.Bot.commandPrefixes {
		if strings.HasPrefix(content, prefix) {
			return success, removePrefix(content, prefix)
		}
	}

	// None of the conditions were met so this is considered a failure
	return failure, content
}

// Removes a substring from the string and cleans up leading & trailing spaces.
func removePrefix(str, remove string) string {
	result := strings.TrimPrefix(str, remove)
	return strings.TrimSpace(result)
}
