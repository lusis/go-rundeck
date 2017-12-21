package rundeck

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTokens(t *testing.T) {
	jsonfile, err := os.Open("assets/test/tokens.json")
	if err != nil {
		t.Fatal(err.Error())
	}
	defer func() { _ = jsonfile.Close() }()
	jsonData, _ := ioutil.ReadAll(jsonfile)
	var s []Token
	_ = json.Unmarshal(jsonData, &s)
	assert.Len(t, s, 2)
	assert.Equal(t, "admin", s[0].User)
	assert.Equal(t, "71b3dfe3-1dde-439f-9fe3-48f2ab0f47d2", s[0].ID)
	assert.Equal(t, "admin", s[0].Creator)
	assert.Len(t, s[0].Roles, 1)
	assert.Equal(t, "user", s[0].Roles[0])
	assert.False(t, s[0].Expired)
	assert.NotEmpty(t, s[0].Expiration)

}

func TestUserToken(t *testing.T) {
	jsonfile, err := os.Open("assets/test/token.json")
	if err != nil {
		t.Fatal(err.Error())
	}
	defer func() { _ = jsonfile.Close() }()
	jsonData, _ := ioutil.ReadAll(jsonfile)
	var s Token
	_ = json.Unmarshal(jsonData, &s)
	assert.Len(t, s.Roles, 1)
	assert.Equal(t, "admin", s.User)
	assert.Equal(t, "lK2iaQLEkf6rINMAYOXfrFNIpuwHRq67", s.Token)
	assert.Equal(t, "admin", s.Creator)
	assert.NotEmpty(t, s.Expiration)
	assert.False(t, s.Expired)
	assert.Equal(t, "54d6839c-7938-46b0-967d-146712a544b8", s.ID)
}
