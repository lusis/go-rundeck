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

func (a ResourceCollectionResponse) minVersion() int  { return AbsoluteMinimumVersion }
func (a ResourceCollectionResponse) maxVersion() int  { return CurrentVersion }
func (a ResourceCollectionResponse) deprecated() bool { return false }

// ResourceResponseTestFile is the testdata user in testing
const ResourceResponseTestFile = "resource.json"

// ResourceResponse is a single resource in a response
type ResourceResponse map[string]ResourceDetailResponse

func (a ResourceResponse) minVersion() int  { return AbsoluteMinimumVersion }
func (a ResourceResponse) maxVersion() int  { return CurrentVersion }
func (a ResourceResponse) deprecated() bool { return false }

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
	HostName          string                                `json:"hostname"`
	FileCopier        string                                `json:"file-copier"`
	NodeExectutor     string                                `json:"node-executor"`
	RemoteURL         string                                `json:"remoteUrl,omitempty"`
	EditURL           string                                `json:"editUrl,omitempty"`
	CustomProperties  *ArtbitraryResourcePropertiesResponse `json:",-"`
}

func (a ResourceDetailResponse) minVersion() int  { return AbsoluteMinimumVersion }
func (a ResourceDetailResponse) maxVersion() int  { return CurrentVersion }
func (a ResourceDetailResponse) deprecated() bool { return false }

// ArtbitraryResourcePropertiesResponse represents custom properties in a resource response
type ArtbitraryResourcePropertiesResponse map[string]string

func (a ArtbitraryResourcePropertiesResponse) minVersion() int  { return AbsoluteMinimumVersion }
func (a ArtbitraryResourcePropertiesResponse) maxVersion() int  { return CurrentVersion }
func (a ArtbitraryResourcePropertiesResponse) deprecated() bool { return false }

// FromReader returns a ResourceCollectionResponse from an io.Reader
func (a *ResourceCollectionResponse) FromReader(i io.Reader) error {
	b, err := ioutil.ReadAll(i)
	if err != nil {
		return err
	}
	return json.Unmarshal(b, a)

}

// FromFile returns a ResourceCollectionResponse from a file
func (a *ResourceCollectionResponse) FromFile(f string) error {
	file, err := os.Open(f)
	defer func() { _ = file.Close() }()
	if err != nil {
		return err
	}
	return a.FromReader(file)
}

// FromBytes returns a ResourceCollectionResponse from a byte slice
func (a *ResourceCollectionResponse) FromBytes(f []byte) error {
	file := bytes.NewReader(f)
	return a.FromReader(file)
}

// FromReader returns a ResourceResponse from an io.Reader
func (a *ResourceResponse) FromReader(i io.Reader) error {
	b, err := ioutil.ReadAll(i)
	if err != nil {
		return err
	}
	return json.Unmarshal(b, a)

}

// FromFile returns a ResourceResponse from a file
func (a *ResourceResponse) FromFile(f string) error {
	file, err := os.Open(f)
	defer func() { _ = file.Close() }()
	if err != nil {
		return err
	}
	return a.FromReader(file)
}

// FromBytes returns a ResourceResponse from a byte slice
func (a *ResourceResponse) FromBytes(f []byte) error {
	file := bytes.NewReader(f)
	return a.FromReader(file)
}
