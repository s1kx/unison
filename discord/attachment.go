package discord

import (
	"gopkg.in/bwmarrin/Discordgo.v0"
)

type Attachment struct {
	ID       uint64 `json:"id"`
	Filename string `json:"filename"`
	Size     uint   `json:"size"`
	URL      string `json:"url"`
	ProxyURL string `json:"proxy_url"`
	Height   uint   `json:"height"`
	Width    uint   `json:"width"`
}

func NewAttachmentFromDiscordgo(a *discordgo.MessageAttachment) *Attachment {
	return &Attachment{
		ID:       discordgoIDStringToUint64(a.ID),
		Filename: a.Filename,
		Size:     uint(a.Size),
		URL:      a.URL,
		ProxyURL: a.ProxyURL,
		Height:   uint(a.Height),
		Width:    uint(a.Width),
	}
}
