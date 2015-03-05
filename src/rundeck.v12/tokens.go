package rundeck

import "encoding/xml"

type Tokens struct {
	XMLName  xml.Name `xml:"tokens"`
	Count    int64    `xml:"count,attr"`
	AllUsers *bool    `xml:"allusers,omitempty"`
	User     *string  `xml:"user,attr"`
	Tokens   []*Token `xml:"token"`
}

type Token struct {
	XMLName xml.Name `xml:"token"`
	ID      string   `xml:"id,attr"`
	User    string   `xml:"user,attr"`
}

func (c *RundeckClient) GetTokens() (Tokens, error) {
	u := make(map[string]string)
	var data Tokens
	err := c.Get(&data, "tokens", u)
	return data, err
}

func (c *RundeckClient) GetUserTokens(user string) (Tokens, error) {
	u := make(map[string]string)
	var data Tokens
	err := c.Get(&data, "tokens/"+user, u)
	return data, err
}

func (c *RundeckClient) GetToken(tokenId string) (Token, error) {
	u := make(map[string]string)
	var data Token
	err := c.Get(&data, "token/"+tokenId, u)
	return data, err
}
