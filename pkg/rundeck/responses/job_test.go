package responses

import (
	"testing"

	"github.com/lusis/go-rundeck/pkg/rundeck/responses/testdata"
	"github.com/stretchr/testify/assert"
)

func TestJobsResponse(t *testing.T) {
	obj := JobsResponse{}
	data, dataErr := testdata.GetBytes(JobsResponseTestFile)
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
	assert.Len(t, obj, 1)
	assert.Equal(t, "[UUID]", obj[0].ID)
	assert.Equal(t, "[name]", obj[0].Name)
	assert.Equal(t, "[group]", obj[0].Group)
	assert.Equal(t, "[project]", obj[0].Project)
	assert.Equal(t, "...", obj[0].Description)
	assert.Equal(t, "[API url]", obj[0].HRef)
	assert.Equal(t, "[GUI url]", obj[0].Permalink)
	assert.Equal(t, "[UUID]", obj[0].ServerNodeUUID)
	assert.True(t, obj[0].Scheduled)
	assert.False(t, obj[0].ScheduleEnabled)
	assert.True(t, obj[0].Enabled)
	assert.True(t, obj[0].ServerOwned)
}

func TestJobMetaDataResponse(t *testing.T) {
	obj := &JobMetaDataResponse{}
	data, dataErr := testdata.GetBytes(JobMetaDataResponseTestFile)
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
	assert.Equal(t, "[UUID]", obj.ID)
	assert.Equal(t, "[name]", obj.Name)
	assert.Equal(t, "[group]", obj.Group)
	assert.Equal(t, "[project]", obj.Project)
	assert.Equal(t, "...", obj.Description)
	assert.Equal(t, "[API url]", obj.HRef)
	assert.Equal(t, "[GUI url]", obj.Permalink)
	assert.False(t, obj.Scheduled)
	assert.False(t, obj.ScheduleEnabled)
	assert.True(t, obj.Enabled)
	assert.Equal(t, int64(1431975278220), obj.AverageDuration)
}

func TestImportedJobResponse(t *testing.T) {
	obj := &ImportedJobResponse{}
	data, dataErr := testdata.GetBytes(ImportedJobResponseTestFile)
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
}

func TestBulkDeleteJobResponse(t *testing.T) {
	obj := &BulkDeleteJobResponse{}
	data, dataErr := testdata.GetBytes(BulkDeleteJobResponseTestFile)
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
}

func TestJobOptionFileUploadResponse(t *testing.T) {
	obj := &JobOptionFileUploadResponse{}
	data, dataErr := testdata.GetBytes(JobOptionFileUploadResponseTestFile)
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
}

func TestUploadedJobInputFilesResponse(t *testing.T) {
	obj := &UploadedJobInputFilesResponse{}
	data, dataErr := testdata.GetBytes(UploadedJobInputFilesResponseTestFile)
	if dataErr != nil {
		t.Error(dataErr.Error())
		t.FailNow()
	}
	err := obj.FromBytes(data)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
}
