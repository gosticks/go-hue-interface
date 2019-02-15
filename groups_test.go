package hue

import (
	"fmt"
	"os"
	"testing"

	"github.com/gosticks/go-hue-interface/utils"
)

func TestParseGroups(t *testing.T) {
	v := make(map[string]*Group)
	err := utils.CompareJSONDecode(groupsTestData, &v)
	if err != nil {
		t.Fail()
	}
}

func TestBridgeGroupsParse(t *testing.T) {
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
	resp, respErr := b.getRawResponse(groupsEndpoint)
	if respErr != nil {
		t.Log(respErr)
		t.Error()
	}

	groups, errGroups := b.GetAllGroups()
	if errGroups != nil {
		t.Log(errGroups)
		t.Error()
	}

	errComp := utils.CompareStructToJSON(groups, string(resp))
	if errComp != nil {
		t.Fail()
	}
}

const groupsTestData = `
{
    "1": {
        "name": "Group 1",
        "lights": [
            "1",
            "2"
        ],
        "type": "LightGroup",
        "action": {
            "on": true,
            "bri": 254,
            "hue": 10000,
            "sat": 254,
            "effect": "none",
            "xy": [
                0.5,
                0.5
            ],
            "ct": 250,
            "alert": "select",
            "colormode": "ct"
        }
    },
    "2": {
        "name": "Group 2",
        "lights": [
            "3",
            "4",
            "5"
        ],
        "type": "LightGroup",
        "action": {
            "on": true,
            "bri": 153,
            "hue": 4345,
            "sat": 254,
            "effect": "none",
            "xy": [
                0.5,
                0.5
            ],
            "ct": 250,
            "alert": "select",
            "colormode": "ct"
        }
    }
}`
