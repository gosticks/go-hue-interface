package hue

type GroupType int

const (
	All GroupType = iota
	Luminaire
	Lightsource
	LightGroup
	Room
	Entertainment
)
