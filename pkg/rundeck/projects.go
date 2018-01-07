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

// GetProjectInfo gets a project by name
// http://rundeck.org/docs/api/index.html#getting-project-info
func (c *Client) GetProjectInfo(name string) (*Project, error) {
	if _, err := c.hasRequiredAPIVersion(minJSONSupportedAPIVersion, maxRundeckVersionInt); err != nil {
		return nil, err
	}
	p := &responses.ProjectInfoResponse{}
	res, err := c.httpGet("project/"+name, requestJSON(), requestExpects(200))
	if err != nil {
		return nil, err
	}
	if jsonErr := json.Unmarshal(res, &p); jsonErr != nil {
		return nil, errDecoding
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
func (c *Client) ListProjects() (*Projects, error) {
	if _, err := c.hasRequiredAPIVersion(minJSONSupportedAPIVersion, maxRundeckVersionInt); err != nil {
		return nil, err
	}
	data := &responses.ListProjectsResponse{}
	res, err := c.httpGet("projects", requestJSON(), requestExpects(200))
	if err != nil {
		return nil, err
	}
	if jsonErr := json.Unmarshal(res, &data); jsonErr != nil {
		return nil, errDecoding
	}
	projects := &Projects{}
	for _, p := range *data {
		*projects = append(*projects, &Project{
			URL:         p.URL,
			Name:        p.Name,
			Description: p.Description,
		})
	}
	return projects, nil
}

// CreateProject makes a project
// http://rundeck.org/docs/api/index.html#project-creation
func (c *Client) CreateProject(name string, properties map[string]string) (*Project, error) {
	if _, err := c.hasRequiredAPIVersion(minJSONSupportedAPIVersion, maxRundeckVersionInt); err != nil {
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
		return nil, errDecoding
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
	if _, err := c.hasRequiredAPIVersion(minJSONSupportedAPIVersion, maxRundeckVersionInt); err != nil {
		return err
	}
	url := fmt.Sprintf("project/%s", p)
	return c.httpDelete(url, requestJSON(), requestExpects(204))
}

// GetProjectConfiguration gets a project's configuration
// http://rundeck.org/docs/api/index.html#get-project-configuration
func (c *Client) GetProjectConfiguration(p string) (*responses.ProjectConfigResponse, error) {
	if _, err := c.hasRequiredAPIVersion(minJSONSupportedAPIVersion, maxRundeckVersionInt); err != nil {
		return nil, err
	}
	data := &responses.ProjectConfigResponse{}
	res, err := c.httpGet("project/"+p+"/config", requestExpects(200), requestJSON())
	if err != nil {
		return nil, err
	}
	if jsonErr := json.Unmarshal(res, data); jsonErr != nil {
		return nil, &UnmarshalError{msg: multierror.Append(errEncoding, jsonErr).Error()}
	}
	return data, nil
}

// PutProjectConfiguration replaces all configuration data with the submitted values
// http://rundeck.org/docs/api/index.html#put-project-configuration
func (c *Client) PutProjectConfiguration() error {
	if _, err := c.hasRequiredAPIVersion(minJSONSupportedAPIVersion, maxRundeckVersionInt); err != nil {
		return err
	}
	return fmt.Errorf("not yet implemented")
}

// GetProjectConfigurationKey gets a specific configuration key
// http://rundeck.org/docs/api/index.html#get-project-configuration-key
func (c *Client) GetProjectConfigurationKey() error {
	if _, err := c.hasRequiredAPIVersion(minJSONSupportedAPIVersion, maxRundeckVersionInt); err != nil {
		return err
	}
	return fmt.Errorf("not yet implemented")
}

// PutProjectConfigurationKey sets a value for a configuration key
// http://rundeck.org/docs/api/index.html#put-project-configuration-key
func (c *Client) PutProjectConfigurationKey() error {
	if _, err := c.hasRequiredAPIVersion(minJSONSupportedAPIVersion, maxRundeckVersionInt); err != nil {
		return err
	}
	return fmt.Errorf("not yet implemented")
}

// DeleteProjectConfigurationKey deletes a configuration key
// http://rundeck.org/docs/api/index.html#delete-project-configuration-key
func (c *Client) DeleteProjectConfigurationKey() error {
	if _, err := c.hasRequiredAPIVersion(minJSONSupportedAPIVersion, maxRundeckVersionInt); err != nil {
		return err
	}
	return fmt.Errorf("not yet implemented")
}

// GetProjectArchiveExport export exports a zip file of the project
// http://rundeck.org/docs/api/index.html#project-archive-export
func (c *Client) GetProjectArchiveExport(p string, w io.Writer, opts ...ProjectExportOption) error {
	if _, err := c.hasRequiredAPIVersion(19, maxRundeckVersionInt); err != nil {
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
	res, resErr := c.httpGet("project/"+p+"/export", requestExpects(200), queryParams(*params))
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
func (c *Client) GetProjectArchiveExportAsync() error {
	if _, err := c.hasRequiredAPIVersion(19, maxRundeckVersionInt); err != nil {
		return err
	}
	return fmt.Errorf("not yet implemented")
}

// GetProjectArchiveExportAsyncStatus gets the status of an async export request
// http://rundeck.org/docs/api/index.html#project-archive-export-async-status
func (c *Client) GetProjectArchiveExportAsyncStatus() error {
	if _, err := c.hasRequiredAPIVersion(19, maxRundeckVersionInt); err != nil {
		return err
	}
	return fmt.Errorf("not yet implemented")
}

// GetProjectArchiveExportAsyncDownload downloads an async project export archive file
// http://rundeck.org/docs/api/index.html#project-archive-export-async-download
func (c *Client) GetProjectArchiveExportAsyncDownload() error {
	if _, err := c.hasRequiredAPIVersion(19, maxRundeckVersionInt); err != nil {
		return err
	}
	return fmt.Errorf("not yet implemented")
}

// ProjectArchiveImport imports a zip archive to a project
// http://rundeck.org/docs/api/index.html#project-archive-import
func (c *Client) ProjectArchiveImport() error {
	if _, err := c.hasRequiredAPIVersion(19, maxRundeckVersionInt); err != nil {
		return err
	}
	return fmt.Errorf("not yet implemented")
}
