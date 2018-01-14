package discord

import (
	"time"
)

type Message struct {
	ID              uint64        `json:"id"`
	ChannelID       uint64        `json:"channel_id"`
	Author          *User         `json:"author"`
	Content         string        `json:"content"`
	Timestamp       time.Time     `json:"timestamp"`
	EditedTimestamp time.Time     `json:"edited_timestamp"` // ?
	Tts             bool          `json:"tts"`
	MentionEveryone bool          `json:"mention_everyone"`
	Mentions        []*User       `json:"mentions"`
	MentionRoles    []uint64      `json:"mention_roles"`
	Attachments     []*Attachment `json:"attachments"`
	Embeds          []*Embed      `json:"embeds"`
	Reactions       []*Reaction   `json:"reactions"` // ?
	Nonce           uint64        `json:"nonce"`     // ?, used for validating a message was sent
	Pinned          bool          `json:"pinned"`
	WebhookID       uint64        `json:"webhook_id"` // ?
	Type            uint          `json:"type"`
}
