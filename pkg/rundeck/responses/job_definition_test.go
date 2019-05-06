package responses

import (
	"testing"

	"github.com/stretchr/testify/require"
	yaml "gopkg.in/yaml.v2"
)

func TestJobYAMLResponse(t *testing.T) {
	obj := &JobYAMLResponse{}
	data, err := getAssetBytes(JobYAMLResponseTestFile)
	require.NoError(t, err)

	err = yaml.UnmarshalStrict(data, &obj)
	require.NoError(t, err)
	require.Implements(t, (*VersionedResponse)(nil), obj)
}
