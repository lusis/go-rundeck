package responses

import (
	"testing"

	"github.com/lusis/go-rundeck/pkg/rundeck/responses/testdata"
	"github.com/stretchr/testify/assert"
)

func TestResourceCollectionResponse(t *testing.T) {
	obj := ResourceCollectionResponse{}
	data, dataErr := testdata.GetBytes(ResourceCollectionResponseTestFile)
	if dataErr != nil {
		t.Error(dataErr.Error())
		t.FailNow()
	}
	err := obj.FromBytes(data)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	assert.Implements(t, (*VersionedResponse)(nil), obj)
	assert.Len(t, obj, 11)
}

func TestResourceResponse(t *testing.T) {
	obj := ResourceResponse{}
	data, dataErr := testdata.GetBytes(ResourceResponseTestFile)
	if dataErr != nil {
		t.Error(dataErr.Error())
		t.FailNow()
	}
	err := obj.FromBytes(data)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	node := obj["node-0-fake"]
	assert.Implements(t, (*VersionedResponse)(nil), obj)
	assert.Equal(t, "node-0-fake", node.NodeName)
	assert.Equal(t, "nodehost-fake", node.HostName)
	assert.Equal(t, "stub", node.NodeExectutor)
	assert.Equal(t, "stub", node.FileCopier)
	assert.Equal(t, "nodeuser-fake", node.UserName)
	assert.Contains(t, "stub", node.Tags)
}
