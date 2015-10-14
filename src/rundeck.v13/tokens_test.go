package rundeck

import (
	"encoding/xml"
	"io/ioutil"
	"os"
	"testing"
)

func TestTokens(t *testing.T) {
	xmlfile, err := os.Open("assets/test/tokens.xml")
	if err != nil {
		t.Fatal(err.Error())
	}
	defer xmlfile.Close()
	xmlData, _ := ioutil.ReadAll(xmlfile)
	var s Tokens
	xml.Unmarshal(xmlData, &s)
	intexpects(int64(len(s.Tokens)), s.Count, t)
	if &s.AllUsers == nil {
		t.Errorf("Missing AllUsers field")
	}
	strexpects(s.Tokens[0].User, "admin", t)
	strexpects(s.Tokens[0].ID, "XXXX", t)
	strexpects(s.Tokens[1].User, "bob", t)
	strexpects(s.Tokens[1].ID, "YYYY", t)

}

func TestUserToken(t *testing.T) {
	xmlfile, err := os.Open("assets/test/user_token.xml")
	if err != nil {
		t.Fatal(err.Error())
	}
	defer xmlfile.Close()
	xmlData, _ := ioutil.ReadAll(xmlfile)
	var s Tokens
	xml.Unmarshal(xmlData, &s)
	intexpects(int64(len(s.Tokens)), s.Count, t)
	usertoken := s.Tokens[0]
	strexpects(*s.User, usertoken.User, t)
	strexpects(usertoken.ID, "XXXX", t)
}
