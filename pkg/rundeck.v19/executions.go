package rundeck

import (
	"encoding/xml"
	"errors"
	"strconv"
	"strings"
)

// Executions represents a collection of `Execution`
type Executions struct {
	Count      int64       `xml:"count,attr"`
	Total      int64       `xml:"total,attr"`
	Max        int64       `xml:"max,attr"`
	Offset     int64       `xml:"offset,attr"`
	Executions []Execution `xml:"execution"`
}

// ExecutionStep represents an execution step
type ExecutionStep struct {
	XMLName        xml.Name    `xml:"step"`
	StepCtx        int64       `xml:"stepctx,attr"`
	ID             int64       `xml:"id,attr"`
	StartTime      string      `xml:"startTime"`
	UpdateTime     string      `xml:"updateTime"`
	EndTime        string      `xml:"endTime"`
	ExecutionState string      `xml:"executionState"`
	NodeStep       bool        `xml:"nodeStep"`
	NodeStates     []NodeState `xml:"nodeStates>nodeState"`
}

// ExecutionsDeleted represents a deleted Executions
type ExecutionsDeleted struct {
	XMLName       xml.Name `xml:"deleteExecutions"`
	RequestCount  int64    `xml:"requestCount,attr"`
	AllSuccessful bool     `xml:"allSuccessful,attr"`
	Successful    struct {
		XMLName xml.Name `xml:"successful"`
		Count   int64    `xml:"count,attr"`
	} `xml:"successful"`
	Failed struct {
		XMLName  xml.Name                `xml:"failed"`
		Count    int64                   `xml:"count,attr"`
		Failures []FailedExecutionDelete `xml:"execution,omitempty"`
	} `xml:"failed"`
}

// FailedExecutionDelete represents a failed execution delete
type FailedExecutionDelete struct {
	XMLName xml.Name `xml:"execution"`
	ID      int64    `xml:"id,attr"`
	Message string   `xml:"message,attr"`
}

// ListProjectExecutions lists a projects executions
func (c *Client) ListProjectExecutions(projectID string, options map[string]string) (Executions, error) {
	var res []byte
	options["project"] = projectID
	var data Executions
	if err := c.Get(&res, "executions", options); err != nil {
		return data, err
	}
	xmlErr := xml.Unmarshal(res, &data)
	return data, xmlErr
}

// ListRunningExecutions lists running executions
func (c *Client) ListRunningExecutions(projectID string) (executions Executions, err error) {
	options := make(map[string]string)
	options["project"] = projectID
	var res []byte
	if err := c.Get(&res, "executions/running", options); err != nil {
		return executions, err
	}
	xmlErr := xml.Unmarshal(res, &executions)
	return executions, xmlErr
}

// DeleteExecutions deletes a list of executions by id
func (c *Client) DeleteExecutions(ids []string) (ExecutionsDeleted, error) {
	var res []byte
	var data ExecutionsDeleted
	opts := make(map[string]string)
	opts["ids"] = strings.Join(ids, ",")
	if err := c.Post(&res, "executions/delete", nil, opts); err != nil {
		return data, err
	}
	xmlErr := xml.Unmarshal(res, &data)
	return data, xmlErr
}

// DeleteAllExecutionsForProject deletes all executions for a project up to the max (default: 10)
func (c *Client) DeleteAllExecutionsForProject(project string, max int64) (ExecutionsDeleted, error) {
	var data ExecutionsDeleted
	eopts := make(map[string]string)
	eopts["max"] = strconv.FormatInt(max, 10)
	e, err := c.ListProjectExecutions(project, eopts)
	if err != nil {
		return data, err
	}

	var toDelete []string
	for _, execution := range e.Executions {
		toDelete = append(toDelete, execution.ID)
	}
	if len(toDelete) == 0 {
		return data, errors.New("No executions found for project: " + project)
	}
	opts := make(map[string]string)
	opts["ids"] = strings.Join(toDelete, ",")
	var res []byte
	if err := c.Post(&res, "executions/delete", nil, opts); err != nil {
		return data, err
	}
	xmlErr := xml.Unmarshal(res, &data)
	return data, xmlErr
}
