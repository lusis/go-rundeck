package rundeck

import (
	"testing"

	"github.com/lusis/go-rundeck/pkg/rundeck.v21/responses"
	"github.com/lusis/go-rundeck/pkg/rundeck.v21/responses/testdata"

	"github.com/stretchr/testify/assert"
)

func TestGetExecutions(t *testing.T) {
	jsonfile, err := testdata.GetBytes(responses.ExecutionResponseTestFile)
	if err != nil {
		t.Fatalf(err.Error())
	}

	client, server, cErr := newTestRundeckClient(jsonfile, "application/json", 200)
	defer server.Close()
	if cErr != nil {
		t.Fatalf(cErr.Error())
	}

	obj, cErr := client.GetExecution("1")
	assert.NoError(t, cErr)
	assert.Equal(t, 1, obj.ID)
	assert.Equal(t, "[url]", obj.HRef)
	assert.Equal(t, "[url]", obj.Permalink)
	assert.Equal(t, "[project]", obj.Project)
	assert.Equal(t, "[user]", obj.User)
	assert.Equal(t, "succeeded/failed/aborted/timedout/retried/other", obj.Status)
	assert.Len(t, obj.FailedNodes, 2)
	assert.Len(t, obj.SuccessfulNodes, 2)
	assert.Equal(t, "echo hello there [... 5 steps]", obj.Description)
	assert.Equal(t, "-opt1 testvalue -opt2 a", obj.ArgString)
	job := obj.Job
	assert.Len(t, job.Options, 2)
	assert.Equal(t, "[uuid]", job.ID)
	assert.Equal(t, "[url]", job.HRef)
	assert.Equal(t, "[url]", job.Permalink)
	assert.Equal(t, int64(6094), job.AverageDuration)
	assert.Equal(t, "[name]", job.Name)
	assert.Equal(t, "[group]", job.Group)
	assert.Equal(t, "[project]", job.Project)
	assert.Equal(t, "[description]", job.Description)

	dateStarted := obj.DateStarted.Date
	dateEnded := obj.DateEnded.Date
	assert.Equal(t, 2015, dateStarted.Year())
	assert.Equal(t, 2016, dateEnded.Year())
}

func TestGetExecutionState(t *testing.T) {
	jsonfile, err := testdata.GetBytes(responses.ExecutionStateResponseTestFile)
	if err != nil {
		t.Fatalf(err.Error())
	}

	client, server, cErr := newTestRundeckClient(jsonfile, "application/json", 200)
	defer server.Close()
	if cErr != nil {
		t.Fatalf(cErr.Error())
	}

	obj, cErr := client.GetExecutionState("1")
	assert.NoError(t, cErr)
	assert.NotNil(t, obj)
}

func TestGetExecutionOutput(t *testing.T) {
	jsonfile, err := testdata.GetBytes("execution_output.txt")
	if err != nil {
		t.Fatalf(err.Error())
	}

	client, server, cErr := newTestRundeckClient(jsonfile, "application/json", 200)
	defer server.Close()
	if cErr != nil {
		t.Fatalf(cErr.Error())
	}

	obj, cErr := client.GetExecutionOutput("1")
	assert.NoError(t, cErr)
	assert.NotEqual(t, "", string(obj))
}
