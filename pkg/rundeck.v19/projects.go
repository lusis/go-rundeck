package rundeck

import (
	"encoding/xml"
	"fmt"
)

// Project represents a project
type Project struct {
	XMLName     xml.Name `xml:"project"`
	Name        string   `xml:"name"`
	Description string   `xml:"description,omitempty"`
	URL         string   `xml:"url,attr"`
}

// Projects is a collection of `Project`
type Projects struct {
	Count    int64     `xml:"count,attr"`
	Projects []Project `xml:"project"`
}

// NewProject represents a new project
type NewProject struct {
	XMLName     xml.Name         `xml:"project"`
	Name        string           `xml:"name"`
	Description string           `xml:"description"`
	Config      []ConfigProperty `xml:"config>property,omitempty"`
}

// ConfigProperty is a configuration property
type ConfigProperty struct {
	XMLName xml.Name `xml:"property"`
	Key     string   `xml:"key,attr"`
	Value   string   `xml:"value,attr"`
}

// GetProject gets a project by name
func (c *Client) GetProject(name string) (p Project, err error) {
	var res []byte
	err = c.Get(&res, "project/"+name, nil)
	if err != nil {
		return p, err
	}
	xmlErr := xml.Unmarshal(res, &p)
	return p, xmlErr
}

// ListProjects lists all projects
func (c *Client) ListProjects() (data Projects, err error) {
	var res []byte
	err = c.Get(&res, "projects", nil)
	if err != nil {
		return data, err
	}
	xmlErr := xml.Unmarshal(res, &data)
	return data, xmlErr
}

// MakeProject makes a project
func (c *Client) MakeProject(p NewProject) error {
	var res []byte
	data, err := xml.Marshal(p)
	if err != nil {
		return err
	}
	err = c.Post(&res, "projects", data, nil)
	return err
}

// DeleteProject deletes a project
func (c *Client) DeleteProject(p string) error {
	url := fmt.Sprintf("project/%s", p)
	err := c.Delete(url, nil)
	return err
}
