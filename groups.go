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
	Name      string         `json:"name"`
	LightIDs  []string       `json:"lights"`
	SensorIDs []string       `json:"sensors"`
	Type      string         `json:"type"`
	State     *GroupState    `json:"state"`
	Recycle   bool           `json:"recycle"`
	ModelID   string         `json:"modelid,omitempty"`
	UniqueID  string         `json:"uniqueid,omitempty"`
	Class     string         `json:"class,omitempty"`
	Action    LightState     `json:"action"`
	Presence  *GroupPresence `json:"presence,omitempty"`
	// Lightlevel..
}

// GroupPresence only exists if sensors array contains a presence sensor of type “ZLLPresence”, “CLIPPresence” or “Geofence”. This object contains a state object which contains the aggregated state of the sensors
type GroupPresence struct {
	// State?
	LastUpdated Time `json:"lastupdated"`
	Presence    bool `json:"presence"`
	PresenceAll bool `json:"presence_all"`
}

// GroupCreateResponse is returned after a create group request
type GroupCreateResponse struct {
	Success struct {
		ID string `json:"id"`
	} `json:"success"`
}

// GroupState describes the state of a group
type GroupState struct {
	AllOn bool `json:"all_on"`
	AnyOn bool `json:"any_on"`
}

// GroupAttributes that can be changed
type GroupAttributes struct {
	Name     string   `json:"name,omitempty"`
	LightIDs []string `json:"lights,omitempty"`
	Class    string   `json:"class"`
}

// GroupAction is struct for changing a state of a hue group
// TODO: merge with light actions maybe?
type GroupAction struct {
	On             string    `json:"on,omitempty"`
	Bri            uint8     `json:"bri,omitempty"`
	Hue            uint16    `json:"hue,omitempty"`
	Sat            uint8     `json:"sat,omitempty"`
	Xy             []float32 `json:"xy,omitempty"`
	Ct             uint16    `json:"ct,omitempty"`
	Alert          string    `json:"alert,omitempty"`
	Effect         string    `json:"effect,omitempty"`
	TransitionTime uint16    `json:"transitiontime,omitempty"`
	BriInc         int16     `json:"bri_inc,omitempty"`
	SatInc         int16     `json:"sat_inc,omitempty"`
	HueInc         int       `json:"hue_inc,omitempty"`
	CtInc          int       `json:"ct_inc,omitempty"`
	XyInc          int8      `json:"xy_inc,omitempty"`
	Scene          string    `json:"scene,omitempty"`
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

// GetAllGroups returns all the groups for a hue bridge
func (b *Bridge) GetAllGroups() (map[string]*Group, error) {
	result := make(map[string]*Group)
	errCom := b.getAndDecode(groupsEndpoint, &result)
	if errCom != nil {
		return nil, errCom
	}

	return result, nil
}

// GetGroupAttributes returns the state of a group by id
func (b *Bridge) GetGroupAttributes(id string) (*Group, error) {
	result := &Group{}
	errCom := b.getAndDecode(groupsEndpoint+"/"+id, &result)
	if errCom != nil {
		return nil, errCom
	}

	return result, nil
}

// SetGroupAttributes updates a groups settings by adding devices or changing name or class
func (b *Bridge) SetGroupAttributes(id string, attributes *GroupAttributes) (*BridgeResponse, error) {
	res, errCom := b.putToBridge(groupsEndpoint+"/"+id, attributes)
	if errCom != nil {
		return nil, errCom
	}

	result := &BridgeResponse{}

	// Unmarshal data
	errDecode := json.NewDecoder(res.Body).Decode(result)
	if errDecode != nil {
		return nil, errDecode
	}

	return result, nil
}

// SetGroupState sets the state of a group by id
func (b *Bridge) SetGroupState(id string, action *GroupAction) ([]*BridgeResponse, error) {
	res, errCom := b.putToBridge(groupsEndpoint+"/"+id+"/action", action)
	if errCom != nil {
		return nil, errCom
	}

	result := []*BridgeResponse{}

	// Unmarshal data
	errDecode := json.NewDecoder(res.Body).Decode(&result)
	if errDecode != nil {
		return nil, errDecode
	}

	return result, nil
}

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
