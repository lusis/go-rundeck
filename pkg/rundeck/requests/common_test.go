package requests

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

type testTS struct {
	DateTime JSONTime `json:"datetime"`
}

func TestMarshalJSONTime(t *testing.T) {
	tstamp := time.Date(2016, time.November, 23, 12, 20, 55, 0, time.UTC)
	s := &testTS{
		DateTime: JSONTime{tstamp},
	}
	str := fmt.Sprintf(`{"datetime":"2016-11-23T12:20:55+0000"}`)
	res, resErr := json.Marshal(s)
	require.NoError(t, resErr)

	require.Equal(t, str, string(res))
}

func TestMarshalJSONTimeNil(t *testing.T) {
	s := &testTS{}
	res, resErr := json.Marshal(s)
	require.NoError(t, resErr)
	require.NotEmpty(t, string(res))
}
