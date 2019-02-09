package main

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type Bridge struct {
	Config *Config
}

func (b *Bridge) postToBridge(endpoint string, payload interface{}) (interface{}, error) {
	data, errMarhshal := json.Marshal(payload)
	if errMarhshal != nil {
		return nil, errMarhshal
	}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}
