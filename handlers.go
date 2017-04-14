package unison

import (
	"strings"

	"github.com/Sirupsen/logrus"
	"github.com/bwmarrin/discordgo"
)

func handleMessageCreate(ctx *Context, ds *discordgo.Session, m *discordgo.MessageCreate) {
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

	// Only handle commands with the right prefix
	legitCommandPrefix, request := identifiesAsCommand(content, ctx)
	if !legitCommandPrefix { // will be changed later, when custom events are implemented.
		return
	}
	// end [SwissCheeze]


	// Find command, if exists
	for name, cmd := range ctx.Bot.commandMap {
		if !strings.HasPrefix(request, name) {
			continue
		}

		request = removePrefix(request, name)

		// Invoke command
		err := cmd.Action(ctx, ds, m.Message, request)
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

	// Bot mention in a string value
	botMention := "<@" + ctx.Bot.User.ID + ">"

	// Considtions that must be met and the content afterwards.
	// as of now this can be switched out with an array that just contains the prefixes
	// But leave this for now to not optimize something that we might regret
	//
	// Optimized version for this function:
	// prefixes := []string {
	//	CommandPrefix,
	//	botMention  // or just: "<@" + ctx.Bot.User.ID + ">"
	// }
	//
	// for prefix := range prefixes {
	//	if strings.HasPrefix(content, prefix) {
	//		return (true, removePrefix(content, prefix))
	//	}		
	// }
	//
	referenced := map[bool]string {
		strings.HasPrefix(content, CommandPrefix): removePrefix(content, CommandPrefix),
		strings.HasPrefix(content, botMention): removePrefix(content, botMention),
	}

	for success, updatedContent := range referenced {
		if (success) {
			return success, updatedContent
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
