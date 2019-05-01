package responses

import (
	"encoding/json"
	"testing"

	"github.com/mitchellh/mapstructure"
	"github.com/stretchr/testify/assert"
)

func TestACLResponse(t *testing.T) {
	obj := &ACLResponse{}
	data, err := getAssetBytes(ACLResponseTestFile)
	if err != nil {
		t.Error(err.Error())
		t.FailNow()
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

func TestFailedACLValidationResponse(t *testing.T) {
	obj := &FailedACLValidationResponse{}
	data, err := getAssetBytes(FailedACLValidationResponseTestFile)
	if err != nil {
		t.Error(err.Error())
		t.FailNow()
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
