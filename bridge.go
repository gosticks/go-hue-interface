package main

import (
	"bytes"
	"encoding/json"
	"net/http"
)

// Bridge is the hue bridge interface
type Bridge struct {
	Config *Config
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

func (b *Bridge) getBridgeAPIURI() string {
	return b.Config.BridgeAddrScheme + "://" + b.Config.BridgeAddr + "/api/" + b.Config.Username
}
