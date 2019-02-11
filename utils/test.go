package utils

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/sergi/go-diff/diffmatchpatch"
)

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
