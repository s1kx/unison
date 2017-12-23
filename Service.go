package unison

// ServiceActionFunc is the action to execute when a service runs.
type ServiceActionFunc func(ctx *Context) error

// Service struct used to create new services
type Service struct {
	Name string

	Description string

	Deactivated bool

	Action ServiceActionFunc

	Data map[string]string // store realtime data here
}
