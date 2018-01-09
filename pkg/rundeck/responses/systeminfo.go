package responses

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"os"
)

// SystemInfoResponse represents a system info response
// http://rundeck.org/docs/api/index.html#system-info
type SystemInfoResponse struct {
	System *SystemsResponse `json:"system"`
}

func (a SystemInfoResponse) minVersion() int  { return AbsoluteMinimumVersion }
func (a SystemInfoResponse) maxVersion() int  { return CurrentVersion }
func (a SystemInfoResponse) deprecated() bool { return false }

// SystemInfoResponseTestFile is test data for a SystemInfoResponse
const SystemInfoResponseTestFile = "systeminfo.json"

// FromReader returns a SystemInfoResponse from an io.Reader
func (a *SystemInfoResponse) FromReader(i io.Reader) error {
	b, err := ioutil.ReadAll(i)
	if err != nil {
		return err
	}
	return json.Unmarshal(b, a)
}

// FromFile returns a SystemInfoResponse from a file
func (a *SystemInfoResponse) FromFile(f string) error {
	file, err := os.Open(f)
	defer func() { _ = file.Close() }()
	if err != nil {
		return err
	}
	return a.FromReader(file)
}

// FromBytes returns a SystemInfoResponse from a byte slice
func (a *SystemInfoResponse) FromBytes(f []byte) error {
	file := bytes.NewReader(f)
	return a.FromReader(file)
}

// SystemsResponse represents a systems response
// http://rundeck.org/docs/api/index.html#system-info
type SystemsResponse struct {
	Timestamp  *SysInfoTimestampResponse  `json:"timestamp"`
	Rundeck    *SysInfoRundeckResponse    `json:"rundeck"`
	Executions *SysInfoExecutionsResponse `json:"executions"`
	OS         *SysInfoOSResponse         `json:"os"`
	JVM        *SysInfoJVMResponse        `json:"jvm"`
	Stats      *SysInfoStatsResponse      `json:"stats"`
	Metrics    *SysInfoMetricsResponse    `json:"metrics"`
	ThreadDump *SysInfoThreadDumpResponse `json:"threadDump"`
}

// SysInfoStatsResponse is a stats response
type SysInfoStatsResponse struct {
	Uptime *struct {
		Duration int64                     `json:"duration"`
		Unit     string                    `json:"unit"`
		Since    *SysInfoTimestampResponse `json:"since"`
	} `json:"uptime"`
	CPU *struct {
		Processors  int `json:"processors"`
		LoadAverage struct {
			Unit    string  `json:"unit"`
			Average float64 `json:"average"`
		} `json:"loadAverage"`
	} `json:"cpu"`
	Memory *struct {
		Unit  string `json:"unit"`
		Max   int64  `json:"max"`
		Free  int64  `json:"free"`
		Total int64  `json:"total"`
	} `json:"memory"`
	Scheduler *struct {
		Running        int `json:"running"`
		ThreadPoolSize int `json:"threadPoolSize"`
	} `json:"scheduler"`
	Threads *struct {
		Active int `json:"active"`
	} `json:"threads"`
}

// SysInfoThreadDumpResponse is a thread dump response
type SysInfoThreadDumpResponse struct {
	HRef        string `json:"href"`
	ContentType string `json:"contentType"`
}

// SysInfoMetricsResponse is a metrics response
type SysInfoMetricsResponse struct {
	HRef        string `json:"href"`
	ContentType string `json:"contentType"`
}

// SysInfoTimestampResponse represents a timestamp response
type SysInfoTimestampResponse struct {
	Epoch    int64     `json:"epoch"`
	Unit     string    `json:"unit"`
	DateTime *JSONTime `json:"datetime"`
}

// SysInfoRundeckResponse represents a rundeck response
type SysInfoRundeckResponse struct {
	Version    string `json:"version"`
	Build      string `json:"build"`
	Node       string `json:"node"`
	Base       string `json:"base"`
	APIVersion int    `json:"apiversion"`
	ServerUUID string `json:"serverUUID"`
}

// SysInfoExecutionsResponse represents an executions response in a systeminfo response
type SysInfoExecutionsResponse struct {
	Active        bool   `json:"active"`
	ExecutionMode string `json:"executionMode"`
}

// SysInfoOSResponse represents an OS response
type SysInfoOSResponse struct {
	Arch    string `json:"arch"`
	Name    string `json:"name"`
	Version string `json:"version"`
}

// SysInfoJVMResponse represents a jvm response
type SysInfoJVMResponse struct {
	Name                  string `json:"name"`
	Vendor                string `json:"vendor"`
	Version               string `json:"version"`
	ImplementationVersion string `json:"implementationVersion"`
}
