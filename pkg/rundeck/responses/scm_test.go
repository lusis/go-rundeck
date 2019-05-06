package responses

import (
	"encoding/json"
	"testing"

	"github.com/mitchellh/mapstructure"
	"github.com/stretchr/testify/require"
)

func TestSCMReponses(t *testing.T) {
	testCases := []struct{
		name string
		placeholder interface{}
		obj interface{}
		testfile string
	}{
		{
			name: "ListSCMPluginsResponseImport",
			placeholder: make([]map[string]interface{},1 ,1),
			obj: &ListSCMPluginsResponse{},
			testfile: ListSCMPluginsResponseImportTestFile,
		},
		{
			name: "ListSCMPluginsResponseExport",
			placeholder: make(map[string]interface{}),
			obj: &ListSCMPluginsResponse{},
			testfile: ListSCMPluginsResponseExportTestFile,
		},
		{
			name: "GetSCMPluginInputFieldsResponseImport",
			placeholder: make(map[string]interface{}),
			obj: &GetSCMPluginInputFieldsResponse{},
			testfile: GetSCMPluginInputFieldsResponseImportTestData,
		},
		{
			name: "GetSCMPluginInputFieldsResponseExport",
			placeholder: make(map[string]interface{}),
			obj: &GetSCMPluginInputFieldsResponse{},
			testfile: GetSCMPluginInputFieldsResponseExportTestData,
		},
		{
			name: "SCMPluginForProjectResponseEnableImport",
			placeholder: make(map[string]interface{}),
			obj: &SCMPluginForProjectResponse{},
			testfile: SCMPluginForProjectResponseEnableImportTestFile,
		},
		{
			name: "SCMPluginForProjectResponseDisableImport",
			placeholder: make(map[string]interface{}),
			obj: &SCMPluginForProjectResponse{},
			testfile: SCMPluginForProjectResponseDisableImportTestFile,
		},
		{
			name: "SCMPluginForProjectResponseEnableExport",
			placeholder: make(map[string]interface{}),
			obj: &SCMPluginForProjectResponse{},
			testfile: SCMPluginForProjectResponseEnableExportTestFile,
		},
		{
			name: "SCMPluginForProjectResponseDisableExport",
			placeholder: make(map[string]interface{}),
			obj: &SCMPluginForProjectResponse{},
			testfile: SCMPluginForProjectResponseDisableExportTestFile,
		},
		{
			name: "GetProjectSCMStatusResponseImport",
			placeholder: make(map[string]interface{}),
			obj: &GetProjectSCMStatusResponse{},
			testfile: GetProjectSCMStatusResponseImportTestFile,
		},
		{
			name: "GetProjectSCMStatusResponseExport",
			placeholder: make(map[string]interface{}),
			obj: &GetProjectSCMStatusResponse{},
			testfile: GetProjectSCMStatusResponseExportTestFile,
		},
		{
			name: "GetProjectSCMConfigResponseImport",
			placeholder: make(map[string]interface{}),
			obj: &GetProjectSCMConfigResponse{},
			testfile: GetProjectSCMConfigResponseImportTestFile,
		},
		{
			name: "GetProjectSCMConfigResponseExport",
			placeholder: make(map[string]interface{}),
			obj: &GetProjectSCMConfigResponse{},
			testfile: GetProjectSCMConfigResponseExportTestFile,
		},
		{
			name: "GetSCMActionInputFieldsResponseProjectImport",
			placeholder: make(map[string]interface{}),
			obj: &GetSCMActionInputFieldsResponse{},
			testfile: GetSCMActionInputFieldsResponseTestFileProjectImport,
		},
		{
			name: "GetSCMActionInputFieldsResponseProjectExport",
			placeholder: make(map[string]interface{}),
			obj: &GetSCMActionInputFieldsResponse{},
			testfile: GetSCMActionInputFieldsResponseTestFileProjectExport,
		},
		{
			name: "GetSCMActionInputFieldsResponseJobImport",
			placeholder: make(map[string]interface{}),
			obj: &GetSCMActionInputFieldsResponse{},
			testfile: GetSCMActionInputFieldsResponseTestFileJobImport,
		},
		{
			name: "GetSCMActionInputFieldsResponseJobExport",
			placeholder: make(map[string]interface{}),
			obj: &GetSCMActionInputFieldsResponse{},
			testfile: GetSCMActionInputFieldsResponseTestFileJobExport,
		},
		{
			name: "GetJobSCMStatusResponseImport",
			placeholder: make(map[string]interface{}),
			obj: &GetJobSCMStatusResponse{},
			testfile: GetJobSCMStatusResponseTestFileImport,
		},
		{
			name: "GetJobSCMStatusResponseExport",
			placeholder: make(map[string]interface{}),
			obj: &GetJobSCMStatusResponse{},
			testfile: GetJobSCMStatusResponseTestFileExport,
		},
		{
			name: "GetJobSCMDiffResponseExport",
			placeholder: make(map[string]interface{}),
			obj: &GetJobSCMDiffResponse{},
			testfile: GetJobSCMDiffResponseTestFileExport,
		},
		{
			name: "GetJobSCMDiffResponseImport",
			placeholder: make(map[string]interface{}),
			obj: &GetJobSCMDiffResponse{},
			testfile: GetJobSCMDiffResponseTestFileImport,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			data, err := getAssetBytes(tc.testfile)
			require.NoError(t, err)
			err = json.Unmarshal(data, &tc.placeholder)
			require.NoError(t, err)
			config := newMSDecoderConfig()
			config.Result = tc.obj
			decoder, err := mapstructure.NewDecoder(config)
			require.NoError(t, err)
			err = decoder.Decode(tc.placeholder)
			require.NoError(t, err)
			require.Implements(t, (*VersionedResponse)(nil), tc.obj)
		})
	}
}


func TestSCMPluginResponse(t *testing.T) {
	obj := SCMPluginResponse{}
	require.Implements(t, (*VersionedResponse)(nil), obj)
}