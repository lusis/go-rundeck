package responses

// LogStorageResponseTestFile is the test data for a LogStorageResponse
const LogStorageResponseTestFile = "logstorage.json"

// LogStorageResponse represents a LogStorageResponse response
type LogStorageResponse struct {
	Enabled         bool   `json:"enabled"`
	PluginName      string `json:"pluginName"`
	SucceededCount  int    `json:"succeededCount"`
	FailedCount     int    `json:"failedCount"`
	QueuedCount     int    `json:"queuedCount"`
	TotalCount      int    `json:"totalCount"`
	IncompleteCount int    `json:"incompleteCount"`
	MissingCount    int    `json:"missingCount"`
}

func (a LogStorageResponse) minVersion() int  { return 17 }
func (a LogStorageResponse) maxVersion() int  { return CurrentVersion }
func (a LogStorageResponse) deprecated() bool { return false }

// IncompleteLogStorageResponseTestFile is test data for an IncompleteLogStorageResponse
const IncompleteLogStorageResponseTestFile = "incomplete_logstorage_executions.json"

// IncompleteLogStorageResponse represents an incomplete log storage response
type IncompleteLogStorageResponse struct {
	Total      int                                      `json:"total"`
	Max        int                                      `json:"max"`
	Offset     int                                      `json:"offset"`
	Executions []*IncompleteLogStorageExecutionResponse `json:"executions"`
}

func (a IncompleteLogStorageResponse) minVersion() int  { return 17 }
func (a IncompleteLogStorageResponse) maxVersion() int  { return CurrentVersion }
func (a IncompleteLogStorageResponse) deprecated() bool { return false }

// IncompleteLogStorageExecutionResponse represents an incomplete log storage execution response
type IncompleteLogStorageExecutionResponse struct {
	ID        int    `json:"id"`
	Project   string `json:"project"`
	HRef      string `json:"href"`
	Permalink string `json:"permalink"`
	Storage   struct {
		LocalFilesPresent   bool      `json:"localFilesPresent"`
		IncompleteFiletypes string    `json:"incompleteFiletypes"`
		Queued              bool      `json:"queued"`
		Failed              bool      `json:"failed"`
		Date                *JSONTime `json:"date"`
	} `json:"storage"`
	Errors []string `json:"errors"`
}

func (a IncompleteLogStorageExecutionResponse) minVersion() int  { return 17 }
func (a IncompleteLogStorageExecutionResponse) maxVersion() int  { return CurrentVersion }
func (a IncompleteLogStorageExecutionResponse) deprecated() bool { return false }
