package rundeck

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"

	multierror "github.com/hashicorp/go-multierror"
	requests "github.com/lusis/go-rundeck/pkg/rundeck/requests"
	responses "github.com/lusis/go-rundeck/pkg/rundeck/responses"
)

// AdHocExecution represents an adhoc execution
type AdHocExecution struct {
	responses.AdHocExecutionResponse
}

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
func CmdKeepGoing(b bool) AdHocRunOption {
	return func(c *requests.AdHocCommandRequest) error {
		c.NodeKeepGoing = b
		return nil
	}
}

// RunAdHocCommand runs an adhoc job - all nodes by default
// http://rundeck.org/docs/api/index.html#running-adhoc-commands
func (c *Client) RunAdHocCommand(projectID string, exec string, opts ...AdHocRunOption) (*AdHocExecution, error) {
	if err := c.checkRequiredAPIVersion(responses.AdHocExecutionResponse{}); err != nil {
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
	body, bodyErr := json.Marshal(req)
	if bodyErr != nil {
		return nil, bodyErr
	}
	res, err := c.httpPost("project/"+projectID+"/run/command", requestExpects(200), requestJSON(), withBody(bytes.NewReader(body)))
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(res, data); err != nil {
		return nil, &UnmarshalError{msg: multierror.Append(errEncoding, err).Error()}
	}
	return data, nil
}

// AdHocScriptOption is a function option type for adhoc commands
type AdHocScriptOption func(c *map[string]string) error

// ScriptRunAs is the option for specifying who to run an adhoc command
func ScriptRunAs(user string) AdHocScriptOption {
	return func(c *map[string]string) error {
		(*c)["asUser"] = user
		return nil
	}
}

// ScriptNodeFilters is the option for passing node filters to an adhoc command
func ScriptNodeFilters(filters string) AdHocScriptOption {
	return func(c *map[string]string) error {
		(*c)["filter"] = filters
		return nil
	}
}

// ScriptThreadCount is the option for number of threads to run an adhoc command
func ScriptThreadCount(count int) AdHocScriptOption {
	return func(c *map[string]string) error {
		(*c)["nodeThreadcount"] = fmt.Sprintf("%d", count)
		return nil
	}
}

// ScriptKeepGoing is the option to keep running even if a node fails
func ScriptKeepGoing(b bool) AdHocScriptOption {
	return func(c *map[string]string) error {
		(*c)["nodeKeepgoing"] = fmt.Sprintf("%t", b)
		return nil
	}

}

// ScriptInterpreter is the option to set the Script interpreter
func ScriptInterpreter(i string) AdHocScriptOption {
	return func(c *map[string]string) error {
		(*c)["scriptInterpreter"] = i
		return nil
	}
}

// ScriptArgString is the option for setting arguments passed to the Script being run
func ScriptArgString(i string) AdHocScriptOption {
	return func(c *map[string]string) error {
		(*c)["argString"] = i
		return nil
	}
}

// ScriptArgsQuoted is the option for setting if Script and args are quoted when passed to the interpreter
func ScriptArgsQuoted(q bool) AdHocScriptOption {
	return func(c *map[string]string) error {
		(*c)["interpreterArgsQuoted"] = fmt.Sprintf("%t", q)
		return nil
	}
}

// ScriptFileExtension is the option for setting the file extension on the remote host
func ScriptFileExtension(e string) AdHocScriptOption {
	return func(c *map[string]string) error {
		(*c)["fileExtension"] = e
		return nil
	}
}

// RunAdHocScript runs a Script ad-hoc
// http://rundeck.org/docs/api/index.html#running-adhoc-scripts
// Because script contents can be overly complicated and large,
// we do not currently run scripts via json body post
func (c *Client) RunAdHocScript(projectID string, scriptData io.Reader, opts ...AdHocScriptOption) (*AdHocExecution, error) {
	if err := c.checkRequiredAPIVersion(responses.AdHocExecutionResponse{}); err != nil {
		return nil, err
	}
	data := &AdHocExecution{}
	qp := &map[string]string{}
	for _, opt := range opts {
		if err := opt(qp); err != nil {
			return nil, &OptionError{msg: multierror.Append(errOption, err).Error()}
		}
	}

	scriptBytes, sbErr := ioutil.ReadAll(scriptData)
	if sbErr != nil {
		return nil, sbErr
	}
	if (*qp)["filter"] == "" {
		(*qp)["filter"] = "name: .*"
	}
	(*qp)["scriptFile"] = string(scriptBytes)
	res, err := c.httpPost("project/"+projectID+"/run/script",
		requestExpects(200),
		accept("application/json"),
		contentType("application/x-www-form-urlencoded"),
		queryParams(*qp))
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(res, data); err != nil {
		return nil, &UnmarshalError{msg: multierror.Append(errEncoding, err).Error()}
	}
	return data, nil
}

// AdHocScriptURLOption is a function option type for adhoc commands
type AdHocScriptURLOption func(c *map[string]string) error

// ScriptURLRunAs is the option for specifying who to run an adhoc command
func ScriptURLRunAs(user string) AdHocScriptURLOption {
	return func(c *map[string]string) error {
		(*c)["asUser"] = user
		return nil
	}
}

// ScriptURLNodeFilters is the option for passing node filters to an adhoc command
func ScriptURLNodeFilters(filters string) AdHocScriptURLOption {
	return func(c *map[string]string) error {
		(*c)["filter"] = filters
		return nil
	}
}

// ScriptURLThreadCount is the option for number of threads to run an adhoc command
func ScriptURLThreadCount(count int) AdHocScriptURLOption {
	return func(c *map[string]string) error {
		(*c)["nodeThreadcount"] = fmt.Sprintf("%d", count)
		return nil
	}
}

// ScriptURLKeepGoing is the option to keep running even if a node fails
func ScriptURLKeepGoing(b bool) AdHocScriptURLOption {
	return func(c *map[string]string) error {
		(*c)["nodeKeepgoing"] = fmt.Sprintf("%t", b)
		return nil
	}
}

// ScriptURLInterpreter is the option to set the ScriptURL interpreter
func ScriptURLInterpreter(i string) AdHocScriptURLOption {
	return func(c *map[string]string) error {
		(*c)["scriptInterpreter"] = i
		return nil
	}
}

// ScriptURLArgString is the option for setting arguments passed to the ScriptURL being run
func ScriptURLArgString(i string) AdHocScriptURLOption {
	return func(c *map[string]string) error {
		(*c)["argString"] = i
		return nil
	}
}

// ScriptURLArgsQuoted is the option for setting if ScriptURL and args are quoted when passed to the interpreter
func ScriptURLArgsQuoted(q bool) AdHocScriptURLOption {
	return func(c *map[string]string) error {
		(*c)["interpreterArgsQuoted"] = fmt.Sprintf("%t", q)
		return nil
	}
}

// ScriptURLFileExtension is the option for setting the file extension on the remote host
func ScriptURLFileExtension(e string) AdHocScriptURLOption {
	return func(c *map[string]string) error {
		(*c)["fileExtension"] = e
		return nil
	}
}

// RunAdHocScriptFromURL runs a ScriptURL ad-hoc from a url
// http://rundeck.org/docs/api/index.html#running-adhoc-ScriptURL-urls
// Due to the fact that we must still provide scriptURL as a query param,
// we do not currently run script urls via json body post
func (c *Client) RunAdHocScriptFromURL(projectID, scriptURL string, opts ...AdHocScriptURLOption) (*AdHocExecution, error) {
	if err := c.checkRequiredAPIVersion(responses.AdHocExecutionResponse{}); err != nil {
		return nil, err
	}
	data := &AdHocExecution{}
	qp := &map[string]string{}
	for _, opt := range opts {
		if err := opt(qp); err != nil {
			return nil, &OptionError{msg: multierror.Append(errOption, err).Error()}
		}
	}

	if (*qp)["filter"] == "" {
		(*qp)["filter"] = defaultNodeFilter
	}
	(*qp)["scriptURL"] = scriptURL
	res, err := c.httpPost("project/"+projectID+"/run/url",
		requestExpects(200),
		accept("application/json"),
		contentType("application/x-www-form-urlencoded"),
		queryParams(*qp))
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(res, data); err != nil {
		return nil, &UnmarshalError{msg: multierror.Append(errEncoding, err).Error()}
	}
	return data, nil
}
