package rundeck

import (
	"encoding/xml"
	"fmt"
)

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

func (c *RundeckClient) GetTokens() (data Tokens, err error) {
	u := make(map[string]string)
	var res []byte
	err = c.Get(&res, "tokens", u)
	if err != nil {
		return data, err
	} else {
		xml.Unmarshal(res, &data)
		return data, nil
	}
}

func (c *RundeckClient) GetUserTokens(user string) (data Tokens, err error) {
	u := make(map[string]string)
	var res []byte
	err = c.Get(&res, "tokens/"+user, u)
	if err != nil {
		return data, err
	} else {
		xml.Unmarshal(res, &data)
		return data, nil
	}
}

func (c *RundeckClient) GetToken(tokenId string) (data Token, err error) {
	u := make(map[string]string)
	var res []byte
	err = c.Get(&res, "token/"+tokenId, u)
	if err != nil {
		return data, err
	} else {
		xml.Unmarshal(res, &data)
		return data, nil
	}
}

func (c *RundeckClient) CreateToken(u string) (token string, e error) {
	var res []byte
	var t Token
	url := fmt.Sprintf("tokens/%s", u)
	err := c.Post(&res, url, nil, nil)
	if err != nil {
		return token, err
	} else {
		xml.Unmarshal(res, &t)
		return t.ID, nil
	}
}

func (c *RundeckClient) DeleteToken(token string) error {
	url := fmt.Sprintf("token/%s", token)
	err := c.Delete(url, nil)
	if err != nil {
		return err
	}
	return nil
}
