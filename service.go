package unison

// CommandActionFunc is the action to execute when a command is called.
type ServiceActionFunc func(ctx *Context) error

type Service struct {
	Name string

	Description string

	Deactivated bool

	Action ServiceActionFunc
	
	Data map[string]string // store realtime data here
}
