package unison

import (
	"os"
	"os/signal"

	"github.com/bwmarrin/discordgo"
)

// Context is a type that is passed to every handler
// in a bot application.
// It can be used to refer back to main components.
type Context struct {
	Bot                *Bot
	Discord            *discordgo.Session
	SystemInteruptChan chan struct{}
}

func NewContext(bot *Bot, ds *discordgo.Session) *Context {
	is := make(chan struct{})

	// shutdown signal
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		<-c
		close(is)
	}()

	return &Context{Bot: bot, Discord: ds, SystemInteruptChan: is}
}
