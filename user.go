package hue

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	username string
	// TODO:
	//key string
}

// CreateNewUser will create a new user. This should be called only of there's none in the yaml config.
func CreateUser(addr string) (u *User, err error) {
	return CreateUserExtended(addr, "go-hue-interface", "Philips hue")
}

// TODO: remove these comments
func CreateUserExtended(addr, application, deviceName string) (u *User, err error) {
	uri := "http://" + addr + "/api"

	var reqBody = []byte(fmt.Sprintf(`{"devicetype": "%s#%s"}`, application, deviceName))
	req, err := http.NewRequest("POST", uri, bytes.NewBuffer(reqBody))

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return
	}
	defer res.Body.Close()

	// Unmarshal data
	err = json.NewDecoder(res.Body).Decode(u)
	if err != nil {
		return
	}

	return
}
