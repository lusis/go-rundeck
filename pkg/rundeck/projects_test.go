package rundeck

import (
	"testing"

	"github.com/lusis/go-rundeck/pkg/rundeck/responses"
	"github.com/lusis/go-rundeck/pkg/rundeck/responses/testdata"

	"github.com/stretchr/testify/assert"
)

func TestGetProject(t *testing.T) {
	jsonfile, err := testdata.GetBytes(responses.ProjectInfoResponseTestFile)
	if err != nil {
		t.Fatalf(err.Error())
	}

	client, server, cErr := newTestRundeckClient(jsonfile, "application/json", 200)
	defer server.Close()
	if cErr != nil {
		t.Fatalf(cErr.Error())
	}
	obj, getErr := client.GetProjectInfo("testproject")
	assert.NoError(t, getErr)
	assert.NotNil(t, obj)
	assert.NotEmpty(t, obj.Properties)
	assert.Len(t, obj.Properties, 32)
	assert.Equal(t, "[API Href]", obj.URL)
	assert.Equal(t, "testproject", obj.Name)
	assert.Equal(t, "test project", obj.Description)
}

func TestGetProjectHTTPError(t *testing.T) {
	jsonfile, err := testdata.GetBytes(responses.ProjectInfoResponseTestFile)
	if err != nil {
		t.Fatalf(err.Error())
	}

	client, server, cErr := newTestRundeckClient(jsonfile, "application/json", 404)
	defer server.Close()
	if cErr != nil {
		t.Fatalf(cErr.Error())
	}
	obj, err := client.GetProjectInfo("testproject")
	assert.Error(t, err)
	assert.EqualError(t, ErrMissingResource, err.Error())
	assert.Nil(t, obj)
}

func TestGetProjectDecodeError(t *testing.T) {
	client, server, cErr := newTestRundeckClient([]byte(""), "application/json", 200)
	defer server.Close()
	if cErr != nil {
		t.Fatalf(cErr.Error())
	}
	obj, oErr := client.GetProjectInfo("testproject")
	assert.Error(t, oErr)
	assert.Nil(t, obj)
	assert.EqualError(t, errDecoding, oErr.Error())
}

func TestListProjects(t *testing.T) {
	jsonfile, err := testdata.GetBytes(responses.ListProjectsResponseTestFile)
	if err != nil {
		t.Fatalf(err.Error())
	}

	client, server, cErr := newTestRundeckClient(jsonfile, "application/json", 200)
	defer server.Close()
	if cErr != nil {
		t.Fatalf(cErr.Error())
	}
	obj, cErr := client.ListProjects()
	assert.NoError(t, cErr)
	assert.NotNil(t, obj)
	assert.Len(t, *obj, 1)
	project := (*obj)[0]
	assert.Equal(t, "[API Href]", project.URL)
	assert.Equal(t, "testproject", project.Name)
	assert.Equal(t, "test project", project.Description)
}

func TestListProjectsDecodeError(t *testing.T) {
	client, server, cErr := newTestRundeckClient([]byte(""), "application/json", 200)
	defer server.Close()
	if cErr != nil {
		t.Fatalf(cErr.Error())
	}
	obj, oErr := client.ListProjects()
	assert.Error(t, oErr)
	assert.Nil(t, obj)
	assert.EqualError(t, errDecoding, oErr.Error())
}

func TestListProjectsNotFound(t *testing.T) {
	client, server, cErr := newTestRundeckClient([]byte(""), "application/json", 404)
	defer server.Close()
	if cErr != nil {
		t.Fatalf(cErr.Error())
	}
	obj, oErr := client.ListProjects()
	assert.Error(t, oErr)
	assert.Nil(t, obj)
	assert.EqualError(t, ErrMissingResource, oErr.Error())
}

func TestDeleteProject(t *testing.T) {
	client, server, cErr := newTestRundeckClient([]byte(""), "application/json", 204)
	defer server.Close()
	if cErr != nil {
		t.Fatalf(cErr.Error())
	}
	err := client.DeleteProject("abc123")
	assert.NoError(t, err)
}

func TestDeleteProjectNotFound(t *testing.T) {
	client, server, cErr := newTestRundeckClient([]byte(""), "application/json", 404)
	defer server.Close()
	if cErr != nil {
		t.Fatalf(cErr.Error())
	}
	err := client.DeleteProject("abc123")
	assert.EqualError(t, ErrMissingResource, err.Error())
}

func TestCreateProject(t *testing.T) {
	jsonfile, err := testdata.GetBytes(responses.ProjectInfoResponseTestFile)
	if err != nil {
		t.Fatalf(err.Error())
	}

	client, server, _ := newTestRundeckClient(jsonfile, "application/json", 201)
	defer server.Close()
	obj, getErr := client.CreateProject("testproject", nil)
	assert.NoError(t, getErr)
	assert.NotNil(t, obj)
	assert.NotEmpty(t, obj.Properties)
	assert.Len(t, obj.Properties, 32)
	assert.Equal(t, "[API Href]", obj.URL)
	assert.Equal(t, "testproject", obj.Name)
	assert.Equal(t, "test project", obj.Description)
}
