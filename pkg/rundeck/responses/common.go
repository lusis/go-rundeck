package responses

import (
	"errors"
	"strings"
	"time"
)

// JSONTime is for custom marshal/unmarshal of rundeck datetime values
type JSONTime struct {
	time.Time
}

const rdTime = "2006-01-02T15:04:05Z"

// UnmarshalJSON parses the rundeck datetime format
func (t *JSONTime) UnmarshalJSON(data []byte) error {
	if t == nil {
		return errors.New("JSONTime: UnmarshalText on nil pointer")
	}
	s := strings.Trim(string(data), "\"")
	if s == "null" {
		t.Time = time.Time{}
		return nil
	}
	tempTime, tErr := time.Parse(rdTime, s)
	if tErr != nil {
		return errors.New("JSONTime: " + tErr.Error())
	}
	t.Time = tempTime
	return nil
}
