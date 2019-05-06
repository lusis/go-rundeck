package rundeck

import (
	"testing"

	"github.com/lusis/go-rundeck/pkg/rundeck/responses"

	"github.com/stretchr/testify/require"
)

func TestListRunningExecutions(t *testing.T) {
	jsonfile, err := responses.GetTestData(responses.ListRunningExecutionsResponseTestFile)
	if err != nil {
		t.Fatalf(err.Error())
	}

	client, server, cErr := newTestRundeckClient(jsonfile, "application/json", 200)
	defer server.Close()
	if cErr != nil {
		t.Fatalf(cErr.Error())
	}
	obj, cErr := client.ListRunningExecutions("testproject")
	require.NoError(t, cErr)
	require.NotNil(t, obj)
}

func TestListRunningExecutionsHTTPError(t *testing.T) {
	jsonfile, err := responses.GetTestData(responses.ListRunningExecutionsResponseTestFile)
	if err != nil {
		t.Fatalf(err.Error())
	}

	client, server, _ := newTestRundeckClient(jsonfile, "application/json", 404)
	defer server.Close()
	obj, cErr := client.ListRunningExecutions("testproject")
	require.Error(t, cErr)
	require.Nil(t, obj)
}

func TestListRunningExecutionsJSONError(t *testing.T) {
	client, server, _ := newTestRundeckClient([]byte(""), "application/json", 200)
	defer server.Close()
	obj, cErr := client.ListRunningExecutions("testproject")
	require.Error(t, cErr)
	require.Nil(t, obj)
}

func TestListProjectExecutions(t *testing.T) {
	jsonfile, err := responses.GetTestData(responses.ListRunningExecutionsResponseTestFile)
	if err != nil {
		t.Fatalf(err.Error())
	}

	client, server, cErr := newTestRundeckClient(jsonfile, "application/json", 200)
	defer server.Close()
	if cErr != nil {
		t.Fatalf(cErr.Error())
	}
	obj, cErr := client.ListProjectExecutions("testproject", nil)
	require.NoError(t, cErr)
	require.NotNil(t, obj)
}

func TestListProjectExecutionsHTTPError(t *testing.T) {
	jsonfile, err := responses.GetTestData(responses.ListRunningExecutionsResponseTestFile)
	if err != nil {
		t.Fatalf(err.Error())
	}

	client, server, _ := newTestRundeckClient(jsonfile, "application/json", 404)
	defer server.Close()
	obj, cErr := client.ListProjectExecutions("testproject", nil)
	require.Error(t, cErr)
	require.Nil(t, obj)
}

func TestListProjectExecutionsJSONError(t *testing.T) {
	client, server, _ := newTestRundeckClient([]byte(""), "application/json", 200)
	defer server.Close()
	obj, cErr := client.ListProjectExecutions("testproject", nil)
	require.Error(t, cErr)
	require.Nil(t, obj)
}

func TestDeleteExecutions(t *testing.T) {
	jsonfile, err := responses.GetTestData(responses.BulkDeleteExecutionsResponseTestFile)
	if err != nil {
		t.Fatalf(err.Error())
	}

	client, server, _ := newTestRundeckClient(jsonfile, "application/json", 200)
	defer server.Close()
	obj, cErr := client.BulkDeleteExecutions(1, 2, 3)
	require.NoError(t, cErr)
	require.NotNil(t, obj)
}

func TestDeleteExecutionsHTTPError(t *testing.T) {
	jsonfile, err := responses.GetTestData(responses.BulkDeleteExecutionsResponseTestFile)
	if err != nil {
		t.Fatalf(err.Error())
	}

	client, server, _ := newTestRundeckClient(jsonfile, "application/json", 500)
	defer server.Close()
	obj, cErr := client.BulkDeleteExecutions(1, 2, 3)
	require.Error(t, cErr)
	require.Nil(t, obj)
}

func TestDeleteExecutionsJSONError(t *testing.T) {
	client, server, _ := newTestRundeckClient([]byte(""), "application/json", 200)
	defer server.Close()
	obj, cErr := client.BulkDeleteExecutions(1, 2, 3)
	require.Error(t, cErr)
	require.Nil(t, obj)
}

func TestDeleteAllExecutionsForProject(t *testing.T) {}

func TestBulkEnableExecution(t *testing.T) {
	jsonfile, err := responses.GetTestData(responses.BulkToggleResponseTestFile)
	if err != nil {
		t.Fatalf(err.Error())
	}

	client, server, _ := newTestRundeckClient(jsonfile, "application/json", 200)
	defer server.Close()
	obj, cErr := client.BulkEnableExecution("a", "b", "c")
	require.NoError(t, cErr)
	require.NotNil(t, obj)
}

func TestBulkEnableExecutionHTTPError(t *testing.T) {
	jsonfile, err := responses.GetTestData(responses.BulkToggleResponseTestFile)
	if err != nil {
		t.Fatalf(err.Error())
	}

	client, server, _ := newTestRundeckClient(jsonfile, "application/json", 500)
	defer server.Close()
	obj, cErr := client.BulkEnableExecution("a", "b", "c")
	require.Error(t, cErr)
	require.Nil(t, obj)
}

func TestBulkEnableExecutionJSONError(t *testing.T) {
	client, server, _ := newTestRundeckClient([]byte(""), "application/json", 200)
	defer server.Close()
	obj, cErr := client.BulkEnableExecution("a", "b", "c")
	require.Error(t, cErr)
	require.Nil(t, obj)
}

func TestBulkDisableExecution(t *testing.T) {
	jsonfile, err := responses.GetTestData(responses.BulkToggleResponseTestFile)
	if err != nil {
		t.Fatalf(err.Error())
	}

	client, server, _ := newTestRundeckClient(jsonfile, "application/json", 200)
	defer server.Close()
	obj, cErr := client.BulkDisableExecution("a", "b", "c")
	require.NoError(t, cErr)
	require.NotNil(t, obj)
}

func TestBulkDisableExecutionHTTPError(t *testing.T) {
	jsonfile, err := responses.GetTestData(responses.BulkToggleResponseTestFile)
	if err != nil {
		t.Fatalf(err.Error())
	}

	client, server, _ := newTestRundeckClient(jsonfile, "application/json", 500)
	defer server.Close()
	obj, cErr := client.BulkDisableExecution("a", "b", "c")
	require.Error(t, cErr)
	require.Nil(t, obj)
}

func TestBulkDisableExecutionJSONError(t *testing.T) {
	client, server, _ := newTestRundeckClient([]byte(""), "application/json", 200)
	defer server.Close()
	obj, cErr := client.BulkDisableExecution("a", "b", "c")
	require.Error(t, cErr)
	require.Nil(t, obj)
}
