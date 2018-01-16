package unison

type Bot struct {
	Commands   []*Command
	EventHooks []*EventHook
	Services   []*Service
}
