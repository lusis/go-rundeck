package rundeck

import (
	"testing"

	"github.com/lusis/go-rundeck/pkg/rundeck/responses"
	"github.com/lusis/go-rundeck/pkg/rundeck/responses/testdata"

	"github.com/stretchr/testify/assert"
)

func TestGetHistory(t *testing.T) {
	jsonfile, err := testdata.GetBytes(responses.HistoryResponseTestFile)
	if err != nil {
		t.Fatalf(err.Error())
	}

	client, server, cErr := newTestRundeckClient(jsonfile, "application/json", 200)
	defer server.Close()
	if cErr != nil {
		t.Fatalf(cErr.Error())
	}
	obj, cErr := client.ListHistory("testproject", nil)
	assert.NoError(t, cErr)
	assert.NotNil(t, obj)
}

func TestGetHistoryNotFound(t *testing.T) {
	jsonfile, err := testdata.GetBytes(responses.HistoryResponseTestFile)
	if err != nil {
		t.Fatalf(err.Error())
	}

	client, server, cErr := newTestRundeckClient(jsonfile, "application/json", 404)
	defer server.Close()
	if cErr != nil {
		t.Fatalf(cErr.Error())
	}
	obj, oErr := client.ListHistory("testproject", nil)
	assert.Error(t, oErr)
	assert.Nil(t, obj)
}

func TestGetHistoryDecodeError(t *testing.T) {
	client, server, cErr := newTestRundeckClient([]byte(""), "application/json", 200)
	defer server.Close()
	if cErr != nil {
		t.Fatalf(cErr.Error())
	}
	obj, oErr := client.ListHistory("testproject", nil)
	assert.Error(t, oErr)
	assert.Nil(t, obj)
}

func TestGetHistoryOptions(t *testing.T) {
	jsonfile, err := testdata.GetBytes(responses.HistoryResponseTestFile)
	if err != nil {
		t.Fatalf(err.Error())
	}

	client, server, cErr := newTestRundeckClient(jsonfile, "application/json", 200)
	defer server.Close()
	if cErr != nil {
		t.Fatalf(cErr.Error())
	}
	obj, cErr := client.ListHistory("testproject", map[string]string{"foo": "bar"})

	assert.NoError(t, cErr)
	assert.NotNil(t, obj)
}
