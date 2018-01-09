package unison

import (
	"encoding/json"
	"errors"
  "gopkg.in/bwmarrin/Discordgo.v0"
)

type AuditLogChange struct {
	NewValue interface{} `json:"new_value"`
	OldValue interface{} `json:"old_value"`
	Key      string      `json:"key"`
}

type AuditLogEntry struct {
	TargetID   string `json:"target_id"`
	UserID     string `json:"user_id"`
	ID         string `json:"id"`
	ActionType uint64 `json:"action_type"`
	//Changes    []*AuditLogChange `json:"changes"`
	//Options []*
	Reason string `json:"reason"`
}

type AuditLog struct {
	Webhooks        []*discordgo.Webhook `json:"webhooks"`
	Users           []*discordgo.User    `json:"users"`
	AuditLogEntries []*AuditLogEntry     `json:"audit_log_entries"`
}

// GetAuditLogs Get the last 50 audit logs for the given guild
func GetAuditLogs(ctx *Context, guildID string) (*AuditLog, error) {
	byteArr, err := ctx.Discord.Request("GET", discordgo.EndpointGuilds+guildID+"/audit-logs", nil)
  
	auditLog := &AuditLog{}
	err = json.Unmarshal(bytes, &auditLog)
	if err != nil {
		return nil, err
	}
  
  return auditLog, nil
}
