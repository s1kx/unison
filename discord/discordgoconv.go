package discord

import (
	"strconv"
	"time"

	"gopkg.in/bwmarrin/Discordgo.v0"
)

const discordgoTimestampLayout string = "2016-08-06T17:20:33.803-0400"

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
	timestamp, err := time.Parse(discordgoTimestampLayout, string(ts))
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

func discordgoCopyTodiscordStruct(discordgoStruct interface{}, discordStruct interface{}) {
	// TODO use reflection to copy over values with similar type and json tag.
}
