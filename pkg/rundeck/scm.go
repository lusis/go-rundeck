package rundeck

import (
	"bytes"
	"encoding/json"
	"fmt"

	multierror "github.com/hashicorp/go-multierror"
	requests "github.com/lusis/go-rundeck/pkg/rundeck/requests"
	"github.com/lusis/go-rundeck/pkg/rundeck/responses"
)

// SCMPluginForProject represents an scm plugin for a project
type SCMPluginForProject struct {
	responses.SCMPluginForProjectResponse
}

// SCMActionInputFields represents project scm action input fields
type SCMActionInputFields struct {
	responses.GetSCMActionInputFieldsResponse
}

// ProjectSCMConfig represents a project's SCM config
type ProjectSCMConfig struct {
	responses.GetProjectSCMConfigResponse
}

// SCMProjectStatus represents a project's scm status
type SCMProjectStatus struct {
	responses.GetProjectSCMStatusResponse
}

// SCMPlugins is a list of SCM plugins grouped by Integration (import or export)
type SCMPlugins struct {
	Import []responses.SCMPluginResponse
	Export []responses.SCMPluginResponse
}

// SCMPluginInputFields are input fields for scm plugins
type SCMPluginInputFields struct {
	responses.GetSCMPluginInputFieldsResponse
}

// SCMJobStatus represents a job's scm status
type SCMJobStatus struct {
	responses.GetJobSCMStatusResponse
}

// SCMJobDiff represent's a job's scm diff
type SCMJobDiff struct {
	responses.GetJobSCMDiffResponse
}

// ListSCMPlugins list the available plugins for the specified integration
// http://rundeck.org/docs/api/index.html#list-scm-plugins
// One minor customization we do here is to return both import and export in this for a nicer bit of sugar
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

// GetProjectSCMPluginInputFields List the input fields for a specific plugin.
// http://rundeck.org/docs/api/index.html#get-scm-plugin-input-fields
func (c *Client) GetProjectSCMPluginInputFields(projectName, integration, pluginType string) (*SCMPluginInputFields, error) {
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
func (c *Client) SetupSCMPluginForProject(project, integration, pluginType string, params map[string]string) (*SCMPluginForProject, error) {
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
	results := &SCMPluginForProject{}
	res, respErr := c.httpPost(u, withBody(bytes.NewReader(reqBody)), requestJSON(), requestExpects(200), requestExpects(400))
	if respErr != nil {
		return nil, respErr
	}
	if jsonErr := json.Unmarshal(res, results); jsonErr != nil {
		return nil, &UnmarshalError{msg: multierror.Append(errDecoding, jsonErr).Error()}
	}
	if results.Success {
		return results, nil
	}
	var errs = []error{}
	for k, v := range results.ValidationErrors {
		errs = append(errs, fmt.Errorf("%s - %s", k, v))
	}
	return nil, &SCMValidationError{msg: multierror.Append(errValidation, errs...).Error()}
}

// EnableSCMPluginForProject enables a plugin for a project
// http://rundeck.org/docs/api/index.html#enable-scm-plugin-for-a-project
func (c *Client) EnableSCMPluginForProject(project, integration, pluginType string) error {
	if err := c.checkRequiredAPIVersion(responses.SCMResponse{}); err != nil {
		return err
	}
	u := fmt.Sprintf("project/%s/scm/%s/plugin/%s/enable",
		project, integration, pluginType)
	res, err := c.httpPost(u, withBody(nil), requestJSON(), requestExpects(200), requestExpects(400))
	if err != nil {
		return err
	}
	results := &SCMPluginForProject{}
	if jsonErr := json.Unmarshal(res, results); jsonErr != nil {
		return &UnmarshalError{msg: multierror.Append(errDecoding, jsonErr).Error()}
	}
	if results.Success {
		return nil
	}
	var errs = []error{}
	for k, v := range results.ValidationErrors {
		errs = append(errs, fmt.Errorf("%s - %s", k, v))
	}
	return &SCMValidationError{msg: multierror.Append(errValidation, errs...).Error()}
}

// DisableSCMPluginForProject disables a plugin for a project
// http://rundeck.org/docs/api/index.html#enable-scm-plugin-for-a-project
func (c *Client) DisableSCMPluginForProject(project, integration, pluginType string) error {
	if err := c.checkRequiredAPIVersion(responses.SCMResponse{}); err != nil {
		return err
	}
	u := fmt.Sprintf("project/%s/scm/%s/plugin/%s/disable",
		project, integration, pluginType)
	res, err := c.httpPost(u, withBody(nil), requestJSON(), requestExpects(200), requestExpects(400))
	if err != nil {
		return err
	}
	results := &SCMPluginForProject{}
	if jsonErr := json.Unmarshal(res, results); jsonErr != nil {
		return &UnmarshalError{msg: multierror.Append(errDecoding, jsonErr).Error()}
	}
	if results.Success {
		return nil
	}
	var errs = []error{}
	for k, v := range results.ValidationErrors {
		errs = append(errs, fmt.Errorf("%s - %s", k, v))
	}
	return &SCMValidationError{msg: multierror.Append(errValidation, errs...).Error()}
}

// GetProjectSCMStatus Get the SCM plugin status and available actions for the project.
// http://rundeck.org/docs/api/index.html#get-project-scm-status
func (c *Client) GetProjectSCMStatus(project, integration string) (*SCMProjectStatus, error) {
	if err := c.checkRequiredAPIVersion(responses.SCMResponse{}); err != nil {
		return nil, err
	}
	// project/[PROJECT]/scm/[INTEGRATION]/status
	u := fmt.Sprintf("project/%s/scm/%s/status", project, integration)
	results := &SCMProjectStatus{}
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
func (c *Client) GetProjectSCMConfig(projectName, integration string) (*ProjectSCMConfig, error) {
	if err := c.checkRequiredAPIVersion(responses.SCMResponse{}); err != nil {
		return nil, err
	}
	u := fmt.Sprintf("project/%s/scm/%s/config", projectName, integration)
	data := &ProjectSCMConfig{}
	res, err := c.httpGet(u, requestJSON(), requestExpects(200))
	if err != nil {
		return nil, err
	}
	if jsonErr := json.Unmarshal(res, data); jsonErr != nil {
		return nil, &UnmarshalError{msg: multierror.Append(errDecoding, jsonErr).Error()}
	}
	return data, nil
}

// GetProjectSCMActionInputFields Get the input fields and selectable items for a specific action.
// http://rundeck.org/docs/api/index.html#get-project-scm-action-input-fields
func (c *Client) GetProjectSCMActionInputFields(project, integration, action string) (*SCMActionInputFields, error) {
	if err := c.checkRequiredAPIVersion(responses.SCMResponse{}); err != nil {
		return nil, err
	}
	// project/[PROJECT]/scm/[INTEGRATION]/action/[ACTION_ID]/input
	u := fmt.Sprintf("project/%s/scm/%s/action/%s/input", project, integration, action)
	resp := &SCMActionInputFields{}
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
func (c *Client) PerformProjectSCMAction(project, integration, action string, opts ...SCMActionOption) (*SCMPluginForProject, error) {
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
	results := &SCMPluginForProject{}
	if jsonErr := json.Unmarshal(post, results); jsonErr != nil {
		return nil, &UnmarshalError{msg: multierror.Append(errDecoding, jsonErr).Error()}
	}
	return results, nil
}

// GetJobSCMStatus gets a job's scm status
// http://rundeck.org/docs/api/index.html#get-job-scm-status
func (c *Client) GetJobSCMStatus(jobid, integration string) (*SCMJobStatus, error) {
	if err := c.checkRequiredAPIVersion(responses.SCMResponse{}); err != nil {
		return nil, err
	}
	u := fmt.Sprintf("job/%s/scm/%s/status", jobid, integration)
	results := &SCMJobStatus{}
	res, resErr := c.httpGet(u, requestExpects(200), accept("application/json"))
	if resErr != nil {
		return nil, resErr
	}
	if jsonErr := json.Unmarshal(res, results); jsonErr != nil {
		return nil, &UnmarshalError{msg: multierror.Append(errDecoding, jsonErr).Error()}
	}
	return results, nil
}

// GetJobSCMDiff Retrieve the file diff for the Job, if there are changes for the integration.
// http://rundeck.org/docs/api/index.html#get-job-scm-diff
func (c *Client) GetJobSCMDiff(jobid, integration string) (*SCMJobDiff, error) {
	if err := c.checkRequiredAPIVersion(responses.SCMResponse{}); err != nil {
		return nil, err
	}
	u := fmt.Sprintf("job/%s/scm/%s/diff", jobid, integration)
	results := &SCMJobDiff{}
	res, resErr := c.httpGet(u, requestExpects(200), accept("application/json"))
	if resErr != nil {
		return nil, resErr
	}
	if jsonErr := json.Unmarshal(res, results); jsonErr != nil {
		return nil, &UnmarshalError{msg: multierror.Append(errDecoding, jsonErr).Error()}
	}
	return results, nil
}

// GetJobSCMActionInputFields Get the input fields and selectable items for a specific action.
// http://rundeck.org/docs/api/index.html#get-job-scm-action-input-fields
func (c *Client) GetJobSCMActionInputFields(jobid, integration, action string) (*SCMActionInputFields, error) {
	if err := c.checkRequiredAPIVersion(responses.SCMResponse{}); err != nil {
		return nil, err
	}
	// /api/15/job/[ID]/scm/[INTEGRATION]/action/[ACTION_ID]/input
	u := fmt.Sprintf("job/%s/scm/%s/action/%s/input", jobid, integration, action)
	resp := &SCMActionInputFields{}
	res, resErr := c.httpGet(u, accept("application/json"), requestExpects(200))
	if resErr != nil {
		return nil, resErr
	}
	if jsonErr := json.Unmarshal(res, resp); jsonErr != nil {
		return nil, &UnmarshalError{msg: multierror.Append(errDecoding, jsonErr).Error()}
	}
	return resp, nil
}

// PerformJobSCMAction Perform the action for the SCM integration plugin, with a set of input parameters, selected Jobs, or Items, or Items to delete.
// http://rundeck.org/docs/api/index.html#perform-job-scm-action
func (c *Client) PerformJobSCMAction() error {
	if err := c.checkRequiredAPIVersion(responses.SCMResponse{}); err != nil {
		return err
	}
	return fmt.Errorf("not yet implemented")
}
