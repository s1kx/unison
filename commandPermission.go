package unison

// import "github.com/bwmarrin/discordgo"

// CommandPermission Designed to handle X guilds per bot instance
type CommandPermission struct {
	// Used for specific User banning on a per command basis
	BannedUserIDs []string // *discordgo.User.ID

	// Users that has been given access to a command even though they do not have the required role.
	AllowedUserIDs []string // *discordgo.User.ID

	// What roles can run this command
	AllowedRoles []string // *discordgo.GuildRole.Role.ID
}

// NewCommandPermission creates a new default filled command permission struct
func NewCommandPermission() CommandPermission {
	cmdP := CommandPermission{}

	cmdP.BannedUserIDs = []string{}
	cmdP.AllowedRoles = []string{"*"}
	cmdP.AllowedUserIDs = []string{"*"}

	return cmdP
}
