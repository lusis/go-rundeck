package rundeck

import (
	"encoding/json"
	"io"

	multierror "github.com/hashicorp/go-multierror"
	"github.com/lusis/go-rundeck/pkg/rundeck/responses"
)

// Resources represents a collection of project resources (usually nodes)
type Resources responses.ResourceCollectionResponse

// Resource represents a project resource (usually a node)
type Resource responses.ResourceResponse

// ListResourcesForProject returns resources for a project (usually nodes)
// http://rundeck.org/docs/api/index.html#list-resources-for-a-project
func (c *Client) ListResourcesForProject(p string) (*Resources, error) {
	if err := c.checkRequiredAPIVersion(responses.ResourceCollectionResponse{}); err != nil {
		return nil, err
	}
	ls := &Resources{}
	data, err := c.httpGet("project/"+p+"/resources", requestJSON(), requestExpects(200))
	if err != nil {
		return nil, err
	}
	if jsonErr := json.Unmarshal(data, &ls); jsonErr != nil {
		return nil, &UnmarshalError{msg: multierror.Append(errDecoding, jsonErr).Error()}
	}
	return ls, nil
}

// GetResourceInfo get a specific resource within a project (usually a node)
// http://rundeck.org/docs/api/index.html#getting-resource-info
func (c *Client) GetResourceInfo(p, n string) (*responses.ResourceDetailResponse, error) {
	if err := c.checkRequiredAPIVersion(responses.ResourceDetailResponse{}); err != nil {
		return nil, err
	}
	r := Resource{}
	data, err := c.httpGet("project/"+p+"/resource/"+n, requestJSON(), requestExpects(200))
	if err != nil {
		return nil, err
	}
	if jsonErr := json.Unmarshal(data, &r); jsonErr != nil {
		return nil, &UnmarshalError{msg: multierror.Append(errDecoding, jsonErr).Error()}
	}
	return r[n], nil

}

// GetProjectReadme gets a project's readme.md
// http://rundeck.org/docs/api/index.html#get-readme-file
func (c *Client) GetProjectReadme(projectName string) (string, error) {
	if err := c.checkRequiredAPIVersion(responses.ResourceResponse{}); err != nil {
		return "", err
	}
	data, err := c.httpGet("project/"+projectName+"/readme.md", accept("text/plain"), requestExpects(200))
	if err != nil {
		return "", err
	}
	return string(data), nil
}

// PutProjectReadme creates or modifies a project's readme.md
// http://rundeck.org/docs/api/index.html#put-readme-file
func (c *Client) PutProjectReadme(projectName string, readme io.Reader) error {
	if err := c.checkRequiredAPIVersion(responses.ResourceResponse{}); err != nil {
		return err
	}
	_, err := c.httpPut("project/"+projectName+"/readme.md", withBody(readme), requestExpects(200), contentType("text/plain"))
	return err
}

// DeleteProjectReadme deletes a project's readme.md
func (c *Client) DeleteProjectReadme(projectName string) error {
	if err := c.checkRequiredAPIVersion(responses.ResourceResponse{}); err != nil {
		return err
	}
	return c.httpDelete("project/"+projectName+"/readme.md", requestExpects(204))
}

// GetProjectMotd gets a project's Motd.md
// http://rundeck.org/docs/api/index.html#get-readme-file
func (c *Client) GetProjectMotd(projectName string) (string, error) {
	if err := c.checkRequiredAPIVersion(responses.ResourceResponse{}); err != nil {
		return "", err
	}
	data, err := c.httpGet("project/"+projectName+"/motd.md", accept("text/plain"), requestExpects(200))
	if err != nil {
		return "", err
	}
	return string(data), nil
}

// PutProjectMotd creates or modifies a project's motd.md
// http://rundeck.org/docs/api/index.html#put-readme-file
func (c *Client) PutProjectMotd(projectName string, motd io.Reader) error {
	if err := c.checkRequiredAPIVersion(responses.ResourceResponse{}); err != nil {
		return err
	}
	_, err := c.httpPut("project/"+projectName+"/motd.md", withBody(motd), requestExpects(200), contentType("text/plain"))
	return err
}

// DeleteProjectMotd deletes a project's motd.md
func (c *Client) DeleteProjectMotd(projectName string) error {
	if err := c.checkRequiredAPIVersion(responses.ResourceResponse{}); err != nil {
		return err
	}
	return c.httpDelete("project/"+projectName+"/motd.md", requestExpects(204))
}
