package unison

import (
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/Sirupsen/logrus"
	"github.com/bwmarrin/discordgo"

	"github.com/s1kx/unison/events"
)

var termSignal chan os.Signal

// BotSettings contains the definition of bot behavior.
// It is used for creating the actual bot.
type BotSettings struct {
	Token string

	Commands   []*Command
	EventHooks []*EventHook
	Services   []*Service

	// Every option here is added to an array of accepted prefixes later
	// So you can set values in CommandPrefix and/or CommandPrefixes at the same time
	// This also regards the option CommandInvokedByMention
	CommandPrefix           string
	CommandPrefixes         []string
	CommandInvokedByMention bool
}

func RunBot(settings *BotSettings) error {
	// TODO: Validate commands

	// discordgo requires "Bot " prefix for Bot applications
	token := settings.Token
	if !strings.HasPrefix(token, "Bot ") {
		token = "Bot " + token
	}

	// Initialize discord client
	ds, err := discordgo.New(token)
	if err != nil {
		return err
	}

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

	// Contains a generated array of accepted prefixes based on BotSettings
	commandPrefixes         []string
	commandInvokedByMention bool

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

		commandPrefixes:         []string{},
		commandInvokedByMention: settings.CommandInvokedByMention,
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

	// Generate the array of accepted command prefixes.
	// Use a channel to generate one for bot mentions, or add it later
	if settings.CommandPrefix != "" {
		settings.CommandPrefixes = append(settings.CommandPrefixes, settings.CommandPrefix)
	}

	for _, prefix := range settings.CommandPrefixes {
		err := bot.RegisterCommandPrefix(prefix)
		if err != nil {
			return nil, err
		}
	}

	// Make sure at least one prefix exists, let's use "!" as default if none was given.
	if len(bot.CommandPrefixes) == 0 && !settings.CommandInvokedByMention {
		err := bot.RegisterCommandPrefix(DefaultCommandPrefix) // See command.go
		if err != nil {
			return nil, err
		}
	}

	return bot, nil
}

// Get a data value from existing services
func (bot *Bot) GetServiceData(srvName string, key string) string {
	if val, ok := bot.serviceMap[srvName]; ok {
		if d, s := val.Data[key]; s {
			// key exist
			return d
		}
	}

	return ""
}

func (bot *Bot) SetServiceData(srvName string, key string, val string) string {
	if v, ok := bot.serviceMap[srvName]; ok {
		if _, s := v.Data[key]; s {
			bot.serviceMap[srvName].Data[key] = val

			return val
		}
	}

	return ""
}

func (bot *Bot) Run() error {
	// Add handler to wait for ready state in order to initialize the bot fully.
	bot.Discord.AddHandler(bot.onReady)

	// Open the websocket and begin listening.
	fmt.Print("Opening WS connection to Discord .. ")
	err := bot.Discord.Open()
	if err != nil {
		return fmt.Errorf("error: %s", err)
	}
	fmt.Println("OK")

	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	termSignal = make(chan os.Signal, 1)
	signal.Notify(termSignal, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-termSignal
	fmt.Println("\nShutting down bot..")

	// Cleanly close down the Discord session.
	fmt.Print("\tClosing WS discord connection .. ")
	err = bot.Discord.Close()
	if err != nil {
		fmt.Println("ERROR")
		return err
	}

	fmt.Println("OK")

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

func (bot *Bot) RegisterCommandPrefix(prefix string) error {
	// The prefix must have a length of minimum 1
	if len(prefix) < 1 {
		return &TooShortCommandPrefixError{Prefix: prefix}
	}

	// Dont add one that already exists
	exists := false
	for _, existingPrefix := range bot.commandPrefixes {
		if existingPrefix == prefix {
			exists = true
			break
		}
	}

	if !exists {
		// Add the new prefix entry.
		bot.commandPrefixes = append(bot.commandPrefixes, prefix)
	} else {
		// Was not able to add the prefix because it already exists.
		return &DuplicateCommandPrefixError{Prefix: prefix}
	}

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

	// Add a command prefix based on the Bot ID if commandInvokedByMention is set to true
	if bot.commandInvokedByMention {
		bot.RegisterCommandPrefix("<@" + bot.User.ID + ">")
		//TODO[BLOCKER]: What if this fails?
	}

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
