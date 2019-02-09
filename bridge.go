package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

// Bridge is the hue bridge interface
type Bridge struct {
	Config *Config
}

type BridgeUserConfig struct {
	Name             string `json:"name"`
	APIVersion       string `json:"apiversion"`
	IPAddress        string `json:"ipaddress"`
	MAC              string `json:"mac"`
	BridgeID         string `json:"bridgeid"`
	DataStoreVersion string `json:"datastoreversion"`
	StarterKitID     string `json:"starterkitid"`
	ReplacesBridgeID string `json:"replacesbridgeid"`
}

// NewBridge creates a new bridge api instance
func NewBridge(conf *Config) *Bridge {
	return &Bridge{
		Config: conf,
	}
}

func (b *Bridge) postToBridge(endpoint string, payload interface{}) (interface{}, error) {
	data, errMarhshal := json.Marshal(payload)
	if errMarhshal != nil {
		return nil, errMarhshal
	}
	uri := b.getBridgeAPIURI() + endpoint
	req, err := http.NewRequest(http.MethodPost, uri, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (b *Bridge) getFromBridge(endpoint string, target interface{}) error {

	uri := b.getBridgeAPIURI() + endpoint

	req, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		return err
	}

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return err
	}

	if res.StatusCode != http.StatusOK {
		return errors.New("Hue responded with error" + res.Status + fmt.Sprint(res.StatusCode))
	}

	// Unmarshal data
	errDecode := json.NewDecoder(res.Body).Decode(target)
	if errDecode != nil {
		return err
	}

	return nil
}

func (b *Bridge) getBridgeAPIURI() string {
	return b.Config.BridgeAddrScheme + "://" + b.Config.BridgeAddr + "/api/" + b.Config.Username
}
