package rundeck

import (
	"bytes"
	"encoding/json"
	"strconv"
	"strings"

	multierror "github.com/hashicorp/go-multierror"
	requests "github.com/lusis/go-rundeck/pkg/rundeck/requests"
	responses "github.com/lusis/go-rundeck/pkg/rundeck/responses"
)

// Executions represents a list of executions for a project
type Executions struct {
	responses.ListRunningExecutionsResponse
}

// DeletedExecutions represents the results of a bulk execution delete
type DeletedExecutions struct {
	responses.BulkDeleteExecutionsResponse
}

// BulkToggleResponse represents the results of a bulk toggle request
type BulkToggleResponse struct {
	responses.BulkToggleResponse
}

// ListProjectExecutions lists a projects executions
// http://rundeck.org/docs/api/index.html#execution-query
func (c *Client) ListProjectExecutions(projectID string, options map[string]string) (*Executions, error) {
	if err := c.checkRequiredAPIVersion(responses.ListRunningExecutionsResponse{}); err != nil {
		return nil, err
	}
	data := &Executions{}
	res, err := c.httpGet("project/"+projectID+"/executions",
		requestJSON(),
		queryParams(options),
		requestExpects(200))
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(res, data); err != nil {
		return nil, &UnmarshalError{msg: multierror.Append(errEncoding, err).Error()}
	}
	return data, nil
}

// ListRunningExecutions lists running executions
// http://rundeck.org/docs/api/index.html#listing-running-executions
func (c *Client) ListRunningExecutions(projectID string) (*Executions, error) {
	if err := c.checkRequiredAPIVersion(responses.ListRunningExecutionsResponse{}); err != nil {
		return nil, err
	}
	options := make(map[string]string)
	data := &Executions{}
	res, err := c.httpGet("project/"+projectID+"/executions/running",
		requestJSON(),
		queryParams(options),
		requestExpects(200))
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(res, data); err != nil {
		return nil, &UnmarshalError{msg: multierror.Append(errEncoding, err).Error()}
	}
	return data, nil
}

// BulkDeleteExecutions deletes a list of executions by id
// http://rundeck.org/docs/api/index.html#bulk-delete-executions
func (c *Client) BulkDeleteExecutions(ids ...int) (*DeletedExecutions, error) {
	if err := c.checkRequiredAPIVersion(responses.BulkDeleteExecutionsResponse{}); err != nil {
		return nil, err
	}
	data := &DeletedExecutions{}
	opts := make(map[string]string)

	toDelete := []string{}
	for _, i := range ids {
		toDelete = append(toDelete, strconv.Itoa(i))
	}
	opts["ids"] = strings.Join(toDelete, ",")

	res, err := c.httpPost("executions/delete",
		accept("application/json"),
		queryParams(opts),
		requestExpects(200))
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(res, data); err != nil {
		return nil, &UnmarshalError{msg: multierror.Append(errEncoding, err).Error()}
	}
	return data, nil
}

// BulkEnableExecution enables job execution in bulk
// http://rundeck.org/docs/api/index.html#bulk-toggle-job-execution
func (c *Client) BulkEnableExecution(ids ...string) (*BulkToggleResponse, error) {
	if err := c.checkRequiredAPIVersion(responses.BulkToggleResponse{}); err != nil {
		return nil, err
	}
	req := &requests.BulkToggleRequest{
		IDs: ids,
	}
	results := &BulkToggleResponse{}
	data, _ := json.Marshal(req)
	res, err := c.httpPost("jobs/execution/enable",
		withBody(bytes.NewReader(data)),
		requestJSON(),
		requestExpects(200))
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(res, results); err != nil {
		return nil, &UnmarshalError{msg: multierror.Append(errEncoding, err).Error()}
	}
	return results, nil
}

// BulkDisableExecution disables job execution in bulk
// http://rundeck.org/docs/api/index.html#bulk-toggle-job-execution
func (c *Client) BulkDisableExecution(ids ...string) (*BulkToggleResponse, error) {
	if err := c.checkRequiredAPIVersion(responses.BulkToggleResponse{}); err != nil {
		return nil, err
	}
	req := &requests.BulkToggleRequest{
		IDs: ids,
	}
	results := &BulkToggleResponse{}
	data, _ := json.Marshal(req)
	res, err := c.httpPost("jobs/execution/disable",
		withBody(bytes.NewReader(data)),
		requestJSON(),
		requestExpects(200))
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(res, results); err != nil {
		return nil, &UnmarshalError{msg: multierror.Append(errEncoding, err).Error()}
	}
	return results, nil
}
