package responses

import (
	"encoding/json"
	"testing"

	"github.com/lusis/go-rundeck/pkg/rundeck/responses/testdata"
	"github.com/mitchellh/mapstructure"
	"github.com/stretchr/testify/assert"
)

func TestListKeysResourceResponse(t *testing.T) {
	obj := &ListKeysResourceResponse{}
	data, dataErr := testdata.GetBytes(ListKeysResourceResponseTestFile)
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

func TestListKeysResponse(t *testing.T) {
	obj := &ListKeysResponse{}
	data, dataErr := testdata.GetBytes(ListKeysResponseTestFile)
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
