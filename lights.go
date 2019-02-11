package hue

import (
	"fmt"
)

// SwUpdate provides the current sw state and last install
type SwUpdate struct {
	State       string `json:"state,omitempty"`
	LastInstall string `json:"lastinstall,omitempty"`
}

// LightCapabilities type providing control and certification settings
type LightCapabilities struct {
	Certified bool `json:"certified"`
}

// Light hue object
type Light struct {
	State            *LightState `json:"state,omitempty"`
	Type             string      `json:"type,omitempty"`
	Name             string      `json:"name,omitempty"`
	ModelID          string      `json:"modelid,omitempty"`
	ManufacturerName string      `json:"manufacturername,omitempty"`
	Productname      string      `json:"productname,omitempty"`
	SwVersion        string      `json:"swversion,omitempty"`
	UniqueID         string      `json:"uniqueID,omitempty"`
	SwUpdate         *SwUpdate   `json:"swupdate,omitempty"`
	// "capabilities": {
	// 	"certified": true,
	// 	"control": {
	// 		"mindimlevel": 5000,
	// 		"maxlumen": 600,
	// 		"colorgamuttype": "B",
	// 		"colorgamut": [
	// 			[
	// 				0.675,
	// 				0.322
	// 			],
	// 			[
	// 				0.409,
	// 				0.518
	// 			],
	// 			[
	// 				0.167,
	// 				0.04
	// 			]
	// 		],
	// 		"ct": {
	// 			"min": 153,
	// 			"max": 500
	// 		}
	// 	},
	// 	"streaming": {
	// 		"renderer": true,
	// 		"proxy": false
	// 	}
	// },
	// "config": {
	// 	"archetype": "sultanbulb",
	// 	"function": "mixed",
	// 	"direction": "omnidirectional"
	// },
}

// LightState is the hue light>state object
type LightState struct {
	On             bool      `json:"on"`
	BridgeID       int       `json:"bridgeid,omitempty"`
	Hue            uint16    `json:"hue,omitempty"`
	Sat            uint8     `json:"sat,omitempty"`
	XY             []float32 `json:"xy,omitempty"`
	Ct             uint16    `json:"ct,omitempty"`
	Alert          string    `json:"alert,omitempty"`
	Effect         string    `json:"effect,omitempty"`
	TransitionTime uint16    `json:"transitiontime,omitempty"`
	ColorMode      string    `json:"colormode,omitempty"`
	Reachable      bool      `json:"reachable,omitempty"`
}

// LightsEndpoint for the lights
const LightsEndpoint = "/lights"

func (l *Light) String() string {
	return fmt.Sprintf("Name=\"%s\" ModelID=\"%s\" ProductName=\"%s\" On=\"%v\" Manu=\"%s\" \n", l.Name, l.ModelID, l.Productname, l.State.On, l.ManufacturerName)
}

// ToggleLight switches light on or off
func (b *Bridge) ToggleLight(id string, on bool) (resp *BridgeResponse, err error) {
	state := &LightState{
		On: on,
	}
	return b.SetLightState(id, state)
}

// SetLightState updates the light state
func (b *Bridge) SetLightState(id string, state *LightState) (resp *BridgeResponse, err error) {
	err = b.putToBridge(LightsEndpoint+"/"+id+"/state", state, resp)
	return resp, err
}

// GetLights returns all the hue lights
func (b *Bridge) GetLights() (map[string]*Light, error) {
	result := make(map[string]*Light)
	err := b.getFromBridge("/lights", &result)
	return result, err
}
