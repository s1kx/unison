package discord

import (
	"fmt"

	"gopkg.in/bwmarrin/Discordgo.v0"
)

type User struct {
	ID            uint64 `json:"id"`
	Email         string `json:"email"`
	Username      string `json:"username"`
	Avatar        string `json:"avatar"`
	Discriminator string `json:"discriminator"`
	Token         string `json:"token"`
	Verified      bool   `json:"verified"`
	MFAEnabled    bool   `json:"mfa_enabled"`
	Bot           bool   `json:"bot"`
}

func NewUser() *User {
	return &User{}
}

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

func (u *User) Mention() string {
	return fmt.Sprintf("<@%d>", u.ID)
}

func (u *User) MentionNickname() string {
	return fmt.Sprintf("<@!%d>", u.ID)
}
