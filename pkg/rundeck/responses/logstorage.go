package responses

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"os"
)

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

// FromReader returns a LogStorageResponse from an io.Reader
func (a *LogStorageResponse) FromReader(i io.Reader) error {
	b, err := ioutil.ReadAll(i)
	if err != nil {
		return err
	}
	return json.Unmarshal(b, a)
}

// FromFile returns a LogStorageResponse from a file
func (a *LogStorageResponse) FromFile(f string) error {
	file, err := os.Open(f)
	defer func() { _ = file.Close() }()
	if err != nil {
		return err
	}
	return a.FromReader(file)
}

// FromBytes returns a LogStorageResponse from a byte slice
func (a *LogStorageResponse) FromBytes(f []byte) error {
	file := bytes.NewReader(f)
	return a.FromReader(file)
}

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

// FromReader returns a IncompleteLogStorageResponse from an io.Reader
func (a *IncompleteLogStorageResponse) FromReader(i io.Reader) error {
	b, err := ioutil.ReadAll(i)
	if err != nil {
		return err
	}
	return json.Unmarshal(b, a)
}

// FromFile returns a IncompleteLogStorageResponse from a file
func (a *IncompleteLogStorageResponse) FromFile(f string) error {
	file, err := os.Open(f)
	defer func() { _ = file.Close() }()
	if err != nil {
		return err
	}
	return a.FromReader(file)
}

// FromBytes returns a IncompleteLogStorageResponse from a byte slice
func (a *IncompleteLogStorageResponse) FromBytes(f []byte) error {
	file := bytes.NewReader(f)
	return a.FromReader(file)
}

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
