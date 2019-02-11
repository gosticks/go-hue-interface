package hue

import (
	"fmt"
)

// SwUpdate provides the current sw state and last install
type SwUpdate struct {
	State       string `json:"state,omitempty"`
	LastInstall Time   `json:"lastinstall,omitempty"`
}

// LightCapabilities type providing control and certification settings
type LightCapabilities struct {
	Certified bool                   `json:"certified"`
	Control   *LightControl          `json:"control"`
	Streaming *StreamingCapabilities `json:"streaming"`
}

type StreamingCapabilities struct {
	Renderer bool `json:"renderer"`
	Proxy    bool `json:"proxy"`
}

type DeviceConfig struct {
	ArcheType string         `json:"archetype"`
	Function  string         `json:"function"`
	Direction string         `json:"direction"`
	Startup   *DeviceStartUp `json:"startup,omitempty"`
}

type DeviceStartUp struct {
	Mode       string `json:"mode"`
	Configured bool   `json:"configured"`
}

type LightControl struct {
	MinDimLevel    uint16      `json:"mindimlevel,omitempty"`
	MaxLumen       uint        `json:"maxlumen,omitempty"`
	ColorGamutType string      `json:"colorgamuttype,omitempty"`
	ColorGamut     [][]float32 `json:"colorgamut,omitempty"`
	Ct             *LightCt    `json:"ct,omitempty"`
}

type LightCt struct {
	Min uint `json:"min"`
	Max uint `json:"max"`
}

// Light hue object
type Light struct {
	State            *LightState        `json:"state,omitempty"`
	SwUpdate         *SwUpdate          `json:"swupdate,omitempty"`
	Type             string             `json:"type,omitempty"`
	Name             string             `json:"name,omitempty"`
	ModelID          string             `json:"modelid,omitempty"`
	ManufacturerName string             `json:"manufacturername,omitempty"`
	Productname      string             `json:"productname,omitempty"`
	Capabilities     *LightCapabilities `json:"capabilities,omitempty"`
	Config           *DeviceConfig      `json:"config"`
	UniqueID         string             `json:"uniqueid,omitempty"`
	SwVersion        string             `json:"swversion,omitempty"`
	SwConfigID       string             `json:"swconfigid,omitempty"`
	ProductID        string             `json:"productid,omitempty"`
}

// LightState is the hue light>state object
type LightState struct {
	On             bool      `json:"on"`
	BridgeID       int       `json:"bri,omitempty"`
	Hue            uint16    `json:"hue,omitempty"`
	Sat            uint8     `json:"sat,omitempty"`
	Effect         string    `json:"effect,omitempty"`
	XY             []float32 `json:"xy,omitempty"`
	Ct             uint16    `json:"ct,omitempty"`
	Alert          string    `json:"alert,omitempty"`
	TransitionTime uint16    `json:"transitiontime,omitempty"`
	ColorMode      string    `json:"colormode,omitempty"`
	Mode           string    `json:"mode,omitempty"`
	Reachable      bool      `json:"reachable"`
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
