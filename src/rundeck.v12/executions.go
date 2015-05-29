package rundeck

import "encoding/xml"

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

func (c *RundeckClient) ListExecutions(projectId string, options map[string]string) (Executions, error) {
	options["project"] = projectId
	var data Executions
	err := c.Get(&data, "executions", options)
	return data, err
}

func (c *RundeckClient) GetExecutionState(executionId string) (ExecutionState, error) {
	u := make(map[string]string)
	var data ExecutionState
	err := c.Get(&data, "execution/"+executionId+"/state", u)
	return data, err
}
