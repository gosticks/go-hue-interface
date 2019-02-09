package main

import "fmt"

// Light hue object
type Light struct {
	State     *LightState `json:"state"`
	Type      string      `json:"type"`
	Name      string      `json:"name"`
	ModelID   string      `json:"modelid"`
	SwVersion string      `json:"swversion"`
}

// LightState is the hue light>state object
type LightState struct {
	On        bool   `json:"on"`
	BridgeID  int    `json:"bridgeid"`
	Hue       int    `json:"hue"`
	XY        []int  `json:"xy"`
	Ct        int    `json:"ct"`
	Alert     string `json:"alert"`
	Effect    string `json:"effect"`
	ColorMode string `json:"colormode"`
	Reachable bool   `json:"reachable"`
}

type cmdResponse struct {
}

const LightsEndpoint = "/lights"

func (l *Light) String() string {
	return fmt.Sprintf("Name=\"%s\" Model=\"%s\" On=\"%x\" XY=\"%x\" \n", l.Name, l.ModelID, l.State.On, l.State.XY)
}

func (b *Bridge) ToggleLight(id string, on bool) error {
	cmd := &LightState{
		On: on,
	}
	return b.putToBridge(LightsEndpoint+"/"+id+"/state", cmd, nil)
}
