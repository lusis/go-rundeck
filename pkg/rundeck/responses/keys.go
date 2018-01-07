package responses

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"os"
)

// ListKeysResponseTestFile is the test data for a ListKeysResponse
const ListKeysResponseTestFile = "list_keys.json"

// ListKeysResponse represents a list keys response
type ListKeysResponse struct {
	Resources []ListKeysResourceResponse `json:"resources"`
	URL       string                     `json:"url"`
	Type      string                     `json:"type"`
	Path      string                     `json:"path"`
}

// FromReader returns a ListKeysResponse from an io.Reader
func (a *ListKeysResponse) FromReader(i io.Reader) error {
	b, err := ioutil.ReadAll(i)
	if err != nil {
		return err
	}
	return json.Unmarshal(b, a)
}

// FromFile returns a ListKeysResponse from a file
func (a *ListKeysResponse) FromFile(f string) error {
	file, err := os.Open(f)
	defer func() { _ = file.Close() }()
	if err != nil {
		return err
	}
	return a.FromReader(file)
}

// FromBytes returns a ListKeysResponse from a byte slice
func (a *ListKeysResponse) FromBytes(f []byte) error {
	file := bytes.NewReader(f)
	return a.FromReader(file)
}

// ListKeysResourceResponse is an individual resource in a list keys response
type ListKeysResourceResponse struct {
	Meta KeyMetaResponse `json:"meta"`
	URL  string          `json:"url"`
	Name string          `json:"name"`
	Type string          `json:"type"`
	Path string          `json:"path"`
}

// ListKeysResourceResponseTestFile is the test data for a KeyMetaResponse
const ListKeysResourceResponseTestFile = "key_metadata.json"

// FromReader returns a KeyMetaResponse from an io.Reader
func (a *ListKeysResourceResponse) FromReader(i io.Reader) error {
	b, err := ioutil.ReadAll(i)
	if err != nil {
		return err
	}
	return json.Unmarshal(b, a)
}

// FromFile returns a KeyMetaResponse from a file
func (a *ListKeysResourceResponse) FromFile(f string) error {
	file, err := os.Open(f)
	defer func() { _ = file.Close() }()
	if err != nil {
		return err
	}
	return a.FromReader(file)
}

// FromBytes returns a ListKeysResourceResponse from a byte slice
func (a *ListKeysResourceResponse) FromBytes(f []byte) error {
	file := bytes.NewReader(f)
	return a.FromReader(file)
}

// KeyMetaResponse is the metadata about an individual list keys resource
type KeyMetaResponse struct {
	RundeckKeyType     string `json:"Rundeck-key-type"`
	RundeckContentMask string `json:"Rundeck-content-mask"`
	RundeckContentSize string `json:"Rundeck-content-size"`
	RundeckContentType string `json:"Rundeck-content-type"`
}
