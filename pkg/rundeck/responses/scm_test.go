package responses

import (
	"encoding/json"
	"testing"

	"github.com/lusis/go-rundeck/pkg/rundeck/responses/testdata"
	"github.com/mitchellh/mapstructure"
	"github.com/stretchr/testify/assert"
)

func TestListSCMPluginsResponseImport(t *testing.T) {
	obj := &ListSCMPluginsResponse{}
	data, dataErr := testdata.GetBytes(ListSCMPluginsResponseImportTestFile)
	if dataErr != nil {
		t.Fatalf(dataErr.Error())
	}

	placeholder := make(map[string]interface{})
	_ = json.Unmarshal(data, &placeholder)
	config := newMSDecoderConfig()
	config.Result = obj
	decoder, newErr := mapstructure.NewDecoder(config)
	assert.NoError(t, newErr)
	dErr := decoder.Decode(placeholder)
	assert.NoError(t, dErr)
	assert.Implements(t, (*VersionedResponse)(nil), obj)
}

func TestListSCMPluginsResponseExport(t *testing.T) {
	obj := &ListSCMPluginsResponse{}
	data, dataErr := testdata.GetBytes(ListSCMPluginsResponseExportTestFile)
	if dataErr != nil {
		t.Fatalf(dataErr.Error())
	}
	placeholder := make(map[string]interface{})
	_ = json.Unmarshal(data, &placeholder)
	config := newMSDecoderConfig()
	config.Result = obj
	decoder, newErr := mapstructure.NewDecoder(config)
	assert.NoError(t, newErr)
	dErr := decoder.Decode(placeholder)
	assert.NoError(t, dErr)
	assert.Implements(t, (*VersionedResponse)(nil), obj)
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
	placeholder := make(map[string]interface{})
	_ = json.Unmarshal(data, &placeholder)
	config := newMSDecoderConfig()
	config.Result = obj
	decoder, newErr := mapstructure.NewDecoder(config)
	assert.NoError(t, newErr)
	dErr := decoder.Decode(placeholder)
	assert.NoError(t, dErr)
	assert.Implements(t, (*VersionedResponse)(nil), obj)
}

func TestGetSCMPluginInputFieldsResponseExport(t *testing.T) {
	obj := &GetSCMPluginInputFieldsResponse{}
	data, dataErr := testdata.GetBytes(GetSCMPluginInputFieldsResponseExportTestData)
	if dataErr != nil {
		t.Fatalf(dataErr.Error())
	}

	placeholder := make(map[string]interface{})
	_ = json.Unmarshal(data, &placeholder)
	config := newMSDecoderConfig()
	config.Result = obj
	decoder, newErr := mapstructure.NewDecoder(config)
	assert.NoError(t, newErr)
	dErr := decoder.Decode(placeholder)
	assert.NoError(t, dErr)
	assert.Implements(t, (*VersionedResponse)(nil), obj)
}

func TestSCMPluginForProjectResponseEnableImport(t *testing.T) {
	obj := &SCMPluginForProjectResponse{}
	data, dataErr := testdata.GetBytes(SCMPluginForProjectResponseEnableImportTestFile)
	if dataErr != nil {
		t.Fatalf(dataErr.Error())
	}

	placeholder := make(map[string]interface{})
	_ = json.Unmarshal(data, &placeholder)
	config := newMSDecoderConfig()
	config.Result = obj
	decoder, newErr := mapstructure.NewDecoder(config)
	assert.NoError(t, newErr)
	dErr := decoder.Decode(placeholder)
	assert.NoError(t, dErr)
	assert.Implements(t, (*VersionedResponse)(nil), obj)
}

func TestSCMPluginForProjectResponseDisableImport(t *testing.T) {
	obj := &SCMPluginForProjectResponse{}
	data, dataErr := testdata.GetBytes(SCMPluginForProjectResponseDisableImportTestFile)
	if dataErr != nil {
		t.Fatalf(dataErr.Error())
	}

	placeholder := make(map[string]interface{})
	_ = json.Unmarshal(data, &placeholder)
	config := newMSDecoderConfig()
	config.Result = obj
	decoder, newErr := mapstructure.NewDecoder(config)
	assert.NoError(t, newErr)
	dErr := decoder.Decode(placeholder)
	assert.NoError(t, dErr)
	assert.Implements(t, (*VersionedResponse)(nil), obj)
}

func TestSCMPluginForProjectResponseEnableExport(t *testing.T) {
	obj := &SCMPluginForProjectResponse{}
	data, dataErr := testdata.GetBytes(SCMPluginForProjectResponseEnableExportTestFile)
	if dataErr != nil {
		t.Fatalf(dataErr.Error())
	}

	placeholder := make(map[string]interface{})
	_ = json.Unmarshal(data, &placeholder)
	config := newMSDecoderConfig()
	config.Result = obj
	decoder, newErr := mapstructure.NewDecoder(config)
	assert.NoError(t, newErr)
	dErr := decoder.Decode(placeholder)
	assert.NoError(t, dErr)
	assert.Implements(t, (*VersionedResponse)(nil), obj)
}

func TestSCMPluginForProjectResponseDisableExport(t *testing.T) {
	obj := &SCMPluginForProjectResponse{}
	data, dataErr := testdata.GetBytes(SCMPluginForProjectResponseDisableExportTestFile)
	if dataErr != nil {
		t.Fatalf(dataErr.Error())
	}

	placeholder := make(map[string]interface{})
	_ = json.Unmarshal(data, &placeholder)
	config := newMSDecoderConfig()
	config.Result = obj
	decoder, newErr := mapstructure.NewDecoder(config)
	assert.NoError(t, newErr)
	dErr := decoder.Decode(placeholder)
	assert.NoError(t, dErr)
	assert.Implements(t, (*VersionedResponse)(nil), obj)
}

func TestGetProjectSCMStatusResponseImport(t *testing.T) {
	obj := &GetProjectSCMStatusResponse{}
	data, dataErr := testdata.GetBytes(GetProjectSCMStatusResponseImportTestFile)
	if dataErr != nil {
		t.Fatalf(dataErr.Error())
	}

	placeholder := make(map[string]interface{})
	_ = json.Unmarshal(data, &placeholder)
	config := newMSDecoderConfig()
	config.Result = obj
	decoder, newErr := mapstructure.NewDecoder(config)
	assert.NoError(t, newErr)
	dErr := decoder.Decode(placeholder)
	assert.NoError(t, dErr)
	assert.Implements(t, (*VersionedResponse)(nil), obj)
}

func TestGetProjectSCMStatusResponseExport(t *testing.T) {
	obj := &GetProjectSCMStatusResponse{}
	data, dataErr := testdata.GetBytes(GetProjectSCMStatusResponseExportTestFile)
	if dataErr != nil {
		t.Fatalf(dataErr.Error())
	}

	placeholder := make(map[string]interface{})
	_ = json.Unmarshal(data, &placeholder)
	config := newMSDecoderConfig()
	config.Result = obj
	decoder, newErr := mapstructure.NewDecoder(config)
	assert.NoError(t, newErr)
	dErr := decoder.Decode(placeholder)
	assert.NoError(t, dErr)
	assert.Implements(t, (*VersionedResponse)(nil), obj)
}

func TestGetProjectSCMConfigResponseImport(t *testing.T) {
	obj := &GetProjectSCMConfigResponse{}
	data, dataErr := testdata.GetBytes(GetProjectSCMConfigResponseImportTestFile)
	if dataErr != nil {
		t.Fatalf(dataErr.Error())
	}

	placeholder := make(map[string]interface{})
	_ = json.Unmarshal(data, &placeholder)
	config := newMSDecoderConfig()
	config.Result = obj
	decoder, newErr := mapstructure.NewDecoder(config)
	assert.NoError(t, newErr)
	dErr := decoder.Decode(placeholder)
	assert.NoError(t, dErr)
	assert.Implements(t, (*VersionedResponse)(nil), obj)
}

func TestGetProjectSCMConfigResponseExport(t *testing.T) {
	obj := &GetProjectSCMConfigResponse{}
	data, dataErr := testdata.GetBytes(GetProjectSCMConfigResponseExportTestFile)
	if dataErr != nil {
		t.Fatalf(dataErr.Error())
	}

	placeholder := make(map[string]interface{})
	_ = json.Unmarshal(data, &placeholder)
	config := newMSDecoderConfig()
	config.Result = obj
	decoder, newErr := mapstructure.NewDecoder(config)
	assert.NoError(t, newErr)
	dErr := decoder.Decode(placeholder)
	assert.NoError(t, dErr)
	assert.Implements(t, (*VersionedResponse)(nil), obj)
}

func TestGetSCMActionInputFieldsResponseProjectImport(t *testing.T) {
	obj := &GetSCMActionInputFieldsResponse{}
	data, dataErr := testdata.GetBytes(GetSCMActionInputFieldsResponseTestFileProjectImport)
	if dataErr != nil {
		t.Fatalf(dataErr.Error())
	}
	placeholder := make(map[string]interface{})
	_ = json.Unmarshal(data, &placeholder)
	config := newMSDecoderConfig()
	config.Result = obj
	decoder, newErr := mapstructure.NewDecoder(config)
	assert.NoError(t, newErr)
	dErr := decoder.Decode(placeholder)
	assert.NoError(t, dErr)
	assert.Implements(t, (*VersionedResponse)(nil), obj)
}

func TestGetSCMActionInputFieldsResponseProjectExport(t *testing.T) {
	obj := &GetSCMActionInputFieldsResponse{}
	data, dataErr := testdata.GetBytes(GetSCMActionInputFieldsResponseTestFileProjectExport)
	if dataErr != nil {
		t.Fatalf(dataErr.Error())
	}
	placeholder := make(map[string]interface{})
	_ = json.Unmarshal(data, &placeholder)
	config := newMSDecoderConfig()
	config.Result = obj
	decoder, newErr := mapstructure.NewDecoder(config)
	assert.NoError(t, newErr)
	dErr := decoder.Decode(placeholder)
	assert.NoError(t, dErr)
	assert.Implements(t, (*VersionedResponse)(nil), obj)
}

func TestGetSCMActionInputFieldsResponseJobImport(t *testing.T) {
	obj := &GetSCMActionInputFieldsResponse{}
	data, dataErr := testdata.GetBytes(GetSCMActionInputFieldsResponseTestFileJobImport)
	if dataErr != nil {
		t.Fatalf(dataErr.Error())
	}
	placeholder := make(map[string]interface{})
	_ = json.Unmarshal(data, &placeholder)
	config := newMSDecoderConfig()
	config.Result = obj
	decoder, newErr := mapstructure.NewDecoder(config)
	assert.NoError(t, newErr)
	dErr := decoder.Decode(placeholder)
	assert.NoError(t, dErr)
	assert.Implements(t, (*VersionedResponse)(nil), obj)
}

func TestGetSCMActionInputFieldsResponseJobExport(t *testing.T) {
	obj := &GetSCMActionInputFieldsResponse{}
	data, dataErr := testdata.GetBytes(GetSCMActionInputFieldsResponseTestFileJobExport)
	if dataErr != nil {
		t.Fatalf(dataErr.Error())
	}
	placeholder := make(map[string]interface{})
	_ = json.Unmarshal(data, &placeholder)
	config := newMSDecoderConfig()
	config.Result = obj
	decoder, newErr := mapstructure.NewDecoder(config)
	assert.NoError(t, newErr)
	dErr := decoder.Decode(placeholder)
	assert.NoError(t, dErr)
	assert.Implements(t, (*VersionedResponse)(nil), obj)
}

func TestGetJobSCMStatusResponseImport(t *testing.T) {
	obj := &GetJobSCMStatusResponse{}
	data, dataErr := testdata.GetBytes(GetJobSCMStatusResponseTestFileImport)
	if dataErr != nil {
		t.Fatalf(dataErr.Error())
	}
	placeholder := make(map[string]interface{})
	_ = json.Unmarshal(data, &placeholder)
	config := newMSDecoderConfig()
	config.Result = obj
	decoder, newErr := mapstructure.NewDecoder(config)
	assert.NoError(t, newErr)
	dErr := decoder.Decode(placeholder)
	assert.NoError(t, dErr)
	assert.Implements(t, (*VersionedResponse)(nil), obj)
}

func TestGetJobSCMStatusResponseExport(t *testing.T) {
	obj := &GetJobSCMStatusResponse{}
	data, dataErr := testdata.GetBytes(GetJobSCMStatusResponseTestFileExport)
	if dataErr != nil {
		t.Fatalf(dataErr.Error())
	}
	placeholder := make(map[string]interface{})
	_ = json.Unmarshal(data, &placeholder)
	config := newMSDecoderConfig()
	config.Result = obj
	decoder, newErr := mapstructure.NewDecoder(config)
	assert.NoError(t, newErr)
	dErr := decoder.Decode(placeholder)
	assert.NoError(t, dErr)
	assert.Implements(t, (*VersionedResponse)(nil), obj)
}

func TestGetJobSCMDiffResponseExport(t *testing.T) {
	obj := &GetJobSCMDiffResponse{}
	data, dataErr := testdata.GetBytes(GetJobSCMDiffResponseTestFileExport)
	if dataErr != nil {
		t.Fatalf(dataErr.Error())
	}
	placeholder := make(map[string]interface{})
	_ = json.Unmarshal(data, &placeholder)
	config := newMSDecoderConfig()
	config.Result = obj
	decoder, newErr := mapstructure.NewDecoder(config)
	assert.NoError(t, newErr)
	dErr := decoder.Decode(placeholder)
	assert.NoError(t, dErr)
	assert.Implements(t, (*VersionedResponse)(nil), obj)
}

func TestGetJobSCMDiffResponseImport(t *testing.T) {
	obj := &GetJobSCMDiffResponse{}
	data, dataErr := testdata.GetBytes(GetJobSCMDiffResponseTestFileImport)
	if dataErr != nil {
		t.Fatalf(dataErr.Error())
	}
	placeholder := make(map[string]interface{})
	_ = json.Unmarshal(data, &placeholder)
	config := newMSDecoderConfig()
	config.Result = obj
	decoder, newErr := mapstructure.NewDecoder(config)
	assert.NoError(t, newErr)
	dErr := decoder.Decode(placeholder)
	assert.NoError(t, dErr)
	assert.Implements(t, (*VersionedResponse)(nil), obj)
}
