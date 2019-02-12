package monitor

import (
	"fmt"
	"time"

	"github.com/gosticks/go-hue-interface"
)

// Monitor is a interface which is used for checking events
type Monitor struct {
	Listeners map[ListenerType]map[string]*Listener
	Interval  time.Duration
	Bridge    *hue.Bridge
	ticker    *time.Ticker
}

// NewMonitor creates a new monitor. The monitor can be used to observe changes to specific lamps, scenes, groups..
func NewMonitor(b *hue.Bridge, interval time.Duration) *Monitor {
	return &Monitor{
		Listeners: make(map[ListenerType]map[string]*Listener),
		Interval:  interval,
		Bridge:    b,
	}
}

// Tick starts a update process by iterating over all the listeners for each type. Tick normally should be called by the provided Start method but can also be use manually.
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

// Start begins a monitoring loop with the interval set on the monitor
func (m *Monitor) Start() {
	m.ticker = time.NewTicker(m.Interval)
	go func() {
		for range m.ticker.C {
			// fmt.Println("Tick: ", t)
			//Call the periodic function here.
			errTick := m.Tick()
			if errTick != nil {
				fmt.Println("[ERROR] failed to perform tick")
			}
		}
	}()
}

// AddListener adds a new listener to the listener map. The listener will only be called if the id and type match and the state has changed. If a listener for a id and type already exists it will be removed
func (m *Monitor) AddListener(id string, t ListenerType, handler func(interface{})) {
	if m.Listeners[t] == nil {
		m.Listeners[t] = make(map[string]*Listener)
	}
	m.Listeners[t][id] = &Listener{Handler: handler}
}

// TODO: add RemoveListener and Stop
