package hue

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
func (b *Bridge) GetState() (*BridgeState, error) {
	state := &BridgeState{}
	err := b.getFromBridge("", state)
	return state, err
}
