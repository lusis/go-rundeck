package rundeck

import (
	"bytes"
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

func TestListSystemAclPoliciesHTTPError(t *testing.T) {
	client, server, cErr := newTestRundeckClient([]byte(""), "application/json", 500)
	defer server.Close()
	if cErr != nil {
		t.Fatalf(cErr.Error())
	}

	s, err := client.GetACLPolicies()
	assert.Error(t, err)
	assert.Nil(t, s)
}

func TestListSystemAclPoliciesJSONError(t *testing.T) {
	client, server, cErr := newTestRundeckClient([]byte(""), "application/json", 200)
	defer server.Close()
	if cErr != nil {
		t.Fatalf(cErr.Error())
	}

	s, err := client.GetACLPolicies()
	assert.Error(t, err)
	assert.Nil(t, s)
}

func TestGetAclPolicy(t *testing.T) {
	jsonfile, err := testdata.GetBytes("foo.aclpolicy")
	if err != nil {
		t.Fatalf(err.Error())
	}

	client, server, cErr := newTestRundeckClient(jsonfile, "application/yaml", 200)
	defer server.Close()
	if cErr != nil {
		t.Fatalf(cErr.Error())
	}

	s, err := client.GetACLPolicy("foo")
	assert.NoError(t, err)
	assert.NotNil(t, s)
}

func TestGetAclPolicyHTTPError(t *testing.T) {
	jsonfile, err := testdata.GetBytes("foo.aclpolicy")
	if err != nil {
		t.Fatalf(err.Error())
	}

	client, server, cErr := newTestRundeckClient(jsonfile, "application/yaml", 500)
	defer server.Close()
	if cErr != nil {
		t.Fatalf(cErr.Error())
	}

	s, err := client.GetACLPolicy("foo")
	assert.Error(t, err)
	assert.Nil(t, s)
}

func TestCreateACLPolicy(t *testing.T) {
	jsonfile, err := testdata.GetBytes("foo.aclpolicy")
	if err != nil {
		t.Fatalf(err.Error())
	}

	client, server, cErr := newTestRundeckClient([]byte(""), "application/yaml", 201)
	defer server.Close()
	if cErr != nil {
		t.Fatalf(cErr.Error())
	}

	pErr := client.CreateACLPolicy("foo", bytes.NewReader(jsonfile))
	assert.NoError(t, pErr)
}

func TestCreateACLPolicyConflict(t *testing.T) {
	jsonfile, err := testdata.GetBytes("foo.aclpolicy")
	if err != nil {
		t.Fatalf(err.Error())
	}

	client, server, cErr := newTestRundeckClient([]byte(""), "application/yaml", 409)
	defer server.Close()
	if cErr != nil {
		t.Fatalf(cErr.Error())
	}

	pErr := client.CreateACLPolicy("foo", bytes.NewReader(jsonfile))
	assert.Error(t, pErr)
	assert.EqualError(t, pErr, ErrResourceConflict.Error())
}

func TestCreateACLPolicyValidationErrors(t *testing.T) {
	jsonfile, err := testdata.GetBytes(responses.FailedACLValidationResponseTestFile)
	if err != nil {
		t.Fatalf(err.Error())
	}

	client, server, cErr := newTestRundeckClient(jsonfile, "application/json", 400)
	defer server.Close()
	if cErr != nil {
		t.Fatalf(cErr.Error())
	}

	pErr := client.CreateACLPolicy("foo", bytes.NewReader([]byte("")))
	assert.Error(t, pErr)
	assert.NotEmpty(t, pErr.Error())
}

func TestUpdateACLPolicy(t *testing.T) {
	jsonfile, err := testdata.GetBytes("foo.aclpolicy")
	if err != nil {
		t.Fatalf(err.Error())
	}

	client, server, cErr := newTestRundeckClient([]byte(""), "application/yaml", 201)
	defer server.Close()
	if cErr != nil {
		t.Fatalf(cErr.Error())
	}

	pErr := client.UpdateACLPolicy("foo", bytes.NewReader(jsonfile))
	assert.NoError(t, pErr)
}

func TestUpdateACLPolicyConflict(t *testing.T) {
	jsonfile, err := testdata.GetBytes("foo.aclpolicy")
	if err != nil {
		t.Fatalf(err.Error())
	}

	client, server, cErr := newTestRundeckClient([]byte(""), "application/yaml", 409)
	defer server.Close()
	if cErr != nil {
		t.Fatalf(cErr.Error())
	}

	pErr := client.UpdateACLPolicy("foo", bytes.NewReader(jsonfile))
	assert.Error(t, pErr)
	assert.EqualError(t, pErr, ErrResourceConflict.Error())
}

func TestUpdateACLPolicyValidationErrors(t *testing.T) {
	jsonfile, err := testdata.GetBytes(responses.FailedACLValidationResponseTestFile)
	if err != nil {
		t.Fatalf(err.Error())
	}

	client, server, cErr := newTestRundeckClient(jsonfile, "application/json", 400)
	defer server.Close()
	if cErr != nil {
		t.Fatalf(cErr.Error())
	}

	pErr := client.UpdateACLPolicy("foo", bytes.NewReader([]byte("")))
	assert.Error(t, pErr)
	assert.NotEmpty(t, pErr.Error())
}
