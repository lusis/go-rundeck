package responses

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"os"
)

// ACLResponseTestFile is the test data for ACLResourceResponse
const ACLResponseTestFile = "acl.json"

// FailedACLValidationResponseTestFile is the test data for ACLResourceResponse
const FailedACLValidationResponseTestFile = "failed_acl_validation.json"

// ACLResponse represents acl response
type ACLResponse struct {
	Path      string                `json:"path"`
	Type      string                `json:"type"`
	Href      string                `json:"href"`
	Resources []ACLResourceResponse `json:"resources,omitempty"`
}

func (a ACLResponse) minVersion() int  { return AbsoluteMinimumVersion }
func (a ACLResponse) maxVersion() int  { return CurrentVersion }
func (a ACLResponse) deprecated() bool { return false }

// FromReader returns an ACLResponse from an io.Reader
func (a *ACLResponse) FromReader(i io.Reader) error {
	b, err := ioutil.ReadAll(i)
	if err != nil {
		return err
	}
	return json.Unmarshal(b, a)
}

// FromFile returns an ACLResponse from a file
func (a *ACLResponse) FromFile(f string) error {
	file, err := os.Open(f)
	defer func() { _ = file.Close() }()
	if err != nil {
		return err
	}
	return a.FromReader(file)
}

// FromBytes returns an ACLResponse from a byte slice
func (a *ACLResponse) FromBytes(f []byte) error {
	file := bytes.NewReader(f)
	return a.FromReader(file)
}

// ACLResourceResponse represent an ACL Resource response
type ACLResourceResponse struct {
	Path string `json:"path"`
	Type string `json:"type"`
	Href string `json:"href"`
	Name string `json:"name,omitempty"`
}

func (a ACLResourceResponse) minVersion() int  { return AbsoluteMinimumVersion }
func (a ACLResourceResponse) maxVersion() int  { return CurrentVersion }
func (a ACLResourceResponse) deprecated() bool { return false }

// FailedACLValidationResponse represents a failed ACL validation response
type FailedACLValidationResponse struct {
	Valid    bool                      `json:"valid"`
	Policies []FailedACLPolicyResponse `json:"policies"`
}

func (a FailedACLValidationResponse) minVersion() int  { return AbsoluteMinimumVersion }
func (a FailedACLValidationResponse) maxVersion() int  { return CurrentVersion }
func (a FailedACLValidationResponse) deprecated() bool { return false }

// FailedACLPolicyResponse represents a failed ACL policy
type FailedACLPolicyResponse struct {
	Policy string   `json:"policy"`
	Errors []string `json:"errors"`
}

func (a FailedACLPolicyResponse) minVersion() int  { return AbsoluteMinimumVersion }
func (a FailedACLPolicyResponse) maxVersion() int  { return CurrentVersion }
func (a FailedACLPolicyResponse) deprecated() bool { return false }

// FromReader returns a FailedACLValidationResponse from an io.Reader
func (a *FailedACLValidationResponse) FromReader(i io.Reader) error {
	b, err := ioutil.ReadAll(i)
	if err != nil {
		return err
	}
	return json.Unmarshal(b, a)
}

// FromFile returns a FailedACLValidationResponse from a file
func (a *FailedACLValidationResponse) FromFile(f string) error {
	file, err := os.Open(f)
	defer func() { _ = file.Close() }()
	if err != nil {
		return err
	}
	return a.FromReader(file)
}

// FromBytes returns a FaileACLValidationResponse from a byte slice
func (a *FailedACLValidationResponse) FromBytes(f []byte) error {
	file := bytes.NewReader(f)
	return a.FromReader(file)
}
