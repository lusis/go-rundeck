package rundeck

import (
	"bytes"
	"testing"

	"github.com/lusis/go-rundeck/pkg/rundeck/responses"

	"github.com/stretchr/testify/require"
)

func TestListSystemAclPolicies(t *testing.T) {
	jsonfile, err := responses.GetTestData(responses.ACLResponseTestFile)
	require.NoError(t, err)

	client, server, err := newTestRundeckClient(jsonfile, "application/json", 200)
	defer server.Close()
	require.NoError(t, err)

	s, err := client.ListSystemACLPolicies()
	require.NoError(t, err)
	require.Equal(t, "", s.Path)
	require.Equal(t, "directory", s.Type)
	require.NotEmpty(t, s.Href)
	require.Len(t, s.Resources, 1)
	require.Equal(t, "name.aclpolicy", s.Resources[0].Path)
	require.Equal(t, "file", s.Resources[0].Type)
	require.Equal(t, "name.aclpolicy", s.Resources[0].Name)
	require.NotEmpty(t, s.Href)

}

func TestListSystemAclPoliciesHTTPError(t *testing.T) {
	client, server, err := newTestRundeckClient([]byte(""), "application/json", 500)
	defer server.Close()
	require.NoError(t, err)

	s, err := client.ListSystemACLPolicies()
	require.Error(t, err)
	require.Nil(t, s)
}

func TestListSystemAclPoliciesJSONError(t *testing.T) {
	client, server, err := newTestRundeckClient([]byte(""), "application/json", 200)
	defer server.Close()
	require.NoError(t, err)

	s, err := client.ListSystemACLPolicies()
	require.Error(t, err)
	require.Nil(t, s)
}

func TestGetSystemAclPolicy(t *testing.T) {
	jsonfile, err := responses.GetTestData("foo.aclpolicy")
	require.NoError(t, err)

	client, server, err := newTestRundeckClient(jsonfile, "application/yaml", 200)
	defer server.Close()
	require.NoError(t, err)

	s, err := client.GetSystemACLPolicy("foo")
	require.NoError(t, err)
	require.NotNil(t, s)
}

func TestGetSystemAclPolicyHTTPError(t *testing.T) {
	jsonfile, err := responses.GetTestData("foo.aclpolicy")
	require.NoError(t, err)

	client, server, err := newTestRundeckClient(jsonfile, "application/yaml", 500)
	defer server.Close()
	require.NoError(t, err)

	s, err := client.GetSystemACLPolicy("foo")
	require.Error(t, err)
	require.Nil(t, s)
}

func TestCreateSystemACLPolicy(t *testing.T) {
	jsonfile, err := responses.GetTestData("foo.aclpolicy")
	require.NoError(t, err)

	client, server, err := newTestRundeckClient([]byte(""), "application/yaml", 201)
	defer server.Close()
	require.NoError(t, err)

	pErr := client.CreateSystemACLPolicy("foo", bytes.NewReader(jsonfile))
	require.NoError(t, pErr)
}

func TestCreateSystemACLPolicyConflict(t *testing.T) {
	jsonfile, err := responses.GetTestData("foo.aclpolicy")
	require.NoError(t, err)

	client, server, err := newTestRundeckClient([]byte(""), "application/yaml", 409)
	defer server.Close()
	require.NoError(t, err)

	pErr := client.CreateSystemACLPolicy("foo", bytes.NewReader(jsonfile))
	require.Error(t, pErr)
	require.EqualError(t, pErr, ErrResourceConflict.Error())
}

func TestCreateSystemACLPolicyValidationErrors(t *testing.T) {
	jsonfile, err := responses.GetTestData(responses.FailedACLValidationResponseTestFile)
	require.NoError(t, err)

	client, server, err := newTestRundeckClient(jsonfile, "application/json", 400)
	defer server.Close()
	require.NoError(t, err)

	pErr := client.CreateSystemACLPolicy("foo", bytes.NewReader([]byte("")))
	require.Error(t, pErr)
	require.NotEmpty(t, pErr.Error())
}

func TestUpdateSystemACLPolicy(t *testing.T) {
	jsonfile, err := responses.GetTestData("foo.aclpolicy")
	require.NoError(t, err)

	client, server, err := newTestRundeckClient([]byte(""), "application/yaml", 200)
	defer server.Close()
	require.NoError(t, err)

	pErr := client.UpdateSystemACLPolicy("foo", bytes.NewReader(jsonfile))
	require.NoError(t, pErr)
}

func TestUpdateSystemACLPolicyConflict(t *testing.T) {
	jsonfile, err := responses.GetTestData("foo.aclpolicy")
	require.NoError(t, err)

	client, server, err := newTestRundeckClient([]byte(""), "application/yaml", 409)
	defer server.Close()
	require.NoError(t, err)

	pErr := client.UpdateSystemACLPolicy("foo", bytes.NewReader(jsonfile))
	require.Error(t, pErr)
	require.EqualError(t, pErr, ErrResourceConflict.Error())
}

func TestUpdateSystemACLPolicyValidationErrors(t *testing.T) {
	jsonfile, err := responses.GetTestData(responses.FailedACLValidationResponseTestFile)
	require.NoError(t, err)

	client, server, err := newTestRundeckClient(jsonfile, "application/json", 400)
	defer server.Close()
	require.NoError(t, err)

	pErr := client.UpdateSystemACLPolicy("foo", bytes.NewReader([]byte("")))
	require.Error(t, pErr)
	require.NotEmpty(t, pErr.Error())
}

func TestDeleteSystemACLPolicy(t *testing.T) {
	client, server, err := newTestRundeckClient([]byte(""), "application/json", 204)
	defer server.Close()
	require.NoError(t, err)

	pErr := client.DeleteSystemACLPolicy("foo")
	require.NoError(t, pErr)
}

func TestDeleteSystemACLPolicyHTTPError(t *testing.T) {
	client, server, err := newTestRundeckClient([]byte(""), "application/json", 404)
	defer server.Close()
	require.NoError(t, err)
	pErr := client.DeleteSystemACLPolicy("foo")
	require.Error(t, pErr)
}

func TestListProjectAclPolicies(t *testing.T) {
	jsonfile, err := responses.GetTestData(responses.ACLResponseTestFile)
	require.NoError(t, err)

	client, server, err := newTestRundeckClient(jsonfile, "application/json", 200)
	defer server.Close()
	require.NoError(t, err)

	s, err := client.ListProjectACLPolicies("testproject")
	require.NoError(t, err)
	require.Equal(t, "", s.Path)
	require.Equal(t, "directory", s.Type)
	require.NotEmpty(t, s.Href)
	require.Len(t, s.Resources, 1)
	require.Equal(t, "name.aclpolicy", s.Resources[0].Path)
	require.Equal(t, "file", s.Resources[0].Type)
	require.Equal(t, "name.aclpolicy", s.Resources[0].Name)
	require.NotEmpty(t, s.Href)

}

func TestListProjectAclPoliciesHTTPError(t *testing.T) {
	client, server, err := newTestRundeckClient([]byte(""), "application/json", 500)
	defer server.Close()
	require.NoError(t, err)

	s, err := client.ListProjectACLPolicies("testproject")
	require.Error(t, err)
	require.Nil(t, s)
}

func TestListProjectAclPoliciesJSONError(t *testing.T) {
	client, server, err := newTestRundeckClient([]byte(""), "application/json", 200)
	defer server.Close()
	require.NoError(t, err)

	s, err := client.ListProjectACLPolicies("testproject")
	require.Error(t, err)
	require.Nil(t, s)
}

func TestGetProjectAclPolicy(t *testing.T) {
	jsonfile, err := responses.GetTestData("foo.aclpolicy")
	require.NoError(t, err)

	client, server, err := newTestRundeckClient(jsonfile, "application/yaml", 200)
	defer server.Close()
	require.NoError(t, err)

	s, err := client.GetProjectACLPolicy("foo", "bar")
	require.NoError(t, err)
	require.NotNil(t, s)
}

func TestGetProjectAclPolicyHTTPError(t *testing.T) {
	jsonfile, err := responses.GetTestData("foo.aclpolicy")
	require.NoError(t, err)

	client, server, err := newTestRundeckClient(jsonfile, "application/yaml", 500)
	defer server.Close()
	require.NoError(t, err)

	s, err := client.GetProjectACLPolicy("foo", "bar")
	require.Error(t, err)
	require.Nil(t, s)
}

func TestDeleteProjectACLPolicy(t *testing.T) {
	client, server, err := newTestRundeckClient([]byte(""), "application/json", 204)
	defer server.Close()
	require.NoError(t, err)
	pErr := client.DeleteProjectACLPolicy("foo", "bar")
	require.NoError(t, pErr)
}

func TestDeleteProjectACLPolicyHTTPError(t *testing.T) {
	client, server, err := newTestRundeckClient([]byte(""), "application/json", 404)
	defer server.Close()
	require.NoError(t, err)
	pErr := client.DeleteProjectACLPolicy("foo", "bar")
	require.Error(t, pErr)
}

func TestCreateProjectACLPolicy(t *testing.T) {
	jsonfile, err := responses.GetTestData("foo.aclpolicy")
	require.NoError(t, err)

	client, server, err := newTestRundeckClient([]byte(""), "application/yaml", 201)
	defer server.Close()
	require.NoError(t, err)

	pErr := client.CreateProjectACLPolicy("foo", "bar", bytes.NewReader(jsonfile))
	require.NoError(t, pErr)
}

func TestCreateProjectACLPolicyConflict(t *testing.T) {
	jsonfile, err := responses.GetTestData("foo.aclpolicy")
	require.NoError(t, err)

	client, server, err := newTestRundeckClient([]byte(""), "application/yaml", 409)
	defer server.Close()
	require.NoError(t, err)

	pErr := client.CreateProjectACLPolicy("foo", "bar", bytes.NewReader(jsonfile))
	require.Error(t, pErr)
	require.EqualError(t, pErr, ErrResourceConflict.Error())
}

func TestCreateProjectACLPolicyValidationErrors(t *testing.T) {
	jsonfile, err := responses.GetTestData(responses.FailedACLValidationResponseTestFile)
	require.NoError(t, err)

	client, server, err := newTestRundeckClient(jsonfile, "application/json", 400)
	defer server.Close()
	require.NoError(t, err)

	pErr := client.CreateProjectACLPolicy("foo", "bar", bytes.NewReader([]byte("")))
	require.Error(t, pErr)
	require.NotEmpty(t, pErr.Error())
}

func TestUpdateProjectACLPolicy(t *testing.T) {
	jsonfile, err := responses.GetTestData("foo.aclpolicy")
	require.NoError(t, err)

	client, server, err := newTestRundeckClient([]byte(""), "application/yaml", 200)
	defer server.Close()
	require.NoError(t, err)

	pErr := client.UpdateProjectACLPolicy("foo", "bar", bytes.NewReader(jsonfile))
	require.NoError(t, pErr)
}

func TestUpdateProjectACLPolicyConflict(t *testing.T) {
	jsonfile, err := responses.GetTestData("foo.aclpolicy")
	require.NoError(t, err)

	client, server, err := newTestRundeckClient([]byte(""), "application/yaml", 409)
	defer server.Close()
	require.NoError(t, err)

	pErr := client.UpdateProjectACLPolicy("foo", "bar", bytes.NewReader(jsonfile))
	require.Error(t, pErr)
	require.EqualError(t, pErr, ErrResourceConflict.Error())
}

func TestUpdateProjectACLPolicyValidationErrors(t *testing.T) {
	jsonfile, err := responses.GetTestData(responses.FailedACLValidationResponseTestFile)
	require.NoError(t, err)

	client, server, err := newTestRundeckClient(jsonfile, "application/json", 400)
	defer server.Close()
	require.NoError(t, err)

	pErr := client.UpdateProjectACLPolicy("foo", "bar", bytes.NewReader([]byte("")))
	require.Error(t, pErr)
	require.NotEmpty(t, pErr.Error())
}
