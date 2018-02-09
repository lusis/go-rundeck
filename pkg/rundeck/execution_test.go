package rundeck

import (
	"fmt"
	"testing"

	"github.com/lusis/go-rundeck/pkg/rundeck/responses"
	"github.com/lusis/go-rundeck/pkg/rundeck/responses/testdata"

	"github.com/stretchr/testify/assert"
)

func TestGetExecution(t *testing.T) {
	jsonfile, err := testdata.GetBytes(responses.ExecutionResponseTestFile)
	if err != nil {
		t.Fatalf(err.Error())
	}

	client, server, cErr := newTestRundeckClient(jsonfile, "application/json", 200)
	defer server.Close()
	if cErr != nil {
		t.Fatalf(cErr.Error())
	}

	obj, oerr := client.GetExecutionInfo(1)
	assert.NoError(t, oerr)
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

func TestGetExecutionInvalidStatusCode(t *testing.T) {
	jsonfile, err := testdata.GetBytes(responses.ExecutionResponseTestFile)
	if err != nil {
		t.Fatalf(err.Error())
	}

	client, server, cErr := newTestRundeckClient(jsonfile, "application/json", 500)
	defer server.Close()
	if cErr != nil {
		t.Fatalf(cErr.Error())
	}

	obj, oerr := client.GetExecutionInfo(1)
	assert.Error(t, oerr)
	assert.Nil(t, obj)
}

func TestGetExecutionJSONError(t *testing.T) {
	client, server, cErr := newTestRundeckClient([]byte(""), "application/json", 200)
	defer server.Close()
	if cErr != nil {
		t.Fatalf(cErr.Error())
	}

	obj, oerr := client.GetExecutionInfo(1)
	assert.Error(t, oerr)
	assert.Nil(t, obj)
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

	obj, oerr := client.GetExecutionState(1)
	assert.NoError(t, oerr)
	assert.NotNil(t, obj)
}

func TestGetExecutionStateJSONError(t *testing.T) {
	client, server, cErr := newTestRundeckClient([]byte(""), "application/json", 200)
	defer server.Close()
	if cErr != nil {
		t.Fatalf(cErr.Error())
	}

	obj, oerr := client.GetExecutionState(1)
	assert.Error(t, oerr)
	assert.Nil(t, obj)
}

func TestGetExecutionInvalidStatus(t *testing.T) {
	client, server, cErr := newTestRundeckClient([]byte(""), "application/json", 500)
	defer server.Close()
	if cErr != nil {
		t.Fatalf(cErr.Error())
	}

	obj, oerr := client.GetExecutionState(1)
	assert.Error(t, oerr)
	assert.Nil(t, obj)
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

	obj, oerr := client.GetExecutionOutput(1)
	assert.NoError(t, oerr)
	assert.NotNil(t, obj)
}

func TestDeleteExecution(t *testing.T) {
	client, server, cErr := newTestRundeckClient([]byte(""), "application/json", 204)
	defer server.Close()
	if cErr != nil {
		t.Fatalf(cErr.Error())
	}

	err := client.DeleteExecution(1)
	assert.NoError(t, err)
}

func TestDisableExecutionSuccess(t *testing.T) {
	client, server, cErr := newTestRundeckClient([]byte(`{"success":true}`), "application/json", 200)
	defer server.Close()
	if cErr != nil {
		t.Fatalf(cErr.Error())
	}

	res, err := client.DisableExecution(1)
	assert.NoError(t, err)
	assert.True(t, res)
}

func TestDisableExecutionJSONError(t *testing.T) {
	client, server, cErr := newTestRundeckClient([]byte(``), "application/json", 200)
	defer server.Close()
	if cErr != nil {
		t.Fatalf(cErr.Error())
	}

	res, err := client.DisableExecution(1)
	assert.Error(t, err)
	assert.False(t, res)
}

func TestDisableExecutionHTTPError(t *testing.T) {
	client, server, cErr := newTestRundeckClient([]byte(``), "application/json", 500)
	defer server.Close()
	if cErr != nil {
		t.Fatalf(cErr.Error())
	}

	res, err := client.DisableExecution(1)
	assert.Error(t, err)
	assert.False(t, res)
}

func TestEnableExecutionSuccess(t *testing.T) {
	client, server, cErr := newTestRundeckClient([]byte(`{"success":true}`), "application/json", 200)
	defer server.Close()
	if cErr != nil {
		t.Fatalf(cErr.Error())
	}

	res, err := client.EnableExecution(1)
	assert.NoError(t, err)
	assert.True(t, res)
}

func TestEnableExecutionJSONError(t *testing.T) {
	client, server, cErr := newTestRundeckClient([]byte(``), "application/json", 200)
	defer server.Close()
	if cErr != nil {
		t.Fatalf(cErr.Error())
	}

	res, err := client.EnableExecution(1)
	assert.Error(t, err)
	assert.False(t, res)
}

func TestEnableExecutionHTTPError(t *testing.T) {
	client, server, cErr := newTestRundeckClient([]byte(``), "application/json", 500)
	defer server.Close()
	if cErr != nil {
		t.Fatalf(cErr.Error())
	}

	res, err := client.EnableExecution(1)
	assert.Error(t, err)
	assert.False(t, res)
}

func TestAbortExecution(t *testing.T) {
	jsonfile, err := testdata.GetBytes(responses.AbortExecutionResponseTestFile)
	if err != nil {
		t.Fatalf(err.Error())
	}

	client, server, _ := newTestRundeckClient(jsonfile, "application/json", 200)
	defer server.Close()

	obj, oerr := client.AbortExecution(1)
	assert.NoError(t, oerr)
	assert.NotNil(t, obj)
}

func TestAbortExecutionAsUser(t *testing.T) {
	jsonfile, err := testdata.GetBytes(responses.AbortExecutionResponseTestFile)
	if err != nil {
		t.Fatalf(err.Error())
	}

	client, server, _ := newTestRundeckClient(jsonfile, "application/json", 200)
	defer server.Close()

	obj, oerr := client.AbortExecution(1, AbortExecutionAsUser("auser"))
	assert.NoError(t, oerr)
	assert.NotNil(t, obj)
}

func TestAbortExecutionHTTPError(t *testing.T) {
	jsonfile, err := testdata.GetBytes(responses.AbortExecutionResponseTestFile)
	if err != nil {
		t.Fatalf(err.Error())
	}

	client, server, _ := newTestRundeckClient(jsonfile, "application/json", 500)
	defer server.Close()

	obj, oerr := client.AbortExecution(1)
	assert.Error(t, oerr)
	assert.Nil(t, obj)
}

func TestAbortExecutionOptionError(t *testing.T) {
	jsonfile, err := testdata.GetBytes(responses.AbortExecutionResponseTestFile)
	if err != nil {
		t.Fatalf(err.Error())
	}

	myopt := func() AbortExecutionOption {
		return func(m *map[string]string) error {
			return fmt.Errorf("option error happened")
		}
	}
	client, server, _ := newTestRundeckClient(jsonfile, "application/json", 200)
	defer server.Close()

	obj, oerr := client.AbortExecution(1, myopt())
	assert.Error(t, oerr)
	assert.Nil(t, obj)
}

func TestAbortExecutionJSONError(t *testing.T) {

	client, server, _ := newTestRundeckClient([]byte("jsonfile"), "application/json", 200)
	defer server.Close()

	obj, oerr := client.AbortExecution(1)
	assert.Error(t, oerr)
	assert.Nil(t, obj)
}
