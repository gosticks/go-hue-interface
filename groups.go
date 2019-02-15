package hue

import (
	"encoding/json"
	"fmt"
)

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
	Name      string      `json:"name"`
	LightIDs  []string    `json:"lights"`
	SensorIDs []string    `json:"sensors"`
	Type      string      `json:"type"`
	State     *GroupState `json:"state"`
	Recycle   bool        `json:"recycle"`
	Class     string      `json:"class,omitempty"`
	Action    LightState  `json:"action"`
}

// GroupCreateResponse is returned after a create group request
type GroupCreateResponse struct {
	Success struct {
		ID string `json:"id"`
	} `json:"success"`
}

type GroupState struct {
	AllOn bool `json:"all_on"`
	AnyOn bool `json:"any_on"`
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
// ~ Public methods
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

// GetAllGroups returns all the groups for a hue bridge
func (b *Bridge) GetAllGroups() (map[string]*Group, error) {
	result := make(map[string]*Group)
	errCom := b.getAndDecode(groupsEndpoint, &result)
	if errCom != nil {
		return nil, errCom
	}

	return result, nil
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

// func (b *Bridge) SetGroupAttributes(id string)

// -------------------------------------------------------------
// ~ Private methods
// -------------------------------------------------------------

func (b *Bridge) createGroup(group *Group) (string, error) {
	res, errCreate := b.postToBridge(groupsEndpoint, group)
	if errCreate != nil {
		return "", errCreate
	}

	result := []*GroupCreateResponse{}

	// Unmarshal data
	errDecode := json.NewDecoder(res.Body).Decode(result)
	if errDecode != nil {
		return "", errDecode
	}

	if len(result) == 1 {
		return result[0].Success.ID, nil
	} else {
		return "", fmt.Errorf("could not create group, bridge did not return new group id")
	}
}
