package hue

import (
	"bytes"
	"encoding/json"
	"fmt"
	"testing"
)

func TestParseLights(t *testing.T) {

	target := make(map[string]*Light)
	buffer := new(bytes.Buffer)
	// Remove all spaces
	json.Compact(buffer, []byte(LightsTestData))
	bytes := buffer.Bytes()
	// Unmarshal the data
	json.Unmarshal(bytes, &target)

	// Marshal it again
	outputData, _ := json.Marshal(target)

	old := string(bytes)
	new := string(outputData)

	if old != new {
		fmt.Println("String do not match!")
		fmt.Println("OLD: \n " + old)
		fmt.Println("----------------------------- ")
		fmt.Println("NEW: \n " + new)
		t.Fail()
	}
	t.Log("Completed")
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
