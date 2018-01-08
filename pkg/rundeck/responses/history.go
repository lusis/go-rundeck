package responses

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"os"
)

// HistoryResponseTestFile is the test data for a HistoryResponse
const HistoryResponseTestFile = "history.json"

// HistoryResponse represents a project history response
// http://rundeck.org/docs/api/index.html#listing-history
type HistoryResponse struct {
	Paging *PagingResponse         `json:"paging"`
	Events []*HistoryEventResponse `json:"events"`
}

// MinVersion is the minimum version of the API required for this response
func (a HistoryResponse) MinVersion() int {
	return 14
}

// MaxVersion is the maximum version of the API that this response supports
func (a HistoryResponse) MaxVersion() int {
	return CurrentVersion
}

// Deprecated is if a given response is deprecated
func (a HistoryResponse) Deprecated() bool {
	return false
}

// FromReader returns a HistoryResponse from an io.Reader
func (a *HistoryResponse) FromReader(i io.Reader) error {
	b, err := ioutil.ReadAll(i)
	if err != nil {
		return err
	}
	return json.Unmarshal(b, a)
}

// FromFile returns a HistoryResponse from a file
func (a *HistoryResponse) FromFile(f string) error {
	file, err := os.Open(f)
	defer func() { _ = file.Close() }()
	if err != nil {
		return err
	}
	return a.FromReader(file)
}

// FromBytes returns a HistoryResponse from a byte slice
func (a *HistoryResponse) FromBytes(f []byte) error {
	file := bytes.NewReader(f)
	return a.FromReader(file)
}

// HistoryEventResponse represents an individual event in a history response
type HistoryEventResponse struct {
	StartTime   int       `json:"starttime"`
	EndTime     int       `json:"endtime"`
	DateStarted *JSONTime `json:"date-started"`
	DateEnded   *JSONTime `json:"date-ended"`
	Title       string    `json:"title"`
	Status      string    `json:"status"`
	Summary     string    `json:"summary"`
	NodeSummary *struct {
		Succeeded int `json:"succeeded"`
		Failed    int `json:"failed"`
		Total     int `json:"total"`
	} `json:"node-summary"`
	User      string `json:"user"`
	Project   string `json:"project"`
	Execution *struct {
		ID        string `json:"id"`
		HRef      string `json:"href"`
		Permalink string `json:"permalink"`
	} `json:"execution"`
}

// MinVersion is the minimum version of the API required for this response
func (a HistoryEventResponse) MinVersion() int {
	return 14
}

// MaxVersion is the maximum version of the API that this response supports
func (a HistoryEventResponse) MaxVersion() int {
	return CurrentVersion
}

// Deprecated is if a given response is deprecated
func (a HistoryEventResponse) Deprecated() bool {
	return false
}
