package responses

import (
	"encoding/json"
	"testing"

	"github.com/mitchellh/mapstructure"
	"github.com/stretchr/testify/require"
)


func TestProjectsReponses(t *testing.T) {
	testCases := []struct{
		name string
		placeholder interface{}
		obj interface{}
		testfile string
	}{
		{
			name: "ListProjectsResponse",
			placeholder: make([]map[string]interface{},1,1),
			obj: &ListProjectsResponse{},
			testfile: ListProjectsResponseTestFile,
		},
		{
			name: "ProjectInfoResponse",
			placeholder: make(map[string]interface{}),
			obj: &ProjectInfoResponse{},
			testfile: ProjectInfoResponseTestFile,
		},
		{
			name: "ProjectConfigItemResponse",
			placeholder: make(map[string]interface{}),
			obj: &ProjectConfigItemResponse{},
			testfile: ProjectConfigItemResponseTestFile,
		},
		{
			name: "ProjectArchiveExportAsyncResponse",
			placeholder: make(map[string]interface{}),
			obj: &ProjectArchiveExportAsyncResponse{},
			testfile: ProjectArchiveExportAsyncResponseTestFile,
		},
		{
			name: "ProjectImportArchiveResponse",
			placeholder: make(map[string]interface{}),
			obj: &ProjectImportArchiveResponse{},
			testfile: ProjectImportArchiveResponseTestFile,
		},
		{
			name: "ProjectArchiveImportFailedResponse",
			placeholder: make(map[string]interface{}),
			obj: &ProjectImportArchiveResponse{},
			testfile: ProjectImportArchiveFailedResponseTestFile,
		},
		{
			name: "ProjectExecutionsMetricsResponse",
			placeholder: make(map[string]interface{}),
			obj: &ProjectExecutionsMetricsResponse{},
			testfile: ProjectExecutionsMetricsResponseTestFile,
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