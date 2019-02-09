package main

import "fmt"

// -------------------------------------------------------------
// ~ Interfaces & Types
// -------------------------------------------------------------

type BridgeState struct {
	Lights map[string]*Light `json:"lights"`
}

type Light struct {
	State     *LightState `json:"state"`
	Type      string      `json:"type"`
	Name      string      `json:"name"`
	ModelID   string      `json:"modelid"`
	SwVersion string      `json:"swversion"`
}

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

func (l *Light) String() string {
	return fmt.Sprintf("Name=\"%s\" Model=\"%s\" On=\"%x\" \n", l.Name, l.ModelID, l.State.On)
}

func (bs *BridgeState) String() string {
	str := ""
	for k, l := range bs.Lights {
		str += k + ": " + l.String()
	}

	return str
}
