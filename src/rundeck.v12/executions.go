package rundeck

import (
	"encoding/xml"
	"errors"
	"strconv"
	"strings"
)

type Execution struct {
	XMLName         xml.Name          `xml:"execution"`
	ID              string            `xml:"id,attr"`
	HRef            string            `xml:"href,attr"`
	Status          string            `xml:"status,attr"`
	Project         string            `xml:"project,attr"`
	User            string            `xml:"user"`
	DateEnded       ExecutionDateTime `xml:"date-ended,omitempty"`
	DateStarted     ExecutionDateTime `xml:"date-started,omitempty"`
	Job             *Job              `xml:"job"`
	Description     string            `xml:"description,omitempty"`
	SuccessfulNodes Nodes             `xml:"successfulNodes,omitempty"`
	FailedNodes     Nodes             `xml:"failedNodes,omitempty"`
}

type ExecutionDateTime struct {
	UnixTime int64 `xml:"unixtime,attr"`
}

type Executions struct {
	Count      int64       `xml:"count,attr"`
	Total      int64       `xml:"total,attr"`
	Max        int64       `xml:"max,attr"`
	Offset     int64       `xml:"offset,attr"`
	Executions []Execution `xml:"execution"`
}

type ExecutionState struct {
	XMLName     xml.Name        `xml:"result"`
	Success     bool            `xml:"success,attr"`
	ApiVersion  int64           `xml:"apiversion,attr"`
	StartTime   string          `xml:"executionState>startTime"`
	StepCount   int64           `xml:"executionState>stepCount"`
	AllNodes    []Node          `xml:"executionState>allNodes>nodes>node,omitempty"`
	TargetNodes []Node          `xml:"executionState>targetNodes>nodes>node,omitempty"`
	ExecutionID int64           `xml:"executionState>executionId"`
	Completed   bool            `xml:"executionState>completed"`
	UpdateTime  string          `xml:"executionState>updateTime,omitempty"`
	Steps       []ExecutionStep `xml:"executionState>steps>step,omitempty"`
	Nodes       []NodeWithSteps `xml:"executionState>nodes>node"`
}

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

type FailedExecutionDelete struct {
	XMLName xml.Name `xml:"execution"`
	ID      int64    `xml:"id,attr"`
	Message string   `xml:"message,attr"`
}

func (c *RundeckClient) ListExecutions(projectId string, options map[string]string) (Executions, error) {
	options["project"] = projectId
	var data Executions
	err := c.Get(&data, "executions", options)
	return data, err
}

func (c *RundeckClient) ListRunningExecutions(projectId string) (executions Executions, err error) {
	options := make(map[string]string)
	options["project"] = projectId
	err = c.Get(executions, "executions/running", options)
	return executions, err
}

func (c *RundeckClient) DeleteExecutions(ids []string) (ExecutionsDeleted, error) {
	var data ExecutionsDeleted
	opts := make(map[string]string)
	opts["ids"] = strings.Join(ids, ",")
	err := c.Post(&data, "executions/delete", nil, opts)
	if err != nil {
		return data, err
	} else {
		return data, nil
	}
}

func (c *RundeckClient) DeleteAllExecutionsFor(project string, max int64) (ExecutionsDeleted, error) {
	var data ExecutionsDeleted
	eopts := make(map[string]string)
	eopts["max"] = strconv.FormatInt(max, 10)
	e, err := c.ListExecutions(project, eopts)
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
	err = c.Post(&data, "executions/delete", nil, opts)
	if err != nil {
		return data, err
	} else {
		return data, nil
	}
}
