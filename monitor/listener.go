package monitor

import "reflect"

// ListenerType matches the Hue API types e.g. Lights, Scenes, Groups...
type ListenerType int

const (
	// Light is a hue light
	Light  ListenerType = 0
	Scene  ListenerType = 1
	Sensor ListenerType = 2
	Group  ListenerType = 3
)

// Listener stores the current state and a ref to the handler on the event of a change
type Listener struct {
	CurrentState interface{}
	Handler      func(interface{})
}

// hasChanged chacks if the value changed compared to the value stored on the listener.
func (l *Listener) hasChanged(newState interface{}) bool {
	return !reflect.DeepEqual(l.CurrentState, newState)
}

// Update is called by the monitor if the result of hasChanged is false. If the value did change state of the listener is updated and the handler is called.
func (l *Listener) Update(newState interface{}) {
	oldState := l.CurrentState
	l.CurrentState = newState
	if oldState != nil {
		l.Handler(newState)
	}
}
