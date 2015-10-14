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

func (c *RundeckClient) ListProjects() (data Projects, err error) {
	var res []byte
	err = c.Get(&res, "projects", nil)
	if err != nil {
		return data, err
	} else {
		xml.Unmarshal(res, &data)
		return data, nil
	}
}
