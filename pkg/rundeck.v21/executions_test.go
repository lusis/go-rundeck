package rundeck

import (
	"testing"

	"github.com/lusis/go-rundeck/pkg/rundeck.v21/responses"
	"github.com/lusis/go-rundeck/pkg/rundeck.v21/responses/testdata"

	"github.com/stretchr/testify/assert"
)

func TestListRunningExecutions(t *testing.T) {
	jsonfile, err := testdata.GetBytes(responses.ListRunningExecutionsResponseTestFile)
	if err != nil {
		t.Fatalf(err.Error())
	}

	client, server, cErr := newTestRundeckClient(jsonfile, "application/json", 200)
	defer server.Close()
	if cErr != nil {
		t.Fatalf(cErr.Error())
	}
	obj, cErr := client.ListRunningExecutions("testproject")
	assert.NoError(t, cErr)
	assert.NotNil(t, obj)
}

func TestListRunningExecutionsHTTPError(t *testing.T) {
	jsonfile, err := testdata.GetBytes(responses.ListRunningExecutionsResponseTestFile)
	if err != nil {
		t.Fatalf(err.Error())
	}

	client, server, _ := newTestRundeckClient(jsonfile, "application/json", 404)
	defer server.Close()
	obj, cErr := client.ListRunningExecutions("testproject")
	assert.Error(t, cErr)
	assert.Nil(t, obj)
}

func TestListRunningExecutionsJSONError(t *testing.T) {
	client, server, _ := newTestRundeckClient([]byte(""), "application/json", 200)
	defer server.Close()
	obj, cErr := client.ListRunningExecutions("testproject")
	assert.Error(t, cErr)
	assert.Nil(t, obj)
}

func TestListProjectExecutions(t *testing.T) {
	jsonfile, err := testdata.GetBytes(responses.ListRunningExecutionsResponseTestFile)
	if err != nil {
		t.Fatalf(err.Error())
	}

	client, server, cErr := newTestRundeckClient(jsonfile, "application/json", 200)
	defer server.Close()
	if cErr != nil {
		t.Fatalf(cErr.Error())
	}
	obj, cErr := client.ListProjectExecutions("testproject", nil)
	assert.NoError(t, cErr)
	assert.NotNil(t, obj)
}

func TestListProjectExecutionsHTTPError(t *testing.T) {
	jsonfile, err := testdata.GetBytes(responses.ListRunningExecutionsResponseTestFile)
	if err != nil {
		t.Fatalf(err.Error())
	}

	client, server, _ := newTestRundeckClient(jsonfile, "application/json", 404)
	defer server.Close()
	obj, cErr := client.ListProjectExecutions("testproject", nil)
	assert.Error(t, cErr)
	assert.Nil(t, obj)
}

func TestListProjectExecutionsJSONError(t *testing.T) {
	client, server, _ := newTestRundeckClient([]byte(""), "application/json", 200)
	defer server.Close()
	obj, cErr := client.ListProjectExecutions("testproject", nil)
	assert.Error(t, cErr)
	assert.Nil(t, obj)
}

func TestDeleteExecutions(t *testing.T) {
	jsonfile, err := testdata.GetBytes(responses.BulkDeleteExecutionsResponseTestFile)
	if err != nil {
		t.Fatalf(err.Error())
	}

	client, server, _ := newTestRundeckClient(jsonfile, "application/json", 200)
	defer server.Close()
	obj, cErr := client.DeleteExecutions(1, 2, 3)
	assert.NoError(t, cErr)
	assert.NotNil(t, obj)
}

func TestDeleteExecutionsHTTPError(t *testing.T) {
	jsonfile, err := testdata.GetBytes(responses.BulkDeleteExecutionsResponseTestFile)
	if err != nil {
		t.Fatalf(err.Error())
	}

	client, server, _ := newTestRundeckClient(jsonfile, "application/json", 500)
	defer server.Close()
	obj, cErr := client.DeleteExecutions(1, 2, 3)
	assert.Error(t, cErr)
	assert.Nil(t, obj)
}

func TestDeleteExecutionsJSONError(t *testing.T) {
	client, server, _ := newTestRundeckClient([]byte(""), "application/json", 200)
	defer server.Close()
	obj, cErr := client.DeleteExecutions(1, 2, 3)
	assert.Error(t, cErr)
	assert.Nil(t, obj)
}

func TestDeleteAllExecutionsForProject(t *testing.T) {}

func TestBulkEnableExecution(t *testing.T) {
	jsonfile, err := testdata.GetBytes(responses.BulkToggleResponseTestFile)
	if err != nil {
		t.Fatalf(err.Error())
	}

	client, server, _ := newTestRundeckClient(jsonfile, "application/json", 200)
	defer server.Close()
	obj, cErr := client.BulkEnableExecution("a", "b", "c")
	assert.NoError(t, cErr)
	assert.NotNil(t, obj)
}

func TestBulkDisableExecution(t *testing.T) {
	jsonfile, err := testdata.GetBytes(responses.BulkToggleResponseTestFile)
	if err != nil {
		t.Fatalf(err.Error())
	}

	client, server, _ := newTestRundeckClient(jsonfile, "application/json", 200)
	defer server.Close()
	obj, cErr := client.BulkDisableExecution("a", "b", "c")
	assert.NoError(t, cErr)
	assert.NotNil(t, obj)
}
