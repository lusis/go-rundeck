package responses

// JobExecutionsResponse is the response for listing the executions for a job
type JobExecutionsResponse ListRunningExecutionsResponse

func (a JobExecutionsResponse) minVersion() int  { return AbsoluteMinimumVersion }
func (a JobExecutionsResponse) maxVersion() int  { return CurrentVersion }
func (a JobExecutionsResponse) deprecated() bool { return false }

// ListRunningExecutionsResponseTestFile is the test data for JobExecutionResponse
const ListRunningExecutionsResponseTestFile = "executions.json"

// ListRunningExecutionsResponse is the response for listing the running executions for a project
type ListRunningExecutionsResponse struct {
	Paging     PagingResponse      `json:"paging"`
	Executions []ExecutionResponse `json:"executions"`
}

func (a ListRunningExecutionsResponse) minVersion() int  { return AbsoluteMinimumVersion }
func (a ListRunningExecutionsResponse) maxVersion() int  { return CurrentVersion }
func (a ListRunningExecutionsResponse) deprecated() bool { return false }

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

func (a ExecutionResponse) minVersion() int  { return AbsoluteMinimumVersion }
func (a ExecutionResponse) maxVersion() int  { return CurrentVersion }
func (a ExecutionResponse) deprecated() bool { return false }

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

func (a ExecutionJobEntryResponse) minVersion() int  { return AbsoluteMinimumVersion }
func (a ExecutionJobEntryResponse) maxVersion() int  { return CurrentVersion }
func (a ExecutionJobEntryResponse) deprecated() bool { return false }

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

func (a ExecutionInputFileResponse) minVersion() int  { return 19 }
func (a ExecutionInputFileResponse) maxVersion() int  { return CurrentVersion }
func (a ExecutionInputFileResponse) deprecated() bool { return false }

// ExecutionInputFilesResponseTestFile is test data for an ExecutionInputFileResponse
const ExecutionInputFilesResponseTestFile = "execution_input_files.json"

// ExecutionInputFilesResponse is a response for listing execution input files
type ExecutionInputFilesResponse struct {
	Files []ExecutionInputFileResponse `json:"files"`
}

func (a ExecutionInputFilesResponse) minVersion() int  { return 19 }
func (a ExecutionInputFilesResponse) maxVersion() int  { return CurrentVersion }
func (a ExecutionInputFilesResponse) deprecated() bool { return false }

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

func (a BulkDeleteExecutionsResponse) minVersion() int  { return AbsoluteMinimumVersion }
func (a BulkDeleteExecutionsResponse) maxVersion() int  { return CurrentVersion }
func (a BulkDeleteExecutionsResponse) deprecated() bool { return false }

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
	Steps          []interface{}                                `json:"steps"`
}

func (a ExecutionStateResponse) minVersion() int  { return AbsoluteMinimumVersion }
func (a ExecutionStateResponse) maxVersion() int  { return CurrentVersion }
func (a ExecutionStateResponse) deprecated() bool { return false }

// ExecutionStepResponse represents an execution step
type ExecutionStepResponse struct {
	ID              string                        `json:"id"`
	StepCTX         string                        `json:"stepctx"`
	ParameterStates map[string]interface{}        `json:"parameterStates"`
	Duration        int                           `json:"duration"`
	NodeStep        bool                          `json:"nodeStep"`
	ExecutionState  string                        `json:"executionState"`
	StartTime       *JSONTime                     `json:"startTime"`
	UpdateTime      *JSONTime                     `json:"updateTime"`
	EndTime         *JSONTime                     `json:"endTime"`
	NodeStates      map[string]*NodeStateResponse `json:"nodeStates"`
}

func (a ExecutionStepResponse) minVersion() int  { return AbsoluteMinimumVersion }
func (a ExecutionStepResponse) maxVersion() int  { return CurrentVersion }
func (a ExecutionStepResponse) deprecated() bool { return false }

// WorkflowStepResponse represents a workflow step response
type WorkflowStepResponse struct {
	ID              string                 `json:"id"`
	StepCTX         string                 `json:"stepctx"`
	ParameterStates map[string]interface{} `json:"parameterStates"`
	Duration        int                    `json:"duration"`
	NodeStep        bool                   `json:"nodeStep"`
	ExecutionState  string                 `json:"executionState"`
	StartTime       *JSONTime              `json:"startTime"`
	UpdateTime      *JSONTime              `json:"updateTime"`
	EndTime         *JSONTime              `json:"endTime"`
	Workflow        *WorkflowResponse      `json:"workflow"`
	HasSubworkFlow  bool                   `json:"hasSubworkFlow"`
}

func (a WorkflowStepResponse) minVersion() int  { return AbsoluteMinimumVersion }
func (a WorkflowStepResponse) maxVersion() int  { return CurrentVersion }
func (a WorkflowStepResponse) deprecated() bool { return false }

// WorkflowResponse represents a workflow response
type WorkflowResponse struct {
	Completed      bool          `json:"completed"`
	EndTime        *JSONTime     `json:"endTime"`
	StartTime      *JSONTime     `json:"startTime"`
	UpdateTime     *JSONTime     `json:"updateTime"`
	StepCount      int           `json:"stepCount"`
	AllNodes       []string      `json:"allNodes"`
	TargetNodes    []string      `json:"targetNodes"`
	ExecutionState string        `json:"executionState"`
	Steps          []interface{} `json:"steps"`
}

func (a WorkflowResponse) minVersion() int  { return AbsoluteMinimumVersion }
func (a WorkflowResponse) maxVersion() int  { return CurrentVersion }
func (a WorkflowResponse) deprecated() bool { return false }

// NodeStateResponse represents a nodeState response
type NodeStateResponse struct {
	Duration       int       `json:"duration"`
	ExecutionState string    `json:"executionState"`
	EndTime        *JSONTime `json:"endTime"`
	UpdateTime     *JSONTime `json:"updateTime"`
	StartTime      *JSONTime `json:"startTime"`
}

func (a NodeStateResponse) minVersion() int  { return AbsoluteMinimumVersion }
func (a NodeStateResponse) maxVersion() int  { return CurrentVersion }
func (a NodeStateResponse) deprecated() bool { return false }

// ExecutionStateNodeEntryResponse represents an individual node entry response
type ExecutionStateNodeEntryResponse struct {
	ExecutionState string `json:"executionState"`
	StepCtx        string `json:"stepctx"`
}

func (a ExecutionStateNodeEntryResponse) minVersion() int  { return AbsoluteMinimumVersion }
func (a ExecutionStateNodeEntryResponse) maxVersion() int  { return CurrentVersion }
func (a ExecutionStateNodeEntryResponse) deprecated() bool { return false }

// AdHocExecutionResponseTestFile is the test data for an AdHocExecutionResponse
const AdHocExecutionResponseTestFile = "execution_adhoc.json"

// AdHocExecutionResponse is the response for an running and adhoc command
type AdHocExecutionResponse struct {
	Message   string                     `json:"message"`
	Execution AdHocExecutionItemResponse `json:"execution"`
}

func (a AdHocExecutionResponse) minVersion() int  { return AbsoluteMinimumVersion }
func (a AdHocExecutionResponse) maxVersion() int  { return CurrentVersion }
func (a AdHocExecutionResponse) deprecated() bool { return false }

// AdHocExecutionItemResponse is an individual adhoc execution response
type AdHocExecutionItemResponse struct {
	ID        int    `json:"id"`
	HRef      string `json:"href"`
	Permalink string `json:"permalink"`
}

func (a AdHocExecutionItemResponse) minVersion() int  { return AbsoluteMinimumVersion }
func (a AdHocExecutionItemResponse) maxVersion() int  { return CurrentVersion }
func (a AdHocExecutionItemResponse) deprecated() bool { return false }

// AbortExecutionResponse is the response for aborting an execution
type AbortExecutionResponse struct {
	Abort struct {
		Status string `json:"status"`
		Reason string `json:"reason"`
	} `json:"abort"`
	Execution struct {
		ID     string `json:"id"`
		Status string `json:"status"`
		HRef   string `json:"href"`
	} `json:"execution"`
}

// AbortExecutionResponseTestFile is test data for aborting an execution
const AbortExecutionResponseTestFile = "execution_aborted.json"

func (a AbortExecutionResponse) minVersion() int  { return AbsoluteMinimumVersion }
func (a AbortExecutionResponse) maxVersion() int  { return CurrentVersion }
func (a AbortExecutionResponse) deprecated() bool { return false }

// ExecutionOutputResponse is the response for getting execution output
type ExecutionOutputResponse struct {
	ID             string  `json:"id"`
	Offset         string  `json:"offset"`
	Completed      bool    `json:"completed"`
	ExecCompleted  bool    `json:"execCompleted"`
	HasFailedNodes bool    `json:"hasFailedNodes"`
	ExecState      string  `json:"execState"`
	LastModified   string  `json:"lastModified"`
	ExecDuration   int     `json:"execDuration"`
	PercentLoaded  float64 `json:"percentLoaded"`
	TotalSize      int     `json:"totalSize"`
	RetryBackoff   int     `json:"retryBackoff"`
	ClusterExec    bool    `json:"clusterExec"`
	ServerNodeUUID string  `json:"serverNodeUUID"`
	Compacted      bool    `json:"compacted"`
	Entries        []struct {
		Time         string    `json:"time"`
		AbsoluteTime *JSONTime `json:"absolute_time"`
		Log          string    `json:"log"`
		Level        string    `json:"level"`
		User         string    `json:"user"`
		StepCTX      string    `json:"stepctx"`
		Node         string    `json:"node"`
	} `json:"entries"`
}

// ExecutionOutputResponseTestFile is test data for getting an output execution response
const ExecutionOutputResponseTestFile = "execution_output.json"

func (a ExecutionOutputResponse) minVersion() int  { return 21 }
func (a ExecutionOutputResponse) maxVersion() int  { return CurrentVersion }
func (a ExecutionOutputResponse) deprecated() bool { return false }
