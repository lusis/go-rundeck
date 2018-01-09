package rundeck

import (
	"bytes"
	"testing"

	"github.com/lusis/go-rundeck/pkg/rundeck/responses"
	"github.com/lusis/go-rundeck/pkg/rundeck/responses/testdata"

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

	s, err := client.ListSystemACLPolicies()
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

	s, err := client.ListSystemACLPolicies()
	assert.Error(t, err)
	assert.Nil(t, s)
}

func TestListSystemAclPoliciesJSONError(t *testing.T) {
	client, server, cErr := newTestRundeckClient([]byte(""), "application/json", 200)
	defer server.Close()
	if cErr != nil {
		t.Fatalf(cErr.Error())
	}

	s, err := client.ListSystemACLPolicies()
	assert.Error(t, err)
	assert.Nil(t, s)
}

func TestGetSystemAclPolicy(t *testing.T) {
	jsonfile, err := testdata.GetBytes("foo.aclpolicy")
	if err != nil {
		t.Fatalf(err.Error())
	}

	client, server, cErr := newTestRundeckClient(jsonfile, "application/yaml", 200)
	defer server.Close()
	if cErr != nil {
		t.Fatalf(cErr.Error())
	}

	s, err := client.GetSystemACLPolicy("foo")
	assert.NoError(t, err)
	assert.NotNil(t, s)
}

func TestGetSystemAclPolicyHTTPError(t *testing.T) {
	jsonfile, err := testdata.GetBytes("foo.aclpolicy")
	if err != nil {
		t.Fatalf(err.Error())
	}

	client, server, cErr := newTestRundeckClient(jsonfile, "application/yaml", 500)
	defer server.Close()
	if cErr != nil {
		t.Fatalf(cErr.Error())
	}

	s, err := client.GetSystemACLPolicy("foo")
	assert.Error(t, err)
	assert.Nil(t, s)
}

func TestCreateSystemACLPolicy(t *testing.T) {
	jsonfile, err := testdata.GetBytes("foo.aclpolicy")
	if err != nil {
		t.Fatalf(err.Error())
	}

	client, server, cErr := newTestRundeckClient([]byte(""), "application/yaml", 201)
	defer server.Close()
	if cErr != nil {
		t.Fatalf(cErr.Error())
	}

	pErr := client.CreateSystemACLPolicy("foo", bytes.NewReader(jsonfile))
	assert.NoError(t, pErr)
}

func TestCreateSystemACLPolicyConflict(t *testing.T) {
	jsonfile, err := testdata.GetBytes("foo.aclpolicy")
	if err != nil {
		t.Fatalf(err.Error())
	}

	client, server, cErr := newTestRundeckClient([]byte(""), "application/yaml", 409)
	defer server.Close()
	if cErr != nil {
		t.Fatalf(cErr.Error())
	}

	pErr := client.CreateSystemACLPolicy("foo", bytes.NewReader(jsonfile))
	assert.Error(t, pErr)
	assert.EqualError(t, pErr, ErrResourceConflict.Error())
}

func TestCreateSystemACLPolicyValidationErrors(t *testing.T) {
	jsonfile, err := testdata.GetBytes(responses.FailedACLValidationResponseTestFile)
	if err != nil {
		t.Fatalf(err.Error())
	}

	client, server, cErr := newTestRundeckClient(jsonfile, "application/json", 400)
	defer server.Close()
	if cErr != nil {
		t.Fatalf(cErr.Error())
	}

	pErr := client.CreateSystemACLPolicy("foo", bytes.NewReader([]byte("")))
	assert.Error(t, pErr)
	assert.NotEmpty(t, pErr.Error())
}

func TestUpdateSystemACLPolicy(t *testing.T) {
	jsonfile, err := testdata.GetBytes("foo.aclpolicy")
	if err != nil {
		t.Fatalf(err.Error())
	}

	client, server, cErr := newTestRundeckClient([]byte(""), "application/yaml", 200)
	defer server.Close()
	if cErr != nil {
		t.Fatalf(cErr.Error())
	}

	pErr := client.UpdateSystemACLPolicy("foo", bytes.NewReader(jsonfile))
	assert.NoError(t, pErr)
}

func TestUpdateSystemACLPolicyConflict(t *testing.T) {
	jsonfile, err := testdata.GetBytes("foo.aclpolicy")
	if err != nil {
		t.Fatalf(err.Error())
	}

	client, server, cErr := newTestRundeckClient([]byte(""), "application/yaml", 409)
	defer server.Close()
	if cErr != nil {
		t.Fatalf(cErr.Error())
	}

	pErr := client.UpdateSystemACLPolicy("foo", bytes.NewReader(jsonfile))
	assert.Error(t, pErr)
	assert.EqualError(t, pErr, ErrResourceConflict.Error())
}

func TestUpdateSystemACLPolicyValidationErrors(t *testing.T) {
	jsonfile, err := testdata.GetBytes(responses.FailedACLValidationResponseTestFile)
	if err != nil {
		t.Fatalf(err.Error())
	}

	client, server, cErr := newTestRundeckClient(jsonfile, "application/json", 400)
	defer server.Close()
	if cErr != nil {
		t.Fatalf(cErr.Error())
	}

	pErr := client.UpdateSystemACLPolicy("foo", bytes.NewReader([]byte("")))
	assert.Error(t, pErr)
	assert.NotEmpty(t, pErr.Error())
}

func TestDeleteSystemACLPolicy(t *testing.T) {
	client, server, _ := newTestRundeckClient([]byte(""), "application/json", 204)
	defer server.Close()
	pErr := client.DeleteSystemACLPolicy("foo")
	assert.NoError(t, pErr)
}

func TestDeleteSystemACLPolicyHTTPError(t *testing.T) {
	client, server, _ := newTestRundeckClient([]byte(""), "application/json", 404)
	defer server.Close()
	pErr := client.DeleteSystemACLPolicy("foo")
	assert.Error(t, pErr)
}

func TestListProjectAclPolicies(t *testing.T) {
	jsonfile, err := testdata.GetBytes(responses.ACLResponseTestFile)
	if err != nil {
		t.Fatalf(err.Error())
	}

	client, server, cErr := newTestRundeckClient(jsonfile, "application/json", 200)
	defer server.Close()
	if cErr != nil {
		t.Fatalf(cErr.Error())
	}

	s, err := client.ListProjectACLPolicies("testproject")
	assert.Equal(t, "", s.Path)
	assert.Equal(t, "directory", s.Type)
	assert.NotEmpty(t, s.Href)
	assert.Len(t, s.Resources, 1)
	assert.Equal(t, "name.aclpolicy", s.Resources[0].Path)
	assert.Equal(t, "file", s.Resources[0].Type)
	assert.Equal(t, "name.aclpolicy", s.Resources[0].Name)
	assert.NotEmpty(t, s.Href)

}

func TestListProjectAclPoliciesHTTPError(t *testing.T) {
	client, server, cErr := newTestRundeckClient([]byte(""), "application/json", 500)
	defer server.Close()
	if cErr != nil {
		t.Fatalf(cErr.Error())
	}

	s, err := client.ListProjectACLPolicies("testproject")
	assert.Error(t, err)
	assert.Nil(t, s)
}

func TestListProjectAclPoliciesJSONError(t *testing.T) {
	client, server, cErr := newTestRundeckClient([]byte(""), "application/json", 200)
	defer server.Close()
	if cErr != nil {
		t.Fatalf(cErr.Error())
	}

	s, err := client.ListProjectACLPolicies("testproject")
	assert.Error(t, err)
	assert.Nil(t, s)
}

func TestGetProjectAclPolicy(t *testing.T) {
	jsonfile, err := testdata.GetBytes("foo.aclpolicy")
	if err != nil {
		t.Fatalf(err.Error())
	}

	client, server, cErr := newTestRundeckClient(jsonfile, "application/yaml", 200)
	defer server.Close()
	if cErr != nil {
		t.Fatalf(cErr.Error())
	}

	s, err := client.GetProjectACLPolicy("foo", "bar")
	assert.NoError(t, err)
	assert.NotNil(t, s)
}

func TestGetProjectAclPolicyHTTPError(t *testing.T) {
	jsonfile, err := testdata.GetBytes("foo.aclpolicy")
	if err != nil {
		t.Fatalf(err.Error())
	}

	client, server, cErr := newTestRundeckClient(jsonfile, "application/yaml", 500)
	defer server.Close()
	if cErr != nil {
		t.Fatalf(cErr.Error())
	}

	s, err := client.GetProjectACLPolicy("foo", "bar")
	assert.Error(t, err)
	assert.Nil(t, s)
}

func TestDeleteProjectACLPolicy(t *testing.T) {
	client, server, _ := newTestRundeckClient([]byte(""), "application/json", 204)
	defer server.Close()
	pErr := client.DeleteProjectACLPolicy("foo", "bar")
	assert.NoError(t, pErr)
}

func TestDeleteProjectACLPolicyHTTPError(t *testing.T) {
	client, server, _ := newTestRundeckClient([]byte(""), "application/json", 404)
	defer server.Close()
	pErr := client.DeleteProjectACLPolicy("foo", "bar")
	assert.Error(t, pErr)
}

func TestCreateProjectACLPolicy(t *testing.T) {
	jsonfile, err := testdata.GetBytes("foo.aclpolicy")
	if err != nil {
		t.Fatalf(err.Error())
	}

	client, server, cErr := newTestRundeckClient([]byte(""), "application/yaml", 201)
	defer server.Close()
	if cErr != nil {
		t.Fatalf(cErr.Error())
	}

	pErr := client.CreateProjectACLPolicy("foo", "bar", bytes.NewReader(jsonfile))
	assert.NoError(t, pErr)
}

func TestCreateProjectACLPolicyConflict(t *testing.T) {
	jsonfile, err := testdata.GetBytes("foo.aclpolicy")
	if err != nil {
		t.Fatalf(err.Error())
	}

	client, server, cErr := newTestRundeckClient([]byte(""), "application/yaml", 409)
	defer server.Close()
	if cErr != nil {
		t.Fatalf(cErr.Error())
	}

	pErr := client.CreateProjectACLPolicy("foo", "bar", bytes.NewReader(jsonfile))
	assert.Error(t, pErr)
	assert.EqualError(t, pErr, ErrResourceConflict.Error())
}

func TestCreateProjectACLPolicyValidationErrors(t *testing.T) {
	jsonfile, err := testdata.GetBytes(responses.FailedACLValidationResponseTestFile)
	if err != nil {
		t.Fatalf(err.Error())
	}

	client, server, cErr := newTestRundeckClient(jsonfile, "application/json", 400)
	defer server.Close()
	if cErr != nil {
		t.Fatalf(cErr.Error())
	}

	pErr := client.CreateProjectACLPolicy("foo", "bar", bytes.NewReader([]byte("")))
	assert.Error(t, pErr)
	assert.NotEmpty(t, pErr.Error())
}

func TestUpdateProjectACLPolicy(t *testing.T) {
	jsonfile, err := testdata.GetBytes("foo.aclpolicy")
	if err != nil {
		t.Fatalf(err.Error())
	}

	client, server, cErr := newTestRundeckClient([]byte(""), "application/yaml", 200)
	defer server.Close()
	if cErr != nil {
		t.Fatalf(cErr.Error())
	}

	pErr := client.UpdateProjectACLPolicy("foo", "bar", bytes.NewReader(jsonfile))
	assert.NoError(t, pErr)
}

func TestUpdateProjectACLPolicyConflict(t *testing.T) {
	jsonfile, err := testdata.GetBytes("foo.aclpolicy")
	if err != nil {
		t.Fatalf(err.Error())
	}

	client, server, cErr := newTestRundeckClient([]byte(""), "application/yaml", 409)
	defer server.Close()
	if cErr != nil {
		t.Fatalf(cErr.Error())
	}

	pErr := client.UpdateProjectACLPolicy("foo", "bar", bytes.NewReader(jsonfile))
	assert.Error(t, pErr)
	assert.EqualError(t, pErr, ErrResourceConflict.Error())
}

func TestUpdateProjectACLPolicyValidationErrors(t *testing.T) {
	jsonfile, err := testdata.GetBytes(responses.FailedACLValidationResponseTestFile)
	if err != nil {
		t.Fatalf(err.Error())
	}

	client, server, cErr := newTestRundeckClient(jsonfile, "application/json", 400)
	defer server.Close()
	if cErr != nil {
		t.Fatalf(cErr.Error())
	}

	pErr := client.UpdateProjectACLPolicy("foo", "bar", bytes.NewReader([]byte("")))
	assert.Error(t, pErr)
	assert.NotEmpty(t, pErr.Error())
}
