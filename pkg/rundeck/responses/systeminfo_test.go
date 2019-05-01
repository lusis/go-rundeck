package responses

import (
	"encoding/json"
	"testing"

	"github.com/mitchellh/mapstructure"
	"github.com/stretchr/testify/assert"
)

func TestSystemInfoResponse(t *testing.T) {
	obj := &SystemInfoResponse{}

	data, dataErr := getAssetBytes(SystemInfoResponseTestFile)
	if dataErr != nil {
		t.Fatalf(dataErr.Error())
	}
	placeholder := make(map[string]interface{})
	_ = json.Unmarshal(data, &placeholder)
	config := newMSDecoderConfig()
	config.Result = obj
	decoder, newErr := mapstructure.NewDecoder(config)
	assert.NoError(t, newErr)
	dErr := decoder.Decode(placeholder)
	assert.NoError(t, dErr)
	assert.Implements(t, (*VersionedResponse)(nil), obj)
}
