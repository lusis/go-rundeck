package responses

import (
	"encoding/json"
	"testing"

	"github.com/mitchellh/mapstructure"
	"github.com/stretchr/testify/require"
)

func TestLogStorageReponses(t *testing.T) {
	testCases := []struct{
		name string
		placeholder interface{}
		obj interface{}
		testfile string
	}{
		{
			name: "LogStorageResponse",
			placeholder: make(map[string]interface{}),
			obj: &LogStorageResponse{},
			testfile: LogStorageResponseTestFile,
		},
		{
			name: "IncompleteLogStorageExecutionResponse",
			placeholder: make(map[string]interface{}),
			obj: &IncompleteLogStorageResponse{},
			testfile: IncompleteLogStorageResponseTestFile,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			data, err := getAssetBytes(tc.testfile)
			require.NoError(t, err)
			err = json.Unmarshal(data, &tc.placeholder)
			require.NoError(t, err)
			config := newMSDecoderConfig()
			config.Result = tc.obj
			decoder, err := mapstructure.NewDecoder(config)
			require.NoError(t, err)
			err = decoder.Decode(tc.placeholder)
			require.NoError(t, err)
			require.Implements(t, (*VersionedResponse)(nil), tc.obj)
		})
	}
}
