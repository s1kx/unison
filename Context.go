package unison

import (
	"os"

	"github.com/bwmarrin/discordgo"
)

// Context is a type that is passed to every handler
// in a bot application.
// It can be used to refer back to main components.
type Context struct {
	Bot                *Bot
	Discord            *discordgo.Session
	SystemInteruptChan chan os.Signal
}

// NewContext Create a new context class for the discord bot
// This will also hold a signal for system interupts
func NewContext(bot *Bot, ds *discordgo.Session, sig chan os.Signal) *Context {
	ctx := new(Context)
	ctx.Bot = bot
	ctx.Discord = ds
	ctx.SystemInteruptChan = sig

	return ctx
}
