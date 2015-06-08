package rundeck

import "encoding/xml"

type Nodes struct {
	Nodes []Node `xml:"node"`
}

// TODO: Convert to a Basic Node that just has "name,attr"
type Node struct {
	XMLName     xml.Name `xml:"node"`
	Name        string   `xml:"name,attr"`
	Description string   `xml:"description,attr,omitempty"`
	Tags        string   `xml:"tags,attr,omitempty"`
	Hostname    string   `xml:"hostname,attr,omitempty"`
	OsArch      string   `xml:"osArch,attr,omitempty"`
	OsFamily    string   `xml:"osFamily,attr,omitempty"`
	OsName      string   `xml:"osName,attr,omitempty"`
	OsVersion   string   `xml:"osVersion,attr,omitempty"`
	Username    string   `xml:"username,attr,omitempty"`
}

type NodeState struct {
	XMLName        xml.Name `xml:"nodeState"`
	Name           string   `xml:"name,attr"`
	StartTime      string   `xml:"startTime"`
	UpdateTime     string   `xml:"updateTime"`
	EndTime        string   `xml:"endTime"`
	ExecutionState string   `xml:"executionState"`
}

type NodeStep struct {
	XMLName        xml.Name `xml:"step"`
	StepCtx        int64    `xml:"stepctx"`
	ExecutionState string   `xml:"executionState"`
}

type NodeWithSteps struct {
	XMLName xml.Name   `xml:"node"`
	Name    string     `xml:"name,attr"`
	Steps   []NodeStep `xml:"steps>step"`
}

func (c *RundeckClient) ListNodes(projectId string) (Nodes, error) {
	options := make(map[string]string)
	options["project"] = projectId
	var data Nodes
	err := c.Get(&data, "resources", options)
	return data, err
}
