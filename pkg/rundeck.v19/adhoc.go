package rundeck

import (
	"encoding/xml"
	"strings"
)

// ExecutionID represents an execution id
type ExecutionID struct {
	ID string `xml:"id,attr"`
}

// RunAdhoc runs an adhoc job
func (c *RundeckClient) RunAdhoc(projectID string, exec string, nodeFilter string) (ExecutionID, error) {
	options := make(map[string]string)
	options["project"] = projectID
	options["exec"] = exec
	n := strings.Split(nodeFilter, " ")
	for _, i := range n {
		f := strings.Split(i, "=")
		k, v := f[0], f[1]
		options[k] = v
	}
	var res []byte
	var data ExecutionID
	err := c.Get(&res, "run/command", options)
	if err != nil {
		return data, err
	}
	xmlErr := xml.Unmarshal(res, &data)
	return data, xmlErr
}
