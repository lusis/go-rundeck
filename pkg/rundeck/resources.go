package rundeck

import (
	"encoding/json"
	"fmt"

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
	if _, err := c.hasRequiredAPIVersion(minJSONSupportedAPIVersion, maxRundeckVersionInt); err != nil {
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
	if _, err := c.hasRequiredAPIVersion(minJSONSupportedAPIVersion, maxRundeckVersionInt); err != nil {
		return nil, err
	}
	r := Resource{}
	data, err := c.httpGet("project/"+p+"/resources/"+n, requestJSON(), requestExpects(200))
	if err != nil {
		return nil, err
	}
	if jsonErr := json.Unmarshal(data, &r); jsonErr != nil {
		return nil, &UnmarshalError{msg: multierror.Append(errDecoding, jsonErr).Error()}
	}
	return r[n], nil

}

// UpdateResource updates a project resource
// http://rundeck.org/docs/api/index.html#list-resources-for-a-project
func (c *Client) UpdateResource() error {
	if _, err := c.hasRequiredAPIVersion(19, maxRundeckVersionInt); err != nil {
		return err
	}
	return fmt.Errorf("not yet implemented")
}

// GetProjectReadme gets a project's readme.md
// http://rundeck.org/docs/api/index.html#get-readme-file
func (c *Client) GetProjectReadme() error {
	if _, err := c.hasRequiredAPIVersion(19, maxRundeckVersionInt); err != nil {
		return err
	}
	return fmt.Errorf("not yet implemented")
}

// PutProjectReadme creates or modifies a project's readme.md
// http://rundeck.org/docs/api/index.html#put-readme-file
func (c *Client) PutProjectReadme() error {
	if _, err := c.hasRequiredAPIVersion(19, maxRundeckVersionInt); err != nil {
		return err
	}
	return fmt.Errorf("not yet implemented")
}

// DeleteProjectReadme deletes a project's readme.md
func (c *Client) DeleteProjectReadme() error {
	if _, err := c.hasRequiredAPIVersion(19, maxRundeckVersionInt); err != nil {
		return err
	}
	return fmt.Errorf("not yet implemented")
}

// GetProjectMotd gets a project's Motd.md
// http://rundeck.org/docs/api/index.html#get-readme-file
func (c *Client) GetProjectMotd() error {
	if _, err := c.hasRequiredAPIVersion(19, maxRundeckVersionInt); err != nil {
		return err
	}
	return fmt.Errorf("not yet implemented")
}

// PutProjectMotd creates or modifies a project's motd.md
// http://rundeck.org/docs/api/index.html#put-readme-file
func (c *Client) PutProjectMotd() error {
	if _, err := c.hasRequiredAPIVersion(19, maxRundeckVersionInt); err != nil {
		return err
	}
	return fmt.Errorf("not yet implemented")
}

// DeleteProjectMotd deletes a project's motd.md
func (c *Client) DeleteProjectMotd() error {
	if _, err := c.hasRequiredAPIVersion(19, maxRundeckVersionInt); err != nil {
		return err
	}
	return fmt.Errorf("not yet implemented")
}
