package client

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/sirupsen/logrus"

	"github.com/s1kx/discordgo"
	"github.com/s1kx/unison"
	"github.com/s1kx/unison/discord"
	"github.com/s1kx/unison/discord/events"
	"github.com/s1kx/unison/state"
)

// Session is an active bot session.
type Client struct {
	bot    *unison.Bot
	config *unison.Config

	session *discordgo.Session

	// name/alias => command
	commands commandRegistry
	// name => hook
	hooks hookRegistry
	// name => service
	services serviceRegistry

	dispatcher *dispatcher

	// Command prefixes
	commandPrefix []string

	readyState *discordgo.Ready
	User       *discordgo.User

	state state.Type // default bot state
}

func (s Client) Session() *discordgo.Session {
	return s.session
}

// This is awful and needs to be handled. It's used in "onGuildJoin" func
var defaultGuildState state.Type

func NewClient(bot *unison.Bot, conf *unison.Config, session *discordgo.Session) (*Client, error) {
	cmdPrefixes := make([]string, len(conf.CommandPrefixes))
	copy(cmdPrefixes, conf.CommandPrefixes)

	// Initialize basic client
	cl := &Client{
		bot:    bot,
		config: conf,

		session: session,

		commands: make(map[string]*unison.Command),
		hooks:    make(map[string]*unison.EventHook),
		services: make(map[string]*unison.Service),

		dispatcher: newEventDispatcher(),

		commandPrefix: []string{},
	}

	// Register commands
	for _, cmd := range bot.Commands {
		err := cl.RegisterCommand(cmd)
		if err != nil {
			return nil, err
		}
	}

	// Register event hooks
	for _, hook := range bot.EventHooks {
		err := cl.RegisterEventHook(hook)
		if err != nil {
			return nil, err
		}
	}

	// Register services
	for _, srv := range bot.Services {
		err := cl.RegisterService(srv)
		if err != nil {
			return nil, err
		}
	}

	return cl, nil
}

// GetServiceData a data value from existing services
func (cl *Client) GetServiceData(srvName string, key string) string {
	if val, ok := cl.services[srvName]; ok {
		if d, s := val.Data[key]; s {
			// key exist
			return d
		}
	}

	return ""
}

// SetServiceData update or set a new value for a given service key
func (cl *Client) SetServiceData(srvName string, key string, val string) string {
	if v, ok := cl.services[srvName]; ok {
		if _, s := v.Data[key]; s {
			cl.services[srvName].Data[key] = val

			return val
		}
	}

	return ""
}

// Run Start the bot instance
func (cl *Client) Run() error {
	// Add handler to wait for ready state in order to initialize the bot fully.
	cl.session.AddHandler(cl.onReady)

	// Add generic handler for event hooks
	// Add command handler
	cl.session.AddHandler(func(ds *discordgo.Session, event interface{}) {
		cl.onEvent(ds, event)
	})

	// Handle joining new guilds
	if !cl.DisableBoltDatabase {
		defaultGuildState = cl.BotState /// ugh...
		cl.Discord.AddHandler(onGuildJoin)
	}

	// Open the websocket and begin listening.
	logrus.Info("Opening WS connection to Discord .. ")
	err := cl.session.Open()
	if err != nil {
		return fmt.Errorf("error: %s", err)
	}
	logrus.Info("OK")

	// add trigger by mention unless disabled in config
	if !cl.Config.DisableMentionTrigger {
		cl.commandPrefix = append(cl.commandPrefix, cl.Discord.State.User.Mention())
		cl.Config.CommandPrefix = cl.commandPrefix
	}

	// Create context for services
	ctx := NewContext(bot, cl.Discord, termSignal)

	// Run services
	for _, srv := range cl.serviceMap {
		if srv.Deactivated {
			continue
		}

		// run service
		go srv.Action(ctx)
	}

	// create a add bot url
	logrus.Info("Add bot using: https://discordapp.com/oauth2/authorize?scope=bot&client_id=" + cl.Discord.State.User.ID)

	logrus.Info("Bot is now running.  Press CTRL-C to exit.")
	termSignal = make(chan os.Signal, 1)
	signal.Notify(termSignal, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-termSignal
	fmt.Println("") // keep the `^C` on it's own line for prettiness
	logrus.Info("Shutting down cl..")

	// Cleanly close down the Discord session.
	logrus.Info("\tClosing WS discord connection .. ")
	err = cl.Discord.Close()
	if err != nil {
		return err
	}
	logrus.Info("\tClosed WS discord connection.")

	logrus.Info("Shutdown successfully")

	return nil
}

// RegisterCommand ...
func (cl *Client) RegisterCommand(cmd *unison.Command) error {
	name := cmd.Name
	if ex, exists := cl.commandMap[name]; exists {
		return &DuplicateCommandError{Existing: ex, New: cmd, Name: name}
	}

	logrus.Info("[unison] Registerred command: " + cmd.Name)
	cl.commandMap[name] = cmd.buildCommand()

	// TODO: Register aliases

	return nil
}

// RegisterEventHook ...
func (cl *Client) RegisterEventHook(hook *unison.EventHook) error {
	name := hook.Name
	if ex, exists := cl.eventHookMap[name]; exists {
		return &DuplicateEventHookError{Existing: ex, New: hook}
	}
	cl.eventHookMap[name] = hook

	if len(hook.Events) == 0 {
		logrus.Warnf("Hook '%s' is not subscribed to any events", name)
	}

	cl.eventDispatcher.AddHook(hook)

	return nil
}

// RegisterService ...
func (cl *Client) RegisterService(srv *unison.Service) error {
	name := srv.Name
	if ex, exists := cl.serviceMap[name]; exists {
		return &DuplicateServiceError{Existing: ex, New: srv, Name: name}
	}
	cl.serviceMap[name] = srv

	return nil
}

func (cl *Client) onEvent(ds *discordgo.Session, dv interface{}) {
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
	cl.eventDispatcher.Dispatch(ctx, ev, self)

	// Invoke command handler on new messages
	if ev.Type == events.MessageCreateEvent {
		handleMessageCreate(ctx, ev.Event.(*discordgo.MessageCreate))
	}
}

// Bot state
//

// GetState retrieves the state for given guild
func (cl *Client) GetState(guildID string) (state.Type, error) {
	if cl.DisableBoltDatabase {
		return state.MissingState, ErrDatabaseDisabled
	}
	return state.GetGuildState(guildID)
}

// SetState updates state for given guild
func (cl *Client) SetState(guildID string, st state.Type) error {
	if cl.DisableBoltDatabase {
		return ErrDatabaseDisabled
	}
	return state.SetGuildState(guildID, st)
}

// GetGuildValue returns a value from the bots key/value database
func (cl *Client) GetGuildValue(guildID, key string) ([]byte, error) {
	if cl.DisableBoltDatabase {
		return nil, ErrDatabaseDisabled
	}
	return state.GetGuildValue(guildID, key)
}

// SetGuildValue updates/inserts a key-value into the given guild bucket
func (cl *Client) SetGuildValue(guildID, key string, val []byte) error {
	if cl.DisableBoltDatabase {
		return ErrDatabaseDisabled
	}
	return state.SetGuildValue(guildID, key, val)
}

// SendMessage sends a string message to a given channel
func (cl *Client) SendMessage(channel *discord.Channel, msg string) (*discord.Message, error) {
	// TODO maybe add this a discord.Channel method, but need to store discord session when creating object
	discordgoMessage, err := cl.Discord.ChannelMessageSend(channel.ID.String(), msg)
	if err != nil {
		return nil, err
	}

	return discord.NewMessageFromDgo(discordgoMessage), nil
}

// Event listeners
//

func (cl *Client) onReady(ds *discordgo.Session, r *discordgo.Ready) {
	// Set bot state
	cl.readyState = r
	cl.User = r.User

	logrus.WithFields(logrus.Fields{
		"ID":       r.User.ID,
		"Username": r.User.Username,
	}).Infof("Websocket connected.")
}
