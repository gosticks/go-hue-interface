package main

// Config hue api config
type Config struct {
	Username         string `yaml:name`
	Password         string `yaml:userpassword`
	BridgeAddr       string `yaml:bridgeAddress`
	BridgeAddrScheme string `yaml:bridgeAddressScheme`
}
