package unison

import (
	"sync"

	"github.com/s1kx/unison/events"
)

// EventHandlerFunc handles a discord event and returns whether it handles the
// event type and if an error occured.
type EventHandlerFunc func(ctx *Context, ev *events.DiscordEvent) (handled bool, err error)

// Hook interface for anything that is supposed to react on a event, besides commands.
type EventHook struct {
	// Name of the hook
	Name string

	// Description of what the hook does
	Description string

	// Events that the hook should react to
	Events []events.EventType

	// Check if this hook is deactivated
	Deactivated bool

	// Command behavior
	OnEvent EventHandlerFunc
}

type eventDispatcher struct {
	mu             sync.RWMutex
	hookMap        map[string]*EventHook
	typeToHooksMap map[events.EventType]map[string]*EventHook
}

func newEventDispatcher() *eventDispatcher {
	return &eventDispatcher{
		hookMap:        make(map[string]*EventHook),
		typeToHooksMap: make(map[events.EventType]map[string]*EventHook),
	}
}

func (d *eventDispatcher) GetHooks() map[string]*EventHook {
	d.mu.RLock()
	defer d.mu.RUnlock()

	// Copy elements to new map rather than passing on a pointer to the internal map
	nm := make(map[string]*EventHook, len(d.hookMap))
	for k, v := range d.hookMap {
		nm[k] = v
	}

	return nm
}

func (d *eventDispatcher) AddHook(hook *EventHook) error {
	d.mu.Lock()
	defer d.mu.Unlock()

	name := hook.Name

	// Add to hookMap
	if ex, exists := d.hookMap[name]; exists {
		return &DuplicateEventHookError{Existing: ex, New: hook}
	}
	d.hookMap[name] = hook

	// Add hook to reverse lookup map for type => hooks
	for _, t := range hook.Events {
		// Create entry in typeToHooksMap if it doesn't exist
		if _, ok := d.typeToHooksMap[t]; !ok {
			d.typeToHooksMap[t] = make(map[string]*EventHook)
		}
		hooks := d.typeToHooksMap[t]

		// Add hook to hook map for given type
		hooks[name] = hook
	}

	return nil
}

func (d *eventDispatcher) Dispatch(ctx *Context, event *events.DiscordEvent) error {
	d.mu.RLock()
	defer d.mu.RUnlock()

	t := event.Type
	if _, ok := d.typeToHooksMap[t]; !ok {
		// No hooks exist for given event type
		return nil
	}
	hooks := d.typeToHooksMap[t]

	for _, hook := range hooks {
		// TODO: Run event handler in goroutine
		hook.OnEvent(ctx, event)
	}

	return nil
}
