package hue

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/gosticks/go-hue-interface/utils"
)

func TestParseLights(t *testing.T) {
	v := make(map[string]*Light)
	err := utils.CompareJSONDecode(LightsTestData, &v)
	if err != nil {
		t.Fail()
	}
}

func TestLightToggle(t *testing.T) {
	userID, okUser := os.LookupEnv("HUE_BRIDGE_USER")
	addr, okAddr := os.LookupEnv("HUE_BRIDGE_ADDR")
	if !okAddr || !okUser {
		fmt.Println("HUE_BRIDGE_USER and HUE_BRIDGE_ADDR must be set in env for this test to work")
		t.Fail()
	}

	conf := &Config{
		Username:         userID,
		BridgeAddr:       addr,
		BridgeAddrScheme: "http",
	}
	b := NewBridge(conf)

	ticker := time.NewTicker(200 * time.Millisecond)
	defer ticker.Stop()
	go func() {
		var state = true
		for range ticker.C {
			b.ToggleLight("2", state)
			state = !state
		}
	}()

	select {}
}
func TestBridgeLightsParse(t *testing.T) {
	userID, okUser := os.LookupEnv("HUE_BRIDGE_USER")
	addr, okAddr := os.LookupEnv("HUE_BRIDGE_ADDR")
	if !okAddr || !okUser {
		fmt.Println("HUE_BRIDGE_USER and HUE_BRIDGE_ADDR must be set in env for this test to work")
		t.Fail()
	}

	conf := &Config{
		Username:         userID,
		BridgeAddr:       addr,
		BridgeAddrScheme: "http",
	}
	b := NewBridge(conf)
	resp, respErr := b.getRawResponse(LightsEndpoint)
	if respErr != nil {
		t.Log(respErr)
		t.Error()
	}

	lights, errLights := b.GetLights()
	if errLights != nil {
		t.Log(errLights)
		t.Error()
	}

	errComp := utils.CompareStructToJSON(lights, string(resp))
	if errComp != nil {
		t.Fail()
	}
}

const LightsTestData = `{
	"1": {
			"state": {
				"on": false,
				"bri": 1,
				"hue": 33761,
				"sat": 254,
				"effect": "none",
				"xy": [
					0.3171,
					0.3366
				],
				"ct": 159,
				"alert": "none",
				"colormode": "xy",
				"mode": "homeautomation",
				"reachable": true
			},
			"swupdate": {
				"state": "noupdates",
				"lastinstall": "2018-01-02T19:24:20"
			},
			"type": "Extended color light",
			"name": "Hue color lamp 7",
			"modelid": "LCT007",
			"manufacturername": "Philips",
			"productname": "Hue color lamp",
			"capabilities": {
				"certified": true,
				"control": {
					"mindimlevel": 5000,
					"maxlumen": 600,
					"colorgamuttype": "B",
					"colorgamut": [
						[
							0.675,
							0.322
						],
						[
							0.409,
							0.518
						],
						[
							0.167,
							0.04
						]
					],
					"ct": {
						"min": 153,
						"max": 500
					}
				},
				"streaming": {
					"renderer": true,
					"proxy": false
				}
			},
			"config": {
				"archetype": "sultanbulb",
				"function": "mixed",
				"direction": "omnidirectional"
			},
			"uniqueid": "00:17:88:01:00:bd:c7:b9-0b",
			"swversion": "5.105.0.21169"
		}
	}`
