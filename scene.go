package hue

import (
	"encoding/json"
	"errors"
	"strconv"
)

const scenesEndpoint = "/scenes"

type AppData struct {
	Version int    `json:"version,omitempty"`
	Data 	string `json:"data,omitempty"`
}

// TODO: maybe change the naming and hierachy of the types.
type BaseScene struct {
	Name    string   `json:"name,omitempty"`
	Type 	string   `json:"type,omitempty"`
	Group   string   `json:"group,omitempty"`
	Lights  []int    `json:"lights,omitempty"`
	Owner   string   `json:"owner,omitempty"`
	Recycle	bool     `json:"recycle,omitempty"`
	Locked  bool     `json:"lcoked,omitempty"`
	AppData *AppData `json:"appdata,omitempty"`
	Picture string   `json:"picture,omitempty"`
	LastUpdated Time `json:"lastupdated,omitempty"`
	Version int `json:"version,omitempty"`
}

type Scene struct {
	BaseScene
	LightStates map[int][]*LightState `json:"lightstates,omitempty"`
}

type CreationScene struct {
	Scene
	TransitionTime int `json:"transitiontime,omitempty"`
}

// GetScenes returns all the hue lights
func (b *Bridge) GetScenes() (map[string]*BaseScene, error) {
	result := make(map[string]*BaseScene)
	err := b.getAndDecode(scenesEndpoint, &result)
	return result, err
}

// TODO: probably support the old version
// CreateScene will create a new Scene on the bridge. It will return the
// id of the new scene if no error occures.
func (b *Bridge) CreateScene(s *CreationScene) (string, error) {
	res, err := b.postToBridge(scenesEndpoint, s)
	if err != nil {
		return "", err
	}

	var result string
	// Unmarshal data
	errDecode := json.NewDecoder(res.Body).Decode(result)
	if errDecode != nil {
		return "", errDecode
	}

	return result, nil
}


// TODO:
// - modify scene

// GetScene will return the scene specified by the id. The result will contain the
//lights states of lights within the scene.
func (b *Bridge) GetScene(id string) (*Scene, error) {
	result := &Scene{}
	err := b.getAndDecode(scenesEndpoint + "/" + id, result)
	return result, err
}

// DeleteScene deletes the scene specified by the id. Additional to communication
// errors a non 2xx status code of the response will trigger an error.
func (b *Bridge) DeleteScene(id string) error {
	res, errCom := b.deleteFromBridge(scenesEndpoint + "/" + id, nil)
	if errCom != nil {
		return errCom
	}

	// TODO: Should we put the check into the bridge methods (or maybe remove it)?
	if res.StatusCode >= 200 && res.StatusCode <= 299 {
		return errors.New("Response has non 2xx status code: " + strconv.Itoa(res.StatusCode))
	}

	return nil
}