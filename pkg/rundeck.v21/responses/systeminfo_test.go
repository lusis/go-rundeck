package responses

import (
	"testing"

	"github.com/lusis/go-rundeck/pkg/rundeck.v21/responses/testdata"
	"github.com/stretchr/testify/assert"
)

func TestSystemInfoResponse(t *testing.T) {
	obj := &SystemInfoResponse{}

	data, dataErr := testdata.GetBytes(SystemInfoResponseTestFile)
	if dataErr != nil {
		t.Error(dataErr.Error())
		t.FailNow()
	}
	err := obj.FromBytes(data)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	assert.NotNil(t, obj.System)
	assert.NotNil(t, obj.System.Timestamp)
	assert.NotNil(t, obj.System.Rundeck)
	assert.NotNil(t, obj.System.Executions)
	assert.NotNil(t, obj.System.OS)
	assert.NotNil(t, obj.System.JVM)
	assert.NotNil(t, obj.System.Stats)
	assert.NotNil(t, obj.System.Stats.Uptime)
	assert.NotNil(t, obj.System.Stats.Uptime.Since)
	assert.NotNil(t, obj.System.Stats.Uptime.Since.DateTime)
	assert.NotNil(t, obj.System.Stats.CPU)
	assert.NotNil(t, obj.System.Stats.Memory)
	assert.NotNil(t, obj.System.Stats.Threads)
	assert.NotNil(t, obj.System.Stats.Scheduler)
	assert.NotNil(t, obj.System.Metrics)
	assert.NotNil(t, obj.System.ThreadDump)
}
