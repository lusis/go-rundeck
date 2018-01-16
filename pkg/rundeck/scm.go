package rundeck

import (
	"bytes"
	"encoding/json"
	"fmt"

	multierror "github.com/hashicorp/go-multierror"
	requests "github.com/lusis/go-rundeck/pkg/rundeck/requests"
	"github.com/lusis/go-rundeck/pkg/rundeck/responses"
)

// SCMPlugins is a list of SCM plugins grouped by Integration (import or export)
type SCMPlugins struct {
	Import []responses.SCMPluginResponse
	Export []responses.SCMPluginResponse
}

// SCMPluginInputFields are input fields for scm plugins
type SCMPluginInputFields struct {
	responses.GetSCMPluginInputFieldsResponse
}

// ListSCMPlugins list the available plugins for the specified integration
// http://rundeck.org/docs/api/index.html#list-scm-plugins
func (c *Client) ListSCMPlugins(projectName string) (*SCMPlugins, error) {
	if err := c.checkRequiredAPIVersion(responses.SCMResponse{}); err != nil {
		return nil, err
	}
	importPluginsURL := fmt.Sprintf("project/%s/scm/import/plugins", projectName)
	exportPluginsURL := fmt.Sprintf("project/%s/scm/export/plugins", projectName)
	plugins := &SCMPlugins{}
	impRes, impErr := c.httpGet(importPluginsURL, accept("application/json"), requestExpects(200))
	if impErr != nil {
		return nil, impErr
	}
	expRes, expErr := c.httpGet(exportPluginsURL, accept("application/json"), requestExpects(200))
	if expErr != nil {
		return nil, expErr
	}
	imports := &responses.ListSCMPluginsResponse{}
	exports := &responses.ListSCMPluginsResponse{}
	if jsonErr := json.Unmarshal(expRes, exports); jsonErr != nil {
		return nil, &UnmarshalError{msg: multierror.Append(errDecoding, jsonErr).Error()}
	}
	if jsonErr := json.Unmarshal(impRes, imports); jsonErr != nil {
		return nil, &UnmarshalError{msg: multierror.Append(errDecoding, jsonErr).Error()}
	}
	plugins.Export = exports.Plugins
	plugins.Import = imports.Plugins
	return plugins, nil
}

// GetSCMPluginInputFields List the input fields for a specific plugin.
// http://rundeck.org/docs/api/index.html#get-scm-plugin-input-fields
func (c *Client) GetSCMPluginInputFields(projectName, integration, pluginType string) (*SCMPluginInputFields, error) {
	if err := c.checkRequiredAPIVersion(responses.SCMResponse{}); err != nil {
		return nil, err
	}
	// [PROJECT]/scm/[INTEGRATION]/plugin/[TYPE]/input
	u := fmt.Sprintf("project/%s/scm/%s/plugin/%s/input", projectName, integration, pluginType)
	fields := &SCMPluginInputFields{}
	res, resErr := c.httpGet(u, accept("application/json"), requestExpects(200))
	if resErr != nil {
		return nil, resErr
	}
	if jsonErr := json.Unmarshal(res, fields); jsonErr != nil {
		return nil, &UnmarshalError{msg: multierror.Append(errDecoding, jsonErr).Error()}
	}
	return fields, nil
}

// SetupSCMPluginForProject configures and enables a plugin for a project
// http://rundeck.org/docs/api/index.html#setup-scm-plugin-for-a-project
func (c *Client) SetupSCMPluginForProject(project, integration, pluginType string, params map[string]string) (*responses.SCMPluginForProjectResponse, error) {
	if err := c.checkRequiredAPIVersion(responses.SCMResponse{}); err != nil {
		return nil, err
	}
	// project/[PROJECT]/scm/[INTEGRATION]/plugin/[TYPE]/setup
	u := fmt.Sprintf("project/%s/scm/%s/plugin/%s/setup", project, integration, pluginType)
	data := &requests.SetupSCMPluginRequest{}
	data.Config = params
	reqBody, reqBodyErr := json.Marshal(data)
	if reqBodyErr != nil {
		return nil, reqBodyErr
	}
	results := &responses.SCMPluginForProjectResponse{}
	res, respErr := c.httpPost(u, withBody(bytes.NewReader(reqBody)), requestJSON(), requestExpects(200), requestExpects(400))
	if respErr != nil {
		return nil, respErr
	}
	if jsonErr := json.Unmarshal(res, results); jsonErr != nil {
		return nil, &UnmarshalError{msg: multierror.Append(errDecoding, jsonErr).Error()}
	}
	return results, nil
}

// EnableSCMPluginForProject enables a plugin for a project
// http://rundeck.org/docs/api/index.html#enable-scm-plugin-for-a-project
func (c *Client) EnableSCMPluginForProject() error {
	if err := c.checkRequiredAPIVersion(responses.SCMResponse{}); err != nil {
		return err
	}
	return fmt.Errorf("not yet implemented")
}

// DisableSCMPluginForProject disables a plugin for a project
// http://rundeck.org/docs/api/index.html#enable-scm-plugin-for-a-project
func (c *Client) DisableSCMPluginForProject() error {
	if err := c.checkRequiredAPIVersion(responses.SCMResponse{}); err != nil {
		return err
	}
	return fmt.Errorf("not yet implemented")
}

// GetProjectSCMStatus Get the SCM plugin status and available actions for the project.
// http://rundeck.org/docs/api/index.html#get-project-scm-status
func (c *Client) GetProjectSCMStatus(project, integration string) (*responses.GetProjectSCMStatusResponse, error) {
	if err := c.checkRequiredAPIVersion(responses.SCMResponse{}); err != nil {
		return nil, err
	}
	// project/[PROJECT]/scm/[INTEGRATION]/status
	u := fmt.Sprintf("project/%s/scm/%s/status", project, integration)
	results := &responses.GetProjectSCMStatusResponse{}
	res, resErr := c.httpGet(u, requestExpects(200), accept("application/json"))
	if resErr != nil {
		return nil, resErr
	}
	if jsonErr := json.Unmarshal(res, results); jsonErr != nil {
		return nil, &UnmarshalError{msg: multierror.Append(errDecoding, jsonErr).Error()}
	}
	return results, nil
}

// GetProjectSCMConfig Get the configuration properties for the current plugin.
// http://rundeck.org/docs/api/index.html#get-project-scm-config
func (c *Client) GetProjectSCMConfig() error {
	if err := c.checkRequiredAPIVersion(responses.SCMResponse{}); err != nil {
		return err
	}
	return fmt.Errorf("not yet implemented")
}

// GetProjectSCMActionInputFields Get the input fields and selectable items for a specific action.
// http://rundeck.org/docs/api/index.html#get-project-scm-action-input-fields
func (c *Client) GetProjectSCMActionInputFields(project, integration, action string) (*responses.GetSCMActionInputFieldsResponse, error) {
	if err := c.checkRequiredAPIVersion(responses.SCMResponse{}); err != nil {
		return nil, err
	}
	// project/[PROJECT]/scm/[INTEGRATION]/action/[ACTION_ID]/input
	u := fmt.Sprintf("project/%s/scm/%s/action/%s/input", project, integration, action)
	resp := &responses.GetSCMActionInputFieldsResponse{}
	res, resErr := c.httpGet(u, accept("application/json"), requestExpects(200))
	if resErr != nil {
		return nil, resErr
	}
	if jsonErr := json.Unmarshal(res, resp); jsonErr != nil {
		return nil, &UnmarshalError{msg: multierror.Append(errDecoding, jsonErr).Error()}
	}
	return resp, nil
}

// SCMAction is the type for scm actions
type SCMAction struct {
	Input   map[string]string
	Jobs    []string
	Items   []string
	Deleted []string
}

// SCMActionOption is the functional option type for setting an SCMAction
type SCMActionOption func(*SCMAction) error

// SCMActionInput sets the input map for an SCMAction
func SCMActionInput(m map[string]string) SCMActionOption {
	return func(a *SCMAction) error {
		a.Input = m
		return nil
	}
}

// SCMActionJobs sets the jobs for an SCMAction
func SCMActionJobs(jobs ...string) SCMActionOption {
	return func(a *SCMAction) error {
		a.Jobs = jobs
		return nil
	}
}

// SCMActionItems sets the itemss for an SCMAction
func SCMActionItems(items ...string) SCMActionOption {
	return func(a *SCMAction) error {
		a.Items = items
		return nil
	}
}

// SCMActionDeleted sets the jobs for an SCMAction
func SCMActionDeleted(deleted ...string) SCMActionOption {
	return func(a *SCMAction) error {
		a.Deleted = deleted
		return nil
	}
}

// PerformProjectSCMAction Perform the action for the SCM integration plugin, with a set of input parameters, selected Jobs, or Items, or Items to delete.
// http://rundeck.org/docs/api/index.html#perform-project-scm-action
func (c *Client) PerformProjectSCMAction(project, integration, action string, opts ...SCMActionOption) (*responses.SCMPluginForProjectResponse, error) {
	if err := c.checkRequiredAPIVersion(responses.SCMResponse{}); err != nil {
		return nil, err
	}
	// project/[PROJECT]/scm/[INTEGRATION]/action/[ACTION_ID]
	// project/scmproject-1515904723766086967/scm/export/action/project-commit
	u := fmt.Sprintf("project/%s/scm/%s/action/%s", project, integration, action)
	scmAction := &SCMAction{}
	for _, opt := range opts {
		if err := opt(scmAction); err != nil {
			return nil, &OptionError{msg: multierror.Append(errOption, err).Error()}
		}
	}
	scmReq := &requests.PerformSCMActionRequest{}
	scmReq.Input = scmAction.Input
	if len(scmAction.Deleted) != 0 {
		scmReq.Deleted = scmAction.Deleted
	}
	if len(scmAction.Jobs) != 0 {
		scmReq.Jobs = scmAction.Jobs
	}
	if len(scmAction.Items) != 0 {
		scmReq.Items = scmAction.Items
	}

	requestBody, marshalErr := json.Marshal(scmReq)
	if marshalErr != nil {
		return nil, marshalErr
	}
	post, postErr := c.httpPost(u, withBody(bytes.NewReader(requestBody)), requestJSON(), requestExpects(200), requestExpects(400))
	if postErr != nil {
		return nil, postErr
	}
	results := &responses.SCMPluginForProjectResponse{}
	if jsonErr := json.Unmarshal(post, results); jsonErr != nil {
		return nil, &UnmarshalError{msg: multierror.Append(errDecoding, jsonErr).Error()}
	}
	return results, nil
}

// GetJobSCMStatus gets a job's scm status
// http://rundeck.org/docs/api/index.html#get-job-scm-status
func (c *Client) GetJobSCMStatus() error {
	if err := c.checkRequiredAPIVersion(responses.SCMResponse{}); err != nil {
		return err
	}
	return fmt.Errorf("not yet implemented")
}

// GetJobSCMDiff Retrieve the file diff for the Job, if there are changes for the integration.
// http://rundeck.org/docs/api/index.html#get-job-scm-diff
func (c *Client) GetJobSCMDiff() error {
	if err := c.checkRequiredAPIVersion(responses.SCMResponse{}); err != nil {
		return err
	}
	return fmt.Errorf("not yet implemented")
}

// GetJobSCMActionInputFields Get the input fields and selectable items for a specific action for a job.
// http://rundeck.org/docs/api/index.html#get-project-scm-action-input-fields
func (c *Client) GetJobSCMActionInputFields() error {
	if err := c.checkRequiredAPIVersion(responses.SCMResponse{}); err != nil {
		return err
	}
	return fmt.Errorf("not yet implemented")
}

// PerformJobSCMAction Perform the action for the SCM integration plugin, with a set of input parameters, selected Jobs, or Items, or Items to delete.
// http://rundeck.org/docs/api/index.html#perform-project-scm-action
func (c *Client) PerformJobSCMAction() error {
	if err := c.checkRequiredAPIVersion(responses.SCMResponse{}); err != nil {
		return err
	}
	return fmt.Errorf("not yet implemented")
}
