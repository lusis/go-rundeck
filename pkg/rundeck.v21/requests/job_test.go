package requests

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestRunJobRequestMarshal(t *testing.T) {
	timeVal := "2018-01-03T11:44:23-0500"
	curTime, _ := time.Parse(rdFormat, timeVal)
	j := &RunJobRequest{
		ArgString: "-opt 1",
		LogLevel:  "DEBUG",
		AsUser:    "user1",
		Filter:    ".*",
		RunAtTime: &JSONTime{curTime},
		Options: map[string]string{
			"opt2": "val2",
		},
	}
	res, err := json.Marshal(j)
	assert.NoError(t, err)
	expected := fmt.Sprintf(`{"argString":"-opt 1","loglevel":"DEBUG","asUser":"user1","filter":".*","runAtTime":"%s","options":{"opt2":"val2"}}`, timeVal)
	assert.Equal(t, expected, string(res))
}
