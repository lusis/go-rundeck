package responses

import (
	"encoding/json"
	"testing"

	"github.com/mitchellh/mapstructure"
	"github.com/stretchr/testify/require"
)

func TestJobReponses(t *testing.T) {
	testCases := []struct{
		name string
		placeholder interface{}
		obj interface{}
		testfile string
	}{
		{
			name: "JobsResponse",
			placeholder: make([]map[string]interface{},1 ,1),
			obj: &JobsResponse{},
			testfile: JobsResponseTestFile,
		},
		{
			name: "JobMetaDataResponse",
			placeholder: make(map[string]interface{}),
			obj: &JobMetaDataResponse{},
			testfile: JobMetaDataResponseTestFile,
		},
		{
			name: "ImportedJobResponse",
			placeholder: make(map[string]interface{}),
			obj: &ImportedJobResponse{},
			testfile: ImportedJobResponseTestFile,
		},
		{
			name: "BulkDeleteJobResponse",
			placeholder: make(map[string]interface{}),
			obj: &BulkDeleteJobResponse{},
			testfile: BulkDeleteJobResponseTestFile,
		},
		{
			name: "UploadedJobInputFilesResponse",
			placeholder: make(map[string]interface{}),
			obj: &UploadedJobInputFilesResponse{},
			testfile: UploadedJobInputFilesResponseTestFile,
		},
		{
			name: "JobOptionFileUpload",
			placeholder: make(map[string]interface{}),
			obj: &JobOptionFileUploadResponse{},
			testfile: JobOptionFileUploadResponseTestFile,
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