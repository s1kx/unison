package unison

import (
	"encoding/json"

	"gopkg.in/bwmarrin/Discordgo.v0"
)

type AuditLogChange struct {
	// will this even work? TODO, NOTE
	NewValue interface{} `json:"new_value"`
	OldValue interface{} `json:"old_value"`
	Key      string      `json:"key"`
}

type AuditLogOption struct {
	DeleteMemberDays string `json:"delete_member_days"`
	MembersRemoved   string `json:"members_removed"`
	ChannelID        string `json:"channel_id"`
	Count            string `json:"count"`
	ID               string `json:"id"`
	Type             string `json:"type"`
	RoleName         string `json:"role_name"`
}

type AuditLogEntry struct {
	TargetID   string            `json:"target_id"`
	UserID     string            `json:"user_id"`
	ID         string            `json:"id"`
	ActionType uint64            `json:"action_type"`
	Changes    []*AuditLogChange `json:"changes"`
	Options    []*AuditLogOption `json:"options"`
	Reason     string            `json:"reason"`
}

type AuditLog struct {
	Webhooks        []*discordgo.Webhook `json:"webhooks"`
	Users           []*discordgo.User    `json:"users"`
	AuditLogEntries []*AuditLogEntry     `json:"audit_log_entries"`
}

// AuditLogParams set params used in endpoint request
type AuditLogParams struct {
	UserID     string `urlparam:"user_id,omitempty"`
	ActionType int    `urlparam:"action_type,omitempty"`
	Before     string `urlparam:"before,omitempty"`
	Limit      int    `urlparam:"limit,omitempty"`
}

//
// func convertAuditLogParamsToStr(params *AuditLogParams) string {
// 	var getParams string
//
// 	v := reflect.ValueOf(*params)
// 	t := reflect.TypeOf(*params)
// 	// Iterate over all available fields and read the tag value
// 	for i := 0; i < t.Elem().NumField(); i++ {
// 		// Get the field, returns https://golang.org/pkg/reflect/#StructField
// 		field := t.Field(i)
//
// 		// Get the field tag value
// 		tag := field.Tag.Get("urlparam")
//
// 		// check if it's omitempty
// 		tags := strings.Split(tag, ",")
// 		if len(tags) > 1 {
// 			var skip bool
// 			for _, tagDetail := range tags {
// 				if tagDetail == "omitempty" && reflect.DeepEqual(field, reflect.Zero(field.Type).Interface()) {
// 					skip = true
// 				}
// 			}
// 			if skip {
// 				continue
// 			}
// 		}
//
// 		getParams += "&" + tags[0] + "=" + v.Field(i).Interface().(string)
// 	}
//
// 	urlParams := ""
// 	if getParams != "" {
// 		urlParams = "?" + getParams[1:len(getParams)]
// 	}
//
// 	return urlParams
// }

// GetAuditLogs Get the last 50 audit logs for the given guild
//	params interface{} is a struct with json tags that are converted into GET url parameters
func GetAuditLogs(ctx *Context, guildID string, params interface{}) (*AuditLog, error) {
	urlParams := "" //convertAuditLogParamsToStr(params)
	byteArr, err := ctx.Discord.Request("GET", discordgo.EndpointGuilds+guildID+"/audit-logs"+urlParams, nil)

	auditLog := &AuditLog{}
	err = json.Unmarshal(byteArr, &auditLog)
	if err != nil {
		return nil, err
	}

	return auditLog, nil
}
