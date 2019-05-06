package rundeck

import (
	"testing"

	"github.com/lusis/go-rundeck/pkg/rundeck/responses"

	"github.com/stretchr/testify/require"
)

func TestGetHistory(t *testing.T) {
	jsonfile, err := responses.GetTestData(responses.HistoryResponseTestFile)
	if err != nil {
		t.Fatalf(err.Error())
	}

	client, server, cErr := newTestRundeckClient(jsonfile, "application/json", 200)
	defer server.Close()
	if cErr != nil {
		t.Fatalf(cErr.Error())
	}
	obj, cErr := client.ListHistory("testproject", nil)
	require.NoError(t, cErr)
	require.NotNil(t, obj)
}

func TestGetHistoryNotFound(t *testing.T) {
	jsonfile, err := responses.GetTestData(responses.HistoryResponseTestFile)
	if err != nil {
		t.Fatalf(err.Error())
	}

	client, server, cErr := newTestRundeckClient(jsonfile, "application/json", 404)
	defer server.Close()
	if cErr != nil {
		t.Fatalf(cErr.Error())
	}
	obj, oErr := client.ListHistory("testproject", nil)
	require.Error(t, oErr)
	require.Nil(t, obj)
}

func TestGetHistoryDecodeError(t *testing.T) {
	client, server, cErr := newTestRundeckClient([]byte(""), "application/json", 200)
	defer server.Close()
	if cErr != nil {
		t.Fatalf(cErr.Error())
	}
	obj, oErr := client.ListHistory("testproject", nil)
	require.Error(t, oErr)
	require.Nil(t, obj)
}

func TestGetHistoryOptions(t *testing.T) {
	jsonfile, err := responses.GetTestData(responses.HistoryResponseTestFile)
	if err != nil {
		t.Fatalf(err.Error())
	}

	client, server, cErr := newTestRundeckClient(jsonfile, "application/json", 200)
	defer server.Close()
	if cErr != nil {
		t.Fatalf(cErr.Error())
	}
	obj, cErr := client.ListHistory("testproject", map[string]string{"foo": "bar"})

	require.NoError(t, cErr)
	require.NotNil(t, obj)
}
