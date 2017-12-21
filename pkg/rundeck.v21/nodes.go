package rundeck

import (
	"encoding/json"
	"encoding/xml"
)

// Nodes is a node collection
type Nodes map[string]*Node

// ArbitraryNodeProperties represents user defined node properties
type ArbitraryNodeProperties map[string]string

// Node represents a node
type Node struct {
	NodeName          string `json:"nodename"`
	Tags              string `json:"tags,omitempty"`
	OsFamily          string `json:"osFamily,omitempty"`
	OsVersion         string `json:"osVersion,omitempty"`
	OsArch            string `json:"osArch,omitempty"`
	OsName            string `json:"osName,omitempty"`
	SSHKeyStoragePath string `json:"ssh-key-storage-path,omitempty"`
	UserName          string `json:"username"`
	Description       string `json:"description,omitempty"`
	HostName          string `json:"hostname,omitempty"`
	FileCopier        string `json:"file-copier,omitempty"`
	NodeExectutor     string `json:"node-executor,omitempty"`
	RemoteURL         string `json:"remoteUrl,omitempty"`
	EditURL           string `json:"editUrl,omitempty"`
	ArbitraryNodeProperties
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
	XMLName xml.Name    `xml:"node"`
	Name    string      `xml:"name,attr"`
	Steps   []*NodeStep `xml:"steps>step"`
}

// ListNodes lists nodes
func (c *Client) ListNodes(projectID string) (*Nodes, error) {
	data := &Nodes{}
	res, err := c.httpGet("project/"+projectID+"/resources", requestJSON())
	if err != nil {
		return nil, err
	}
	jsonErr := json.Unmarshal(res, &data)
	return data, jsonErr
}
