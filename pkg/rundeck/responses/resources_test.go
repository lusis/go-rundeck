package responses

import (
	"encoding/json"
	"testing"

	"github.com/mitchellh/mapstructure"
	"github.com/stretchr/testify/require"
)

func TestResourceCollectionResponse(t *testing.T) {
	obj := &ResourceCollectionResponse{}
	data, err := getAssetBytes(ResourceCollectionResponseTestFile)
	require.NoError(t, err)
	placeholder := make(map[string]interface{})
	err = json.Unmarshal(data, &placeholder)
	require.NoError(t, err)
	config := newMSDecoderConfig()
	config.Result = obj
	// Because of the nature of the resource response, we're going to loosen the checks
	// and validate the uncaptured fields
	md := &mapstructure.Metadata{}
	config.Metadata = md
	config.ErrorUnused = false
	decoder, err := mapstructure.NewDecoder(config)
	require.NoError(t, err)
	err = decoder.Decode(placeholder)
	require.NoError(t, err)
	require.Implements(t, (*VersionedResponse)(nil), obj)
	// our test data has 10 nodes with 2 extra fields each
	require.Len(t, md.Unused, 20)
}

func TestResourceResponse(t *testing.T) {
	obj := &ResourceResponse{}
	data, err := getAssetBytes(ResourceResponseTestFile)
	require.NoError(t, err)
	placeholder := make(map[string]interface{})
	err = json.Unmarshal(data, &placeholder)
	require.NoError(t, err)
	config := newMSDecoderConfig()
	config.Result = obj
	// Because of the nature of the resource response, we're going to loosen the checks
	// and validate the uncaptured fields
	md := &mapstructure.Metadata{}
	config.Metadata = md
	config.ErrorUnused = false
	decoder, err := mapstructure.NewDecoder(config)
	require.NoError(t, err)
	err = decoder.Decode(placeholder)
	require.NoError(t, err)
	node := (*obj)["node-0-fake"]
	require.Implements(t, (*VersionedResponse)(nil), obj)
	require.Equal(t, "node-0-fake", node.NodeName)
	require.Equal(t, "nodehost-fake", node.HostName)
	require.Equal(t, "stub", node.NodeExectutor)
	require.Equal(t, "stub", node.FileCopier)
	require.Equal(t, "nodeuser-fake", node.UserName)
	require.Contains(t, "stub", node.Tags)
	// our test data has 10 nodes with 2 extra fields each
	require.Len(t, md.Unused, 2)
}
