package monitor

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/gosticks/go-hue-interface"
)

func TestLightMonitor(t *testing.T) {
	userID, okUser := os.LookupEnv("HUE_BRIDGE_USER")
	addr, okAddr := os.LookupEnv("HUE_BRIDGE_ADDR")
	if !okAddr || !okUser {
		fmt.Println("HUE_BRIDGE_USER and HUE_BRIDGE_ADDR must be set in env for this test to work")
		t.Fail()
	}

	conf := &hue.Config{
		Username:         userID,
		BridgeAddr:       addr,
		BridgeAddrScheme: "http",
	}
	b := hue.NewBridge(conf)

	m := NewMonitor(b, 1*time.Second)

	// Get first available light
	ls, lsErr := b.GetLights()
	if lsErr != nil {
		t.Fail()
		return
	}

	// get any light
	var (
		light *hue.Light
		id    string
	)
	for k, l := range ls {
		if light != nil {
			break
		}
		light = l
		id = k
	}

	fmt.Println("Adding listener to light: ", light)

	m.AddListener(id, Light, func(s interface{}) {
		fmt.Println("State changed: ", s.(*hue.LightState))
	})

	m.Start()

	select {}
}
