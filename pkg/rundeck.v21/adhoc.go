package rundeck

import (
	"bytes"
	"encoding/json"
	"fmt"

	multierror "github.com/hashicorp/go-multierror"
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

// RunAdHocCommand runs an adhoc job - all nodes by default
// http://rundeck.org/docs/api/index.html#running-adhoc-commands
func (c *Client) RunAdHocCommand(projectID string, exec string, opts ...AdHocRunOption) (*AdHocExecution, error) {
	if _, err := c.hasRequiredAPIVersion(minJSONSupportedAPIVersion, maxRundeckVersionInt); err != nil {
		return nil, err
	}
	data := &AdHocExecution{}
	req := &requests.AdHocCommandRequest{}
	for _, opt := range opts {
		if err := opt(req); err != nil {
			return nil, &OptionError{msg: multierror.Append(errOption, err).Error()}
		}
	}
	req.Project = projectID
	req.Exec = exec
	if req.Filter == "" {
		req.Filter = "name: .*"
	}
	body, _ := json.Marshal(req)
	res, err := c.httpGet("project/"+projectID+"/run/command", requestExpects(200), requestJSON(), withBody(bytes.NewReader(body)))
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(res, data); err != nil {
		return nil, &UnmarshalError{msg: multierror.Append(errEncoding, err).Error()}
	}
	return data, nil
}

// RunAdHocScript runs a script ad-hoc
// http://rundeck.org/docs/api/index.html#running-adhoc-scripts
func (c *Client) RunAdHocScript() error {
	if _, err := c.hasRequiredAPIVersion(19, maxRundeckVersionInt); err != nil {
		return err
	}
	return fmt.Errorf("not yet implemented")
}

// RunAdHocScriptFromURL runs a script ad-hoc from a url
// http://rundeck.org/docs/api/index.html#running-adhoc-script-urls
func (c *Client) RunAdHocScriptFromURL() error {
	if _, err := c.hasRequiredAPIVersion(19, maxRundeckVersionInt); err != nil {
		return err
	}
	return fmt.Errorf("not yet implemented")
}
