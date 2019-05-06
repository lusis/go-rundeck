package responses

import (
	"encoding/json"
	"testing"

	"github.com/mitchellh/mapstructure"
	"github.com/stretchr/testify/require"
)

func TestSystemInfoResponse(t *testing.T) {
	obj := &SystemInfoResponse{}

	data, err := getAssetBytes(SystemInfoResponseTestFile)
	require.NoError(t, err)
	placeholder := make(map[string]interface{})
	err = json.Unmarshal(data, &placeholder)
	require.NoError(t, err)
	config := newMSDecoderConfig()
	config.Result = obj
	decoder, err := mapstructure.NewDecoder(config)
	require.NoError(t, err)

	err = decoder.Decode(placeholder)
	require.NoError(t, err)
	require.Implements(t, (*VersionedResponse)(nil), obj)
}
