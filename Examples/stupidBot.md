# Creating a bot with no abilities
This is all the code required to create a bot that has no commands, hooks, services. And reads its configuration from environment variables (discord bot token, command prefix, etc.).

```Go
package main

import (
	"github.com/Sirupsen/logrus"
	"github.com/s1kx/unison"
)

func main() {
	settings := &unison.BotSettings{
		Commands:   []*unison.Command{},    // No commands added
		EventHooks: []*unison.EventHook{},  // No hooks added
		Services:   []*unison.Service{},    // No services added
	}

	// Start the bot
	err := unison.Run(settings)
	if err != nil {
		logrus.Error(err)
	}
}
```

Terminal output of running built code:
```markdown
github.com/andersfylling/unisonTest
▶ ./unisonTest
INFO[0000] Using bot token from environment variable.
INFO[0000] Commands are triggered by ``
INFO[0000] Opening WS connection to Discord ..
INFO[0000] OK
INFO[0000] Add bot using: https://discordapp.com/oauth2/authorize?client_id=395653573847326811&scope=bot
INFO[0000] Bot is now running.  Press CTRL-C to exit.
INFO[0000] Websocket connected.                          ID=395653573847326811 Username=UnisonTest
INFO[0000] Joined Guild `someGuildName`, and set state to `1`
```
The application list its interaction configuation and notes when the connection to discord is established (goes online). It also generates a url for adding the bot to your server and eventually logs every server it has joined (and creates an independent state entry for given server).

The bot also uses graceful shutdown by listening to interupt signals.
```markdown
^C
INFO[0307] Shutting down bot..
INFO[0307]      Closing WS discord connection ..
INFO[0308]      Closed WS discord connection.
INFO[0308] Shutdown successfully
```

Remember the discord bot token is added using `export UNISON_DISCORD_TOKEN="kdsfjskfhgk4wh8h"` in this example.
