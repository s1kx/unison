package main

import (
	"errors"
	"os"
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"

	"github.com/s1kx/discordgo"

	"github.com/s1kx/unison/client"
	"github.com/s1kx/unison/constant"
	"github.com/s1kx/unison/state"
)

var logFormatter = logrus.TextFormatter{
	FullTimestamp:   true,
	TimestampFormat: "2006-01-02 15:04:05",
}

var (
	ErrMissingDiscordToken = errors.New("discord token is not configured")
)

func main() {
	// Call Run()!
}

// Run is a convenience function for starting a session with the given Bot and coniguration.
func Run(bot *Bot, conf *Config) error {
	// Configure logger format.
	logrus.SetFormatter(&logFormatter)

	// add a underscore suffix to environment prefix
	if conf.EnvironmentPrefix != "" {
		conf.EnvironmentPrefix = conf.EnvironmentPrefix + "_"
	}

	// TODO: Validate commands

	// three steps are done before setting up a connection.
	// 1. Make sure the discord token exists.
	// 2. Set a prefered way of triggering commands.
	//		This must be done after establishin a discord socket. See bot.onReady()
	// 3. Decide the bot state.

	// 1. Make sure the discord token exists.
	//
	token := conf.Token
	// if it was not specified in the Settings struct, check the environment variable
	if token == "" {
		token = os.Getenv(conf.EnvironmentPrefix + constant.EnvUnisonDiscordToken)

		// if the env var was empty as well, crash the bot as this is required.
		if token == "" {
			return ErrMissingDiscordToken
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
	envPrefix := os.Getenv(conf.EnvironmentPrefix + constant.EnvUnisonCommandPrefix)
	if envPrefix != "" {
		conf.CommandPrefix = append(conf.CommandPrefix, envPrefix)
	}
	var cmdPrefixes string
	for _, prefix := range conf.CommandPrefix {
		cmdPrefixes += "; " + prefix
	}
	if !conf.DisableMentionTrigger {
		cmdPrefixes += "; And by @mention."
	}
	logrus.Info("Commands are triggered by" + cmdPrefixes)

	// 3. Decide the bot state.
	//
	uState := conf.BotState
	// check if valid state
	if uState == state.MissingState {
		// check environment variable
		uStateStr := os.Getenv(conf.EnvironmentPrefix + constant.EnvUnisonState)

		if uStateStr == "" {
			uState = state.Normal // uint8(1)
		} else {
			i, e := strconv.ParseInt(uStateStr, 10, 16)
			if e != nil {
				return e
			}

			uState = state.Type(i)

			// fallback
			if uState == state.MissingState {
				uState = state.Normal
			}
		}
	}
	conf.BotState = uState

	// Initialize and start bot
	s, err := client.New(bot, conf, ds)
	if err != nil {
		return err
	}

	return s.Run() // returns nil on successfull shutdown
}
