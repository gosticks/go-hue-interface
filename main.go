package hue

import (
	"fmt"
	"time"
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

	// errCom := bridge.ToggleLight("2", false)
	// if errCom != nil {
	// 	fmt.Println("[ERROR]" + errCom.Error())
	// }
	//fmt.Println(test)
}

func strobeLight(b *Bridge, id string) {
	ticker := time.NewTicker(200 * time.Millisecond)
	quit := make(chan struct{})
	go func() {
		state := false
		for {
			select {
			case <-ticker.C:
				resp, errCom := b.ToggleLight(id, state)
				if errCom != nil {
					fmt.Println("[ERROR]" + errCom.Error())
					ticker.Stop()
					return
				}
				if resp.Error != nil {
					fmt.Println("[ERROR]" + resp.Error.String())
					ticker.Stop()
					return
				}
				state = !state
				// do stuff
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()
}
