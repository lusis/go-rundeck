package rundeck

import (
	"encoding/json"
	"fmt"

	multierror "github.com/hashicorp/go-multierror"
	responses "github.com/lusis/go-rundeck/pkg/rundeck/responses"
)

// Execution represents a job execution
type Execution responses.ExecutionResponse

// ExecutionState represents a job execution state
type ExecutionState responses.ExecutionStateResponse

// GetExecutionInfo returns the details of a job execution
// http://rundeck.org/docs/api/index.html#execution-info
func (c *Client) GetExecutionInfo(executionID string) (*Execution, error) {
	if _, err := c.hasRequiredAPIVersion(minJSONSupportedAPIVersion, maxRundeckVersionInt); err != nil {
		return nil, err
	}
	exec := &Execution{}
	res, err := c.httpGet("execution/"+executionID, requestJSON(), requestExpects(200))
	if err != nil {
		return nil, err
	}
	if jsonErr := json.Unmarshal(res, exec); jsonErr != nil {
		return nil, &UnmarshalError{msg: multierror.Append(errDecoding, jsonErr).Error()}
	}
	return exec, nil
}

// GetExecutionState returns the state of an execution
// http://rundeck.org/docs/api/index.html#execution-state
func (c *Client) GetExecutionState(executionID string) (*ExecutionState, error) {
	if _, err := c.hasRequiredAPIVersion(minJSONSupportedAPIVersion, maxRundeckVersionInt); err != nil {
		return nil, err
	}
	data := &ExecutionState{}
	res, err := c.httpGet("execution/"+executionID+"/state", requestJSON(), requestExpects(200))
	if err != nil {
		return nil, err
	}
	if jsonErr := json.Unmarshal(res, data); jsonErr != nil {
		return nil, &UnmarshalError{msg: multierror.Append(errDecoding, jsonErr).Error()}
	}
	return data, nil
}

// GetExecutionOutput returns the output of an execution
// http://rundeck.org/docs/api/index.html#execution-output
func (c *Client) GetExecutionOutput(executionID string) ([]byte, error) {
	if _, err := c.hasRequiredAPIVersion(minJSONSupportedAPIVersion, maxRundeckVersionInt); err != nil {
		return nil, err
	}
	return c.httpGet("execution/"+executionID+"/output", accept("text/plain"), requestExpects(200))
}

// DeleteExecution deletes an execution
// http://rundeck.org/docs/api/index.html#delete-an-execution
func (c *Client) DeleteExecution(id string) error {
	if _, err := c.hasRequiredAPIVersion(minJSONSupportedAPIVersion, maxRundeckVersionInt); err != nil {
		return err
	}
	return c.httpDelete("execution/"+id, requestJSON(), requestExpects(204))
}

// DisableExecution disables an execution
// http://rundeck.org/docs/api/index.html#disable-executions-for-a-job
func (c *Client) DisableExecution(id string) (bool, error) {
	if _, err := c.hasRequiredAPIVersion(minJSONSupportedAPIVersion, maxRundeckVersionInt); err != nil {
		return false, err
	}
	t := &responses.ToggleResponse{}
	res, err := c.httpPost("job/"+id+"/execution/disable", requestJSON(), requestExpects(200))
	if err != nil {
		return false, err
	}
	if jsonErr := json.Unmarshal(res, t); jsonErr != nil {
		return false, &UnmarshalError{msg: multierror.Append(errDecoding, jsonErr).Error()}
	}
	return t.Success, nil
}

// EnableExecution enables an execution
// http://rundeck.org/docs/api/index.html#enable-executions-for-a-job
func (c *Client) EnableExecution(id string) (bool, error) {
	if _, err := c.hasRequiredAPIVersion(minJSONSupportedAPIVersion, maxRundeckVersionInt); err != nil {
		return false, err
	}
	t := &responses.ToggleResponse{}
	res, err := c.httpPost("job/"+id+"/execution/enable", requestExpects(200), requestJSON())
	if err != nil {
		return false, err
	}
	if jsonErr := json.Unmarshal(res, t); jsonErr != nil {
		return false, &UnmarshalError{msg: multierror.Append(errDecoding, jsonErr).Error()}
	}
	return t.Success, nil
}

// ListInputFilesForExecution lists input files used for an execution
// http://rundeck.org/docs/api/index.html#list-input-files-for-an-execution
func (c *Client) ListInputFilesForExecution() error {
	if _, err := c.hasRequiredAPIVersion(19, maxRundeckVersionInt); err != nil {
		return err
	}
	return fmt.Errorf("not yet implemented")
}

// AbortExecution lists input files used for an execution
// http://rundeck.org/docs/api/index.html#aborting-executions
func (c *Client) AbortExecution() error {
	if _, err := c.hasRequiredAPIVersion(19, maxRundeckVersionInt); err != nil {
		return err
	}
	return fmt.Errorf("not yet implemented")
}
