package hue

import "reflect"

// Monitor is a interface which is used for checking events
type Monitor struct {
	CurrentState interface{}
	Handler      func(interface{})
}

func (m *Monitor) hasChanged(newState interface{}) bool {
	return reflect.DeepEqual(m.CurrentState, newState)
}

func (m *Monitor) Update(newState interface{}) {
	if m.hasChanged(newState) {
		m.Handler(newState)
	}
}

func (b *Bridge) AddMonitor() {

}
