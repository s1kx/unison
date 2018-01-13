package unison

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/s1kx/unison/state"
	"github.com/sirupsen/logrus"

	"github.com/s1kx/unison/events"
)

// Bot is an active bot session.
type Bot struct {
	*Config
	Discord *discordgo.Session

	// Lookup map for name/alias => command
	commandMap map[string]*Command
	// Lookup map for name => hook
	eventHookMap map[string]*EventHook
	// Lookup map for name => service
	serviceMap map[string]*Service

	eventDispatcher *eventDispatcher

	// Command prefixes
	commandPrefix []string

	readyState *discordgo.Ready
	User       *discordgo.User

	state state.Type // default bot state
}

// This is awful and needs to be handled. It's used in "onGuildJoin" func
var defaultGuildState state.Type

func newBot(config *Config, ds *discordgo.Session) (*Bot, error) {
	commandPrefixes := []string{}

	// add desired prefixes
	for _, prefix := range config.CommandPrefix {
		commandPrefixes = append(commandPrefixes, prefix)
	}

	// Initialize bot
	bot := &Bot{
		Config:  config,
		Discord: ds,

		commandMap:      make(map[string]*Command),
		eventHookMap:    make(map[string]*EventHook),
		serviceMap:      make(map[string]*Service),
		eventDispatcher: newEventDispatcher(),

		commandPrefix: commandPrefixes,
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

	// Add generic handler for event hooks
	// Add command handler
	bot.Discord.AddHandler(func(ds *discordgo.Session, event interface{}) {
		bot.onEvent(ds, event)
	})

	// Handle joining new guilds
	defaultGuildState = bot.BotState /// ugh...
	bot.Discord.AddHandler(onGuildJoin)

	// Open the websocket and begin listening.
	logrus.Info("Opening WS connection to Discord .. ")
	err := bot.Discord.Open()
	if err != nil {
		return fmt.Errorf("error: %s", err)
	}
	logrus.Info("OK")

	// add trigger by mention unless disabled in config
	if !bot.Config.DisableMentionTrigger {
		bot.commandPrefix = append(bot.commandPrefix, bot.Discord.State.User.Mention())
		bot.Config.CommandPrefix = bot.commandPrefix
	}

	// Create context for services
	ctx := NewContext(bot, bot.Discord, termSignal)

	// Run services
	for _, srv := range bot.serviceMap {
		if srv.Deactivated {
			continue
		}

		// run service
		go srv.Action(ctx)
	}

	// create a add bot url
	logrus.Info("Add bot using: https://discordapp.com/oauth2/authorize?scope=bot&client_id=" + bot.Discord.State.User.ID)

	logrus.Info("Bot is now running.  Press CTRL-C to exit.")
	termSignal = make(chan os.Signal, 1)
	signal.Notify(termSignal, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-termSignal
	fmt.Println("") // keep the `^C` on it's own line for prettiness
	logrus.Info("Shutting down bot..")

	// Cleanly close down the Discord session.
	logrus.Info("\tClosing WS discord connection .. ")
	err = bot.Discord.Close()
	if err != nil {
		return err
	}
	logrus.Info("\tClosed WS discord connection.")

	logrus.Info("Shutdown successfully")

	return nil
}

// RegisterCommand ...
func (bot *Bot) RegisterCommand(cmd *Command) error {
	name := cmd.Name
	if ex, exists := bot.commandMap[name]; exists {
		return &DuplicateCommandError{Existing: ex, New: cmd, Name: name}
	}

	logrus.Info("[unison] Registerred command: " + cmd.Name)
	bot.commandMap[name] = cmd.buildCommand()

	// TODO: Register aliases

	return nil
}

// RegisterEventHook ...
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

// RegisterService ...
func (bot *Bot) RegisterService(srv *Service) error {
	name := srv.Name
	if ex, exists := bot.serviceMap[name]; exists {
		return &DuplicateServiceError{Existing: ex, New: srv, Name: name}
	}
	bot.serviceMap[name] = srv

	return nil
}

func (bot *Bot) onEvent(ds *discordgo.Session, dv interface{}) {
	// Inspect and wrap event
	ev, err := events.NewDiscordEvent(dv)
	if err != nil {
		logrus.Errorf("event handler: %s", err)
	}

	// Create context for handlers
	ctx := NewContext(bot, ds, termSignal)

	// check if event was triggered by bot itself
	self := false
	// TODO: check

	// Invoke event hooks for the hooks that are subscribed to the event type
	bot.eventDispatcher.Dispatch(ctx, ev, self)

	// Invoke command handler on new messages
	if ev.Type == events.MessageCreateEvent {
		handleMessageCreate(ctx, ev.Event.(*discordgo.MessageCreate))
	}
}

// Bot state
//

// GetState retrieves the state for given guild
func (bot *Bot) GetState(guildID string) (state.Type, error) {
	return state.GetGuildState(guildID)
}

// SetState updates state for given guild
func (bot *Bot) SetState(guildID string, st state.Type) error {
	return state.SetGuildState(guildID, st)
}

// Event listeners
//

func (bot *Bot) onReady(ds *discordgo.Session, r *discordgo.Ready) {
	// Set bot state
	bot.readyState = r
	bot.User = r.User

	logrus.WithFields(logrus.Fields{
		"ID":       r.User.ID,
		"Username": r.User.Username,
	}).Infof("Websocket connected.")
}

func onGuildJoin(s *discordgo.Session, event *discordgo.GuildCreate) {

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
			logrus.Info("Joined Guild `" + event.Guild.Name + "`, and set state to `" + state.ToStr(selectedState) + "`")
		}
	} else {
		logrus.Info("Checked Guild `" + event.Guild.Name + "`, with state `" + state.ToStr(st) + "`")
	}
}
