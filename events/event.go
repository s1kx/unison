package events

type DiscordEvent struct {
	Type  EventType
	Event interface{}
}

func NewDiscordEvent(v interface{}) (*DiscordEvent, error) {
	t, err := GetEventType(v)
	if err != nil {
		return nil, err
	}

	ev := &DiscordEvent{
		Type:  t,
		Event: v,
	}
	return ev, nil
}
