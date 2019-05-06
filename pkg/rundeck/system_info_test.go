package rundeck

import (
	"testing"

	"github.com/lusis/go-rundeck/pkg/rundeck/responses"

	"github.com/stretchr/testify/require"
)

func TestGetSystemInfo(t *testing.T) {
	jsonfile, err := responses.GetTestData(responses.SystemInfoResponseTestFile)
	if err != nil {
		t.Fatalf(err.Error())
	}

	client, server, cErr := newTestRundeckClient(jsonfile, "application/json", 200)
	defer server.Close()
	if cErr != nil {
		t.Fatalf(cErr.Error())
	}

	s, err := client.GetSystemInfo()
	require.NoError(t, err)
	require.NotNil(t, s)
}

func TestGetSystemJSONError(t *testing.T) {
	client, server, cErr := newTestRundeckClient([]byte(""), "application/json", 200)
	defer server.Close()
	if cErr != nil {
		t.Fatalf(cErr.Error())
	}

	s, err := client.GetSystemInfo()
	require.Error(t, err)
	require.Nil(t, s)
}

func TestGetSystemInvalidStatus(t *testing.T) {
	client, server, cErr := newTestRundeckClient([]byte(""), "application/json", 500)
	defer server.Close()
	if cErr != nil {
		t.Fatalf(cErr.Error())
	}

	s, err := client.GetSystemInfo()
	require.Error(t, err)
	require.Nil(t, s)
}
