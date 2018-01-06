package rundeck

import (
	"testing"

	"github.com/lusis/go-rundeck/pkg/rundeck.v21/responses"
	"github.com/lusis/go-rundeck/pkg/rundeck.v21/responses/testdata"

	"github.com/stretchr/testify/assert"
)

func TestGetLogStorage(t *testing.T) {
	jsonfile, err := testdata.GetBytes(responses.LogStorageResponseTestFile)
	if err != nil {
		t.Fatalf(err.Error())
	}

	client, server, _ := newTestRundeckClient(jsonfile, "application/json", 200)
	defer server.Close()
	obj, cErr := client.GetLogStorage()
	assert.NoError(t, cErr)
	assert.NotNil(t, obj)
}

func TestGetLogStorageJSONError(t *testing.T) {
	client, server, _ := newTestRundeckClient([]byte(""), "application/json", 200)
	defer server.Close()
	obj, cErr := client.GetLogStorage()
	assert.Error(t, cErr)
	assert.Nil(t, obj)
}

func TestGetLogStorageHTTPError(t *testing.T) {
	client, server, _ := newTestRundeckClient([]byte(""), "application/json", 201)
	defer server.Close()
	obj, cErr := client.GetLogStorage()
	assert.Error(t, cErr)
	assert.Nil(t, obj)
}

func TestGetIncompleteLogStorage(t *testing.T) {
	jsonfile, err := testdata.GetBytes(responses.IncompleteLogStorageResponseTestFile)
	if err != nil {
		t.Fatalf(err.Error())
	}

	client, server, _ := newTestRundeckClient(jsonfile, "application/json", 200)
	defer server.Close()
	obj, cErr := client.GetIncompleteLogStorage()
	assert.NoError(t, cErr)
	assert.NotNil(t, obj)
}

func TestIncompleteGetLogStorageJSONError(t *testing.T) {
	client, server, _ := newTestRundeckClient([]byte(""), "application/json", 200)
	defer server.Close()
	obj, cErr := client.GetIncompleteLogStorage()
	assert.Error(t, cErr)
	assert.Nil(t, obj)
}

func TestIncompleteGetLogStorageHTTPError(t *testing.T) {
	client, server, _ := newTestRundeckClient([]byte(""), "application/json", 201)
	defer server.Close()
	obj, cErr := client.GetIncompleteLogStorage()
	assert.Error(t, cErr)
	assert.Nil(t, obj)
}
