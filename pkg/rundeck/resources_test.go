package rundeck

import (
	"testing"

	"github.com/lusis/go-rundeck/pkg/rundeck/responses"

	"github.com/stretchr/testify/require"
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
	require.NoError(t, cErr)
	require.NotNil(t, obj)
	require.Len(t, *obj, 11)
}

func TestGetResourcesInvalidJSON(t *testing.T) {
	client, server, cErr := newTestRundeckClient([]byte(""), "application/json", 200)
	defer server.Close()
	if cErr != nil {
		t.Fatalf(cErr.Error())
	}
	obj, cErr := client.ListResourcesForProject("testproject")
	require.Error(t, cErr)
	require.Nil(t, obj)
}

func TestGetResourcesInvalidStatusCode(t *testing.T) {
	client, server, cErr := newTestRundeckClient([]byte(""), "application/json", 500)
	defer server.Close()
	if cErr != nil {
		t.Fatalf(cErr.Error())
	}
	obj, cErr := client.ListResourcesForProject("testproject")
	require.Error(t, cErr)
	require.Nil(t, obj)
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
	require.NoError(t, cErr)
	require.NotNil(t, obj)
	require.Equal(t, "node-0-fake", obj.NodeName)
}

func TestGetResourceInvalidJSON(t *testing.T) {
	client, server, cErr := newTestRundeckClient([]byte(""), "application/json", 200)
	defer server.Close()
	if cErr != nil {
		t.Fatalf(cErr.Error())
	}
	obj, cErr := client.GetResourceInfo("testproject", "node-0-fake")
	require.Error(t, cErr)
	require.Nil(t, obj)
}

func TestGetResourceInvalidStatusCode(t *testing.T) {
	client, server, cErr := newTestRundeckClient([]byte(""), "application/json", 500)
	defer server.Close()
	if cErr != nil {
		t.Fatalf(cErr.Error())
	}
	obj, cErr := client.GetResourceInfo("testproject", "node-0-fake")
	require.Error(t, cErr)
	require.Nil(t, obj)
}
