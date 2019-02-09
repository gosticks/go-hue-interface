package main

import (
	"fmt"
)

const VERSION = "0.1.2"

func main() {
	fmt.Printf("go-hue-interface v. %s \n", VERSION)

	config := &Config{
		Username: "nX1ye7AMQoQswiiJdxyw-92-RNhIwicXiQRg7AtF",
	}

	bridge := NewBridge(config)

	fmt.Println("Created bridge ", bridge)
}
