package hue

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"io/ioutil"
)

// -------------------------------------------------------------
// Interfaces
// -------------------------------------------------------------

// Bridge is the hue bridge interface
type Bridge struct {
	Config *Config
}

// BridgeUserConfig is the config provided for hue for a user
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

// BridgeResponse is the response object returned to a bridge command
type BridgeResponse struct {
	Success map[string]interface{} `json:"success"`
	Error   *BridgeResponseError   `json:"error"`
}

// BridgeResponseError provides info about a bridge api error
type BridgeResponseError struct {
	Type        uint   `json:"type"`
	Address     string `json:"address"`
	Description string `json:"description"`
}

func (err *BridgeResponseError) String() string {
	return fmt.Sprintf("Type=\"%d\" Addr=\"%s\" Desc=\"%s\" \n", err.Type, err.Address, err.Description)
}

// -------------------------------------------------------------
// Methods
// -------------------------------------------------------------

// NewBridge creates a new bridge api instance
func NewBridge(conf *Config) *Bridge {
	return &Bridge{
		Config: conf,
	}
}

func (b *Bridge) postToBridge(endpoint string, payload interface{}) (*http.Response, error) {
	return b.sendToBridge(endpoint, http.MethodPost, payload)
}

func (b *Bridge) putToBridge(endpoint string, payload interface{}) (*http.Response, error) {
	return b.sendToBridge(endpoint, http.MethodPost, payload)
}

func (b *Bridge) sendToBridge(endpoint string, method string, payload interface{}) (*http.Response, error) {
	data, errMarhshal := json.Marshal(payload)
	if errMarhshal != nil {
		return nil, errMarhshal
	}
	uri := b.getBridgeAPIURI() + endpoint
	req, err := http.NewRequest(method, uri, bytes.NewBuffer(data))
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

func (b *Bridge) getRawResponse(endpoint string) ([]byte, error) {
	res, errCom := b.getFromBridge(endpoint)
	if errCom != nil {
		return nil, errCom
	}
	defer res.Body.Close()

	respBytes, errRead := ioutil.ReadAll(res.Body)
	if errRead != nil {
		return nil, errRead
	}
	return respBytes, nil
}

func (b *Bridge) getFromBridge(endpoint string) (*http.Response, error) {

	uri := b.getBridgeAPIURI() + endpoint

	req, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	// check http responses
	if res.StatusCode != http.StatusOK {
		return nil, errors.New("Hue responded with error" + res.Status + fmt.Sprint(res.StatusCode))

	}

	return res, nil
}

// getAndDecode performs a get request and unmarshals the result into target
func (b *Bridge) getAndDecode(endpoint string, target interface{}) error {
	res, errCom := b.getFromBridge(endpoint)
	if errCom != nil {
		return errCom
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return errors.New("Hue responded with error" + res.Status + fmt.Sprint(res.StatusCode))
	}

	// Unmarshal data
	errDecode := json.NewDecoder(res.Body).Decode(target)
	if errDecode != nil {
		return errDecode
	}

	return nil
}

func (b *Bridge) getBridgeAPIURI() string {
	return b.Config.BridgeAddrScheme + "://" + b.Config.BridgeAddr + "/api/" + b.Config.Username
}

// func decodeResponse(r *htpp.Response, target interface{}) {
// 	defer r.Body.Close()

// 	// TODO: remove
// 	fmt.Println("response Status:", res.Status)
// 	fmt.Println("response Headers:", res.Header)
// 	body, _ := ioutil.ReadAll(res.Body)
// 	fmt.Println("response Body:", string(body))

// 	if res.StatusCode != http.StatusOK {
// 		return errors.New("Hue responded with error" + res.Status + fmt.Sprint(res.StatusCode))
// 	}

// 	// Unmarshal data
// 	if respData != nil {
// 		errDecode := json.NewDecoder(res.Body).Decode(respData)
// 		if errDecode != nil {
// 			return nil, errDecode
// 		}
// 	}
// }
