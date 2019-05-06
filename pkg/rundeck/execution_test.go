package rundeck

import (
	"fmt"
	"testing"

	"github.com/lusis/go-rundeck/pkg/rundeck/responses"

	"github.com/stretchr/testify/require"
)

func TestGetExecution(t *testing.T) {
	jsonfile, err := responses.GetTestData(responses.ExecutionResponseTestFile)
	if err != nil {
		t.Fatalf(err.Error())
	}

	client, server, cErr := newTestRundeckClient(jsonfile, "application/json", 200)
	defer server.Close()
	if cErr != nil {
		t.Fatalf(cErr.Error())
	}

	obj, oerr := client.GetExecutionInfo(1)
	require.NoError(t, oerr)
	require.Equal(t, 1, obj.ID)
	require.Equal(t, "[url]", obj.HRef)
	require.Equal(t, "[url]", obj.Permalink)
	require.Equal(t, "[project]", obj.Project)
	require.Equal(t, "[user]", obj.User)
	require.Equal(t, "succeeded/failed/aborted/timedout/retried/other", obj.Status)
	require.Len(t, obj.FailedNodes, 2)
	require.Len(t, obj.SuccessfulNodes, 2)
	require.Equal(t, "echo hello there [... 5 steps]", obj.Description)
	require.Equal(t, "-opt1 testvalue -opt2 a", obj.ArgString)
	job := obj.Job
	require.Len(t, job.Options, 2)
	require.Equal(t, "[uuid]", job.ID)
	require.Equal(t, "[url]", job.HRef)
	require.Equal(t, "[url]", job.Permalink)
	require.Equal(t, int64(6094), job.AverageDuration)
	require.Equal(t, "[name]", job.Name)
	require.Equal(t, "[group]", job.Group)
	require.Equal(t, "[project]", job.Project)
	require.Equal(t, "[description]", job.Description)

	dateStarted := obj.DateStarted.Date
	dateEnded := obj.DateEnded.Date
	require.Equal(t, 2015, dateStarted.Year())
	require.Equal(t, 2016, dateEnded.Year())
}

func TestGetExecutionInvalidStatusCode(t *testing.T) {
	jsonfile, err := responses.GetTestData(responses.ExecutionResponseTestFile)
	if err != nil {
		t.Fatalf(err.Error())
	}

	client, server, cErr := newTestRundeckClient(jsonfile, "application/json", 500)
	defer server.Close()
	if cErr != nil {
		t.Fatalf(cErr.Error())
	}

	obj, oerr := client.GetExecutionInfo(1)
	require.Error(t, oerr)
	require.Nil(t, obj)
}

func TestGetExecutionJSONError(t *testing.T) {
	client, server, cErr := newTestRundeckClient([]byte(""), "application/json", 200)
	defer server.Close()
	if cErr != nil {
		t.Fatalf(cErr.Error())
	}

	obj, oerr := client.GetExecutionInfo(1)
	require.Error(t, oerr)
	require.Nil(t, obj)
}

func TestGetExecutionState(t *testing.T) {
	jsonfile, err := responses.GetTestData(responses.ExecutionStateResponseTestFile)
	if err != nil {
		t.Fatalf(err.Error())
	}

	client, server, cErr := newTestRundeckClient(jsonfile, "application/json", 200)
	defer server.Close()
	if cErr != nil {
		t.Fatalf(cErr.Error())
	}

	obj, oerr := client.GetExecutionState(1)
	require.NoError(t, oerr)
	require.NotNil(t, obj)
}

func TestGetExecutionStateJSONError(t *testing.T) {
	client, server, cErr := newTestRundeckClient([]byte(""), "application/json", 200)
	defer server.Close()
	if cErr != nil {
		t.Fatalf(cErr.Error())
	}

	obj, oerr := client.GetExecutionState(1)
	require.Error(t, oerr)
	require.Nil(t, obj)
}

func TestGetExecutionInvalidStatus(t *testing.T) {
	client, server, cErr := newTestRundeckClient([]byte(""), "application/json", 500)
	defer server.Close()
	if cErr != nil {
		t.Fatalf(cErr.Error())
	}

	obj, oerr := client.GetExecutionState(1)
	require.Error(t, oerr)
	require.Nil(t, obj)
}
func TestGetExecutionOutput(t *testing.T) {
	jsonfile, err := responses.GetTestData("execution_output.json")
	if err != nil {
		t.Fatalf(err.Error())
	}

	client, server, cErr := newTestRundeckClient(jsonfile, "application/json", 200)
	defer server.Close()
	if cErr != nil {
		t.Fatalf(cErr.Error())
	}

	obj, oerr := client.GetExecutionOutput(1)
	require.NoError(t, oerr)
	require.NotNil(t, obj)
}

func TestDeleteExecution(t *testing.T) {
	client, server, cErr := newTestRundeckClient([]byte(""), "application/json", 204)
	defer server.Close()
	if cErr != nil {
		t.Fatalf(cErr.Error())
	}

	err := client.DeleteExecution(1)
	require.NoError(t, err)
}

func TestDisableExecutionSuccess(t *testing.T) {
	client, server, cErr := newTestRundeckClient([]byte(`{"success":true}`), "application/json", 200)
	defer server.Close()
	if cErr != nil {
		t.Fatalf(cErr.Error())
	}

	res, err := client.DisableExecution(1)
	require.NoError(t, err)
	require.True(t, res)
}

func TestDisableExecutionJSONError(t *testing.T) {
	client, server, cErr := newTestRundeckClient([]byte(``), "application/json", 200)
	defer server.Close()
	if cErr != nil {
		t.Fatalf(cErr.Error())
	}

	res, err := client.DisableExecution(1)
	require.Error(t, err)
	require.False(t, res)
}

func TestDisableExecutionHTTPError(t *testing.T) {
	client, server, cErr := newTestRundeckClient([]byte(``), "application/json", 500)
	defer server.Close()
	if cErr != nil {
		t.Fatalf(cErr.Error())
	}

	res, err := client.DisableExecution(1)
	require.Error(t, err)
	require.False(t, res)
}

func TestEnableExecutionSuccess(t *testing.T) {
	client, server, cErr := newTestRundeckClient([]byte(`{"success":true}`), "application/json", 200)
	defer server.Close()
	if cErr != nil {
		t.Fatalf(cErr.Error())
	}

	res, err := client.EnableExecution(1)
	require.NoError(t, err)
	require.True(t, res)
}

func TestEnableExecutionJSONError(t *testing.T) {
	client, server, cErr := newTestRundeckClient([]byte(``), "application/json", 200)
	defer server.Close()
	if cErr != nil {
		t.Fatalf(cErr.Error())
	}

	res, err := client.EnableExecution(1)
	require.Error(t, err)
	require.False(t, res)
}

func TestEnableExecutionHTTPError(t *testing.T) {
	client, server, cErr := newTestRundeckClient([]byte(``), "application/json", 500)
	defer server.Close()
	if cErr != nil {
		t.Fatalf(cErr.Error())
	}

	res, err := client.EnableExecution(1)
	require.Error(t, err)
	require.False(t, res)
}

func TestAbortExecution(t *testing.T) {
	jsonfile, err := responses.GetTestData(responses.AbortExecutionResponseTestFile)
	if err != nil {
		t.Fatalf(err.Error())
	}

	client, server, _ := newTestRundeckClient(jsonfile, "application/json", 200)
	defer server.Close()

	obj, oerr := client.AbortExecution(1)
	require.NoError(t, oerr)
	require.NotNil(t, obj)
}

func TestAbortExecutionAsUser(t *testing.T) {
	jsonfile, err := responses.GetTestData(responses.AbortExecutionResponseTestFile)
	if err != nil {
		t.Fatalf(err.Error())
	}

	client, server, _ := newTestRundeckClient(jsonfile, "application/json", 200)
	defer server.Close()

	obj, oerr := client.AbortExecution(1, AbortExecutionAsUser("auser"))
	require.NoError(t, oerr)
	require.NotNil(t, obj)
}

func TestAbortExecutionHTTPError(t *testing.T) {
	jsonfile, err := responses.GetTestData(responses.AbortExecutionResponseTestFile)
	if err != nil {
		t.Fatalf(err.Error())
	}

	client, server, _ := newTestRundeckClient(jsonfile, "application/json", 500)
	defer server.Close()

	obj, oerr := client.AbortExecution(1)
	require.Error(t, oerr)
	require.Nil(t, obj)
}

func TestAbortExecutionOptionError(t *testing.T) {
	jsonfile, err := responses.GetTestData(responses.AbortExecutionResponseTestFile)
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
	require.Error(t, oerr)
	require.Nil(t, obj)
}

func TestAbortExecutionJSONError(t *testing.T) {

	client, server, _ := newTestRundeckClient([]byte("jsonfile"), "application/json", 200)
	defer server.Close()

	obj, oerr := client.AbortExecution(1)
	require.Error(t, oerr)
	require.Nil(t, obj)
}
