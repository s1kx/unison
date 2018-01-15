package session

import (
	"encoding/json"

	"github.com/s1kx/discordgo"
	"github.com/s1kx/unison/discord"
)

// GetAuditLogs Get the last 50 audit logs for the given guild
//	params interface{} is a struct with json tags that are converted into GET url parameters
func GetAuditLogs(ctx *Context, guildID string, params interface{}) (*discord.AuditLog, error) {
	urlParams := "" //convertAuditLogParamsToStr(params)
	byteArr, err := ctx.Discord.Request("GET", discordgo.EndpointGuilds+guildID+"/audit-logs"+urlParams, nil)
	if err != nil {
		return nil, err
	}

	auditLog := &discord.AuditLog{}
	err = json.Unmarshal(byteArr, &auditLog)
	if err != nil {
		return nil, err
	}

	return auditLog, nil
}
