package responses

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"os"
)

// ResourceCollectionResponseTestFile is the testdata used in testing
const ResourceCollectionResponseTestFile = "resources.json"

// ResourceCollectionResponse represents a collection of resource response
type ResourceCollectionResponse ResourceResponse

// ResourceResponseTestFile is the testdata user in testing
const ResourceResponseTestFile = "resource.json"

// ResourceResponse is a single resource in a response
type ResourceResponse map[string]*ResourceDetailResponse

// ResourceDetailResponse represents a project resource response
type ResourceDetailResponse struct {
	NodeName          string                                `json:"nodename"`
	Tags              string                                `json:"tags,omitempty"`
	OsFamily          string                                `json:"osFamily,omitempty"`
	OsVersion         string                                `json:"osVersion,omitempty"`
	OsArch            string                                `json:"osArch,omitempty"`
	OsName            string                                `json:"osName,omitempty"`
	SSHKeyStoragePath string                                `json:"ssh-key-storage-path,omitempty"`
	UserName          string                                `json:"username"`
	Description       string                                `json:"description,omitempty"`
	HostName          string                                `json:"hostname,omitempty"`
	FileCopier        string                                `json:"file-copier,omitempty"`
	NodeExectutor     string                                `json:"node-executor,omitempty"`
	RemoteURL         string                                `json:"remoteUrl,omitempty"`
	EditURL           string                                `json:"editUrl,omitempty"`
	CustomProperties  *ArtbitraryResourcePropertiesResponse `json:",-"`
}

// ArtbitraryResourcePropertiesResponse represents custom properties in a resource response
type ArtbitraryResourcePropertiesResponse map[string]string

// FromReader returns a ResourceCollectionResponse from an io.Reader
func (r *ResourceCollectionResponse) FromReader(i io.Reader) error {
	b, err := ioutil.ReadAll(i)
	if err != nil {
		return err
	}
	return json.Unmarshal(b, r)

}

// FromFile returns a ResourceCollectionResponse from a file
func (r *ResourceCollectionResponse) FromFile(f string) error {
	file, err := os.Open(f)
	defer func() { _ = file.Close() }()
	if err != nil {
		return err
	}
	return r.FromReader(file)
}

// FromBytes returns a ResourceCollectionResponse from a byte slice
func (r *ResourceCollectionResponse) FromBytes(f []byte) error {
	file := bytes.NewReader(f)
	return r.FromReader(file)
}

// FromReader returns a ResourceResponse from an io.Reader
func (r *ResourceResponse) FromReader(i io.Reader) error {
	b, err := ioutil.ReadAll(i)
	if err != nil {
		return err
	}
	return json.Unmarshal(b, r)

}

// FromFile returns a ResourceResponse from a file
func (r *ResourceResponse) FromFile(f string) error {
	file, err := os.Open(f)
	defer func() { _ = file.Close() }()
	if err != nil {
		return err
	}
	return r.FromReader(file)
}

// FromBytes returns a ResourceResponse from a byte slice
func (r *ResourceResponse) FromBytes(f []byte) error {
	file := bytes.NewReader(f)
	return r.FromReader(file)
}
