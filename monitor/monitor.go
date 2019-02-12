package monitor

import (
	"fmt"
	"reflect"
	"time"

	"github.com/gosticks/go-hue-interface"
)

type ListenerType int

const (
	Light  ListenerType = 0
	Scene  ListenerType = 1
	Sensor ListenerType = 2
	Group  ListenerType = 3
)

// Monitor is a interface which is used for checking events
type Monitor struct {
	Listeners map[ListenerType]map[string]*Listener
	Interval  time.Duration
	Bridge    *hue.Bridge
	ticker    *time.Ticker
}

func NewMonitor(b *hue.Bridge, interval time.Duration) *Monitor {
	return &Monitor{
		Listeners: make(map[ListenerType]map[string]*Listener),
		Interval:  interval,
		Bridge:    b,
	}
}

func (m *Monitor) Tick() error {
	st, errState := m.Bridge.GetState()
	if errState != nil {
		return errState
	}

	ls, hasLightListener := m.Listeners[Light]
	if hasLightListener {
		// Loop over lights
		for id, listener := range ls {
			light, found := st.Lights[id]
			if found && listener.hasChanged(light.State) {
				listener.Update(light.State)
			}
		}

	}
	return nil
}

func (m *Monitor) Start() {
	m.ticker = time.NewTicker(m.Interval)
	go func() {
		for _ = range m.ticker.C {
			// fmt.Println("Tick: ", t)
			//Call the periodic function here.
			errTick := m.Tick()
			if errTick != nil {
				fmt.Println("[ERROR] failed to perform tick")
			}
		}
	}()
}

func (m *Monitor) AddListener(id string, t ListenerType, handler func(interface{})) {
	if m.Listeners[t] == nil {
		m.Listeners[t] = make(map[string]*Listener)
	}
	m.Listeners[t][id] = &Listener{Handler: handler}
}

type Listener struct {
	CurrentState interface{}
	Handler      func(interface{})
}

func (l *Listener) hasChanged(newState interface{}) bool {
	return !reflect.DeepEqual(l.CurrentState, newState)
}

func (l *Listener) Update(newState interface{}) {
	oldState := l.CurrentState
	l.CurrentState = newState
	if oldState != nil {
		l.Handler(newState)
	}
}

// func (b *hue.Bridge) AddMonitor() {

// }
