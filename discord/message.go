package discord

import (
	"time"
)

type Message struct {
	ID              uint64        `json:"id"`
	ChannelID       uint64        `json:"channel_id"`
	Content         string        `json:"content"`
	Timestamp       time.Time     `json:"timestamp"`
	EditedTimestamp time.Time     `json:"edited_timestamp"`
	MentionRoles    []uint64      `json:"mention_roles"`
	Tts             bool          `json:"tts"`
	MentionEveryone bool          `json:"mention_everyone"`
	Author          *User         `json:"author"`
	Attachments     []*Attachment `json:"attachments"`
	Embeds          []*Embed      `json:"embeds"`
	Mentions        []*User       `json:"mentions"`
	Reactions       []*Reaction   `json:"reactions"`
	Type            uint8         `json:"type"`
}
