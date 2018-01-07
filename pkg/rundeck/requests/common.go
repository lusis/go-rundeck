package requests

import (
	"fmt"
	"time"
)

// JSONTime is for custom marshal/unmarshal of rundeck datetime values
type JSONTime struct {
	time.Time
}

var rdFormat = "2006-01-02T15:04:05-0700"

// MarshalJSON converts a JSONTime to the format supported for rundeck requests
// if empty, we specify the current time
func (t *JSONTime) MarshalJSON() ([]byte, error) {
	var jsonTime string
	if !t.IsZero() {
		jsonTime = fmt.Sprintf("\"%s\"", t.Format(rdFormat))
		return []byte(jsonTime), nil
	}
	jsonTime = fmt.Sprintf("\"%s\"", time.Now().Format(rdFormat))
	return []byte(jsonTime), nil
}
