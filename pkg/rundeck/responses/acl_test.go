package responses

import (
	"testing"

	"github.com/lusis/go-rundeck/pkg/rundeck/responses/testdata"
	"github.com/stretchr/testify/assert"
)

func TestACLResponse(t *testing.T) {
	obj := &ACLResponse{}
	data, err := testdata.GetBytes(ACLResponseTestFile)
	if err != nil {
		t.Error(err.Error())
		t.FailNow()
	}
	bErr := obj.FromBytes(data)
	if bErr != nil {
		t.Error(bErr.Error())
		t.FailNow()
	}

	assert.Equal(t, "", obj.Path)
	assert.Equal(t, "directory", obj.Type)
	assert.Equal(t, "[API Href]", obj.Href)
	assert.Len(t, obj.Resources, 1)

	resource := obj.Resources[0]
	assert.Equal(t, "name.aclpolicy", resource.Name)
	assert.Equal(t, "file", resource.Type)
	assert.Equal(t, "name.aclpolicy", resource.Path)
	assert.Equal(t, "[API Href]", resource.Href)
}

func TestFailedACLValidationResponse(t *testing.T) {
	obj := &FailedACLValidationResponse{}
	data, err := testdata.GetBytes(FailedACLValidationResponseTestFile)
	if err != nil {
		t.Error(err.Error())
		t.FailNow()
	}
	bErr := obj.FromBytes(data)
	if bErr != nil {
		t.Error(bErr.Error())
		t.FailNow()
	}
	assert.False(t, obj.Valid)
	assert.Len(t, obj.Policies, 2)
	first := obj.Policies[0]
	second := obj.Policies[1]
	assert.Equal(t, "file1.aclpolicy[1]", first.Policy)
	assert.Len(t, first.Errors, 2)
	assert.Equal(t, "file1.aclpolicy[2]", second.Policy)
	assert.Len(t, second.Errors, 2)
}
