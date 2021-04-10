// Package eventbus provides a simple event bus which can emit events and hold listeners for the events triggered by your application.
package eventbus

import (
	"sort"
	"sync"
)

// EventData represents an event and its associated data.
type EventData struct {
	// Name of the event.
	Name string

	// Optionally, you could add tags here to add additional
	// information for receivers.
	// For example, when you have multiple services you could
	// specify the targets here.
	Tags []string

	// The content carried by the event.
	Content interface{}
}

// Listener represents an event listener.
//
// It contains the pointer to a function that accepts [EventData], and the priority of the event.
// Setting a higher priority means that the listener will be called earlier.
type Listener struct {
	Func     func(EventData)
	Priority uint
}

// EventBus holds the event listeners.
type EventBus struct {
	listeners map[string][]*Listener
	mutex     sync.Mutex
}

// New returns a new EventBus.
func New() *EventBus {
	return &EventBus{
		listeners: make(map[string][]*Listener),
	}
}

// AddListener registers a new listener for the specified event.
func (e *EventBus) AddListener(eventName string, listenerFunc func(EventData), priority uint) {
	listener := &Listener{
		Func:     listenerFunc,
		Priority: priority,
	}

	e.mutex.Lock()
	defer e.mutex.Unlock()

	e.listeners[eventName] = append(e.listeners[eventName], listener)

	// Sort listeners by priority.
	sort.Slice(e.listeners[eventName], func(i, j int) bool {
		return e.listeners[eventName][i].Priority > e.listeners[eventName][j].Priority
	})
}

// Reset removes all the listeners for all events.
func (e *EventBus) Reset() {
	e.mutex.Lock()
	defer e.mutex.Unlock()

	e.listeners = make(map[string][]*Listener)
}

// Emit emits an event with the specified data.
func (e *EventBus) Emit(data EventData) {
	e.mutex.Lock()
	defer e.mutex.Unlock()

	listeners, ok := e.listeners[data.Name]
	if !ok {
		return
	}

	for _, listener := range listeners {
		listener.Func(data)
	}
}
