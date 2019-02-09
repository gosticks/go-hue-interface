package main

import (
	"bytes"
	"fmt"
    "gopkg.in/yaml.v2"
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

// TODO: Rename if this will be placed in a seperate package
// ReadConfig ...
func ReadConfig(path string) (conf *Config, err error) {
    f, err := ioutil.ReadFile(path)
    if err != nil {
        return
	}
	
    err = yaml.Unmarshal(f, conf)
    if err != nil {
		return
	}

	// TODO: check wether a user is already created and if not create one.

    return
}

func (c *Config) WriteConfig(path string) (err error) {
	b, err := yaml.Marshal(c)

	err = ioutil.WriteFile(path, b, 0644)
	return
}