package rundeck

import "encoding/xml"

// Nodes is a collection of `Node`
type Nodes struct {
	Nodes []Node `xml:"node"`
}

// Node represents a node
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

// NodeState represents a node's state
type NodeState struct {
	XMLName        xml.Name `xml:"nodeState"`
	Name           string   `xml:"name,attr"`
	StartTime      string   `xml:"startTime"`
	UpdateTime     string   `xml:"updateTime"`
	EndTime        string   `xml:"endTime"`
	ExecutionState string   `xml:"executionState"`
}

// NodeStep represents a node step
type NodeStep struct {
	XMLName        xml.Name `xml:"step"`
	StepCtx        int64    `xml:"stepctx"`
	ExecutionState string   `xml:"executionState"`
}

// NodeWithSteps represents a node with its steps
type NodeWithSteps struct {
	XMLName xml.Name   `xml:"node"`
	Name    string     `xml:"name,attr"`
	Steps   []NodeStep `xml:"steps>step"`
}

// ListNodes lists nodes
func (c *Client) ListNodes(projectID string) (Nodes, error) {
	options := make(map[string]string)
	options["project"] = projectID
	var res []byte
	var data Nodes
	err := c.Get(&res, "resources", options)
	if err != nil {
		return data, err
	}
	xmlErr := xml.Unmarshal(res, &data)
	return data, xmlErr
}
