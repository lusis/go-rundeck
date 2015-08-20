package rundeck

import (
	"strings"
)

type ExecutionId struct {
	ID string `xml:"id,attr"`
}

func (c *RundeckClient) RunAdhoc(projectId string, exec string, node_filter string) (ExecutionId, error) {
	options := make(map[string]string)
	options["project"] = projectId
	options["exec"] = exec
	n := strings.Split(node_filter, " ")
	for _, i := range n {
		f := strings.Split(i, "=")
		k, v := f[0], f[1]
		options[k] = v
	}
	var data ExecutionId
	err := c.Get(&data, "run/command", options)
	return data, err
}
