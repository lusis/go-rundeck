package responses

import (
	"encoding/json"
	"testing"

	"github.com/mitchellh/mapstructure"
	"github.com/stretchr/testify/require"
)

func TestExecutionReponses(t *testing.T) {
	testCases := []struct{
		name string
		placeholder interface{}
		obj interface{}
		testfile string
	}{
		{
			name: "ExecutionResponse",
			placeholder: make(map[string]interface{}),
			obj: &ExecutionResponse{},
			testfile: ExecutionResponseTestFile,
		},
		{
			name: "ListRunningExecutionsResponse",
			placeholder: make(map[string]interface{}),
			obj: &ListRunningExecutionsResponse{},
			testfile: ListRunningExecutionsResponseTestFile,
		},
		{
			name: "ExecutionInputFilesResponse",
			placeholder: make(map[string]interface{}),
			obj: &ExecutionInputFilesResponse{},
			testfile: ExecutionInputFilesResponseTestFile,
		},
		{
			name: "BulkDeleteExecutionsResponse",
			placeholder: make(map[string]interface{}),
			obj: &BulkDeleteExecutionsResponse{},
			testfile: BulkDeleteExecutionsResponseTestFile,
		},
		{
			name: "AdHocExecutionResponse",
			placeholder: make(map[string]interface{}),
			obj: &AdHocExecutionResponse{},
			testfile: AdHocExecutionResponseTestFile,
		},
		{
			name: "ExecutionOutputResponse",
			placeholder: make(map[string]interface{}),
			obj: &ExecutionOutputResponse{},
			testfile: ExecutionOutputResponseTestFile,
		},
		{
			name: "ExecutionsMetricsResponse",
			placeholder: make(map[string]interface{}),
			obj: &ExecutionsMetricsResponse{},
			testfile: ExecutionsMetricsResponseTestFile,
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




func TestExecutionStateResponse(t *testing.T) {
	obj := &ExecutionStateResponse{}
	data, err := getAssetBytes(ExecutionStateResponseTestFile)
	require.NoError(t, err)
	placeholder := make(map[string]interface{})
	err = json.Unmarshal(data, &placeholder)
	require.NoError(t, err)
	config := newMSDecoderConfig()
	config.Result = obj
	// because of Steps, we need to be lax for this test and we'll check the Steps themselves
	config.ErrorUnused = false
	decoder, newErr := mapstructure.NewDecoder(config)
	require.NoError(t, newErr)
	dErr := decoder.Decode(placeholder)
	require.NoError(t, dErr)
	require.Implements(t, (*VersionedResponse)(nil), obj)
	require.Len(t, obj.Steps, obj.StepCount)
}

func TestExecutionStateExecutionStepResponse(t *testing.T) {
	esr := &ExecutionStateResponse{}
	data, err := getAssetBytes(ExecutionStateResponseTestFile)
	require.NoError(t, err)
	// first we want to actually json unmarshal here
	err = json.Unmarshal(data, esr)
	require.NoError(t, err)

	sr := &ExecutionStepResponse{}
	config := newMSDecoderConfig()
	config.Result = sr
	decoder, newErr := mapstructure.NewDecoder(config)
	require.NoError(t, newErr)
	dErr := decoder.Decode(esr.Steps[0])
	require.NoError(t, dErr)
}

func TestExecutionStateWorkflowStepResponse(t *testing.T) {
	esr := &ExecutionStateResponse{}
	data, err := getAssetBytes(ExecutionStateResponseTestFile)
	require.NoError(t, err)
	// first we want to actually json unmarshal here
	err = json.Unmarshal(data, esr)
	require.NoError(t, err)

	sr := &WorkflowStepResponse{}
	config := newMSDecoderConfig()
	config.Result = sr
	decoder, newErr := mapstructure.NewDecoder(config)
	require.NoError(t, newErr)
	dErr := decoder.Decode(esr.Steps[1])
	require.NoError(t, dErr)
}