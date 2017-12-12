package rundeck

import (
	"encoding/xml"
	"fmt"
)

// Tokens is a collection of `Token`
type Tokens struct {
	XMLName  xml.Name `xml:"tokens"`
	Count    int64    `xml:"count,attr"`
	AllUsers *bool    `xml:"allusers,omitempty"`
	User     *string  `xml:"user,attr"`
	Tokens   []*Token `xml:"token"`
}

// Token represents a user and token
type Token struct {
	XMLName xml.Name `xml:"token"`
	ID      string   `xml:"id,attr"`
	User    string   `xml:"user,attr"`
}

// GetTokens gets all tokens
func (c *Client) GetTokens() (data Tokens, err error) {
	u := make(map[string]string)
	var res []byte
	err = c.Get(&res, "tokens", u)
	if err != nil {
		return data, err
	}
	xmlErr := xml.Unmarshal(res, &data)
	return data, xmlErr
}

// GetUserTokens gets the api tokens for a user
func (c *Client) GetUserTokens(user string) (data Tokens, err error) {
	u := make(map[string]string)
	var res []byte
	err = c.Get(&res, "tokens/"+user, u)
	if err != nil {
		return data, err
	}
	xmlErr := xml.Unmarshal(res, &data)
	return data, xmlErr
}

// GetToken gets a token
func (c *Client) GetToken(tokenID string) (data Token, err error) {
	u := make(map[string]string)
	var res []byte
	err = c.Get(&res, "token/"+tokenID, u)
	if err != nil {
		return data, err
	}
	xmlErr := xml.Unmarshal(res, &data)
	return data, xmlErr
}

// CreateToken creates a token
func (c *Client) CreateToken(u string) (token string, e error) {
	var res []byte
	var t Token
	url := fmt.Sprintf("tokens/%s", u)
	err := c.Post(&res, url, nil, nil)
	if err != nil {
		return token, err
	}
	xmlErr := xml.Unmarshal(res, &t)
	return t.ID, xmlErr
}

// DeleteToken deletes a token
func (c *Client) DeleteToken(token string) error {
	url := fmt.Sprintf("token/%s", token)
	err := c.Delete(url, nil)
	return err
}
