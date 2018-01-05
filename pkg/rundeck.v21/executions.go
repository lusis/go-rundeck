package rundeck

import (
	"bytes"
	"encoding/json"
	"strconv"
	"strings"

	multierror "github.com/hashicorp/go-multierror"
	requests "github.com/lusis/go-rundeck/pkg/rundeck.v21/requests"
	responses "github.com/lusis/go-rundeck/pkg/rundeck.v21/responses"
)

// Executions represents a list of executions for a project
type Executions responses.ListRunningExecutionsResponse

// DeletedExecutions represents the results of a bulk execution delete
type DeletedExecutions responses.BulkDeleteExecutionsResponse

// ListProjectExecutions lists a projects executions
func (c *Client) ListProjectExecutions(projectID string, options map[string]string) (*Executions, error) {
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
func (c *Client) ListRunningExecutions(projectID string) (*Executions, error) {
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

// DeleteExecutions deletes a list of executions by id
func (c *Client) DeleteExecutions(ids ...int) (*DeletedExecutions, error) {
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
func (c *Client) BulkEnableExecution(ids ...string) (*responses.BulkToggleResponse, error) {

	req := &requests.BulkToggleRequest{
		IDs: ids,
	}
	results := &responses.BulkToggleResponse{}
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
func (c *Client) BulkDisableExecution(ids ...string) (*responses.BulkToggleResponse, error) {

	req := &requests.BulkToggleRequest{
		IDs: ids,
	}
	results := &responses.BulkToggleResponse{}
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
