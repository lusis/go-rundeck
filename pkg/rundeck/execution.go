package rundeck

import (
	"encoding/json"
	"fmt"
	"strconv"

	multierror "github.com/hashicorp/go-multierror"
	responses "github.com/lusis/go-rundeck/pkg/rundeck/responses"
)

// Execution represents a job execution
type Execution struct {
	responses.ExecutionResponse
}

// ExecutionState represents a job execution state
type ExecutionState struct {
	responses.ExecutionStateResponse
}

// AbortedExecution represents the results of aborting an execution
type AbortedExecution struct {
	responses.AbortExecutionResponse
}

// ExecutionOutput represents the output of an execution
type ExecutionOutput struct {
	responses.ExecutionOutputResponse
}

// GetExecutionInfo returns the details of a job execution
// http://rundeck.org/docs/api/index.html#execution-info
func (c *Client) GetExecutionInfo(executionID int) (*Execution, error) {
	if err := c.checkRequiredAPIVersion(responses.ExecutionResponse{}); err != nil {
		return nil, err
	}
	exec := &Execution{}
	u := fmt.Sprintf("execution/%d", executionID)
	res, err := c.httpGet(u, requestJSON(), requestExpects(200))
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
func (c *Client) GetExecutionState(executionID int) (*ExecutionState, error) {
	if err := c.checkRequiredAPIVersion(responses.ExecutionStateResponse{}); err != nil {
		return nil, err
	}
	data := &ExecutionState{}
	u := fmt.Sprintf("execution/%d/state", executionID)
	res, err := c.httpGet(u, requestJSON(), requestExpects(200))
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
func (c *Client) GetExecutionOutput(executionID int) (*ExecutionOutput, error) {
	if err := c.checkRequiredAPIVersion(responses.GenericVersionedResponse{}); err != nil {
		return nil, err
	}
	return c.GetExecutionOutputWithOffset(executionID, 0)
}

// GetExecutionOutputWithOffset gets the output of an execution at the given offset
func (c *Client) GetExecutionOutputWithOffset(executionID int, offset int) (*ExecutionOutput, error) {
	if err := c.checkRequiredAPIVersion(responses.GenericVersionedResponse{}); err != nil {
		return nil, err
	}
	t := &ExecutionOutput{}
	params := map[string]string{
		"offset": strconv.Itoa(offset),
	}
	u := fmt.Sprintf("execution/%d/output", executionID)
	res, err := c.httpGet(u, requestJSON(), requestExpects(200), queryParams(params))
	if err != nil {
		return nil, err
	}
	if jsonErr := json.Unmarshal(res, t); jsonErr != nil {
		return nil, &UnmarshalError{msg: multierror.Append(errDecoding, jsonErr).Error()}
	}
	return t, nil
}

// DeleteExecution deletes an execution
// http://rundeck.org/docs/api/index.html#delete-an-execution
func (c *Client) DeleteExecution(executionID int) error {
	if err := c.checkRequiredAPIVersion(responses.GenericVersionedResponse{}); err != nil {
		return err
	}
	u := fmt.Sprintf("execution/%d", executionID)
	_, err := c.httpDelete(u, requestJSON(), requestExpects(204))
	return err
}

// DisableExecution disables an execution
// http://rundeck.org/docs/api/index.html#disable-executions-for-a-job
func (c *Client) DisableExecution(executionID int) (bool, error) {
	if err := c.checkRequiredAPIVersion(responses.ToggleResponse{}); err != nil {
		return false, err
	}
	t := &responses.ToggleResponse{}
	u := fmt.Sprintf("job/%d/execution/disable", executionID)
	res, err := c.httpPost(u, requestJSON(), requestExpects(200))
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
func (c *Client) EnableExecution(executionID int) (bool, error) {
	if err := c.checkRequiredAPIVersion(responses.ToggleResponse{}); err != nil {
		return false, err
	}
	t := &responses.ToggleResponse{}
	u := fmt.Sprintf("job/%d/execution/enable", executionID)
	res, err := c.httpPost(u, requestExpects(200), requestJSON())
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
	if err := c.checkRequiredAPIVersion(responses.ExecutionInputFilesResponse{}); err != nil {
		return err
	}
	return fmt.Errorf("not yet implemented")
}

// AbortExecutionOption is a function option type for AbortExection options
type AbortExecutionOption func(m *map[string]string) error

// AbortExecutionAsUser sets the user for the abort execution call
func AbortExecutionAsUser(runAsUser string) AbortExecutionOption {
	return func(m *map[string]string) error {
		(*m)["runAsUser"] = runAsUser
		return nil
	}
}

// AbortExecution lists input files used for an execution
// http://rundeck.org/docs/api/index.html#aborting-executions
func (c *Client) AbortExecution(executionID int, opts ...AbortExecutionOption) (*AbortedExecution, error) {
	if err := c.checkRequiredAPIVersion(responses.AbortExecutionResponse{}); err != nil {
		return nil, err
	}

	data := AbortedExecution{}
	jobOpts := &map[string]string{}
	for _, opt := range opts {
		if err := opt(jobOpts); err != nil {
			return nil, &OptionError{msg: multierror.Append(errOption, err).Error()}
		}
	}
	u := fmt.Sprintf("execution/%d/abort", executionID)
	if val, ok := (*jobOpts)["runAsUser"]; ok {
		u = fmt.Sprintf("%s?asUser=%s", u, val)
	}
	res, err := c.httpGet(u, requestJSON(), requestExpects(200))
	if err != nil {
		return nil, err
	}
	if jsonErr := json.Unmarshal(res, &data); jsonErr != nil {
		return nil, &UnmarshalError{msg: multierror.Append(errDecoding, jsonErr).Error()}
	}
	return &data, nil
}
