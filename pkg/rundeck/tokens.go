package rundeck

import (
	"bytes"
	"encoding/json"
	"fmt"

	multierror "github.com/hashicorp/go-multierror"
	responses "github.com/lusis/go-rundeck/pkg/rundeck/responses"
)

// Token is a token
type Token struct {
	responses.TokenResponse
}

// TokenOption is a type for options in creating new tokens
type TokenOption func(*Token) error

// TokenDuration is an option for setting the duration of a new token
func TokenDuration(duration string) TokenOption {
	return func(t *Token) error {
		t.Duration = duration
		return nil
	}
}

// TokenRoles is an option to set the roles for a new token
func TokenRoles(roles ...string) TokenOption {
	return func(t *Token) error {
		t.Roles = append(t.Roles, roles...)
		return nil
	}
}

// ListTokens gets all tokens for the current user
// http://rundeck.org/docs/api/index.html#list-tokens
func (c *Client) ListTokens() ([]*Token, error) {
	if _, err := c.hasRequiredAPIVersion(19, maxRundeckVersionInt); err != nil {
		return nil, err
	}
	tokens := make([]*Token, 0)
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
	if _, err := c.hasRequiredAPIVersion(19, maxRundeckVersionInt); err != nil {
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
func (c *Client) CreateToken(u string, opts ...TokenOption) (*Token, error) {
	if _, err := c.hasRequiredAPIVersion(19, maxRundeckVersionInt); err != nil {
		return nil, err
	}
	tokenRequest := &Token{}
	for _, opt := range opts {
		if err := opt(tokenRequest); err != nil {
			return nil, &OptionError{msg: multierror.Append(errOption, err).Error()}
		}
	}
	if len(tokenRequest.Roles) == 0 {
		tokenRequest.Roles = []string{"*"}
	}
	newToken, _ := json.Marshal(tokenRequest)
	url := fmt.Sprintf("tokens/%s", u)
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
	if _, err := c.hasRequiredAPIVersion(19, maxRundeckVersionInt); err != nil {
		return err
	}
	url := fmt.Sprintf("token/%s", token)
	return c.httpDelete(url, requestJSON(), requestExpects(204))
}
