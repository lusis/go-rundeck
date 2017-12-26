package rundeck

import (
	"testing"

	"github.com/lusis/go-rundeck/pkg/rundeck.v21/responses"

	"github.com/lusis/go-rundeck/pkg/rundeck.v21/responses/testdata"
	"github.com/stretchr/testify/assert"
)

func TestGetSystemInfo(t *testing.T) {
	jsonfile, err := testdata.GetBytes(responses.SystemInfoResponseTestFile)
	if err != nil {
		t.Fatalf(err.Error())
	}

	client, server, cErr := newTestRundeckClient(jsonfile, "application/json", 200)
	defer server.Close()
	if cErr != nil {
		t.Fatalf(cErr.Error())
	}

	s, err := client.GetSystemInfo()
	assert.NoError(t, err)
	assert.NotNil(t, s)
}
