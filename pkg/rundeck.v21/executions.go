package rundeck

import (
	"encoding/xml"
	"errors"
	"strconv"
	"strings"
)

// Executions represents a collection of `Execution`
type Executions struct {
	Count      int64        `xml:"count,attr"`
	Total      int64        `xml:"total,attr"`
	Max        int64        `xml:"max,attr"`
	Offset     int64        `xml:"offset,attr"`
	Executions []*Execution `xml:"execution"`
}

// ExecutionStep represents an execution step
type ExecutionStep struct {
	XMLName        xml.Name     `xml:"step"`
	StepCtx        int64        `xml:"stepctx,attr"`
	ID             int64        `xml:"id,attr"`
	StartTime      string       `xml:"startTime"`
	UpdateTime     string       `xml:"updateTime"`
	EndTime        string       `xml:"endTime"`
	ExecutionState string       `xml:"executionState"`
	NodeStep       bool         `xml:"nodeStep"`
	NodeStates     []*NodeState `xml:"nodeStates>nodeState"`
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
		XMLName  xml.Name                 `xml:"failed"`
		Count    int64                    `xml:"count,attr"`
		Failures []*FailedExecutionDelete `xml:"execution,omitempty"`
	} `xml:"failed"`
}

// FailedExecutionDelete represents a failed execution delete
type FailedExecutionDelete struct {
	XMLName xml.Name `xml:"execution"`
	ID      int64    `xml:"id,attr"`
	Message string   `xml:"message,attr"`
}

// ListProjectExecutions lists a projects executions
func (c *Client) ListProjectExecutions(projectID string, options map[string]string) (*Executions, error) {
	options["project"] = projectID
	data := &Executions{}
	res, err := c.httpGet("executions", requestXML(), queryParams(options))
	if err != nil {
		return nil, err
	}
	xmlErr := xml.Unmarshal(res, &data)
	return data, xmlErr
}

// ListRunningExecutions lists running executions
func (c *Client) ListRunningExecutions(projectID string) (*Executions, error) {
	options := make(map[string]string)
	options["project"] = projectID
	executions := &Executions{}
	res, err := c.httpGet("executions/running", requestXML(), queryParams(options))
	if err != nil {
		return nil, err
	}
	xmlErr := xml.Unmarshal(res, &executions)
	return executions, xmlErr
}

// DeleteExecutions deletes a list of executions by id
func (c *Client) DeleteExecutions(ids []string) (*ExecutionsDeleted, error) {
	data := &ExecutionsDeleted{}
	opts := make(map[string]string)
	opts["ids"] = strings.Join(ids, ",")

	res, err := c.httpPost("executions/delete", requestXML(), queryParams(opts))
	if err != nil {
		return nil, err
	}
	xmlErr := xml.Unmarshal(res, &data)
	return data, xmlErr
}

// DeleteAllExecutionsForProject deletes all executions for a project up to the max (default: 10)
func (c *Client) DeleteAllExecutionsForProject(project string, max int64) (*ExecutionsDeleted, error) {
	data := &ExecutionsDeleted{}
	eopts := make(map[string]string)
	eopts["max"] = strconv.FormatInt(max, 10)
	e, err := c.ListProjectExecutions(project, eopts)
	if err != nil {
		return nil, err
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
	res, resErr := c.httpPost("executions/delete", requestXML(), queryParams(opts))
	if resErr != nil {
		return nil, resErr
	}
	xmlErr := xml.Unmarshal(res, &data)
	return data, xmlErr
}
