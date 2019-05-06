package responses

import (
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/require"
)

type testTS struct {
	DateTime *JSONTime `json:"datetime"`
}

type testDuration struct {
	Duration *JSONDuration `json:"duration"`
}

func TestUnmarshal(t *testing.T) {
	str := `{"datetime":"2015-05-13T16:58:59Z"}`
	obj := &testTS{}
	err := json.Unmarshal([]byte(str), &obj)
	require.NoError(t, err)
	require.Equal(t, 2015, obj.DateTime.Year())
}

func TestUnmarshalNil(t *testing.T) {
	str := `{"datetime":null}`
	obj := &testTS{}
	err := json.Unmarshal([]byte(str), &obj)
	require.NoError(t, err)
}

func getAssetBytes(fileName string) ([]byte, error) {
	data, err := assets.Open(fileName)
	if err != nil {
		return nil, err
	}
	return ioutil.ReadAll(data)
}

func TestUnmarshalJSONDuration(t *testing.T) {
	str := `{"duration": "1h"}`
	obj := &testDuration{}
	err := json.Unmarshal([]byte(str), &obj)
	require.NoError(t, err)
	require.Equal(t, float64(3600), obj.Duration.Seconds())
}
