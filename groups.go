package hue

import "fmt"

const groupsEndpoint = "/groups"

// -------------------------------------------------------------
// ~ Types
// -------------------------------------------------------------

// GroupType hue group type
type GroupType int

// RoomClasses in hue
type RoomClasses int

const (
	// All is a group containing all devices cannot be created
	All GroupType = iota
	Luminaire
	Lightsource
	LightGroup
	Room
	Entertainment
)

const (
	LivingRoom RoomClasses = iota
	Kitchen
	Dining
	Bedroom
	KidsBedroom
	Bathroom
	Nursery
	Recreation
	Office
	Gym
	Hallway
	Toilet
	FrontDoor
	Garage
	Terrace
	Garden
	Driveway
	Carport
	Other
)

// Group hue type
type Group struct {
	Name     string     `json:"name"`
	LightIDs []string   `json:"Lights"`
	Type     string     `json:"type"`
	Action   LightState `json:"action"`
	Class    string     `json:"class,omitempty"`
}

// GroupCreateResponse is returned after a create group request
type GroupCreateResponse struct {
	Success struct {
		ID string `json:"id"`
	} `json:"success"`
}

// -------------------------------------------------------------
// ~ String conversions
// -------------------------------------------------------------

func (g GroupType) String() string {
	return [...]string{"0", "Luminaire", "Lightsource", "LightGroup", "Room", "Entertainment"}[g]
}

func (r RoomClasses) String() string {
	return [...]string{
		"Living room",
		"Kitchen",
		"Dining",
		"Bedroom",
		"Kids bedroom",
		"Bathroom",
		"Nursery",
		"Recreation",
		"Office",
		"Gym",
		"Hallway",
		"Toilet",
		"Front door",
		"Garage",
		"Terrace",
		"Garden",
		"Driveway",
		"Carport",
		"Other",
	}[r]
}

// -------------------------------------------------------------
// ~ Private methods
// -------------------------------------------------------------

// CreateGroup creates a new hue group. For rooms please use the CreateRoom call since it also needs a class
func (b *Bridge) CreateGroup(name string, groupType GroupType, lights []string) (string, error) {
	// perform some checks
	if groupType != Lightsource && groupType != LightGroup && groupType != Entertainment {
		return "", fmt.Errorf("only Lightsource, LightGroup or Entertainment type groups can be created, to create a room group please use CreateRoom (As of now other groups cannot be created manually)")
	}

	groupConfig := &Group{
		LightIDs: lights,
		Name:     name,
		Type:     groupType.String(),
	}
	return b.createGroup(groupConfig)
}

// CreateRoom creates a new hue room.
func (b *Bridge) CreateRoom(name string, class RoomClasses, lights []string) (string, error) {
	groupConfig := &Group{
		LightIDs: lights,
		Name:     name,
		Type:     Room.String(),
		Class:    class.String(),
	}
	return b.createGroup(groupConfig)
}

// -------------------------------------------------------------
// ~ Private methods
// -------------------------------------------------------------

func (b *Bridge) createGroup(group *Group) (string, error) {
	res, errCreate := b.postToBridge(groupsEndpoint, group)
	if errCreate != nil {
		return "", errCreate
	}

	newGroups, ok := res.([]*GroupCreateResponse)
	if ok && len(newGroups) == 1 {
		return newGroups[0].Success.ID, nil
	} else {
		return "", fmt.Errorf("could not create group, bridge did not return new group id")
	}
}
