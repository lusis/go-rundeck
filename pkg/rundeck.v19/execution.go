package rundeck

import "encoding/xml"

// Execution represents a job execution
type Execution struct {
	XMLName         xml.Name `xml:"execution"`
	ID              string   `xml:"id,attr"`
	HRef            string   `xml:"href,attr"`
	Status          string   `xml:"status,attr"`
	CustomStatus    string   `xml:"customStatus,omitempty"`
	Project         string   `xml:"project,attr"`
	User            string   `xml:"user"`
	DateEnded       string   `xml:"date-ended,omitempty"`
	UnixTimeEnded   int64    `xml:"date-ended,unixtime,attr,omitempty"`
	DateStarted     string   `xml:"date-started,omitempty"`
	UnixTimeStarted int64    `xml:"date-started,unixtime,attr,omitempty"`
	Job             *Job     `xml:"job"`
	Description     string   `xml:"description,omitempty"`
	SuccessfulNodes Nodes    `xml:"successfulNodes,omitempty"`
	FailedNodes     Nodes    `xml:"failedNodes,omitempty"`
}

// ExecutionDateTime represents an execution timestamp in unixtime format
type ExecutionDateTime struct {
	UnixTime int64 `xml:"unixtime,attr"`
}

// ExecutionOutput represents the output of an execution
type ExecutionOutput struct {
	XMLName        xml.Name               `xml:"output"`
	ID             int64                  `xml:"id"`
	Offset         int64                  `xml:"offset"`
	Completed      bool                   `xml:"completed"`
	ExecCompleted  bool                   `xml:"execCompleted"`
	HasFailedNodes bool                   `xml:"hasFailedNodes"`
	ExecState      string                 `xml:"execState"`
	LastModified   ExecutionDateTime      `xml:"lastModified"`
	ExecDuration   int64                  `xml:"execDuration"`
	TotalSize      int64                  `xml:"totalSize"`
	Entries        ExecutionOutputEntries `xml:"entries"`
}

// ExecutionOutputEntries is a collection of `ExecutionOutputEntry`
type ExecutionOutputEntries struct {
	Entry []ExecutionOutputEntry `xml:"entry"`
}

// ExecutionOutputEntry represents a single execution output entry
type ExecutionOutputEntry struct {
	XMLName      xml.Name
	Time         string `xml:"time,attr"`
	AbsoluteTime string `xml:"absolute_time,attr"`
	Log          string `xml:"log,attr"`
	Level        string `xml:"level,attr"`
	User         string `xml:"user,attr"`
	Command      string `xml:"command,attr"`
	Stepctx      string `xml:"stepctx,attr"`
	Node         string `xml:"node,attr"`
}

// ExecutionState represents an execution state
type ExecutionState struct {
	XMLName     xml.Name        `xml:"result"`
	Success     bool            `xml:"success,attr"`
	APIVersion  int64           `xml:"apiversion,attr"`
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

// GetExecution returns the details of a job execution
func (c *Client) GetExecution(executionID string) (exec Execution, err error) {
	var res []byte
	var execs Executions
	err = c.Get(&res, "execution/"+executionID, nil)
	xmlerr := xml.Unmarshal(res, &execs)
	if xmlerr != nil {
		return exec, xmlerr
	}
	return execs.Executions[0], err
}

// GetExecutionState returns the state of an execution
func (c *Client) GetExecutionState(executionID string) (ExecutionState, error) {
	u := make(map[string]string)
	var res []byte
	var data ExecutionState
	err := c.Get(&res, "execution/"+executionID+"/state", u)
	if xmlErr := xml.Unmarshal(res, &data); xmlErr != nil {
		return data, xmlErr
	}
	return data, err
}

// GetExecutionOutput returns the output of an execution
func (c *Client) GetExecutionOutput(executionID string) (ExecutionOutput, error) {
	u := make(map[string]string)
	var res []byte
	var data ExecutionOutput
	err := c.Get(&res, "execution/"+executionID+"/output", u)
	if xmlErr := xml.Unmarshal(res, &data); xmlErr != nil {
		return data, xmlErr
	}
	return data, err
}

// DeleteExecution deletes an execution
func (c *Client) DeleteExecution(id string) error {
	return c.Delete("execution/"+id, nil)
}

// DisableExecution disables an execution
func (c *Client) DisableExecution(id string) error {
	var res []byte
	err := c.Post(&res, "job/"+id+"/execution/disable", nil, nil)
	return err

}

// EnableExecution enables an execution
func (c *Client) EnableExecution(id string) error {
	var res []byte
	err := c.Post(&res, "job/"+id+"/execution/enable", nil, nil)
	return err
}
