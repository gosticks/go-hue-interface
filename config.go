package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// Config hue api config
type Config struct {
	Username         string `yaml:name`
	Password         string `yaml:userpassword`
	BridgeAddr       string `yaml:bridgeAddress`
	BridgeAddrScheme string `yaml:bridgeAddressScheme`
}

// createNewUser will create a new user. This should be called only of there's none in the yaml config.
func (c *Config) createNewUser() {
	// TODO: read/create the url
	url := "http://192.168.178.46/api"

    var reqBody = []byte(`{"devicetype": "go-hue-interface#Philips hue"}`)
    req, err := http.NewRequest("POST", url, bytes.NewBuffer(reqBody))
	
    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()

    fmt.Println("response Status:", resp.Status)
    fmt.Println("response Headers:", resp.Header)
    body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))

    json.Unmarshal([]byte(str), &res)
    fmt.Println(res)
    fmt.Println(res.Fruits[0])
}
