# Creating a bot that responds to messages

There are two way of doing this; using `Command` or `Hook`.

## Command

A command is triggered by users/members with access rights, and cannot be executed otherwise. At least for now, to create a simple conceptual model.

To create a ping command that responds with pong, we can do the following.

```go
var PingCommand = &unison.Command{
  Name: "ping",
  Usage: "Causes a reply with `pong`",
  Action: pingCommandAction,
}

func pingCommandAction(ctx *unison.Context, m *discordgo.Message, request string) error {
  _, err := ctx.Bot.Discord.ChannelMessageSend(m.ChannelID, "pong")

  return err
}

// Build example configuration with ping command
func buildConfig() *unison.Config {
   return &unison.Config{
        Commands:   []*unison.Command{
          PingCommand,                     // Add ping command
        },
        EventHooks: []*unison.EventHook{}, // No hooks added
        Services:   []*unison.Service{},   // No services added
    }
}
```

The bot then replies with `pong` whenever the command `ping` is issued: `@botname ping` => `replies with pong`.

## Hook

Is not triggered by a user/member. It's executed when a event is fired. The hook logic can subscribe to one or multiple events.

(The Hooks needs a rewrite as it can be simplified..)

```go
var PingHook = &unison.EventHook{
  Name: "ping",
  Usage: "Causes a reply with `pong` whenever someone write just the word `ping`",
  Action: unison.EventHandlerFunc(pingHookAction),
  Events: []events.EventType{
    events.MessageCreateEvent,
  },
}

func pingHookAction(ctx *unison.Context, event *events.DiscordEvent, self bool) (handled bool, err error) {
  var m *discordgo.Message

  // Make sure it's a new message
  if event.Type != events.MessageCreateEvent {
      return true, nil
  }

  // Convert the interface into its correct Message type
  switch ev := event.Event.(type) {
  default:
      return true, nil
  case *discordgo.MessageCreate:
      m = ev.Message
  }

  // if message is from the bot itself, don't respond as this causes an eternal loop.
  if m.Author.ID == ctx.Bot.Discord.State.User.ID {
      return true, nil
  }

  // make sure the message contains just `ping`
  if m.content != "ping" {
    return true, nil
  }

  // Respond with pong
  _, err := ctx.Bot.Discord.ChannelMessageSend(m.ChannelID, "pong")

  return true, err
}

// Build example configuration with ping hook
func buildConfig() *unison.Config {
    return &unison.Config{
        Commands:   []*unison.Command{},    // No commands added
        EventHooks: []*unison.EventHook{
          PingHook,
        },
        Services:   []*unison.Service{},    // No services added
    }
}
```

Note that we don't use hooks if the user/member is directly requesting a response. Hooks are more useful to react to certain events taking place as these can be anything. Say you wan't to create a database record whenever you join a new guild. This is where you would be forced to use a hook as there is "no" other way.
