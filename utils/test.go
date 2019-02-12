package utils

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/sergi/go-diff/diffmatchpatch"
)

// CompareJSONDecode decodes a JSON string into a provided interface and then back to a json string. If the output json string does not match the compact version of the input string a error is returned
func CompareJSONDecode(j string, v interface{}) error {
	bytes, errCompact := CompactJSON(j)
	if errCompact != nil {
		return errCompact
	}
	// Unmarshal the data
	json.Unmarshal(bytes, v)

	// Marshal it again
	outputData, _ := json.Marshal(v)

	old := string(bytes)
	new := string(outputData)

	return DiffStrings(old, new)
}

// CompareStructToJSON compares Marshaled interface to a JSON string and returns an error if they do not match
func CompareStructToJSON(v interface{}, j string) error {
	// Remove all extra spaces from json string
	bytes, errCompact := CompactJSON(j)
	if errCompact != nil {
		return errCompact
	}

	newData, errMarsh := json.Marshal(v)
	if errMarsh != nil {
		return errMarsh
	}

	old := string(bytes)
	new := string(newData)
	return DiffStrings(old, new)
}

// DiffStrings prints a nice compare between strings. Error is returned if strings are not equal
func DiffStrings(s1, s2 string) error {
	if s1 != s2 {
		fmt.Println("String do not match!")
		fmt.Println("OLD: \n " + s1)
		fmt.Println("----------------------------- ")
		fmt.Println("NEW: \n " + s2)
		fmt.Println("----------------------------- ")
		fmt.Println("DIFF:")

		dmp := diffmatchpatch.New()

		diffs := dmp.DiffMain(s1, s2, true)

		fmt.Println(dmp.DiffPrettyText(diffs))
		return errors.New("output is not the same as input")
	}
	return nil
}

// CompactJSON removes all whitespaces between json keys and values that is not required
func CompactJSON(j string) ([]byte, error) {

	buffer := new(bytes.Buffer)
	// Remove all spaces
	errCompact := json.Compact(buffer, []byte(j))
	if errCompact != nil {
		return nil, errCompact
	}
	bytes := buffer.Bytes()

	return bytes, nil
}

// BridgeFromEnv creates a bridge from the config data stored in env. This is primarily designed for automated tests
