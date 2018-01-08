package responses

import (
	"testing"

	"github.com/lusis/go-rundeck/pkg/rundeck/responses/testdata"
	"github.com/stretchr/testify/assert"
)

func TestListProjectsResponse(t *testing.T) {
	obj := ListProjectsResponse{}
	data, dataErr := testdata.GetBytes(ListProjectsResponseTestFile)
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
	assert.Equal(t, "[API Href]", obj[0].URL)
	assert.Equal(t, "testproject", obj[0].Name)
	assert.Equal(t, "test project", obj[0].Description)
}

func TestProjectInfoResponse(t *testing.T) {
	obj := ProjectInfoResponse{}
	data, dataErr := testdata.GetBytes(ProjectInfoResponseTestFile)
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
	assert.Equal(t, "[API Href]", obj.URL)
	assert.Equal(t, "testproject", obj.Name)
	assert.Equal(t, "test project", obj.Description)
	assert.NotNil(t, obj.Config)
	assert.Len(t, *obj.Config, 32)
}

func TestProjectConfigItemResponse(t *testing.T) {
	obj := ProjectConfigItemResponse{}
	data, dataErr := testdata.GetBytes(ProjectConfigItemResponseTestFile)
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
	assert.Equal(t, "project.ssh-connect-timeout", obj.Key)
	assert.Equal(t, "0", obj.Value)
}

func TestProjectArchiveExportAsyncResponse(t *testing.T) {
	obj := ProjectArchiveExportAsyncResponse{}
	data, dataErr := testdata.GetBytes(ProjectArchiveExportAsyncResponseTestFile)
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
	assert.Equal(t, "[TOKEN]", obj.Token)
	assert.False(t, obj.Ready)
	assert.Equal(t, 75, obj.Percentage)
}

func TestProjectImportArchiveResponse(t *testing.T) {
	obj := ProjectImportArchiveResponse{}
	data, dataErr := testdata.GetBytes(ProjectImportArchiveResponseTestFile)
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
	assert.Equal(t, "successful", obj.ImportStatus)
	assert.Nil(t, obj.Errors)
	assert.Nil(t, obj.ExecutionErrors)
	assert.Nil(t, obj.ACLErrors)
}

func TestProjectArchiveImportFailedResponse(t *testing.T) {
	obj := ProjectImportArchiveResponse{}
	data, dataErr := testdata.GetBytes("project_archive_import_failed.json")
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
	assert.Equal(t, "failed", obj.ImportStatus)
	assert.NotNil(t, obj.Errors)
	assert.Len(t, *obj.Errors, 2)
	assert.NotNil(t, obj.ExecutionErrors)
	assert.Len(t, *obj.ExecutionErrors, 2)
	assert.NotNil(t, obj.ACLErrors)
	assert.Len(t, *obj.ACLErrors, 2)
}
