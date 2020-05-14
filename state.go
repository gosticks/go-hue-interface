package hue

import "encoding/json"

// -------------------------------------------------------------
// ~ Interfaces & Types
// -------------------------------------------------------------

// BridgeState provides all data for a bridge
type BridgeState struct {
	Lights map[string]*Light `json:"lights"`
}

func (bs *BridgeState) String() string {
	str := ""
	for k, l := range bs.Lights {
		str += k + ": " + l.String()
	}

	return str
}

// GetState returns the current hue state
func (b *Bridge) GetState() (state *BridgeState, err error) {
	state = &BridgeState{}
	res, err := b.getFromBridge("")
	if err != nil {
		return
	}

	// Unmarshal data
	errDecode := json.NewDecoder(res.Body).Decode(state)
	if errDecode != nil {
		return nil, errDecode
	}

	return state, err
}
