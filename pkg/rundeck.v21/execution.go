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
func (c *Client) GetExecution(executionID string) (*Execution, error) {
	execs := &Executions{}
	res, err := c.httpGet("execution/"+executionID, requestXML())
	if err != nil {
		return nil, err
	}
	xmlerr := xml.Unmarshal(res, &execs)
	if xmlerr != nil {
		return nil, xmlerr
	}
	return execs.Executions[0], nil
}

// GetExecutionState returns the state of an execution
func (c *Client) GetExecutionState(executionID string) (*ExecutionState, error) {
	data := &ExecutionState{}
	res, err := c.httpGet("execution/"+executionID+"/state", requestXML())
	if err != nil {
		return nil, err
	}
	if xmlErr := xml.Unmarshal(res, &data); xmlErr != nil {
		return nil, xmlErr
	}
	return data, nil
}

// GetExecutionOutput returns the output of an execution
func (c *Client) GetExecutionOutput(executionID string) (*ExecutionOutput, error) {
	data := &ExecutionOutput{}
	res, err := c.httpGet("execution/"+executionID+"/output", requestXML())
	if err != nil {
		return nil, err
	}
	if xmlErr := xml.Unmarshal(res, &data); xmlErr != nil {
		return nil, xmlErr
	}
	return data, err
}

// DeleteExecution deletes an execution
func (c *Client) DeleteExecution(id string) error {
	return c.httpDelete("execution/" + id)
}

// DisableExecution disables an execution
func (c *Client) DisableExecution(id string) error {
	_, err := c.httpPost("job/" + id + "/execution/disable")
	return err

}

// EnableExecution enables an execution
func (c *Client) EnableExecution(id string) error {
	_, err := c.httpPost("job/" + id + "/execution/enable")
	return err
}
