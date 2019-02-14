[![Go Report Card](https://goreportcard.com/badge/github.com/gosticks/go-hue-bridge)](https://goreportcard.com/report/github.com/gosticks/go-mite) [![License: GPL v3](https://img.shields.io/badge/License-GPLv3-blue.svg)](https://www.gnu.org/licenses/gpl-3.0)

# go-hue-interface

## Connect to a hue bridge

In your code create a hue bridge instance

```
    func main() {
        // ...
        // create a bridge config
        bridgeConf := &Config{
            Username:         "YOUR_DEVICE_USERNAME", // check below for detailed instructions
            BridgeAddr:       "YOUR_BRIDGE_IP_ADDR", // check below where to find the ip address of the hue
            BridgeAddrScheme: "http",
	    }

        // create bridge
        bridge := NewBridge(config)

        // now do stuff with bridge
        // get ligts as a map of lights over ids
        lights, errLights := bridge.GetLights()
        if errLights {
            fmt.Println("Failed to get lights")
        }

        // Turn all lights on
        for id, _ := range lights {
            bridge.ToggleLight(id, true)
        }
    }
```
