package discord

import (
	"strconv"
	"time"

	"gopkg.in/bwmarrin/Discordgo.v0"
)

const dgoTimestampLayout string = "2016-08-06T17:20:33.803-0400" // Needs checking

// Small purpose specific functionality
//

func dgoIDToSnowflake(id string) Snowflake {
	if id == "" {
		return Snowflake(0)
	}

	u, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		panic(err)
	}

	return Snowflake(u)
}

func dgoIDStringsToUint64s(ids []string) []Snowflake {
	newIDS := make([]Snowflake, 0, len(ids))
	for i, id := range ids {
		newIDS[i] = dgoIDToSnowflake(id)
	}

	return newIDS
}

func uint64ToString(id uint64) string {
	return strconv.FormatUint(id, 10)
}

func dgoAttachmentsToDiscordAttachments(as []*discordgo.MessageAttachment) []*Attachment {
	attachments := make([]*Attachment, 0, len(as))
	for i, a := range as {
		attachments[i] = NewAttachmentFromDgo(a)
	}

	return attachments
}

func dgoTimestampToTime(ts discordgo.Timestamp) time.Time {
	return dgoTimestampStringToTime(string(ts))
}

func dgoTimestampStringToTime(ts string) time.Time {
	timestamp, err := time.Parse(dgoTimestampLayout, ts)
	if err != nil {
		panic(err)
		//return time.Now() // this is so bad..
	}

	return timestamp
}

func dgoUsersTODiscordUsers(users []*discordgo.User) []*User {
	discordUsers := make([]*User, 0, len(users))
	for i, user := range users {
		discordUsers[i] = NewUserFromDgo(user)
	}

	return discordUsers
}

func dgoRolesToDiscordRoles(rs []*discordgo.Role) []*Role {
	roles := make([]*Role, 0, len(rs))
	for i, r := range rs {
		roles[i] = NewRoleFromDgo(r)
	}

	return roles
}

func dgoEmojisToDiscordEmojis(es []*discordgo.Emoji) []*Emoji {
	emojis := make([]*Emoji, 0, len(es))
	for i, e := range es {
		emojis[i] = NewEmojiFromDgo(e)
	}

	return emojis
}

func dgoGuildMembersToDiscordGuildMembers(ms []*discordgo.Member) []*GuildMember {
	guildMembers := make([]*GuildMember, 0, len(ms))
	for i, m := range ms {
		guildMembers[i] = NewGuildMemberFromDgo(m)
	}

	return guildMembers
}

func dgoPresencesToDiscordPresences(ps []*discordgo.Presence) []*Presence {
	presences := make([]*Presence, 0, len(ps))
	for i, p := range ps {
		presences[i] = NewPresenceFromDgo(p)
	}

	return presences
}

func dgoChannelsToDiscordChannels(cs []*discordgo.Channel) []*Channel {
	channels := make([]*Channel, 0, len(cs))
	for i, c := range cs {
		channels[i] = NewChannelFromDgo(c)
	}

	return channels
}

func dgoVoiceStatesToDiscordVoiceStates(vss []*discordgo.VoiceState) []*VoiceState {
	voiceStates := make([]*VoiceState, 0, len(vss))
	for i, vs := range vss {
		voiceStates[i] = NewVoiceStateFromDgo(vs)
	}

	return voiceStates
}

func dgoMessagesToDiscordMessages(ms []*discordgo.Message) []*Message {
	messages := make([]*Message, 0, len(ms))
	for i, m := range ms {
		messages[i] = NewMessageFromDgo(m)
	}

	return messages
}

func dgoPermissionOverwritesToDiscordPermissionOverwrites(pms []*discordgo.PermissionOverwrite) []*PermissionOverwrite {
	permissionOverwrites := make([]*PermissionOverwrite, 0, len(pms))
	for i, pm := range pms {
		permissionOverwrites[i] = NewPermissionOverwriteFromDgo(pm)
	}

	return permissionOverwrites
}

func dgoCopyTodiscordStruct(discordgoStruct interface{}, discordStruct interface{}) {
	// TODO use reflection to copy over values with similar type and json tag.
}

func dgoMessageTypeToUint(t discordgo.MessageType) uint {
	return uint(t)
}

func dgoVerificationLevelToUint(vl discordgo.VerificationLevel) uint {
	return uint(vl)
}

func dgoChannelTypeToUint(ct discordgo.ChannelType) uint {
	return uint(ct)
}

func dgoStatusToString(s discordgo.Status) string {
	return string(s)
}

// Struct converters
//

func NewUserFromDgo(user *discordgo.User) *User {
	return &User{
		ID:            dgoIDToSnowflake(user.ID),
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

func NewEmbedFromDgoEmbed(e *discordgo.MessageEmbed) *Embed {
	return &Embed{
		Title:       e.Title,
		Type:        e.Type,
		Description: e.Description,
		URL:         e.URL,
		Timestamp:   dgoTimestampStringToTime(e.Timestamp),
		Color:       e.Color,
		// Footer: NewEmbedFooterFromDiscordgo(e.Footer),
		// Image: NewEmbedImageFromDiscordgo(e.Image),
		// Thumbnail: NewEmbedThumbnailFromDiscordgo(e.Thumbnail),
		// Video: NewEmbedVideoFromDiscordgo(e.Video),
		// Provider: NewEmbedProviderFromDiscordgo(e.Provider),
		// Author: NewEmbedAuthorFromDiscordgo(e.Author),
		// Fields: dgoFieldArrayToDiscordEmbedFieldArray(e.Fields),
	}
	// TODO
}

func NewGuildFromDgo(g *discordgo.Guild) *Guild {
	return &Guild{
		ID:                dgoIDToSnowflake(g.ID),
		Name:              g.Name,
		Icon:              g.Icon,
		Region:            g.Region,
		AfkChannelID:      dgoIDToSnowflake(g.AfkChannelID),
		EmbedChannelID:    dgoIDToSnowflake(g.EmbedChannelID),
		OwnerID:           dgoIDToSnowflake(g.OwnerID),
		JoinedAt:          dgoTimestampToTime(g.JoinedAt),
		Splash:            g.Splash,
		AfkTimeout:        uint(g.AfkTimeout),
		MemberCount:       uint(g.MemberCount),
		VerificationLevel: dgoVerificationLevelToUint(g.VerificationLevel),
		EmbedEnabled:      g.EmbedEnabled,
		Large:             g.Large,
		DefaultMessageNotifications: g.DefaultMessageNotifications, // TODO: review type
		Roles:       dgoRolesToDiscordRoles(g.Roles),
		Emojis:      dgoEmojisToDiscordEmojis(g.Emojis),
		Members:     dgoGuildMembersToDiscordGuildMembers(g.Members),
		Presences:   dgoPresencesToDiscordPresences(g.Presences),
		Channels:    dgoChannelsToDiscordChannels(g.Channels),
		VoiceStates: dgoVoiceStatesToDiscordVoiceStates(g.VoiceStates),
		Unavailable: g.Unavailable,
	}
}

func NewMessageFromDgo(msg *discordgo.Message) *Message {
	return &Message{
		ID:              dgoIDToSnowflake(msg.ID),
		ChannelID:       dgoIDToSnowflake(msg.ChannelID),
		Content:         msg.Content,
		Timestamp:       dgoTimestampToTime(msg.Timestamp),
		Tts:             msg.Tts,
		MentionEveryone: msg.MentionEveryone,
		Mentions:        dgoUsersTODiscordUsers(msg.Mentions),
		MentionRoles:    dgoIDStringsToUint64s(msg.MentionRoles),
		Attachments:     dgoAttachmentsToDiscordAttachments(msg.Attachments),
		// Embeds: dgoMessageEmbedsToDiscordEmbeds(msg.Embeds),
		// Reactions: dgoReactionsToDiscordReactions(msg.Reactions),
		// Nonce: dgoIDStringToUint64(msg.Nonce),
		// Pinned: msg.Pinned, // not implemented by discordgo..
		// WebhookID: dgoIDStringToUint64(msg.WebhookID), // Not implemented by discordgo...
		Type: dgoMessageTypeToUint(msg.Type),
	}
	// TODO
}

func NewRoleFromDgo(r *discordgo.Role) *Role {
	return &Role{
		ID:          dgoIDToSnowflake(r.ID),
		Name:        r.Name,
		Managed:     r.Managed,
		Mentionable: r.Mentionable,
		Hoist:       r.Hoist,
		Color:       r.Color,
		Position:    r.Position,
		Permissions: uint64(r.Permissions),
	}
}

func NewEmojiFromDgo(e *discordgo.Emoji) *Emoji {
	return &Emoji{
		ID:            dgoIDToSnowflake(e.ID),
		Name:          e.Name,
		Roles:         dgoIDStringsToUint64s(e.Roles),
		RequireColons: e.RequireColons,
		Managed:       e.Managed,
		// User: NewUserFromDiscordgo(e.User), // Not implemented by discordgo
	}
}

func NewGuildMemberFromDgo(m *discordgo.Member) *GuildMember {
	return &GuildMember{
		GuildID:  dgoIDToSnowflake(m.GuildID),
		JoinedAt: dgoTimestampStringToTime(m.JoinedAt),
		Nick:     m.Nick,
		Deaf:     m.Deaf,
		Mute:     m.Mute,
		User:     NewUserFromDgo(m.User),
		Roles:    dgoIDStringsToUint64s(m.Roles),
	}
}

func NewPresenceFromDgo(p *discordgo.Presence) *Presence {
	return &Presence{
		User:  NewUserFromDgo(p.User),
		Roles: dgoIDStringsToUint64s(p.Roles),
		// Game: NewActivityFromDiscordgo(p.Activity), // not implemented by discordgo...
		// GuildID: dgoIDStringToUint64(p.GuildID), // not implemented by discordgo..
		Status: dgoStatusToString(p.Status),
	}
}

func NewChannelFromDgo(c *discordgo.Channel) *Channel {
	return &Channel{
		ID:                   dgoIDToSnowflake(c.ID),
		GuildID:              dgoIDToSnowflake(c.GuildID),
		Name:                 c.Name,
		Topic:                c.Topic,
		Type:                 dgoChannelTypeToUint(c.Type),
		LastMessageID:        dgoIDToSnowflake(c.ID),
		NSFW:                 c.NSFW,
		Position:             uint(c.Position),
		Bitrate:              c.Bitrate,
		Recipients:           dgoUsersTODiscordUsers(c.Recipients),
		Messages:             dgoMessagesToDiscordMessages(c.Messages),
		PermissionOverwrites: dgoPermissionOverwritesToDiscordPermissionOverwrites(c.PermissionOverwrites),
	}
}

func NewVoiceStateFromDgo(vs *discordgo.VoiceState) *VoiceState {
	return &VoiceState{
		UserID:    dgoIDToSnowflake(vs.UserID),
		SessionID: dgoIDToSnowflake(vs.SessionID),
		ChannelID: dgoIDToSnowflake(vs.ChannelID),
		GuildID:   dgoIDToSnowflake(vs.GuildID),
		Suppress:  vs.Suppress,
		SelfMute:  vs.SelfMute,
		SelfDeaf:  vs.SelfDeaf,
		Mute:      vs.Mute,
		Deaf:      vs.Deaf,
	}
}

func NewPermissionOverwriteFromDgo(pm *discordgo.PermissionOverwrite) *PermissionOverwrite {
	return &PermissionOverwrite{
		ID:    dgoIDToSnowflake(pm.ID),
		Type:  pm.Type,
		Deny:  pm.Deny,
		Allow: pm.Allow,
	}
}

func NewAttachmentFromDgo(a *discordgo.MessageAttachment) *Attachment {
	return &Attachment{
		ID:       dgoIDToSnowflake(a.ID),
		Filename: a.Filename,
		Size:     uint(a.Size),
		URL:      a.URL,
		ProxyURL: a.ProxyURL,
		Height:   uint(a.Height),
		Width:    uint(a.Width),
	}
}
