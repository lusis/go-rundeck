package rundeck

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"strings"

	multierror "github.com/hashicorp/go-multierror"
	requests "github.com/lusis/go-rundeck/pkg/rundeck/requests"
	responses "github.com/lusis/go-rundeck/pkg/rundeck/responses"
)

// Project is a rundeck project
type Project struct {
	URL         string
	Name        string
	Description string
	Properties  map[string]string
}

// Projects is a collection of `Project`
type Projects []*Project

// ProjectExportOption is a functional option for project exports
type ProjectExportOption func(p *map[string]string) error

// ProjectExportExecutionIDs specifies the execution ids to export
func ProjectExportExecutionIDs(ids ...string) ProjectExportOption {
	return func(p *map[string]string) error {
		(*p)["executionIds"] = strings.Join(ids, ",")
		return nil
	}
}

// ProjectExportAll toggles exporting everything
func ProjectExportAll(b bool) ProjectExportOption {
	return func(p *map[string]string) error {
		(*p)["exportAll"] = fmt.Sprintf("%t", b)
		return nil
	}
}

// ProjectExportJobs toggles exporting jobs
func ProjectExportJobs(b bool) ProjectExportOption {
	return func(p *map[string]string) error {
		(*p)["exportJobs"] = fmt.Sprintf("%t", b)
		return nil
	}
}

// ProjectExportExecutions toggles exporting executions
func ProjectExportExecutions(b bool) ProjectExportOption {
	return func(p *map[string]string) error {
		(*p)["exportExecutions"] = fmt.Sprintf("%t", b)
		return nil
	}
}

// ProjectExportConfigs toggles exporting configs
func ProjectExportConfigs(b bool) ProjectExportOption {
	return func(p *map[string]string) error {
		(*p)["exportConfigs"] = fmt.Sprintf("%t", b)
		return nil
	}
}

// ProjectExportReadmes toggles exporting readmes and motds
func ProjectExportReadmes(b bool) ProjectExportOption {
	return func(p *map[string]string) error {
		(*p)["exportReadmes"] = fmt.Sprintf("%t", b)
		return nil
	}
}

// ProjectExportAcls toggles exporting project acls
func ProjectExportAcls(b bool) ProjectExportOption {
	return func(p *map[string]string) error {
		(*p)["exportAcls"] = fmt.Sprintf("%t", b)
		return nil
	}
}

// ProjectImportOption is a functional option for project imports
type ProjectImportOption func(p *map[string]string) error

// ProjectImportAcls toggles importing acls with the project
func ProjectImportAcls(b bool) ProjectImportOption {
	return func(p *map[string]string) error {
		(*p)["importACL"] = fmt.Sprintf("%t", b)
		return nil
	}
}

// ProjectImportConfigs toggles importing configs with the project
func ProjectImportConfigs(b bool) ProjectImportOption {
	return func(p *map[string]string) error {
		(*p)["importConfig"] = fmt.Sprintf("%t", b)
		return nil
	}
}

// ProjectImportExecutions toggles importing executions with the project
func ProjectImportExecutions(b bool) ProjectImportOption {
	return func(p *map[string]string) error {
		(*p)["importExecutions"] = fmt.Sprintf("%t", b)
		return nil
	}
}

// ProjectImportJobUUIDs toggles importing job uuids with the project
func ProjectImportJobUUIDs(s string) ProjectImportOption {
	return func(p *map[string]string) error {
		(*p)["jobUuidOption"] = s
		return nil
	}
}

// GetProjectInfo gets a project by name
// http://rundeck.org/docs/api/index.html#getting-project-info
func (c *Client) GetProjectInfo(name string) (*Project, error) {
	if err := c.checkRequiredAPIVersion(responses.ProjectInfoResponse{}); err != nil {
		return nil, err
	}
	p := &responses.ProjectInfoResponse{}
	res, err := c.httpGet("project/"+name, requestJSON(), requestExpects(200))
	if err != nil {
		return nil, err
	}
	if jsonErr := json.Unmarshal(res, &p); jsonErr != nil {
		return nil, &UnmarshalError{msg: multierror.Append(errEncoding, jsonErr).Error()}
	}

	project := &Project{
		URL:         p.URL,
		Name:        p.Name,
		Description: p.Description,
		Properties:  *p.Config,
	}
	return project, nil
}

// ListProjects lists all projects
// http://rundeck.org/docs/api/index.html#listing-projects
func (c *Client) ListProjects() (Projects, error) {
	if err := c.checkRequiredAPIVersion(responses.ListProjectsResponse{}); err != nil {
		return nil, err
	}
	data := &responses.ListProjectsResponse{}
	res, err := c.httpGet("projects", requestJSON(), requestExpects(200))
	if err != nil {
		return nil, err
	}
	if jsonErr := json.Unmarshal(res, &data); jsonErr != nil {
		return nil, &UnmarshalError{msg: multierror.Append(errEncoding, jsonErr).Error()}
	}
	projects := &Projects{}
	for _, p := range *data {
		*projects = append(*projects, &Project{
			URL:         p.URL,
			Name:        p.Name,
			Description: p.Description,
		})
	}
	return *projects, nil
}

// CreateProject makes a project
// http://rundeck.org/docs/api/index.html#project-creation
func (c *Client) CreateProject(name string, properties map[string]string) (*Project, error) {
	if err := c.checkRequiredAPIVersion(responses.ProjectInfoResponse{}); err != nil {
		return nil, err
	}
	req := &requests.ProjectCreationRequest{
		Name:   name,
		Config: &properties,
	}
	data, _ := json.Marshal(req)
	info := &responses.ProjectInfoResponse{}
	res, postErr := c.httpPost("projects", requestJSON(), withBody(bytes.NewReader(data)), requestExpects(201))
	if postErr != nil {
		return nil, postErr
	}
	if jsonErr := json.Unmarshal(res, &info); jsonErr != nil {
		return nil, &UnmarshalError{msg: multierror.Append(errEncoding, jsonErr).Error()}
	}
	project := &Project{
		URL:         info.URL,
		Name:        info.Name,
		Description: info.Description,
		Properties:  *info.Config,
	}
	return project, nil
}

// DeleteProject deletes a project
// http://rundeck.org/docs/api/index.html#project-deletion
func (c *Client) DeleteProject(p string) error {
	if err := c.checkRequiredAPIVersion(responses.ProjectInfoResponse{}); err != nil {
		return err
	}
	url := fmt.Sprintf("project/%s", p)
	_, err := c.httpDelete(url, requestJSON(), requestExpects(204))
	return err
}

// GetProjectConfiguration gets a project's configuration
// http://rundeck.org/docs/api/index.html#get-project-configuration
func (c *Client) GetProjectConfiguration(p string) (map[string]string, error) {
	if err := c.checkRequiredAPIVersion(responses.ProjectConfigResponse{}); err != nil {
		return nil, err
	}
	data := map[string]string{}
	res, err := c.httpGet("project/"+p+"/config", requestExpects(200), requestJSON())
	if err != nil {
		return nil, err
	}
	if jsonErr := json.Unmarshal(res, &data); jsonErr != nil {
		return nil, &UnmarshalError{msg: multierror.Append(errEncoding, jsonErr).Error()}
	}
	return data, nil
}

// PutProjectConfiguration replaces all configuration data with the submitted values
// http://rundeck.org/docs/api/index.html#put-project-configuration
func (c *Client) PutProjectConfiguration(projectName string, config map[string]string) (map[string]string, error) {
	if err := c.checkRequiredAPIVersion(responses.ProjectConfigResponse{}); err != nil {
		return nil, err
	}

	body, mErr := json.Marshal(config)
	if mErr != nil {
		return nil, mErr
	}
	res, err := c.httpPut("project/"+projectName+"/config", withBody(bytes.NewReader(body)), requestExpects(200), requestJSON())
	if err != nil {
		return nil, err
	}
	data := map[string]string{}
	if jsonErr := json.Unmarshal(res, &data); jsonErr != nil {
		return nil, &UnmarshalError{msg: multierror.Append(errEncoding, jsonErr).Error()}
	}
	return data, nil
}

// GetProjectConfigurationKey gets a specific configuration key
// http://rundeck.org/docs/api/index.html#get-project-configuration-key
func (c *Client) GetProjectConfigurationKey(projectName, key string) (string, error) {
	if err := c.checkRequiredAPIVersion(responses.ProjectConfigItemResponse{}); err != nil {
		return "", err
	}
	u := fmt.Sprintf("project/%s/config/%s", projectName, key)
	res, err := c.httpGet(u, requestExpects(200), contentType("text/plain"))
	return string(res), err
}

// PutProjectConfigurationKey sets a value for a configuration key
// http://rundeck.org/docs/api/index.html#put-project-configuration-key
func (c *Client) PutProjectConfigurationKey(projectName, key, value string) error {
	if err := c.checkRequiredAPIVersion(responses.ProjectConfigItemResponse{}); err != nil {
		return err
	}
	u := fmt.Sprintf("project/%s/config/%s", projectName, key)
	_, err := c.httpPut(u, withBody(bytes.NewReader([]byte(value))), requestExpects(200), contentType("text/plain"))
	return err
}

// DeleteProjectConfigurationKey deletes a configuration key
// http://rundeck.org/docs/api/index.html#delete-project-configuration-key
func (c *Client) DeleteProjectConfigurationKey(projectName, key string) error {
	if err := c.checkRequiredAPIVersion(responses.ProjectConfigItemResponse{}); err != nil {
		return err
	}
	u := fmt.Sprintf("project/%s/config/%s", projectName, key)
	_, err := c.httpDelete(u, requestExpects(204))
	return err
}

// GetProjectArchiveExport export exports a zip file of the project
// http://rundeck.org/docs/api/index.html#project-archive-export
func (c *Client) GetProjectArchiveExport(p string, w io.Writer, opts ...ProjectExportOption) error {
	if err := c.checkRequiredAPIVersion(responses.ProjectInfoResponse{}); err != nil {
		return err
	}
	params := &map[string]string{}
	if len(opts) == 0 {
		opts = append(opts, ProjectExportAll(true))
	}
	for _, opt := range opts {
		if err := opt(params); err != nil {
			return &OptionError{msg: multierror.Append(errOption, err).Error()}
		}
	}

	u := fmt.Sprintf("project/%s/export", p)
	res, resErr := c.httpGet(u, requestExpects(200), accept("application/zip"), queryParams(*params))
	if resErr != nil {
		return resErr
	}
	if _, wErr := w.Write(res); wErr != nil {
		return wErr
	}
	return nil
}

// GetProjectArchiveExportAsync export a zip archive of a project async
// http://rundeck.org/docs/api/index.html#project-archive-export-async
func (c *Client) GetProjectArchiveExportAsync(p string, opts ...ProjectExportOption) (string, error) {
	if err := c.checkRequiredAPIVersion(responses.ProjectInfoResponse{}); err != nil {
		return "", err
	}
	params := &map[string]string{}
	if len(opts) == 0 {
		opts = append(opts, ProjectExportAll(true))
	}
	for _, opt := range opts {
		if err := opt(params); err != nil {
			return "", &OptionError{msg: multierror.Append(errOption, err).Error()}
		}
	}

	u := fmt.Sprintf("project/%s/export/async", p)
	res, resErr := c.httpGet(u, requestExpects(200), contentType("application/x-www-form-urlencoded"), accept("application/json"), queryParams(*params))
	if resErr != nil {
		return "", resErr
	}
	data := responses.ProjectArchiveExportAsyncResponse{}
	if jsonErr := json.Unmarshal(res, &data); jsonErr != nil {
		return "", &UnmarshalError{msg: multierror.Append(errEncoding, jsonErr).Error()}
	}
	return data.Token, nil
}

// GetProjectArchiveExportAsyncStatus gets the status of an async export request
// http://rundeck.org/docs/api/index.html#project-archive-export-async-status
func (c *Client) GetProjectArchiveExportAsyncStatus(p, token string) (*responses.ProjectArchiveExportAsyncResponse, error) {
	if err := c.checkRequiredAPIVersion(responses.ProjectInfoResponse{}); err != nil {
		return nil, err
	}

	u := fmt.Sprintf("project/%s/export/status/%s", p, token)
	res, resErr := c.httpGet(u, requestExpects(200), requestJSON())
	if resErr != nil {
		return nil, resErr
	}
	data := responses.ProjectArchiveExportAsyncResponse{}
	if jsonErr := json.Unmarshal(res, &data); jsonErr != nil {
		return nil, &UnmarshalError{msg: multierror.Append(errEncoding, jsonErr).Error()}
	}
	return &data, nil
}

// GetProjectArchiveExportAsyncDownload downloads an async project export archive file
// http://rundeck.org/docs/api/index.html#project-archive-export-async-download
func (c *Client) GetProjectArchiveExportAsyncDownload(p, token string, w io.Writer) error {
	if err := c.checkRequiredAPIVersion(responses.ProjectInfoResponse{}); err != nil {
		return err
	}

	u := fmt.Sprintf("project/%s/export/download/%s", p, token)
	res, resErr := c.httpGet(u, requestExpects(200), accept("application/zip"))
	if resErr != nil {
		return resErr
	}
	if _, wErr := w.Write(res); wErr != nil {
		return wErr
	}
	return nil
}

// ProjectArchiveImport imports a zip archive to a project
// http://rundeck.org/docs/api/index.html#project-archive-import
func (c *Client) ProjectArchiveImport(projectName string, f io.Reader, opts ...ProjectImportOption) (*responses.ProjectImportArchiveResponse, error) {
	if err := c.checkRequiredAPIVersion(responses.ProjectImportArchiveResponse{}); err != nil {
		return nil, err
	}
	// path: project/[PROJECT]/import{?jobUuidOption,importExecutions,importConfig,importACL}
	u := "project/" + projectName + "/import"
	params := &map[string]string{}
	for _, opt := range opts {
		if err := opt(params); err != nil {
			return nil, &OptionError{msg: multierror.Append(errOption, err).Error()}
		}
	}
	res, resErr := c.httpPut(u,
		withBody(f),
		contentType("application/zip"),
		accept("application/json"),
		queryParams(*params),
		requestExpects(200))
	if resErr != nil {
		return nil, resErr
	}
	results := &responses.ProjectImportArchiveResponse{}
	if jsonErr := json.Unmarshal(res, results); jsonErr != nil {
		return nil, &UnmarshalError{msg: multierror.Append(errEncoding, jsonErr).Error()}
	}
	return results, nil
}
