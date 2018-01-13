package unison

import (
	"strconv"

	"github.com/sirupsen/logrus"
)

// DiscordPermissionFlags limits member accessability to discord guilds/channels functionality
// 	https://discordapp.com/developers/docs/topics/permissions
// type DiscordPermissionFlags uint64

func init() {
	// make sure the ToStr and other methods
	if uint64(0x8000040000200001) != uint64(0x8000040000200001) {
		logrus.Fatal("Did the DiscordPermissionFalgs type change? update ToStr() method.")
	}
}

type DiscordPermissions struct {
	flags uint64
}

func NewDiscordPermissionsDiscordgoWrapper(flags uint64, err error) (*DiscordPermissions, error) {
	return &DiscordPermissions{flags: flags}, err
}

func (dp *DiscordPermissions) Add(flags uint64) {
	dp.flags |= flags
}

func (dp *DiscordPermissions) Remove(flags uint64) {
	// be sure the flags exist
	dp.Add(flags)

	// Remove the desired flags
	dp.flags ^= flags
}

func (dp *DiscordPermissions) Set(flags uint64) {
	dp.flags = flags
}

func (dp *DiscordPermissions) Get() uint64 {
	return dp.flags
}

func (dp *DiscordPermissions) HasRequiredPermissions(permission *DiscordPermissions) bool {
	return (permission.Get() & dp.flags) == dp.flags
}

func (dp *DiscordPermissions) String() string {
	return strconv.FormatUint(dp.flags, 10)
}
