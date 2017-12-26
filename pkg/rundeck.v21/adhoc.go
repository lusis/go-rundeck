package rundeck

import (
	"bytes"
	"encoding/json"

	requests "github.com/lusis/go-rundeck/pkg/rundeck.v21/requests"
	responses "github.com/lusis/go-rundeck/pkg/rundeck.v21/responses"
)

// AdHocExecution represents an adhoc execution
type AdHocExecution responses.AdHocExecutionResponse

// AdHocRunOption is a function option type for adhoc commands
type AdHocRunOption func(c *requests.AdHocCommandRequest) error

// CmdRunAs is the option for specifying who to run an adhoc command
func CmdRunAs(user string) AdHocRunOption {
	return func(c *requests.AdHocCommandRequest) error {
		c.AsUser = user
		return nil
	}
}

// CmdNodeFilters is the option for passing node filters to an adhoc command
func CmdNodeFilters(filters string) AdHocRunOption {
	return func(c *requests.AdHocCommandRequest) error {
		c.Filter = filters
		return nil
	}
}

// CmdThreadCount is the option for number of threads to run an adhoc command
func CmdThreadCount(count int) AdHocRunOption {
	return func(c *requests.AdHocCommandRequest) error {
		c.NodeThreadCount = count
		return nil
	}
}

// CmdKeepGoing is the option to keep running even if a node fails
func CmdKeepGoing() AdHocRunOption {
	return func(c *requests.AdHocCommandRequest) error {
		c.NodeKeepGoing = true
		return nil
	}
}

// RunAdhoc runs an adhoc job - all nodes by default
func (c *Client) RunAdhoc(projectID string, exec string, opts ...AdHocRunOption) (*AdHocExecution, error) {
	data := &AdHocExecution{}
	req := &requests.AdHocCommandRequest{}
	for _, opt := range opts {
		if err := opt(req); err != nil {
			return nil, err
		}
	}
	req.Project = projectID
	req.Exec = exec
	if req.Filter == "" {
		req.Filter = "name: .*"
	}
	body, bErr := json.Marshal(req)
	if bErr != nil {
		return nil, bErr
	}
	res, err := c.httpGet("project/"+projectID+"/run/command", requestJSON(), withBody(bytes.NewReader(body)))
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(res, data); err != nil {
		return nil, err
	}
	return data, nil
}
