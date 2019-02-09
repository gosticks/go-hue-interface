package main

import "fmt"

// -------------------------------------------------------------
// ~ Interfaces & Types
// -------------------------------------------------------------

type BridgeState struct {
	Lights map[string]*Light `json:"lights"`
}

func (l *Light) String() string {
	return fmt.Sprintf("Name=\"%s\" Model=\"%s\" On=\"%x\" XY=\"%x\" \n", l.Name, l.ModelID, l.State.On, l.State.XY)
}

func (bs *BridgeState) String() string {
	str := ""
	for k, l := range bs.Lights {
		str += k + ": " + l.String()
	}

	return str
}
