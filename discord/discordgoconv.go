package discord

import (
	"strconv"
	"time"

	"gopkg.in/bwmarrin/Discordgo.v0"
)

const discordgoTimestampLayout string = "2016-08-06T17:20:33.803-0400"

// Small purpose specific functionality
//

func discordgoIDStringToUint64(id string) uint64 {
	if id == "" {
		return 0
	}

	u, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		panic(err)
	}

	return u
}

func discordgoIDStringArrayToUint64Array(ids []string) []uint64 {
	newIDS := make([]uint64, 0, len(ids))
	for i, id := range ids {
		newIDS[i] = discordgoIDStringToUint64(id)
	}

	return newIDS
}

func uint64ToString(id uint64) string {
	return strconv.FormatUint(id, 10)
}

func discordgoAttachmentArrayToDiscordAttachmentArray(as []*discordgo.MessageAttachment) []*Attachment {
	attachments := make([]*Attachment, 0, len(as))
	for i, a := range as {
		attachments[i] = NewAttachmentFromDiscordgo(a)
	}

	return attachments
}

func discordgoTimestampToTime(ts discordgo.Timestamp) time.Time {
	return discordgoTimestampStringToTime(string(ts))
}

func discordgoTimestampStringToTime(ts string) time.Time {
	timestamp, err := time.Parse(discordgoTimestampLayout, ts)
	if err != nil {
		panic(err)
		//return time.Now() // this is so bad..
	}

	return timestamp
}

func discordgoUserArrayTODiscordUserArray(users []*discordgo.User) []*User {
	discordUsers := make([]*User, 0, len(users))
	for i, user := range users {
		discordUsers[i] = NewUserFromDiscordgo(user)
	}

	return discordUsers
}

func discordgoRolesToDiscordRoles(rs []*discordgo.Role) []*Role {
	roles := make([]*Role, 0, len(rs))
	for i, r := range rs {
		roles[i] = NewRoleFromDiscordgo(r)
	}

	return roles
}

func discordgoEmojisToDiscordEmojis(es []*discordgo.Emoji) []*Emoji {
	emojis := make([]*Emoji, 0, len(es))
	for i, e := range es {
		emojis[i] = NewEmojiFromDiscordgo(e)
	}

	return emojis
}

func discordgoGuildMembersToDiscordGuildMembers(ms []*discordgo.Member) []*GuildMember {
	guildMembers := make([]*GuildMember, 0, len(ms))
	for i, m := range ms {
		guildMembers[i] = NewGuildMemberFromDiscordgo(m)
	}

	return guildMembers
}

func discordgoPresencesToDiscordPresences(ps []*discordgo.Presence) []*Presence {
	presences := make([]*Presence, 0, len(ps))
	for i, p := range ps {
		presences[i] = NewPresenceFromDiscordgo(p)
	}

	return presences
}

func discordgoChannelsToDiscordChannels(cs []*discordgo.Channel) []*Channel {
	channels := make([]*Channel, 0, len(cs))
	for i, c := range cs {
		channels[i] = NewChannelFromDiscordgo(c)
	}

	return channels
}

func discordgoVoiceStatesToDiscordVoiceStates(vss []*discordgo.VoiceState) []*VoiceState {
	voiceStates := make([]*VoiceState, 0, len(vss))
	for i, vs := range vss {
		voiceStates[i] = NewVoiceStateFromDiscordgo(vs)
	}

	return voiceStates
}

func discordgoCopyTodiscordStruct(discordgoStruct interface{}, discordStruct interface{}) {
	// TODO use reflection to copy over values with similar type and json tag.
}

func discordgoMessageTypeToUint8(t discordgo.MessageType) uint8 {
	return uint8(t)
}

func discordgoVerificationLevelToUint8(vl discordgo.VerificationLevel) uint8 {
	return uint8(vl)
}

func discordgoStatusToString(s discordgo.Status) string {
	return string(s)
}

// Struct converters
//

func NewUserFromDiscordgo(user *discordgo.User) *User {
	return &User{
		ID:            discordgoIDStringToUint64(user.ID),
		Email:         user.Email,
		Username:      user.Username,
		Avatar:        user.Avatar,
		Discriminator: user.Discriminator,
		Token:         user.Token,
		Verified:      user.Verified,
		MFAEnabled:    user.MFAEnabled,
		Bot:           user.Bot,
	}
}

func NewEmbedFromDiscordgoEmbed(e *discordgo.MessageEmbed) *Embed {
	return &Embed{
		Title:       e.Title,
		Type:        e.Type,
		Description: e.Description,
		URL:         e.URL,
		Timestamp:   discordgoTimestampStringToTime(e.Timestamp),
		Color:       e.Color,
		// Footer: NewEmbedFooterFromDiscordgo(e.Footer),
		// Image: NewEmbedImageFromDiscordgo(e.Image),
		// Thumbnail: NewEmbedThumbnailFromDiscordgo(e.Thumbnail),
		// Video: NewEmbedVideoFromDiscordgo(e.Video),
		// Provider: NewEmbedProviderFromDiscordgo(e.Provider),
		// Author: NewEmbedAuthorFromDiscordgo(e.Author),
		// Fields: discordgoFieldArrayToDiscordEmbedFieldArray(e.Fields),
	}
	// TODO
}

func NewGuildFromDiscordgo(g *discordgo.Guild) *Guild {
	return &Guild{
		ID:                discordgoIDStringToUint64(g.ID),
		Name:              g.Name,
		Icon:              g.Icon,
		Region:            g.Region,
		AfkChannelID:      discordgoIDStringToUint64(g.AfkChannelID),
		EmbedChannelID:    discordgoIDStringToUint64(g.EmbedChannelID),
		OwnerID:           discordgoIDStringToUint64(g.OwnerID),
		JoinedAt:          discordgoTimestampToTime(g.JoinedAt),
		Splash:            g.Splash,
		AfkTimeout:        uint(g.AfkTimeout),
		MemberCount:       uint(g.MemberCount),
		VerificationLevel: discordgoVerificationLevelToUint8(g.VerificationLevel),
		EmbedEnabled:      g.EmbedEnabled,
		Large:             g.Large,
		DefaultMessageNotifications: g.DefaultMessageNotifications, // TODO: review type
		Roles:       discordgoRolesToDiscordRoles(g.Roles),
		Emojis:      discordgoEmojisToDiscordEmojis(g.Emojis),
		Members:     discordgoGuildMembersToDiscordGuildMembers(g.Members),
		Presences:   discordgoPresencesToDiscordPresences(g.Presences),
		Channels:    discordgoChannelsToDiscordChannels(g.Channels),
		VoiceStates: discordgoVoiceStatesToDiscordVoiceStates(g.VoiceStates),
		Unavailable: g.Unavailable,
	}
}

func NewMessageFromDiscordgo(msg *discordgo.Message) *Message {
	return &Message{
		ID:              discordgoIDStringToUint64(msg.ID),
		ChannelID:       discordgoIDStringToUint64(msg.ChannelID),
		Content:         msg.Content,
		Timestamp:       discordgoTimestampToTime(msg.Timestamp),
		Tts:             msg.Tts,
		MentionEveryone: msg.MentionEveryone,
		Mentions:        discordgoUserArrayTODiscordUserArray(msg.Mentions),
		MentionRoles:    discordgoIDStringArrayToUint64Array(msg.MentionRoles),
		Attachments:     discordgoAttachmentArrayToDiscordAttachmentArray(msg.Attachments),
		// Embeds: discordgoMessageEmbedsToDiscordEmbeds(msg.Embeds),
		// Reactions: discordgoReactionsToDiscordReactions(msg.Reactions),
		// Nonce: discordgoIDStringToUint64(msg.Nonce),
		// Pinned: msg.Pinned, // not implemented by discordgo..
		// WebhookID: discordgoIDStringToUint64(msg.WebhookID), // Not implemented by discordgo...
		Type: discordgoMessageTypeToUint8(msg.Type),
	}
	// TODO
}

func NewRoleFromDiscordgo(r *discordgo.Role) *Role {
	return &Role{
		ID:          discordgoIDStringToUint64(r.ID),
		Name:        r.Name,
		Managed:     r.Managed,
		Mentionable: r.Mentionable,
		Hoist:       r.Hoist,
		Color:       r.Color,
		Position:    r.Position,
		Permissions: uint64(r.Permissions),
	}
}

func NewEmojiFromDiscordgo(e *discordgo.Emoji) *Emoji {
	return &Emoji{
		ID:            discordgoIDStringToUint64(e.ID),
		Name:          e.Name,
		Roles:         discordgoIDStringArrayToUint64Array(e.Roles),
		RequireColons: e.RequireColons,
		Managed:       e.Managed,
		// User: NewUserFromDiscordgo(e.User), // Not implemented by discordgo
	}
}

func NewGuildMemberFromDiscordgo(m *discordgo.Member) *GuildMember {
	return &GuildMember{
		GuildID:  discordgoIDStringToUint64(m.GuildID),
		JoinedAt: discordgoTimestampStringToTime(m.JoinedAt),
		Nick:     m.Nick,
		Deaf:     m.Deaf,
		Mute:     m.Mute,
		User:     NewUserFromDiscordgo(m.User),
		Roles:    discordgoIDStringArrayToUint64Array(m.Roles),
	}
}

func NewPresenceFromDiscordgo(p *discordgo.Presence) *Presence {
	return &Presence{
		User:  NewUserFromDiscordgo(p.User),
		Roles: discordgoIDStringArrayToUint64Array(p.Roles),
		// Game: NewActivityFromDiscordgo(p.Activity), // not implemented by discordgo...
		// GuildID: discordgoIDStringToUint64(p.GuildID), // not implemented by discordgo..
		Status: discordgoStatusToString(p.Status),
	}
}

func NewChannelFromDiscordgo(c *discordgo.Channel) *Channel {
	return &Channel{}
	// TODO
}

func NewVoiceStateFromDiscordgo(vs *discordgo.VoiceState) *VoiceState {
	return &VoiceState{}
	// TODO
}
