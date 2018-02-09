package responses

import (
	"encoding/json"
	"testing"

	"github.com/lusis/go-rundeck/pkg/rundeck/responses/testdata"
	"github.com/mitchellh/mapstructure"
	"github.com/stretchr/testify/assert"
)

func TestExecutionResponse(t *testing.T) {
	obj := &ExecutionResponse{}
	data, dataErr := testdata.GetBytes(ExecutionResponseTestFile)
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

func TestListRunningExecutionsResponse(t *testing.T) {
	obj := &ListRunningExecutionsResponse{}
	data, dataErr := testdata.GetBytes(ListRunningExecutionsResponseTestFile)
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

func TestExecutionInputFilesResponse(t *testing.T) {
	obj := &ExecutionInputFilesResponse{}
	data, dataErr := testdata.GetBytes(ExecutionInputFilesResponseTestFile)
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

func TestBulkDeleteExecutionsResponse(t *testing.T) {
	obj := &BulkDeleteExecutionsResponse{}
	data, dataErr := testdata.GetBytes(BulkDeleteExecutionsResponseTestFile)
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

func TestExecutionStateResponse(t *testing.T) {
	obj := &ExecutionStateResponse{}
	data, dataErr := testdata.GetBytes(ExecutionStateResponseTestFile)
	if dataErr != nil {
		t.Fatalf(dataErr.Error())
	}
	placeholder := make(map[string]interface{})
	_ = json.Unmarshal(data, &placeholder)
	config := newMSDecoderConfig()
	config.Result = obj
	// because of Steps, we need to be lax for this test and we'll check the Steps themselves
	config.ErrorUnused = false
	decoder, newErr := mapstructure.NewDecoder(config)
	assert.NoError(t, newErr)
	dErr := decoder.Decode(placeholder)
	assert.NoError(t, dErr)
	assert.Implements(t, (*VersionedResponse)(nil), obj)
	assert.Len(t, obj.Steps, obj.StepCount)
}

func TestExecutionStateExecutionStepResponse(t *testing.T) {
	esr := &ExecutionStateResponse{}
	data, dataErr := testdata.GetBytes(ExecutionStateResponseTestFile)
	if dataErr != nil {
		t.Fatalf(dataErr.Error())
	}
	// first we want to actually json unmarshal here
	jerr := json.Unmarshal(data, esr)
	assert.NoError(t, jerr)

	sr := &ExecutionStepResponse{}
	config := newMSDecoderConfig()
	config.Result = sr
	decoder, newErr := mapstructure.NewDecoder(config)
	assert.NoError(t, newErr)
	dErr := decoder.Decode(esr.Steps[0])
	assert.NoError(t, dErr)
}

func TestExecutionStateWorkflowStepResponse(t *testing.T) {
	esr := &ExecutionStateResponse{}
	data, dataErr := testdata.GetBytes(ExecutionStateResponseTestFile)
	if dataErr != nil {
		t.Fatalf(dataErr.Error())
	}
	// first we want to actually json unmarshal here
	jerr := json.Unmarshal(data, esr)
	assert.NoError(t, jerr)

	sr := &WorkflowStepResponse{}
	config := newMSDecoderConfig()
	config.Result = sr
	decoder, newErr := mapstructure.NewDecoder(config)
	assert.NoError(t, newErr)
	dErr := decoder.Decode(esr.Steps[1])
	assert.NoError(t, dErr)
}
func TestAdHocExecutionResponse(t *testing.T) {
	obj := &AdHocExecutionResponse{}
	data, dataErr := testdata.GetBytes(AdHocExecutionResponseTestFile)
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

func TestExecutionOutputResponse(t *testing.T) {
	obj := &ExecutionOutputResponse{}
	data, dataErr := testdata.GetBytes(ExecutionOutputResponseTestFile)
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
