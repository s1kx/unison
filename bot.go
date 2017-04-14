package unison

import (
	"fmt"
	"strings"

	"github.com/Sirupsen/logrus"
	"github.com/bwmarrin/discordgo"
)

// BotSettings contains the definition of bot behavior.
// It is used for creating the actual bot.
type BotSettings struct {
	Token string

	Commands   []*Command
	EventHooks []*EventHook
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

	readyState *discordgo.Ready
	User       *discordgo.User
}

func newBot(settings *BotSettings, ds *discordgo.Session) (*Bot, error) {
	// Initialize bot
	bot := &Bot{
		BotSettings: settings,
		Discord:     ds,

		commandMap:   make(map[string]*Command),
		eventHookMap: make(map[string]*EventHook),
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

	return bot, nil
}

func (bot *Bot) Run() error {
	// Add handler to wait for ready state in order to initialize the bot fully.
	bot.Discord.AddHandler(bot.onReady)

	// Open the websocket and begin listening.
	err := bot.Discord.Open()
	if err != nil {
		return fmt.Errorf("error opening connection: %s", err)
	}

	logrus.Info("Bot is now running.  Press CTRL-C to exit.")

	// Simple way to keep program running until CTRL-C is pressed.
	// TODO: Add signal handler to exit gracefully.
	<-make(chan struct{})

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

	// Create context for handlers
	ctx := NewContext(bot)

	// Add command handler
	bot.Discord.AddHandler(func(ds *discordgo.Session, m *discordgo.MessageCreate) {
		handleMessageCreate(ctx, ds, m)
	})

	// Add generic handler for event hooks
	// Add command handler
	bot.Discord.AddHandler(func(ds *discordgo.Session, event interface{}) {
		handleDiscordEvent(ctx, ds, event)
	})
}
