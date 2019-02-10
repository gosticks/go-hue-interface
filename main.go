package main

import (
	"fmt"
)

const VERSION = "0.1.2"

func main() {
	fmt.Printf("go-hue-interface v. %s \n", VERSION)

	config := &Config{
		Username:         "nX1ye7AMQoQswiiJdxyw-92-RNhIwicXiQRg7AtF",
		BridgeAddr:       "192.168.178.46",
		BridgeAddrScheme: "http",
	}

	bridge := NewBridge(config)

	fmt.Println("Created bridge ", bridge)

	test := &BridgeState{}
	errCom := bridge.ToggleLight("2", false)
	if errCom != nil {
		fmt.Println("[ERROR]" + errCom.Error())
	}
	fmt.Println(test)
}
