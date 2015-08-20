package rundeck

import "encoding/xml"

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

type ExecutionOutputEntries struct {
	Entry []ExecutionOutputEntry `xml:"entry"`
}

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

func (c *RundeckClient) GetExecutionState(executionId string) (ExecutionState, error) {
	u := make(map[string]string)
	var data ExecutionState
	err := c.Get(&data, "execution/"+executionId+"/state", u)
	return data, err
}

func (c *RundeckClient) GetExecutionOutput(executionId string) (ExecutionOutput, error) {
	u := make(map[string]string)
	var data ExecutionOutput
	err := c.Get(&data, "execution/"+executionId+"/output", u)
	return data, err
}
