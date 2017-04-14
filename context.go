package unison

import "github.com/bwmarrin/discordgo"

// Context is a type that is passed to every handler
// in a bot application.
// It can be used to refer back to main components.
type Context struct {
	Bot     *Bot
	Discord *discordgo.Session
}

func NewContext(bot *Bot, ds *discordgo.Session) *Context {
	return &Context{Bot: bot, Discord: ds}
}
