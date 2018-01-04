package requests

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type testTS struct {
	DateTime JSONTime `json:"datetime"`
}

func TestMarshalJSONTime(t *testing.T) {
	tstamp := time.Date(2016, time.November, 23, 12, 20, 55, 0, time.Local)
	s := &testTS{
		DateTime: JSONTime{tstamp},
	}
	str := fmt.Sprintf(`{"datetime":"2016-11-23T12:20:55-0500"}`)
	res, resErr := json.Marshal(s)
	assert.NoError(t, resErr)
	assert.Equal(t, str, string(res))
}

func TestMarshalJSONTimeNil(t *testing.T) {
	s := &testTS{}
	res, resErr := json.Marshal(s)
	assert.NoError(t, resErr)
	assert.NotEmpty(t, string(res))
}
