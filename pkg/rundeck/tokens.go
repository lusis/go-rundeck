package rundeck

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"

	multierror "github.com/hashicorp/go-multierror"
	requests "github.com/lusis/go-rundeck/pkg/rundeck/requests"
	responses "github.com/lusis/go-rundeck/pkg/rundeck/responses"
)

// Token is a token
type Token struct {
	responses.TokenResponse
}

// TokenRequest is a new token request
type TokenRequest struct {
	requests.TokenRequest
}

// Tokens is a collection of Token
type Tokens responses.ListTokensResponse

// TokenOption is a type for options in creating new tokens
type TokenOption func(*TokenRequest) error

// TokenDuration is an option for setting the duration of a new token
func TokenDuration(duration string) TokenOption {
	return func(t *TokenRequest) error {
		t.Duration = duration
		return nil
	}
}

// TokenRoles is an option to set the roles for a new token
func TokenRoles(roles ...string) TokenOption {
	return func(t *TokenRequest) error {
		t.Roles = strings.Join(roles, ",")
		return nil
	}
}

// ListTokens gets all tokens for the current user
// http://rundeck.org/docs/api/index.html#list-tokens
func (c *Client) ListTokens() ([]*Token, error) {
	if err := c.checkRequiredAPIVersion(responses.ListTokensResponse{}); err != nil {
		return nil, err
	}
	tokens := []*Token{}
	data, err := c.httpGet("tokens", requestJSON(), requestExpects(200))
	if err != nil {
		return nil, err
	}
	jsonErr := json.Unmarshal(data, &tokens)
	if jsonErr != nil {
		return nil, &UnmarshalError{msg: multierror.Append(errDecoding, jsonErr).Error()}
	}
	return tokens, nil
}

// ListTokensForUser gets the api tokens for a user
// http://rundeck.org/docs/api/index.html#list-tokens
func (c *Client) ListTokensForUser(user string) ([]*Token, error) {
	if err := c.checkRequiredAPIVersion(responses.ListTokensResponse{}); err != nil {
		return nil, err
	}
	data, err := c.httpGet("tokens/"+user, requestJSON(), requestExpects(200))
	if err != nil {
		return nil, err
	}
	tokens := make([]*Token, 0)
	jsonErr := json.Unmarshal(data, &tokens)
	if jsonErr != nil {
		return nil, &UnmarshalError{msg: multierror.Append(errDecoding, jsonErr).Error()}
	}
	return tokens, nil
}

// GetToken gets a token
// http://rundeck.org/docs/api/index.html#get-a-token
func (c *Client) GetToken(tokenID string) (*Token, error) {
	if err := c.checkRequiredAPIVersion(responses.TokenResponse{}); err != nil {
		return nil, err
	}

	data, err := c.httpGet("token/"+tokenID, requestJSON(), requestExpects(200))
	if err != nil {
		return nil, err
	}
	token := &Token{}
	jsonErr := json.Unmarshal(data, &token)
	if jsonErr != nil {
		return nil, &UnmarshalError{msg: multierror.Append(errDecoding, jsonErr).Error()}
	}
	return token, nil
}

// CreateToken creates a token
// http://rundeck.org/docs/api/index.html#create-a-token
func (c *Client) CreateToken(username string, opts ...TokenOption) (*Token, error) {
	if err := c.checkRequiredAPIVersion(responses.TokenResponse{}); err != nil {
		return nil, err
	}

	tokenRequest := &TokenRequest{}
	for _, opt := range opts {
		if err := opt(tokenRequest); err != nil {
			return nil, &OptionError{msg: multierror.Append(errOption, err).Error()}
		}
	}
	if len(tokenRequest.Roles) == 0 {
		tokenRequest.Roles = "*"
	}
	tokenRequest.User = username
	newToken, marshalErr := json.Marshal(tokenRequest)
	if marshalErr != nil {
		return nil, marshalErr
	}
	url := "tokens"
	data, err := c.httpPost(url, requestJSON(), withBody(bytes.NewReader(newToken)), requestExpects(201))
	if err != nil {
		return nil, err
	}
	token := &Token{}
	jsonErr := json.Unmarshal(data, &token)
	if jsonErr != nil {
		return nil, &UnmarshalError{msg: multierror.Append(errDecoding, jsonErr).Error()}
	}
	return token, nil
}

// DeleteToken deletes a token
// http://rundeck.org/docs/api/index.html#delete-a-token
func (c *Client) DeleteToken(token string) error {
	if err := c.checkRequiredAPIVersion(responses.TokenResponse{}); err != nil {
		return err
	}
	url := fmt.Sprintf("token/%s", token)
	_, err := c.httpDelete(url, requestJSON(), requestExpects(204))
	return err
}
