package rundeck

import (
	"encoding/xml"
)

type Project struct {
	XMLName     xml.Name `xml:"project"`
	Name        string   `xml:"name"`
	Description string   `xml:"description,omitempty"`
	Url         string   `xml:"url,attr"`
}

type Projects struct {
	Count    int64     `xml:"count,attr"`
	Projects []Project `xml:"project"`
}

type NewProject struct {
	XMLName xml.Name         `xml:"project"`
	Name    string           `xml:"name"`
	Config  []ConfigProperty `xml:"config,omitempty"`
}

type ConfigProperty struct {
	XMLName xml.Name `xml:"property"`
	Key     string   `xml:"key,attr"`
	Value   string   `xml:"value,attr"`
}

func (c *RundeckClient) ListProjects() (Projects, error) {
	options := make(map[string]string)
	var data Projects
	err := c.Get(&data, "projects", options)
	return data, err
}
