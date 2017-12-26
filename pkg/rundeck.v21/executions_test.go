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
