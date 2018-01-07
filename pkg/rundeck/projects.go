package rundeck

import (
	"bytes"
	"encoding/json"
	"fmt"

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
func (c *Client) GetProjectConfiguration() error {
	if _, err := c.hasRequiredAPIVersion(minJSONSupportedAPIVersion, maxRundeckVersionInt); err != nil {
		return err
	}
	return fmt.Errorf("not yet implemented")
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
func (c *Client) GetProjectArchiveExport() error {
	if _, err := c.hasRequiredAPIVersion(19, maxRundeckVersionInt); err != nil {
		return err
	}
	return fmt.Errorf("not yet implemented")
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
