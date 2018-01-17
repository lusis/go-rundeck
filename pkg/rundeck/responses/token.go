package responses

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"os"
)

// TokenResponse represents a user and token
// http://rundeck.org/docs/api/index.html#get-a-token
type TokenResponse struct {
	ID         string    `json:"id,omitempty"`
	User       string    `json:"user,omitempty"`
	Token      string    `json:"token,omitempty"`
	Creator    string    `json:"creator,omitempty"`
	Expiration *JSONTime `json:"expiration,omitempty"`
	Roles      []string  `json:"roles,omitempty"`
	Expired    bool      `json:"expired,omitempty"`
}

func (t TokenResponse) minVersion() int  { return 19 }
func (t TokenResponse) maxVersion() int  { return CurrentVersion }
func (t TokenResponse) deprecated() bool { return false }

// TokenResponseTestFile is test data for a TokenResponse
const TokenResponseTestFile = "token.json"

// FromReader returns a TokenResponse from an io.Reader
func (t *TokenResponse) FromReader(i io.Reader) error {
	b, err := ioutil.ReadAll(i)
	if err != nil {
		return err
	}
	return json.Unmarshal(b, t)
}

// FromFile returns a TokenResponse from a file
func (t *TokenResponse) FromFile(f string) error {
	file, err := os.Open(f)
	defer func() { _ = file.Close() }()
	if err != nil {
		return err
	}
	return t.FromReader(file)
}

// FromBytes returns a TokenResponse from a byte slice
func (t *TokenResponse) FromBytes(f []byte) error {
	file := bytes.NewReader(f)
	return t.FromReader(file)
}

// ListTokensResponse is a collection of `Token`
// http://rundeck.org/docs/api/index.html#list-tokens
type ListTokensResponse []TokenResponse

func (t ListTokensResponse) minVersion() int  { return 19 }
func (t ListTokensResponse) maxVersion() int  { return CurrentVersion }
func (t ListTokensResponse) deprecated() bool { return false }

// ListTokensResponseTestFile is test data for a TokensResponse
const ListTokensResponseTestFile = "tokens.json"
