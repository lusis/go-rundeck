package responses

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"os"
)

// TokenResponse represents a user and token
type TokenResponse struct {
	ID         string    `json:"id,omitempty"`
	User       string    `json:"user,omitempty"`
	Token      string    `json:"token,omitempty"`
	Creator    string    `json:"creator,omitempty"`
	Expiration *JSONTime `json:"expiration,omitempty"`
	Roles      []string  `json:"roles,omitempty"`
	Expired    bool      `json:"expired,omitempty"`
	Duration   string    `json:"duration,omitempty"`
}

// TokenResponseTestFile is test data for a TokenResponse
const TokenResponseTestFile = "token.json"

// FromReader returns a TokenResponse from an io.Reader
func (a *TokenResponse) FromReader(i io.Reader) error {
	b, err := ioutil.ReadAll(i)
	if err != nil {
		return err
	}
	return json.Unmarshal(b, a)
}

// FromFile returns a TokenResponse from a file
func (a *TokenResponse) FromFile(f string) error {
	file, err := os.Open(f)
	defer func() { _ = file.Close() }()
	if err != nil {
		return err
	}
	return a.FromReader(file)
}

// FromBytes returns a TokenResponse from a byte slice
func (a *TokenResponse) FromBytes(f []byte) error {
	file := bytes.NewReader(f)
	return a.FromReader(file)
}

// TokensResponse is a collection of `Token`
type TokensResponse []*TokenResponse

// TokensResponseTestFile is test data for a TokensResponse
const TokensResponseTestFile = "tokens.json"
