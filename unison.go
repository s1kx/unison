package unison

import (
	"encoding/json"
	"errors"
	"os"
	"strconv"
	"strings"

	"github.com/s1kx/discordgo"
	"github.com/s1kx/unison/constant"
	"github.com/s1kx/unison/discord"
	"github.com/s1kx/unison/state"
	"github.com/sirupsen/logrus"
)

var logFormatter = logrus.TextFormatter{
	FullTimestamp:   true,
	TimestampFormat: "2006-01-02 15:04:05",
}

var (
	ErrMissingDiscordToken = errors.New("discord token is not configured")
)

// Run start the bot. Connect to discord, setup commands, hooks and services.
func Run(settings *Config) error {
	// Configure logger format.
	logrus.SetFormatter(&logFormatter)

	// add a underscore suffix to environment prefix
	if settings.EnvironmentPrefix != "" {
		settings.EnvironmentPrefix = settings.EnvironmentPrefix + "_"
	}

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
		token = os.Getenv(settings.EnvironmentPrefix + constant.EnvUnisonDiscordToken)

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
	envPrefix := os.Getenv(settings.EnvironmentPrefix + constant.EnvUnisonCommandPrefix)
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
		uStateStr := os.Getenv(settings.EnvironmentPrefix + constant.EnvUnisonState)

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
	settings.BotState = uState

	// Initialize and start bot
	bot, err := newBot(settings, ds)
	if err != nil {
		return err
	}

	return bot.Run() // returns nil on successfull shutdown
}

// GetAuditLogs Get the last 50 audit logs for the given guild
//	params interface{} is a struct with json tags that are converted into GET url parameters
func GetAuditLogs(ctx *Context, guildID string, params interface{}) (*discord.AuditLog, error) {
	urlParams := "" //convertAuditLogParamsToStr(params)
	byteArr, err := ctx.Discord.Request("GET", discordgo.EndpointGuilds+guildID+"/audit-logs"+urlParams, nil)
	if err != nil {
		return nil, err
	}

	auditLog := &discord.AuditLog{}
	err = json.Unmarshal(byteArr, &auditLog)
	if err != nil {
		return nil, err
	}

	return auditLog, nil
}
