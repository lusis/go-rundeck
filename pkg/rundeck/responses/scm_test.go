package responses

import (
	"testing"

	"github.com/lusis/go-rundeck/pkg/rundeck/responses/testdata"
	"github.com/stretchr/testify/assert"
)

func TestListSCMPluginsResponseImport(t *testing.T) {
	obj := &ListSCMPluginsResponse{}
	data, dataErr := testdata.GetBytes(ListSCMPluginsResponseImportTestFile)
	if dataErr != nil {
		t.Fatalf(dataErr.Error())
	}

	err := obj.FromBytes(data)
	if err != nil {
		t.Fatalf(err.Error())
	}
	assert.Implements(t, (*VersionedResponse)(nil), obj)
	assert.NotNil(t, obj)
}

func TestListSCMPluginsResponseExport(t *testing.T) {
	obj := &ListSCMPluginsResponse{}
	data, dataErr := testdata.GetBytes(ListSCMPluginsResponseExportTestFile)
	if dataErr != nil {
		t.Fatalf(dataErr.Error())
	}

	err := obj.FromBytes(data)
	if err != nil {
		t.Fatalf(err.Error())
	}
	assert.Implements(t, (*VersionedResponse)(nil), obj)
	assert.NotNil(t, obj)
}

func TestSCMPluginResponse(t *testing.T) {
	obj := SCMPluginResponse{}
	assert.Implements(t, (*VersionedResponse)(nil), obj)
}

func TestGetSCMPluginInputFieldsResponseImport(t *testing.T) {
	obj := &GetSCMPluginInputFieldsResponse{}
	data, dataErr := testdata.GetBytes(GetSCMPluginInputFieldsResponseImportTestData)
	if dataErr != nil {
		t.Fatalf(dataErr.Error())
	}

	err := obj.FromBytes(data)
	if err != nil {
		t.Fatalf(err.Error())
	}
	assert.Implements(t, (*VersionedResponse)(nil), obj)
	assert.NotNil(t, obj)
}

func TestGetSCMPluginInputFieldsResponseExport(t *testing.T) {
	obj := &GetSCMPluginInputFieldsResponse{}
	data, dataErr := testdata.GetBytes(GetSCMPluginInputFieldsResponseExportTestData)
	if dataErr != nil {
		t.Fatalf(dataErr.Error())
	}

	err := obj.FromBytes(data)
	if err != nil {
		t.Fatalf(err.Error())
	}
	assert.Implements(t, (*VersionedResponse)(nil), obj)
	assert.NotNil(t, obj)
}

func TestSCMPluginForProjectResponseEnableImport(t *testing.T) {
	obj := &SCMPluginForProjectResponse{}
	data, dataErr := testdata.GetBytes(SCMPluginForProjectResponseEnableImportTestFile)
	if dataErr != nil {
		t.Fatalf(dataErr.Error())
	}

	err := obj.FromBytes(data)
	if err != nil {
		t.Fatalf(err.Error())
	}
	assert.Implements(t, (*VersionedResponse)(nil), obj)
	assert.NotNil(t, obj)
}

func TestSCMPluginForProjectResponseDisableImport(t *testing.T) {
	obj := &SCMPluginForProjectResponse{}
	data, dataErr := testdata.GetBytes(SCMPluginForProjectResponseDisableImportTestFile)
	if dataErr != nil {
		t.Fatalf(dataErr.Error())
	}

	err := obj.FromBytes(data)
	if err != nil {
		t.Fatalf(err.Error())
	}
	assert.Implements(t, (*VersionedResponse)(nil), obj)
	assert.NotNil(t, obj)
}

func TestSCMPluginForProjectResponseEnableExport(t *testing.T) {
	obj := &SCMPluginForProjectResponse{}
	data, dataErr := testdata.GetBytes(SCMPluginForProjectResponseEnableExportTestFile)
	if dataErr != nil {
		t.Fatalf(dataErr.Error())
	}

	err := obj.FromBytes(data)
	if err != nil {
		t.Fatalf(err.Error())
	}
	assert.Implements(t, (*VersionedResponse)(nil), obj)
	assert.NotNil(t, obj)
}

func TestSCMPluginForProjectResponseDisableExport(t *testing.T) {
	obj := &SCMPluginForProjectResponse{}
	data, dataErr := testdata.GetBytes(SCMPluginForProjectResponseDisableExportTestFile)
	if dataErr != nil {
		t.Fatalf(dataErr.Error())
	}

	err := obj.FromBytes(data)
	if err != nil {
		t.Fatalf(err.Error())
	}
	assert.Implements(t, (*VersionedResponse)(nil), obj)
	assert.NotNil(t, obj)
}

func TestGetProjectSCMStatusResponseImport(t *testing.T) {
	obj := &GetProjectSCMStatusResponse{}
	data, dataErr := testdata.GetBytes(GetProjectSCMStatusResponseImportTestFile)
	if dataErr != nil {
		t.Fatalf(dataErr.Error())
	}

	err := obj.FromBytes(data)
	if err != nil {
		t.Fatalf(err.Error())
	}
	assert.Implements(t, (*VersionedResponse)(nil), obj)
	assert.NotNil(t, obj)
}

func TestGetProjectSCMStatusResponseExport(t *testing.T) {
	obj := &GetProjectSCMStatusResponse{}
	data, dataErr := testdata.GetBytes(GetProjectSCMStatusResponseExportTestFile)
	if dataErr != nil {
		t.Fatalf(dataErr.Error())
	}

	err := obj.FromBytes(data)
	if err != nil {
		t.Fatalf(err.Error())
	}
	assert.Implements(t, (*VersionedResponse)(nil), obj)
	assert.NotNil(t, obj)
}
