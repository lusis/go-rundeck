package rundeck

import (
	"encoding/json"
	"errors"
	"strconv"
	"strings"

	responses "github.com/lusis/go-rundeck/pkg/rundeck.v21/responses"
)

// Executions represents a list of executions for a project
type Executions responses.ListRunningExecutionsResponse

// DeletedExecutions represents the results of a bulk execution delete
type DeletedExecutions responses.BulkDeleteExecutionsResponse

// ListProjectExecutions lists a projects executions
func (c *Client) ListProjectExecutions(projectID string, options map[string]string) (*Executions, error) {
	data := &Executions{}
	res, err := c.httpGet("project/"+projectID+"/executions", requestJSON(), queryParams(options))
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(res, data); err != nil {
		return nil, err
	}
	return data, nil
}

// ListRunningExecutions lists running executions
func (c *Client) ListRunningExecutions(projectID string) (*Executions, error) {
	options := make(map[string]string)
	data := &Executions{}
	res, err := c.httpGet("project/"+projectID+"/executions/running", requestJSON(), queryParams(options))
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(res, data); err != nil {
		return nil, err
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

	res, err := c.httpPost("executions/delete", accept("application/json"), queryParams(opts))
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(res, data); err != nil {
		return nil, err
	}
	return data, nil
}

// DeleteAllExecutionsForProject deletes all executions for a project up to the max (default: 10)
func (c *Client) DeleteAllExecutionsForProject(project string, max int) (*DeletedExecutions, error) {
	data := &DeletedExecutions{}

	eopts := make(map[string]string)
	eopts["max"] = strconv.Itoa(max)
	e, err := c.ListProjectExecutions(project, eopts)
	if err != nil {
		return nil, err
	}

	var toDelete []int
	for _, execution := range e.Executions {
		toDelete = append(toDelete, execution.ID)
	}
	if len(toDelete) == 0 {
		return data, errors.New("No executions found for project: " + project)
	}
	return c.DeleteExecutions(toDelete...)
}
