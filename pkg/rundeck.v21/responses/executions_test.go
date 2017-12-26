package responses

import (
	"encoding/json"
	"testing"

	"github.com/lusis/go-rundeck/pkg/rundeck.v21/responses/testdata"
	"github.com/stretchr/testify/assert"
)

func TestExecutionResponse(t *testing.T) {
	obj := &ExecutionResponse{}
	data, dataErr := testdata.GetBytes(ExecutionResponseTestFile)
	if dataErr != nil {
		t.Error(dataErr.Error())
		t.FailNow()
	}
	err := obj.FromBytes(data)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
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

func TestListRunningExecutionsResponse(t *testing.T) {
	obj := &ListRunningExecutionsResponse{}
	data, dataErr := testdata.GetBytes(ListRunningExecutionsResponseTestFile)
	if dataErr != nil {
		t.Error(dataErr.Error())
		t.FailNow()
	}
	err := obj.FromBytes(data)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	assert.Len(t, obj.Executions, 1)
	assert.Len(t, obj.Executions[0].FailedNodes, 2)
	assert.Len(t, obj.Executions[0].SuccessfulNodes, 1)
	assert.Len(t, obj.Executions[0].Job.Options, 2)

	dateStarted := obj.Executions[0].DateStarted.Date
	dateEnded := obj.Executions[0].DateEnded.Date
	assert.Equal(t, 2015, dateStarted.Year())
	assert.Equal(t, 2016, dateEnded.Year())
}

func TestExecutionInputFilesResponse(t *testing.T) {
	obj := &ExecutionInputFilesResponse{}
	data, dataErr := testdata.GetBytes(ExecutionInputFilesResponseTestFile)
	if dataErr != nil {
		t.Error(dataErr.Error())
		t.FailNow()
	}
	err := obj.FromBytes(data)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	assert.Len(t, obj.Files, 1)
	assert.Equal(t, 2014, obj.Files[0].DateCreated.Year())
	assert.Equal(t, 2017, obj.Files[0].ExpirationDate.Year())
}

func TestBulkDeleteExecutionsResponse(t *testing.T) {
	obj := &BulkDeleteExecutionsResponse{}
	data, dataErr := testdata.GetBytes(BulkDeleteExecutionsResponseTestFile)
	if dataErr != nil {
		t.Error(dataErr.Error())
		t.FailNow()
	}
	err := obj.FromBytes(data)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	assert.Len(t, obj.Failures, 3)
	for _, f := range obj.Failures {
		assert.NotEmpty(t, f.ID)
		assert.NotEmpty(t, f.Message)
	}
	assert.Equal(t, 3, obj.FailedCount)
	assert.Equal(t, 2, obj.SuccessCount)
	assert.False(t, obj.AllSuccessful)
	assert.Equal(t, 5, obj.RequestCount)
}

func TestExecutionStateResponse(t *testing.T) {
	obj := &ExecutionStateResponse{}
	data, dataErr := testdata.GetBytes(ExecutionStateResponseTestFile)
	if dataErr != nil {
		t.Error(dataErr.Error())
		t.FailNow()
	}
	err := obj.FromBytes(data)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	assert.True(t, obj.Completed)
	assert.Equal(t, "SUCCEEDED", obj.ExecutionState)
	assert.NotNil(t, obj.EndTime)
	assert.NotNil(t, obj.StartTime)
	assert.NotNil(t, obj.UpdateTime)
	assert.Len(t, obj.AllNodes, 1)
	assert.Len(t, obj.TargetNodes, 1)
	assert.Len(t, obj.Nodes, 1)
	assert.Equal(t, 2, obj.StepCount)
	assert.Equal(t, 134, obj.ExecutionID)
	assert.Len(t, obj.Nodes["dignan"], 2)

	// test decoding the first step
	stdStep := &ExecutionStepResponse{}
	stdStepErr := json.Unmarshal(obj.Steps[0], stdStep)
	assert.NoError(t, stdStepErr)

	// test decoding the workflow step
	wfStep := &WorkflowStepResponse{}
	wfStepErr := json.Unmarshal(obj.Steps[1], wfStep)
	assert.NoError(t, wfStepErr)
	assert.True(t, wfStep.HasSubworkFlow)
	assert.Len(t, wfStep.Workflow.AllNodes, 1)
	assert.Len(t, wfStep.Workflow.TargetNodes, 1)
	wfSubStep := &ExecutionStepResponse{}
	wfSubStepErr := json.Unmarshal(wfStep.Workflow.Steps[0], wfSubStep)
	assert.NoError(t, wfSubStepErr)
	assert.NotNil(t, wfSubStep)
}

func TestAdHocExecutionResponse(t *testing.T) {
	obj := &AdHocExecutionResponse{}
	data, dataErr := testdata.GetBytes(AdHocExecutionResponseTestFile)
	if dataErr != nil {
		t.Error(dataErr.Error())
		t.FailNow()
	}
	err := obj.FromBytes(data)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	assert.Equal(t, "Immediate execution scheduled (X)", obj.Message)
	assert.Equal(t, 1, obj.Execution.ID)
	assert.Equal(t, "[API Href]", obj.Execution.HRef)
	assert.Equal(t, "[GUI Href]", obj.Execution.Permalink)
}
