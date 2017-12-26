package rundeck

import (
	"testing"

	"github.com/lusis/go-rundeck/pkg/rundeck.v21/responses"
	"github.com/lusis/go-rundeck/pkg/rundeck.v21/responses/testdata"

	"github.com/stretchr/testify/assert"
)

func TestListSystemAclPolicies(t *testing.T) {
	jsonfile, err := testdata.GetBytes(responses.ACLResponseTestFile)
	if err != nil {
		t.Fatalf(err.Error())
	}

	client, server, cErr := newTestRundeckClient(jsonfile, "application/json", 200)
	defer server.Close()
	if cErr != nil {
		t.Fatalf(cErr.Error())
	}

	s, err := client.GetACLPolicies()
	assert.Equal(t, "", s.Path)
	assert.Equal(t, "directory", s.Type)
	assert.NotEmpty(t, s.Href)
	assert.Len(t, s.Resources, 1)
	assert.Equal(t, "name.aclpolicy", s.Resources[0].Path)
	assert.Equal(t, "file", s.Resources[0].Type)
	assert.Equal(t, "name.aclpolicy", s.Resources[0].Name)
	assert.NotEmpty(t, s.Href)

}
