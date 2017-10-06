package events

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

type InvalidEventTypeError struct {
	Event interface{}
}

func (e InvalidEventTypeError) Error() string {
	return fmt.Sprintf("events: %T is not a known event type", e.Event)
}

// GetEventType returns the EventType for a given discordgo event.
func GetEventType(v interface{}) (t EventType, err error) {
	switch v.(type) {
	default:
		err = InvalidEventTypeError{Event: v}

	// Websocket events
	case *discordgo.Event:
		t = WebsocketEvent

	// Connection events
	case *discordgo.Connect:
		t = ConnectEvent
	case *discordgo.Disconnect:
		t = DisconnectEvent

	// Connection state events
	case *discordgo.Ready:
		t = ReadyEvent
	case *discordgo.Resumed:
		t = ResumedEvent

	// Channel events
	case *discordgo.ChannelCreate:
		t = ChannelCreateEvent
	case *discordgo.ChannelUpdate:
		t = ChannelUpdateEvent
	case *discordgo.ChannelDelete:
		t = ChannelDeleteEvent
	case *discordgo.ChannelPinsUpdate:
		t = ChannelPinsUpdateEvent

	// Guild events
	case *discordgo.GuildCreate:
		t = GuildCreateEvent
	case *discordgo.GuildUpdate:
		t = GuildUpdateEvent
	case *discordgo.GuildDelete:
		t = GuildDeleteEvent

	// Guild ban events
	case *discordgo.GuildBanAdd:
		t = GuildBanAddEvent
	case *discordgo.GuildBanRemove:
		t = GuildBanRemoveEvent

	// Guild member events
	case *discordgo.GuildMemberAdd:
		t = GuildMemberAddEvent
	case *discordgo.GuildMemberUpdate:
		t = GuildMemberUpdateEvent
	case *discordgo.GuildMemberRemove:
		t = GuildMemberRemoveEvent

	// Guild role events 
	case *discordgo.GuildRoleCreate:
		t = GuildRoleCreateEvent
	case *discordgo.GuildRoleUpdate:
		t = GuildRoleUpdateEvent
	case *discordgo.GuildRoleDelete:
		t = GuildRoleDeleteEvent

	// Guild misc events
	case *discordgo.GuildEmojisUpdate:
		t = GuildEmojisUpdateEvent
	case *discordgo.GuildMembersChunk:
		t = GuildMembersChunkEvent
	case *discordgo.GuildIntegrationsUpdate:
		t = GuildIntegrationsUpdateEvent

	// Message events
	case *discordgo.MessageAck:
		t = MessageAckEvent
	case *discordgo.MessageCreate:
		t = MessageCreateEvent
	case *discordgo.MessageUpdate:
		t = MessageUpdateEvent
	case *discordgo.MessageDelete:
		t = MessageDeleteEvent

	// Message reaction events
	case *discordgo.MessageReactionAdd:
		t = MessageReactionAddEvent
	case *discordgo.MessageReactionRemove:
		t = MessageReactionRemoveEvent

	// Presence events
	case *discordgo.PresencesReplace:
		t = PresencesReplaceEvent
	case *discordgo.PresenceUpdate:
		t = PresenceUpdateEvent

		// Relationship events
	case *discordgo.RelationshipAdd:
		t = RelationshipAddEvent
	case *discordgo.RelationshipRemove:
		t = RelationshipRemoveEvent

	// User events
	case *discordgo.TypingStart:
		t = TypingStartEvent
	case *discordgo.UserUpdate:
		t = UserUpdateEvent
	case *discordgo.UserSettingsUpdate:
		t = UserSettingsUpdateEvent
	case *discordgo.UserGuildSettingsUpdate:
		t = UserGuildSettingsUpdateEvent

	// Voice events
	case *discordgo.VoiceServerUpdate:
		t = VoiceServerUpdateEvent
	case *discordgo.VoiceStateUpdate:
		t = VoiceStateUpdateEvent
	}

	return t, err
}
