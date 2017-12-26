package rundeck

import (
	"testing"

	"github.com/lusis/go-rundeck/pkg/rundeck.v21/responses"
	"github.com/lusis/go-rundeck/pkg/rundeck.v21/responses/testdata"

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
	obj, cErr := client.GetProject("testproject")
	assert.NoError(t, cErr)
	assert.NotNil(t, obj)
	assert.NotEmpty(t, obj.Properties)
	assert.Len(t, obj.Properties, 32)
	assert.Equal(t, "[API Href]", obj.URL)
	assert.Equal(t, "testproject", obj.Name)
	assert.Equal(t, "test project", obj.Description)
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
