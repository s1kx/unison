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

// NewContext Create a new context class for the discord bot
// This will also hold a signal for system interupts
func NewContext(bot *Bot, ds *discordgo.Session) *Context {
	ctx := new(Context)
	ctx.Bot = bot
	ctx.Discord = ds
	ctx.SystemInteruptChan = make(chan struct{})

	// shutdown signal
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		<-c
		close(ctx.SystemInteruptChan)
	}()

	return ctx
}
