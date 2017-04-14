package unison

import (
	"strings"

	"github.com/Sirupsen/logrus"
	"github.com/bwmarrin/discordgo"
)

func handleMessageCreate(ctx *Context, ds *discordgo.Session, m *discordgo.MessageCreate) {
	content := m.Content

	// Check if message is a command and should be handled
	switch {
	// Don't communicate with other bots
	case m.Author.Bot:
		return
	// Don't communicate with myself
	case m.Author.ID == ctx.Bot.User.ID:
		return
	// Only handle commands with the right prefix
	case !strings.HasPrefix(content, CommandPrefix):
		return
	}

	// request is the message without command prefix/bot mention/extra whitespaces.
	request := cleanUpRequest(content, CommandPrefix)

	// Find command, if exists
	for name, cmd := range ctx.Bot.commandMap {
		if !strings.HasPrefix(request, name) {
			continue
		}

		request = cleanUpRequest(request, name)

		// Invoke command
		err := cmd.Action(ctx, ds, m.Message, request)
		if err != nil {
			logrus.Errorf("Command [%s]: %s", name, err)
		}

		// command was found, stop looping
		break
	}
}

// Removes a substring from the string and cleans up leading & trailing spaces.
func cleanUpRequest(str, remove string) string {
	result := strings.TrimPrefix(str, remove)
	return strings.TrimSpace(result)
}

func handleDiscordEvent(ctx *Context, ds *discordgo.Session, event interface{}) {
	runEventHooks(ctx, ds, event)
}

// To run a hook.
// This needs to be updated to handle multiple different event types... should be somewhat of a generic, but will work for now........
func runEventHooks(ctx *Context, ds *discordgo.Session, event interface{}) {
	for _, hook := range ctx.Bot.EventHooks {
		hook.OnEvent(ctx, ds, event)
	}
}
