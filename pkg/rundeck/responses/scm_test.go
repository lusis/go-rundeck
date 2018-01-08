package responses

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListSCMPluginsResponse(t *testing.T) {
	obj := ListSCMPluginsResponse{}
	assert.Implements(t, (*VersionedResponse)(nil), obj)
}

func TestSCMPluginResponse(t *testing.T) {
	obj := SCMPluginResponse{}
	assert.Implements(t, (*VersionedResponse)(nil), obj)
}

func TestGetSCMPluginInputFieldsResponse(t *testing.T) {
	obj := GetSCMPluginInputFieldsResponse{}
	assert.Implements(t, (*VersionedResponse)(nil), obj)
}
