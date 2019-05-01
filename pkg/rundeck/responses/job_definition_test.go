package responses

import (
	"testing"

	"github.com/stretchr/testify/assert"
	yaml "gopkg.in/yaml.v2"
)

func TestJobYAMLResponse(t *testing.T) {
	obj := &JobYAMLResponse{}
	data, dataErr := getAssetBytes(JobYAMLResponseTestFile)
	if dataErr != nil {
		t.Fatalf(dataErr.Error())
	}

	yErr := yaml.UnmarshalStrict(data, &obj)
	assert.NoError(t, yErr)
	assert.Implements(t, (*VersionedResponse)(nil), obj)
}
