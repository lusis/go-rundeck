package responses

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"os"
)

// UserProfileResponse represents a user info response
// http://rundeck.org/docs/api/index.html#get-user-profile
type UserProfileResponse struct {
	Login     string `json:"login"`
	FirstName string `json:"firstName,omitempty"`
	LastName  string `json:"lastName,omitempty"`
	Email     string `json:"email,omitempty"`
}

// MinVersion is the minimum version of the API required for this response
func (u UserProfileResponse) MinVersion() int {
	return 21
}

// MaxVersion is the maximum version of the API that this response supports
func (u UserProfileResponse) MaxVersion() int {
	return CurrentVersion
}

// Deprecated is if a given response is deprecated
func (u UserProfileResponse) Deprecated() bool {
	return false
}

// UserProfileResponseTestFile is test data for a UserInfoResponse
const UserProfileResponseTestFile = "user.json"

// FromReader returns a UserInfoResponse from an io.Reader
func (u *UserProfileResponse) FromReader(i io.Reader) error {
	b, err := ioutil.ReadAll(i)
	if err != nil {
		return err
	}
	return json.Unmarshal(b, u)
}

// FromFile returns a UserInfoResponse from a file
func (u *UserProfileResponse) FromFile(f string) error {
	file, err := os.Open(f)
	defer func() { _ = file.Close() }()
	if err != nil {
		return err
	}
	return u.FromReader(file)
}

// FromBytes returns a UserInfoResponse from a byte slice
func (u *UserProfileResponse) FromBytes(f []byte) error {
	file := bytes.NewReader(f)
	return u.FromReader(file)
}

// ListUsersResponse is a collection of `UserInfo`
// http://rundeck.org/docs/api/index.html#list-users
type ListUsersResponse []*UserProfileResponse

// ListUsersResponseTestFile is test data for a UsersInfoResponse
const ListUsersResponseTestFile = "users.json"

// MinVersion is the minimum version of the API required for this response
func (u ListUsersResponse) MinVersion() int {
	return 21
}

// MaxVersion is the maximum version of the API that this response supports
func (u ListUsersResponse) MaxVersion() int {
	return CurrentVersion
}

// Deprecated is if a given response is deprecated
func (u ListUsersResponse) Deprecated() bool {
	return false
}

// FromReader returns a UsersInfoResponse from an io.Reader
func (u *ListUsersResponse) FromReader(i io.Reader) error {
	b, err := ioutil.ReadAll(i)
	if err != nil {
		return err
	}
	return json.Unmarshal(b, u)
}

// FromFile returns a UsersInfoResponse from a file
func (u *ListUsersResponse) FromFile(f string) error {
	file, err := os.Open(f)
	defer func() { _ = file.Close() }()
	if err != nil {
		return err
	}
	return u.FromReader(file)
}

// FromBytes returns a UsersInfoResponse from a byte slice
func (u *ListUsersResponse) FromBytes(f []byte) error {
	file := bytes.NewReader(f)
	return u.FromReader(file)
}
