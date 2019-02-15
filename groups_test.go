package hue

import (
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
