package rundeck

import (
	"testing"

	"github.com/lusis/go-rundeck/pkg/rundeck/responses"

	"github.com/stretchr/testify/require"
)

func TestGetLogStorage(t *testing.T) {
	jsonfile, err := responses.GetTestData(responses.LogStorageResponseTestFile)
	if err != nil {
		t.Fatalf(err.Error())
	}

	client, server, _ := newTestRundeckClient(jsonfile, "application/json", 200)
	defer server.Close()
	obj, cErr := client.GetLogStorageInfo()
	require.NoError(t, cErr)
	require.NotNil(t, obj)
}

func TestGetLogStorageJSONError(t *testing.T) {
	client, server, _ := newTestRundeckClient([]byte(""), "application/json", 200)
	defer server.Close()
	obj, cErr := client.GetLogStorageInfo()
	require.Error(t, cErr)
	require.Nil(t, obj)
}

func TestGetLogStorageHTTPError(t *testing.T) {
	client, server, _ := newTestRundeckClient([]byte(""), "application/json", 201)
	defer server.Close()
	obj, cErr := client.GetLogStorageInfo()
	require.Error(t, cErr)
	require.Nil(t, obj)
}

func TestGetIncompleteLogStorage(t *testing.T) {
	jsonfile, err := responses.GetTestData(responses.IncompleteLogStorageResponseTestFile)
	if err != nil {
		t.Fatalf(err.Error())
	}

	client, server, _ := newTestRundeckClient(jsonfile, "application/json", 200)
	defer server.Close()
	obj, cErr := client.GetIncompleteLogStorage()
	require.NoError(t, cErr)
	require.NotNil(t, obj)
}

func TestIncompleteGetLogStorageJSONError(t *testing.T) {
	client, server, _ := newTestRundeckClient([]byte(""), "application/json", 200)
	defer server.Close()
	obj, cErr := client.GetIncompleteLogStorage()
	require.Error(t, cErr)
	require.Nil(t, obj)
}

func TestIncompleteGetLogStorageHTTPError(t *testing.T) {
	client, server, _ := newTestRundeckClient([]byte(""), "application/json", 201)
	defer server.Close()
	obj, cErr := client.GetIncompleteLogStorage()
	require.Error(t, cErr)
	require.Nil(t, obj)
}
