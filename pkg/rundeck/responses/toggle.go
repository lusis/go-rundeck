package responses

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"os"
)

// BulkToggleResponseTestFile is test data for an ExecutionInputFileResponse
const BulkToggleResponseTestFile = "bulk_toggle.json"

// BulkToggleResponse represents a bulk toggle response
/*
{
  "requestCount": #integer#,
  "enabled": true/false,
  "allsuccessful": true/false,
  "succeeded": [...],
  "failed":[...]
}
The list of succeeded/failed will contain objects of this form:

{
  "id": "[UUID]",
  "errorCode": "(error code, see above)",
  "message": "(success or failure message)"
}
*/
type BulkToggleResponse struct {
	Enabled       bool                      `json:"enabled"`
	AllSuccessful bool                      `json:"allsuccessful"`
	RequestCount  int                       `json:"requestCount"`
	Failed        []BulkToggleEntryResponse `json:"failed"`
	Succeeded     []BulkToggleEntryResponse `json:"succeeded"`
}

// BulkToggleEntryResponse represents an individual entry in a BulkToggleResponse
type BulkToggleEntryResponse struct {
	ID        string `json:"id"`
	ErrorCode string `json:"errorCode,omitempty"`
	Message   string `json:"message"`
}

// FromReader returns an BulkToggleResponse from an io.Reader
func (a *BulkToggleResponse) FromReader(i io.Reader) error {
	b, err := ioutil.ReadAll(i)
	if err != nil {
		return err
	}
	return json.Unmarshal(b, a)
}

// FromFile returns an BulkToggleResponse from a file
func (a *BulkToggleResponse) FromFile(f string) error {
	file, err := os.Open(f)
	defer func() { _ = file.Close() }()
	if err != nil {
		return err
	}
	return a.FromReader(file)
}

// FromBytes returns a BulkResponse from a byte slice
func (a *BulkToggleResponse) FromBytes(f []byte) error {
	file := bytes.NewReader(f)
	return a.FromReader(file)
}
