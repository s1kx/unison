package discord

type AuditLogChange struct {
	// will this even work? TODO, NOTE
	NewValue interface{} `json:"new_value"`
	OldValue interface{} `json:"old_value"`
	Key      string      `json:"key"`
}

type AuditLogOption struct {
	DeleteMemberDays string `json:"delete_member_days"`
	MembersRemoved   string `json:"members_removed"`
	ChannelID        uint64 `json:"channel_id"`
	Count            string `json:"count"`
	ID               uint64 `json:"id"`
	Type             string `json:"type"` // type of overwritten entity ("member" or "role")
	RoleName         string `json:"role_name"`
}

type AuditLogEntry struct {
	TargetID   uint64            `json:"target_id"`
	UserID     uint64            `json:"user_id"`
	ID         uint64            `json:"id"`
	ActionType uint8             `json:"action_type"`
	Changes    []*AuditLogChange `json:"changes"`
	Options    []*AuditLogOption `json:"options"`
	Reason     string            `json:"reason"`
}

type AuditLog struct {
	Webhooks        []*Webhook       `json:"webhooks"`
	Users           []*User          `json:"users"`
	AuditLogEntries []*AuditLogEntry `json:"audit_log_entries"`
}

// AuditLogParams set params used in endpoint request
type AuditLogParams struct {
	UserID     uint64 `urlparam:"user_id,omitempty"`
	ActionType uint8  `urlparam:"action_type,omitempty"`
	Before     uint64 `urlparam:"before,omitempty"`
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
