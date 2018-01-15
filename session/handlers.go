package session

import (
	"strings"

	"github.com/sirupsen/logrus"

	"github.com/s1kx/discordgo"
	"github.com/s1kx/unison/state"
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
		permissions := &DiscordPermissions{}
		permissions.Set(uint64(memberPermissions))
		if err != nil || !cmd.invokableWithPermissions(permissions) {
			logrus.Info("[unison] User " + m.Author.String() + " tried to invoke unaccessable command: " + cmd.Name + ". User permission: " + permissions.String())
			break //command was found but permission was denied, so just stop looking for another command
		}

		logrus.Info("[unison] Invoking command " + cmd.Name)
		cmd.invoke(ctx, m.Message, request)

		// command was found, stop looping
		break
	}
}

func onGuildJoin(s *discordgo.Session, event *discordgo.GuildCreate) {
	// NOTE! don't run if the config.DisableBoltDatabase is true
	if event.Guild.Unavailable {
		return
	}

	// Add this guild to the database
	guildID := event.Guild.ID
	st, err := state.GetGuildState(guildID)
	if err != nil {
		// should this be handled? 0.o
	}
	if st == state.MissingState {
		selectedState := defaultGuildState
		err := state.SetGuildState(guildID, selectedState)
		if err != nil {
			logrus.Error("Unable to set default state for guild " + event.Guild.Name)
		} else {
			logrus.Infof("Joined Guild `%s`, and set state to `%s`", event.Guild.Name, state.ToStr(selectedState))
		}
	} else {
		logrus.Infof("Checked Guild `%s`, with state `%s`", event.Guild.Name, state.ToStr(st))
	}
}
