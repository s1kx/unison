/*
Package events implements event mapping for the discordgo library.

Since discordgo does not have a common interface amongst Events or a name/ID
attached to them, it's only possible to handle events selectively with
type switches or reflection. Type switches get messy, and reflection is slow.

This package aims to improve event handling by providing types, wrappers and
interfaces for events
*/
package events

type EventType int

// Curated from discordgo/events.go
// TODO: Use go:generate
const (
	// NoEvent and AllEvents are special constants that don't map to a discord event
	NoEvent EventType = iota
	//AllEvents // not supported yet

	// Basic Websocket Event
	WebsocketEvent

	// Discord Events
	ConnectEvent
	DisconnectEvent

	RateLimitEvent

	ReadyEvent
	ResumedEvent

	ChannelCreateEvent
	ChannelUpdateEvent
	ChannelDeleteEvent
	ChannelPinsUpdateEvent

	GuildCreateEvent
	GuildUpdateEvent
	GuildDeleteEvent

	GuildBanAddEvent
	GuildBanRemoveEvent

	GuildMemberAddEvent
	GuildMemberUpdateEvent
	GuildMemberRemoveEvent

	GuildRoleCreateEvent
	GuildRoleUpdateEvent
	GuildRoleDeleteEvent

	GuildEmojisUpdateEvent
	GuildMembersChunkEvent
	GuildIntegrationsUpdateEvent

	MessageAckEvent
	MessageCreateEvent
	MessageUpdateEvent
	MessageDeleteEvent

	MessageReactionAddEvent
	MessageReactionRemoveEvent

	PresencesReplaceEvent
	PresenceUpdateEvent

	RelationshipAddEvent
	RelationshipRemoveEvent

	TypingStartEvent

	UserUpdateEvent
	UserSettingsUpdateEvent
	UserGuildSettingsUpdateEvent

	VoiceServerUpdateEvent
	VoiceStateUpdateEvent
)
