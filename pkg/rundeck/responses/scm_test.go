package responses

import (
	"encoding/json"
	"testing"

	"github.com/mitchellh/mapstructure"
	"github.com/stretchr/testify/assert"
)

func TestListSCMPluginsResponseImport(t *testing.T) {
	obj := &ListSCMPluginsResponse{}
	data, dataErr := getAssetBytes(ListSCMPluginsResponseImportTestFile)
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
	data, dataErr := getAssetBytes(ListSCMPluginsResponseExportTestFile)
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
	data, dataErr := getAssetBytes(GetSCMPluginInputFieldsResponseImportTestData)
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
	data, dataErr := getAssetBytes(GetSCMPluginInputFieldsResponseExportTestData)
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
	data, dataErr := getAssetBytes(SCMPluginForProjectResponseEnableImportTestFile)
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
	data, dataErr := getAssetBytes(SCMPluginForProjectResponseDisableImportTestFile)
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
	data, dataErr := getAssetBytes(SCMPluginForProjectResponseEnableExportTestFile)
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
	data, dataErr := getAssetBytes(SCMPluginForProjectResponseDisableExportTestFile)
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
	data, dataErr := getAssetBytes(GetProjectSCMStatusResponseImportTestFile)
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
	data, dataErr := getAssetBytes(GetProjectSCMStatusResponseExportTestFile)
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
	data, dataErr := getAssetBytes(GetProjectSCMConfigResponseImportTestFile)
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
	data, dataErr := getAssetBytes(GetProjectSCMConfigResponseExportTestFile)
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
	data, dataErr := getAssetBytes(GetSCMActionInputFieldsResponseTestFileProjectImport)
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
	data, dataErr := getAssetBytes(GetSCMActionInputFieldsResponseTestFileProjectExport)
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
	data, dataErr := getAssetBytes(GetSCMActionInputFieldsResponseTestFileJobImport)
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
	data, dataErr := getAssetBytes(GetSCMActionInputFieldsResponseTestFileJobExport)
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
	data, dataErr := getAssetBytes(GetJobSCMStatusResponseTestFileImport)
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
	data, dataErr := getAssetBytes(GetJobSCMStatusResponseTestFileExport)
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
	data, dataErr := getAssetBytes(GetJobSCMDiffResponseTestFileExport)
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
	data, dataErr := getAssetBytes(GetJobSCMDiffResponseTestFileImport)
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
