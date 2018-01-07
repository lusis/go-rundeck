package responses

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"os"
)

// UserInfoResponse represents a user info response
type UserInfoResponse struct {
	Login     string `json:"login"`
	FirstName string `json:"firstName,omitempty"`
	LastName  string `json:"lastName,omitempty"`
	Email     string `json:"email,omitempty"`
}

// UserInfoResponseTestFile is test data for a UserInfoResponse
const UserInfoResponseTestFile = "user.json"

// FromReader returns a UserInfoResponse from an io.Reader
func (a *UserInfoResponse) FromReader(i io.Reader) error {
	b, err := ioutil.ReadAll(i)
	if err != nil {
		return err
	}
	return json.Unmarshal(b, a)
}

// FromFile returns a UserInfoResponse from a file
func (a *UserInfoResponse) FromFile(f string) error {
	file, err := os.Open(f)
	defer func() { _ = file.Close() }()
	if err != nil {
		return err
	}
	return a.FromReader(file)
}

// FromBytes returns a UserInfoResponse from a byte slice
func (a *UserInfoResponse) FromBytes(f []byte) error {
	file := bytes.NewReader(f)
	return a.FromReader(file)
}

// UsersInfoResponse is a collection of `UserInfo`
type UsersInfoResponse []*UserInfoResponse

// UsersInfoResponseTestFile is test data for a UsersInfoResponse
const UsersInfoResponseTestFile = "users.json"

// FromReader returns a UsersInfoResponse from an io.Reader
func (a *UsersInfoResponse) FromReader(i io.Reader) error {
	b, err := ioutil.ReadAll(i)
	if err != nil {
		return err
	}
	return json.Unmarshal(b, a)
}

// FromFile returns a UsersInfoResponse from a file
func (a *UsersInfoResponse) FromFile(f string) error {
	file, err := os.Open(f)
	defer func() { _ = file.Close() }()
	if err != nil {
		return err
	}
	return a.FromReader(file)
}

// FromBytes returns a UsersInfoResponse from a byte slice
func (a *UsersInfoResponse) FromBytes(f []byte) error {
	file := bytes.NewReader(f)
	return a.FromReader(file)
}
