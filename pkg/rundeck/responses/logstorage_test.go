package responses

import (
	"testing"

	"github.com/lusis/go-rundeck/pkg/rundeck/responses/testdata"
	"github.com/stretchr/testify/assert"
)

func TestLogStorageResponse(t *testing.T) {
	obj := &LogStorageResponse{}
	data, dataErr := testdata.GetBytes(LogStorageResponseTestFile)
	if dataErr != nil {
		t.Error(dataErr.Error())
		t.FailNow()
	}
	err := obj.FromBytes(data)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	assert.Implements(t, (*VersionedResponse)(nil), obj)
	assert.True(t, obj.Enabled)
	assert.Equal(t, "NAME", obj.PluginName)
	assert.Equal(t, 369, obj.SucceededCount)
	assert.Equal(t, 0, obj.FailedCount)
	assert.Equal(t, 0, obj.QueuedCount)
	assert.Equal(t, 369, obj.TotalCount)
	assert.Equal(t, 0, obj.IncompleteCount)
	assert.Equal(t, 0, obj.MissingCount)
}

func TestIncompleteLogStorageResponse(t *testing.T) {
	obj := &IncompleteLogStorageResponse{}
	data, dataErr := testdata.GetBytes(IncompleteLogStorageResponseTestFile)
	if dataErr != nil {
		t.Error(dataErr.Error())
		t.FailNow()
	}
	err := obj.FromBytes(data)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	assert.Implements(t, (*VersionedResponse)(nil), obj)
	assert.NotNil(t, obj.Executions)
	assert.Len(t, obj.Executions, 1)
	assert.Equal(t, 1, obj.Total)
	assert.Equal(t, 20, obj.Max)
	assert.Equal(t, 0, obj.Offset)
	execution := obj.Executions[0]
	assert.Equal(t, 1, execution.ID)
	assert.Equal(t, "myProject", execution.Project)
	assert.Equal(t, "[API Href]", execution.HRef)
	assert.Equal(t, "[GUI Href]", execution.Permalink)
	assert.True(t, execution.Storage.LocalFilesPresent)
	assert.True(t, execution.Storage.Queued)
	assert.True(t, execution.Storage.Failed)
	assert.NotNil(t, execution.Storage.Date)
	assert.Equal(t, "rdlog,state.json", execution.Storage.IncompleteFiletypes)
	assert.Len(t, execution.Errors, 2)

}
