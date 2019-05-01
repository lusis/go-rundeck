package rundeck

import (
	"testing"

	"github.com/lusis/go-rundeck/pkg/rundeck/responses"

	"github.com/stretchr/testify/assert"
)

func TestGetResources(t *testing.T) {
	jsonfile, err := responses.GetTestData(responses.ResourceCollectionResponseTestFile)
	if err != nil {
		t.Fatalf(err.Error())
	}

	client, server, cErr := newTestRundeckClient(jsonfile, "application/json", 200)
	defer server.Close()
	if cErr != nil {
		t.Fatalf(cErr.Error())
	}
	obj, cErr := client.ListResourcesForProject("testproject")
	assert.NoError(t, cErr)
	assert.NotNil(t, obj)
	assert.Len(t, *obj, 11)
}

func TestGetResourcesInvalidJSON(t *testing.T) {
	client, server, cErr := newTestRundeckClient([]byte(""), "application/json", 200)
	defer server.Close()
	if cErr != nil {
		t.Fatalf(cErr.Error())
	}
	obj, cErr := client.ListResourcesForProject("testproject")
	assert.Error(t, cErr)
	assert.Nil(t, obj)
}

func TestGetResourcesInvalidStatusCode(t *testing.T) {
	client, server, cErr := newTestRundeckClient([]byte(""), "application/json", 500)
	defer server.Close()
	if cErr != nil {
		t.Fatalf(cErr.Error())
	}
	obj, cErr := client.ListResourcesForProject("testproject")
	assert.Error(t, cErr)
	assert.Nil(t, obj)
}

func TestGetResource(t *testing.T) {
	jsonfile, err := responses.GetTestData(responses.ResourceResponseTestFile)
	if err != nil {
		t.Fatalf(err.Error())
	}

	client, server, cErr := newTestRundeckClient(jsonfile, "application/json", 200)
	defer server.Close()
	if cErr != nil {
		t.Fatalf(cErr.Error())
	}
	obj, cErr := client.GetResourceInfo("testproject", "node-0-fake")
	assert.NoError(t, cErr)
	assert.NotNil(t, obj)
	assert.Equal(t, "node-0-fake", obj.NodeName)
}

func TestGetResourceInvalidJSON(t *testing.T) {
	client, server, cErr := newTestRundeckClient([]byte(""), "application/json", 200)
	defer server.Close()
	if cErr != nil {
		t.Fatalf(cErr.Error())
	}
	obj, cErr := client.GetResourceInfo("testproject", "node-0-fake")
	assert.Error(t, cErr)
	assert.Nil(t, obj)
}

func TestGetResourceInvalidStatusCode(t *testing.T) {
	client, server, cErr := newTestRundeckClient([]byte(""), "application/json", 500)
	defer server.Close()
	if cErr != nil {
		t.Fatalf(cErr.Error())
	}
	obj, cErr := client.GetResourceInfo("testproject", "node-0-fake")
	assert.Error(t, cErr)
	assert.Nil(t, obj)
}
