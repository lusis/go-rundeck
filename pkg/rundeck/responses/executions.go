package responses

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"os"
)

// JobExecutionsResponse is the response for listing the executions for a job
type JobExecutionsResponse ListRunningExecutionsResponse

// ListRunningExecutionsResponseTestFile is the test data for JobExecutionResponse
const ListRunningExecutionsResponseTestFile = "executions.json"

// SuccessToggleResponseTestFile is the test data for a successful toggle
const SuccessToggleResponseTestFile = "success.json"

// FailToggleResponseTestFile is the test data for a successful toggle
const FailToggleResponseTestFile = "failed.json"

// ToggleResponse is the response for a toggled job, exeuction or schedule
type ToggleResponse struct {
	Success bool `json:"success"`
}

// ListRunningExecutionsResponse is the response for listing the running executions for a project
type ListRunningExecutionsResponse struct {
	Paging     PagingResponse      `json:"paging"`
	Executions []ExecutionResponse `json:"executions"`
}

// FromReader returns an ListRunningExecutionsResponse from an io.Reader
func (a *ListRunningExecutionsResponse) FromReader(i io.Reader) error {
	b, err := ioutil.ReadAll(i)
	if err != nil {
		return err
	}
	return json.Unmarshal(b, a)
}

// FromFile returns an ListRunningExectionsResponse from a file
func (a *ListRunningExecutionsResponse) FromFile(f string) error {
	file, err := os.Open(f)
	defer func() { _ = file.Close() }()
	if err != nil {
		return err
	}
	return a.FromReader(file)
}

// FromBytes returns a ListRunningExecutionsResponse from a byte slice
func (a *ListRunningExecutionsResponse) FromBytes(f []byte) error {
	file := bytes.NewReader(f)
	return a.FromReader(file)
}

// ExecutionResponseTestFile is the test data for ExecutionResponse
const ExecutionResponseTestFile = "execution.json"

// ExecutionResponse represents an individual execution response
type ExecutionResponse struct {
	ID           int    `json:"id"`
	HRef         string `json:"href"`
	Permalink    string `json:"permalink"`
	Status       string `json:"status"`
	CustomStatus string `json:"customStatus"`
	Project      string `json:"project"`
	User         string `json:"user"`
	ServerUUID   string `json:"serverUUID"`
	DateStarted  struct {
		UnixTime int64     `json:"unixtime"`
		Date     *JSONTime `json:"date"`
	} `json:"date-started"`
	DateEnded struct {
		UnixTime int64     `json:"unixtime"`
		Date     *JSONTime `json:"date"`
	} `json:"date-ended"`
	Job             ExecutionJobEntryResponse `json:"job"`
	Description     string                    `json:"description"`
	ArgString       string                    `json:"argstring"`
	SuccessfulNodes []string                  `json:"successfulNodes"`
	FailedNodes     []string                  `json:"failedNodes"`
}

// FromReader returns an ExecutionResponse from an io.Reader
func (a *ExecutionResponse) FromReader(i io.Reader) error {
	b, err := ioutil.ReadAll(i)
	if err != nil {
		return err
	}
	return json.Unmarshal(b, a)
}

// FromFile returns an ExectionResponse from a file
func (a *ExecutionResponse) FromFile(f string) error {
	file, err := os.Open(f)
	defer func() { _ = file.Close() }()
	if err != nil {
		return err
	}
	return a.FromReader(file)
}

// FromBytes returns a ExecutionResponse from a byte slice
func (a *ExecutionResponse) FromBytes(f []byte) error {
	file := bytes.NewReader(f)
	return a.FromReader(file)
}

// ExecutionJobEntryResponse represents an individual job execution entry response
type ExecutionJobEntryResponse struct {
	ID              string            `json:"id"`
	AverageDuration int64             `json:"averageDuration"`
	Name            string            `json:"name"`
	Group           string            `json:"group"`
	Project         string            `json:"project"`
	Description     string            `json:"description"`
	HRef            string            `json:"href"`
	Permalink       string            `json:"permalink"`
	Options         map[string]string `json:"options"`
}

// ExecutionInputFileResponse is an individual execution input file entry response
type ExecutionInputFileResponse struct {
	ID             string    `json:"id"`
	User           string    `json:"user"`
	FileState      string    `json:"fileState"`
	SHA            string    `json:"sha"`
	JobID          string    `json:"jobId"`
	DateCreated    *JSONTime `json:"dateCreated"`
	ServerNodeUUID string    `json:"serverNodeUUID"`
	FileName       string    `json:"fileName"`
	Size           int64     `json:"size"`
	ExpirationDate *JSONTime `json:"expirationDate"`
	ExecID         int       `json:"execId"`
}

// ExecutionInputFilesResponseTestFile is test data for an ExecutionInputFileResponse
const ExecutionInputFilesResponseTestFile = "execution_input_files.json"

// ExecutionInputFilesResponse is a response for listing execution input files
type ExecutionInputFilesResponse struct {
	Files []ExecutionInputFileResponse `json:"files"`
}

// FromReader returns an ExecutionInputFilesResponse from an io.Reader
func (a *ExecutionInputFilesResponse) FromReader(i io.Reader) error {
	b, err := ioutil.ReadAll(i)
	if err != nil {
		return err
	}
	return json.Unmarshal(b, a)
}

// FromFile returns an ExecutionInputFilesResponse from a file
func (a *ExecutionInputFilesResponse) FromFile(f string) error {
	file, err := os.Open(f)
	defer func() { _ = file.Close() }()
	if err != nil {
		return err
	}
	return a.FromReader(file)
}

// FromBytes returns a ExecutionInputFilesResponse from a byte slice
func (a *ExecutionInputFilesResponse) FromBytes(f []byte) error {
	file := bytes.NewReader(f)
	return a.FromReader(file)
}

// BulkDeleteExecutionsResponseTestFile is test data for an ExecutionInputFileResponse
const BulkDeleteExecutionsResponseTestFile = "bulk_delete_executions.json"

// BulkDeleteExecutionsResponse represents a bulk delete execution response
type BulkDeleteExecutionsResponse struct {
	FailedCount   int                                  `json:"failedCount"`
	SuccessCount  int                                  `json:"successCount"`
	AllSuccessful bool                                 `json:"allsuccessful"`
	RequestCount  int                                  `json:"requestCount"`
	Failures      []BulkDeleteExecutionFailureResponse `json:"failures"`
}

// FromReader returns an BulkDeleteExecutionsResponse from an io.Reader
func (a *BulkDeleteExecutionsResponse) FromReader(i io.Reader) error {
	b, err := ioutil.ReadAll(i)
	if err != nil {
		return err
	}
	return json.Unmarshal(b, a)
}

// FromFile returns an BulkDeleteExecutionsResponse from a file
func (a *BulkDeleteExecutionsResponse) FromFile(f string) error {
	file, err := os.Open(f)
	defer func() { _ = file.Close() }()
	if err != nil {
		return err
	}
	return a.FromReader(file)
}

// FromBytes returns a BulkDeleteExecutionsResponse from a byte slice
func (a *BulkDeleteExecutionsResponse) FromBytes(f []byte) error {
	file := bytes.NewReader(f)
	return a.FromReader(file)
}

// BulkDeleteExecutionFailureResponse represents an individual bulk delete executions failure entry
type BulkDeleteExecutionFailureResponse struct {
	ID      string `json:"id"`
	Message string `json:"message"`
}

// ExecutionStateResponseTestFile is the test data for ExecutionStateResponse
const ExecutionStateResponseTestFile = "execution_state.json"

// ExecutionStateResponse is an execution state response
type ExecutionStateResponse struct {
	Completed      bool                                         `json:"completed"`
	ExecutionState string                                       `json:"executionState"`
	EndTime        *JSONTime                                    `json:"endTime"`
	ServerNode     string                                       `json:"serverNode"`
	StartTime      *JSONTime                                    `json:"startTime"`
	UpdateTime     *JSONTime                                    `json:"updateTime"`
	StepCount      int                                          `json:"stepCount"`
	AllNodes       []string                                     `json:"allNodes"`
	TargetNodes    []string                                     `json:"targetNodes"`
	Nodes          map[string][]ExecutionStateNodeEntryResponse `json:"nodes"`
	ExecutionID    int                                          `json:"executionId"`
	Steps          []json.RawMessage                            `json:"steps"`
}

// FromReader returns an ExecutionStateResponse from an io.Reader
func (a *ExecutionStateResponse) FromReader(i io.Reader) error {
	b, err := ioutil.ReadAll(i)
	if err != nil {
		return err
	}
	return json.Unmarshal(b, a)
}

// FromFile returns an ExecutionStateResponse from a file
func (a *ExecutionStateResponse) FromFile(f string) error {
	file, err := os.Open(f)
	defer func() { _ = file.Close() }()
	if err != nil {
		return err
	}
	return a.FromReader(file)
}

// FromBytes returns a ExecutionStateResponse from a byte slice
func (a *ExecutionStateResponse) FromBytes(f []byte) error {
	file := bytes.NewReader(f)
	return a.FromReader(file)
}

// ExecutionStepResponse represents an execution step
type ExecutionStepResponse struct {
	ExecutionState string                        `json:"executionState"`
	EndTime        *JSONTime                     `json:"endTime"`
	NodeStates     map[string]*NodeStateResponse `json:"nodeStates"`
	UpdateTime     *JSONTime                     `json:"updateTime"`
	NodeStep       bool                          `json:"nodeStep"`
	ID             string                        `json:"id"`
	StartTime      *JSONTime                     `json:"startTime"`
}

// WorkflowResponse represents a workflow response
type WorkflowResponse struct {
	Completed      bool              `json:"completed"`
	EndTime        *JSONTime         `json:"endTime"`
	StartTime      *JSONTime         `json:"startTime"`
	UpdateTime     *JSONTime         `json:"updateTime"`
	StepCount      int               `json:"stepCount"`
	AllNodes       []string          `json:"allNodes"`
	TargetNodes    []string          `json:"targetNodes"`
	ExecutionState string            `json:"executionState"`
	Steps          []json.RawMessage `json:"steps"`
}

// WorkflowStepResponse represents a workflow step response
type WorkflowStepResponse struct {
	Workflow       *WorkflowResponse `json:"workflow"`
	ExecutionState string            `json:"executionState"`
	EndTime        *JSONTime         `json:"endTime"`
	StartTime      *JSONTime         `json:"startTime"`
	UpdateTime     *JSONTime         `json:"updateTime"`
	HasSubworkFlow bool              `json:"hasSubworkFlow"`
	NodeStep       bool              `json:"nodeStep"`
	ID             string            `json:"id"`
}

// NodeStateResponse represents a nodeState response
type NodeStateResponse struct {
	ExecutionState string    `json:"executionState"`
	EndTime        *JSONTime `json:"endTime"`
	UpdateTime     *JSONTime `json:"updateTime"`
	StartTime      *JSONTime `json:"startTime"`
}

// ExecutionStateNodeEntryResponse represents an individual node entry response
type ExecutionStateNodeEntryResponse struct {
	ExecutionState string `json:"executionState"`
	StepCtx        string `json:"stepctx"`
}

// AdHocExecutionResponseTestFile is the test data for an AdHocExecutionResponse
const AdHocExecutionResponseTestFile = "execution_adhoc.json"

// AdHocExecutionResponse is the response for an running and adhoc command
type AdHocExecutionResponse struct {
	Message   string                     `json:"message"`
	Execution AdHocExecutionItemResponse `json:"execution"`
}

// FromReader returns an AdHocExecutionResponse from an io.Reader
func (a *AdHocExecutionResponse) FromReader(i io.Reader) error {
	b, err := ioutil.ReadAll(i)
	if err != nil {
		return err
	}
	return json.Unmarshal(b, a)
}

// FromFile returns an AdHocExecutionResponse from a file
func (a *AdHocExecutionResponse) FromFile(f string) error {
	file, err := os.Open(f)
	defer func() { _ = file.Close() }()
	if err != nil {
		return err
	}
	return a.FromReader(file)
}

// FromBytes returns a AdHocExecutionResponse from a byte slice
func (a *AdHocExecutionResponse) FromBytes(f []byte) error {
	file := bytes.NewReader(f)
	return a.FromReader(file)
}

// AdHocExecutionItemResponse is an individual adhoc execution response
type AdHocExecutionItemResponse struct {
	ID        int    `json:"id"`
	HRef      string `json:"href"`
	Permalink string `json:"permalink"`
}
