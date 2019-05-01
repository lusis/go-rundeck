package responses

import (
	"encoding/json"
	"testing"

	"github.com/mitchellh/mapstructure"
	"github.com/stretchr/testify/assert"
)

func TestResourceCollectionResponse(t *testing.T) {
	obj := &ResourceCollectionResponse{}
	data, dataErr := getAssetBytes(ResourceCollectionResponseTestFile)
	if dataErr != nil {
		t.Fatalf(dataErr.Error())
	}
	placeholder := make(map[string]interface{})
	_ = json.Unmarshal(data, &placeholder)
	config := newMSDecoderConfig()
	config.Result = obj
	// Because of the nature of the resource response, we're going to loosen the checks
	// and validate the uncaptured fields
	md := &mapstructure.Metadata{}
	config.Metadata = md
	config.ErrorUnused = false
	decoder, newErr := mapstructure.NewDecoder(config)
	assert.NoError(t, newErr)
	dErr := decoder.Decode(placeholder)
	assert.Implements(t, (*VersionedResponse)(nil), obj)
	assert.NoError(t, dErr)
	// our test data has 10 nodes with 2 extra fields each
	assert.Len(t, md.Unused, 20)
}

func TestResourceResponse(t *testing.T) {
	obj := &ResourceResponse{}
	data, dataErr := getAssetBytes(ResourceResponseTestFile)
	if dataErr != nil {
		t.Fatalf(dataErr.Error())
	}
	placeholder := make(map[string]interface{})
	_ = json.Unmarshal(data, &placeholder)
	config := newMSDecoderConfig()
	config.Result = obj
	// Because of the nature of the resource response, we're going to loosen the checks
	// and validate the uncaptured fields
	md := &mapstructure.Metadata{}
	config.Metadata = md
	config.ErrorUnused = false
	decoder, newErr := mapstructure.NewDecoder(config)
	assert.NoError(t, newErr)
	dErr := decoder.Decode(placeholder)
	assert.NoError(t, dErr)
	node := (*obj)["node-0-fake"]
	assert.Implements(t, (*VersionedResponse)(nil), obj)
	assert.Equal(t, "node-0-fake", node.NodeName)
	assert.Equal(t, "nodehost-fake", node.HostName)
	assert.Equal(t, "stub", node.NodeExectutor)
	assert.Equal(t, "stub", node.FileCopier)
	assert.Equal(t, "nodeuser-fake", node.UserName)
	assert.Contains(t, "stub", node.Tags)
	// our test data has 10 nodes with 2 extra fields each
	assert.Len(t, md.Unused, 2)
}
