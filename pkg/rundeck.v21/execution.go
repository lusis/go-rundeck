package rundeck

import (
	"encoding/json"

	multierror "github.com/hashicorp/go-multierror"
	responses "github.com/lusis/go-rundeck/pkg/rundeck.v21/responses"
)

// Execution represents a job execution
type Execution responses.ExecutionResponse

// ExecutionState represents a job execution state
type ExecutionState responses.ExecutionStateResponse

// GetExecution returns the details of a job execution
func (c *Client) GetExecution(executionID string) (*Execution, error) {
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
func (c *Client) GetExecutionState(executionID string) (*ExecutionState, error) {
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
func (c *Client) GetExecutionOutput(executionID string) ([]byte, error) {
	return c.httpGet("execution/"+executionID+"/output", accept("text/plain"), requestExpects(200))
}

// DeleteExecution deletes an execution
func (c *Client) DeleteExecution(id string) error {
	return c.httpDelete("execution/"+id, requestJSON(), requestExpects(204))
}

// DisableExecution disables an execution
func (c *Client) DisableExecution(id string) (bool, error) {
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
func (c *Client) EnableExecution(id string) (bool, error) {
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
