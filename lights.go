package main

import "fmt"

// Light hue object
type Light struct {
	State     *LightState `json:"state,omitempty"`
	Type      string      `json:"type,omitempty"`
	Name      string      `json:"name,omitempty"`
	ModelID   string      `json:"modelid,omitempty"`
	SwVersion string      `json:"swversion,omitempty"`
}

// LightState is the hue light>state object
type LightState struct {
	On        bool   `json:"on"`
	BridgeID  int    `json:"bridgeid,omitempty"`
	Hue       int    `json:"hue,omitempty"`
	XY        []int  `json:"xy,omitempty"`
	Ct        int    `json:"ct,omitempty"`
	Alert     string `json:"alert,omitempty"`
	Effect    string `json:"effect,omitempty"`
	ColorMode string `json:"colormode,omitempty"`
	Reachable bool   `json:"reachable,omitempty"`
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
