package rundeck

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListSystemAclPolicies(t *testing.T) {
	jsonfile, err := os.Open("assets/test/aclpolicies.json")
	if err != nil {
		t.Fatalf(err.Error())
	}
	defer func() { _ = jsonfile.Close() }()

	client, server, err := newTestRundeckClient(jsonfile, "application/json", 200)
	defer server.Close()
	if err != nil {
		t.FailNow()
	}
	s, err := client.GetACLPolicies()
	assert.Equal(t, s.Path, "")
	assert.Equal(t, s.Type, "directory")
	assert.NotEmpty(t, s.Href)
	assert.Len(t, s.Resources, 1)
	assert.Equal(t, s.Resources[0].Path, "foo.aclpolicy")
	assert.Equal(t, s.Resources[0].Type, "file")
	assert.NotEmpty(t, s.Href)
}
