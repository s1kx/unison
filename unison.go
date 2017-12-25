package unison

import (
	"errors"
	"os"
	"strconv"
	"strings"

	"github.com/Sirupsen/logrus"
	"github.com/andersfylling/unison/constant"
	"github.com/andersfylling/unison/state"
	"github.com/bwmarrin/discordgo"
)

// Run start the bot. Connect to discord, setup commands, hooks and services.
func Run(settings *Config) error {
	// TODO: Validate commands

	// three steps are done before setting up a connection.
	// 1. Make sure the discord token exists.
	// 2. Set a prefered way of triggering commands.
	//		This must be done after establishin a discord socket. See bot.onReady()
	// 3. Decide the bot state.

	// 1. Make sure the discord token exists.
	//
	token := settings.Token
	// if it was not specified in the Settings struct, check the environment variable
	if token == "" {
		token = os.Getenv(constant.EnvUnisonDiscordToken)

		// if the env var was empty as well, crash the bot as this is required.
		if token == "" {
			return errors.New("Missing env var " + constant.EnvUnisonDiscordToken + ". This is required. Specify in either Settings struct or env var.")
		}

		logrus.Info("Using bot token from environment variable.")
	}
	// discordgo requires "Bot " prefix for Bot applications
	if !strings.HasPrefix(token, constant.DiscordGoBotTokenPrefix) {
		token = constant.DiscordGoBotTokenPrefix + token
	}

	// Initialize discord client
	ds, err := discordgo.New(token)
	if err != nil {
		return err
	}

	// 2. Set a prefered way of triggering commands
	//
	envPrefix := os.Getenv(constant.EnvUnisonCommandPrefix)
	if envPrefix != "" {
		settings.CommandPrefix = append(settings.CommandPrefix, envPrefix)
	}
	var cmdPrefixes string
	for _, prefix := range settings.CommandPrefix {
		cmdPrefixes += "; " + prefix
	}
	if !settings.DisableMentionTrigger {
		cmdPrefixes += "; And by @mention."
	}
	logrus.Info("Commands are triggered by" + cmdPrefixes)

	// 3. Decide the bot state.
	//
	uState := settings.BotState
	// check if valid state
	if uState == state.MissingState {
		// chjeck environment variable
		uStateStr := os.Getenv(constant.EnvUnisonState)

		if uStateStr == "" {
			uState = state.Normal // uint8(1)
		} else {
			i, e := strconv.ParseInt(uStateStr, 10, 16)
			if e != nil {
				return e
			}

			uState = state.Type(i)
		}
	}
	state.DefaultState = uState

	// Initialize and start bot
	bot, err := newBot(settings, ds)
	if err != nil {
		return err
	}
	bot.Run()

	return nil
}
