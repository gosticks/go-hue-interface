package hue

import (
	"fmt"
	"strings"
	"time"
)

type Time struct {
	time.Time
}

const TimeFormat = "2006-01-02T15:04:05"

// UnmarshalJSON unmarshals mite time from json
func (t *Time) UnmarshalJSON(b []byte) (err error) {
	s := strings.Trim(string(b), "\"")
	if s == "null" {
		t.Time = time.Time{}
		return
	}
	t.Time, err = time.Parse(TimeFormat, s)
	return
}

// MarshalJSON marshals mite time to json
func (t *Time) MarshalJSON() ([]byte, error) {
	if t.Time.UnixNano() == (time.Time{}).UnixNano() {
		return []byte("null"), nil
	}
	return []byte(fmt.Sprintf("\"%s\"", t.Time.Format(TimeFormat))), nil
}
