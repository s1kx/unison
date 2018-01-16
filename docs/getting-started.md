# Creating a bot with no abilities

This is all the code required to create a bot that has no commands, hooks, services. And reads its configuration from environment variables (discord bot token, command prefix, etc.).

```go
settings := &unison.Config{
  Commands:   []*unison.Command{},   // No commands added
  EventHooks: []*unison.EventHook{}, // No hooks added
  Services:   []*unison.Service{},   // No services added
}

// Start the bot
err := unison.Run(settings)
if err != nil {
  return err
}
```

Example output:

```
INFO[2018-01-05 17:25:37] Using bot token from environment variable.
INFO[2018-01-05 17:25:37] Commands are triggered by; And by @mention.
INFO[2018-01-05 17:25:37] Opening WS connection to Discord ..
INFO[2018-01-05 17:25:37] OK
INFO[2018-01-05 17:25:37] Add bot using: https://discordapp.com/oauth2/authorize?scope=bot&client_id=395653573847326811
INFO[2018-01-05 17:25:37] Bot is now running.  Press CTRL-C to exit.
INFO[2018-01-05 17:25:37] Websocket connected.                          ID=395653573847326811 Username=unisonTester1
INFO[2018-01-05 17:25:37] Joined Guild `someGuildName`, and set state to `1`
```

The application list its interaction configuation and notes when the connection to discord is established (goes online). It also generates a url for adding the bot to your server and eventually logs every server it has joined (and creates an independent state entry for given server).

The bot also uses graceful shutdown by listening to interupt signals.

```
^C
INFO[2018-01-05 17:28:00] Shutting down bot..
INFO[2018-01-05 17:28:00]       Closing WS discord connection ..
INFO[2018-01-05 17:28:01]       Closed WS discord connection.
INFO[2018-01-05 17:28:01] Shutdown successfully
```

Remember the discord bot token is added using `export UNISON_DISCORD_TOKEN="kdsfjskfhgk4wh8h"` in this example. If the token is not exported you get a warning.

```
ERRO[2018-01-05 17:22:23] Missing env var UNISON_DISCORD_TOKEN. This is required. Specify in either Settings struct or env var.
```

If the discord token is incorrect the output will give an error and say authentication failed.

```go
INFO[2018-01-05 17:28:35] Using bot token from environment variable.
INFO[2018-01-05 17:28:35] Commands are triggered by; And by @mention.
INFO[2018-01-05 17:28:35] Opening WS connection to Discord ..
ERRO[2018-01-05 17:28:35] error: websocket: close 4004: Authentication failed.
```
