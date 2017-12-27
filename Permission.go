package unison

import (
	"strconv"

	"github.com/Sirupsen/logrus"
)

// DiscordPermissionFlags limits member accessability to discord guilds/channels functionality
// 	https://discordapp.com/developers/docs/topics/permissions
type DiscordPermissionFlags uint64

func init() {
	// make sure the ToStr and other methods
	if uint64(DiscordPermissionFlags(0x8000040000200001)) != uint64(0x8000040000200001) {
		logrus.Fatal("Did the DiscordPermissionFalgs type change? update ToStr() method.")
	}
}

type DiscordPermissions struct {
	flags DiscordPermissionFlags
}

func NewDiscordPermissionsDiscordgoWrapper(flags int, err error) (*DiscordPermissions, error) {
	return &DiscordPermissions{flags: DiscordPermissionFlags(flags)}, err
}

func (dp *DiscordPermissions) Add(flags DiscordPermissionFlags) {
	dp.flags |= flags
}

func (dp *DiscordPermissions) Remove(flags DiscordPermissionFlags) {
	// be sure the flags exist
	dp.Add(flags)

	// Remove the desired flags
	dp.flags ^= flags
}

func (dp *DiscordPermissions) Set(flags DiscordPermissionFlags) {
	dp.flags = flags
}

func (dp *DiscordPermissions) Get() DiscordPermissionFlags {
	return dp.flags
}

func (dp *DiscordPermissions) HasRequiredPermissions(permission *DiscordPermissions) bool {
	return (permission.Get() & dp.flags) == dp.flags
}

func (dp *DiscordPermissions) ToStr() string {
	return strconv.FormatUint(uint64(dp.flags), 10)
}
