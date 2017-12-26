package rundeck

import (
	"encoding/json"

	"github.com/lusis/go-rundeck/pkg/rundeck.v21/responses"
)

// Resources represents a collection of project resources (usually nodes)
type Resources responses.ResourceCollectionResponse

// Resource represents a project resource (usually a node)
type Resource responses.ResourceResponse

// GetResources returns resources for a project (usually nodes)
func (c *Client) GetResources(p string) (*Resources, error) {
	ls := &Resources{}
	data, err := c.httpGet("project/"+p+"/resources", requestJSON())
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(data, &ls); err != nil {
		return nil, err
	}
	return ls, nil
}

// GetResource returns a single resource for the named project by resource name
func (c *Client) GetResource(p, n string) (*responses.ResourceDetailResponse, error) {
	r := Resource{}
	data, err := c.httpGet("project/"+p+"/resources/"+n, requestJSON())
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(data, &r); err != nil {
		return nil, err
	}
	return r[n], nil

}
