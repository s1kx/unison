package unison

import (
	"errors"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/Sirupsen/logrus"
	"github.com/andersfylling/unison/state"
	"github.com/bwmarrin/discordgo"

	"github.com/s1kx/unison/events"
)

// EnvUnisonDiscordToken environment string to collect discord bot token.
const EnvUnisonDiscordToken = "UNISON_DISCORD_TOKEN"

// EnvUnisonCommandPrefix The command prefix to trigger commands. Defaults to mention. @botname
const EnvUnisonCommandPrefix = "UNISON_COMMAND_PREFIX"

// EnvUnisonState the default bot state. Defaults to 0. "normal"
const EnvUnisonState = "UNISON_STATE"

// DiscordGoBotTokenPrefix discordgo requires this token prefix
const DiscordGoBotTokenPrefix = "Bot "

// used to detect interupt signals and handle graceful shut down
var termSignal chan os.Signal

// BotSettings contains the definition of bot behavior.
// It is used for creating the actual bot.
type BotSettings struct {
	Token         string
	CommandPrefix string
	BotState      state.Type

	Commands   []*Command
	EventHooks []*EventHook
	Services   []*Service
}

// Run start the bot. Connect to discord, setup commands, hooks and services.
func Run(settings *BotSettings) error {
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
		token = os.Getenv(EnvUnisonDiscordToken)

		// if the env var was empty as well, crash the bot as this is required.
		if token == "" {
			return errors.New("Missing env var " + EnvUnisonDiscordToken + ". This is required. Specify in either Settings struct or env var.")
		}

		logrus.Info("Using bot token from environment variable.")
	}
	// discordgo requires "Bot " prefix for Bot applications
	if !strings.HasPrefix(token, DiscordGoBotTokenPrefix) {
		token = DiscordGoBotTokenPrefix + token
	}

	// Initialize discord client
	ds, err := discordgo.New(token)
	if err != nil {
		return err
	}

	// 2. Please see bot.onReady()
	//

	// 3. Decide the bot state.
	//
	uState := settings.BotState
	// check if valid
	// TODO: create a key=value database to handle states.
	//			 This must be guild related. so each key represents a guild ID and value of its current state.

	// Initialize and start bot
	bot, err := newBot(settings, ds)
	if err != nil {
		return err
	}
	bot.Run()

	return nil
}

// Bot is an active bot session.
type Bot struct {
	*BotSettings
	Discord *discordgo.Session

	// Lookup map for name/alias => command
	commandMap map[string]*Command
	// Lookup map for name => hook
	eventHookMap map[string]*EventHook
	// Lookup map for name => service
	serviceMap map[string]*Service

	eventDispatcher *eventDispatcher

	// Command prefix
	commandPrefix string

	readyState *discordgo.Ready
	User       *discordgo.User
}

func newBot(settings *BotSettings, ds *discordgo.Session) (*Bot, error) {
	// Initialize bot
	bot := &Bot{
		BotSettings: settings,
		Discord:     ds,

		commandMap:      make(map[string]*Command),
		eventHookMap:    make(map[string]*EventHook),
		serviceMap:      make(map[string]*Service),
		eventDispatcher: newEventDispatcher(),

		commandPrefix: settings.CommandPrefix,
	}

	// Register commands
	for _, cmd := range bot.Commands {
		err := bot.RegisterCommand(cmd)
		if err != nil {
			return nil, err
		}
	}

	// Register event hooks
	for _, hook := range bot.EventHooks {
		err := bot.RegisterEventHook(hook)
		if err != nil {
			return nil, err
		}
	}

	// Register services
	for _, srv := range bot.Services {
		err := bot.RegisterService(srv)
		if err != nil {
			return nil, err
		}
	}

	return bot, nil
}

// GetServiceData a data value from existing services
func (bot *Bot) GetServiceData(srvName string, key string) string {
	if val, ok := bot.serviceMap[srvName]; ok {
		if d, s := val.Data[key]; s {
			// key exist
			return d
		}
	}

	return ""
}

// SetServiceData update or set a new value for a given service key
func (bot *Bot) SetServiceData(srvName string, key string, val string) string {
	if v, ok := bot.serviceMap[srvName]; ok {
		if _, s := v.Data[key]; s {
			bot.serviceMap[srvName].Data[key] = val

			return val
		}
	}

	return ""
}

// Run Start the bot instance
func (bot *Bot) Run() error {
	// Add handler to wait for ready state in order to initialize the bot fully.
	bot.Discord.AddHandler(bot.onReady)

	// Open the websocket and begin listening.
	logrus.Info("Opening WS connection to Discord .. ")
	err := bot.Discord.Open()
	if err != nil {
		return fmt.Errorf("error: %s", err)
	}
	logrus.Info("OK")

	logrus.Info("Bot is now running.  Press CTRL-C to exit.")
	termSignal = make(chan os.Signal, 1)
	signal.Notify(termSignal, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-termSignal
	logrus.Info("\nShutting down bot..")

	// Cleanly close down the Discord session.
	logrus.Info("\tClosing WS discord connection .. ")
	err = bot.Discord.Close()
	if err != nil {
		return err
	}

	logrus.Info("Shutdown successfully")

	return nil
}

func (bot *Bot) RegisterCommand(cmd *Command) error {
	name := cmd.Name
	if ex, exists := bot.commandMap[name]; exists {
		return &DuplicateCommandError{Existing: ex, New: cmd, Name: name}
	}
	bot.commandMap[name] = cmd

	// TODO: Register aliases

	return nil
}

func (bot *Bot) RegisterEventHook(hook *EventHook) error {
	name := hook.Name
	if ex, exists := bot.eventHookMap[name]; exists {
		return &DuplicateEventHookError{Existing: ex, New: hook}
	}
	bot.eventHookMap[name] = hook

	if len(hook.Events) == 0 {
		logrus.Warnf("Hook '%s' is not subscribed to any events", name)
	}

	bot.eventDispatcher.AddHook(hook)

	return nil
}

func (bot *Bot) RegisterService(srv *Service) error {
	name := srv.Name
	if ex, exists := bot.serviceMap[name]; exists {
		return &DuplicateServiceError{Existing: ex, New: srv, Name: name}
	}
	bot.serviceMap[name] = srv

	return nil
}

func (bot *Bot) onReady(ds *discordgo.Session, r *discordgo.Ready) {
	// Set bot state
	bot.readyState = r
	bot.User = r.User

	logrus.WithFields(logrus.Fields{
		"ID":       r.User.ID,
		"Username": r.User.Username,
	}).Infof("Bot is connected and running.")

	// 2. Set a prefered way of triggering commands
	//
	cprefix := bot.CommandPrefix
	// if not given, check the environment variable
	if cprefix == "" {
		cprefix = os.Getenv(EnvUnisonCommandPrefix)

		// in case this was not set, we trigger by mention
		if cprefix == "" {
			cprefix = ds.State.User.Mention()
		}

		// update Settings
		bot.CommandPrefix = cprefix
	}
	logrus.Info("Commands are triggered by `" + cprefix + "`")

	// Create context for services
	ctx := NewContext(bot, ds, termSignal)

	// Run services
	for _, srv := range bot.serviceMap {
		if srv.Deactivated {
			continue
		}

		// run service
		go srv.Action(ctx)
	}

	// Add generic handler for event hooks
	// Add command handler
	bot.Discord.AddHandler(func(ds *discordgo.Session, event interface{}) {
		bot.onEvent(ds, event)
	})
}

func (bot *Bot) onEvent(ds *discordgo.Session, dv interface{}) {
	// Inspect and wrap event
	ev, err := events.NewDiscordEvent(dv)
	if err != nil {
		logrus.Errorf("event handler: %s", err)
	}

	// Create context for handlers
	ctx := NewContext(bot, ds, termSignal)

	// Invoke event hooks for the hooks that are subscribed to the event type
	bot.eventDispatcher.Dispatch(ctx, ev)

	// Invoke command handler on new messages
	if ev.Type == events.MessageCreateEvent {
		handleMessageCreate(ctx, ev.Event.(*discordgo.MessageCreate))
	}
}
