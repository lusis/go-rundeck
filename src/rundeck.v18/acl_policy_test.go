package rundeck

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"testing"
)

func TestListSystemAclPolicies(t *testing.T) {
	xmlfile, err := os.Open("assets/test/system_acl.xml")
	if err != nil {
		t.Fatalf(err.Error())
	}
	defer xmlfile.Close()
	xmlData, _ := ioutil.ReadAll(xmlfile)
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Header().Set("Content-Type", "application/xml")
		fmt.Fprintf(w, string(xmlData))
	}))
	defer server.Close()

	transport := &http.Transport{
		Proxy: func(req *http.Request) (*url.URL, error) {
			return url.Parse(server.URL)
		},
	}

	conf := &ClientConfig{
		BaseURL:    "http://127.0.0.1:8080/",
		Token:      "XXXXXXXXXXXXX",
		VerifySSL:  false,
		AuthMethod: "token",
		Transport:  transport,
	}
	client := NewClient(conf)
	//var s AclPolicies
	//err = xml.Unmarshal(xmlData, &s)
	//if err != nil {
	//	t.Fatalf(err.Error())
	//}
	s, err := client.GetSystemAclPolicies()
	assert.Equal(t, s.Path, "")
	assert.Equal(t, s.Type, "directory")
	assert.Equal(t, s.Href, "http://server/api/14/system/acl/")
	assert.Len(t, s.Contents.Resources, 2)
	assert.Equal(t, s.Contents.Resources[0].Path, "name.aclpolicy")
	assert.Equal(t, s.Contents.Resources[0].Type, "file")
	assert.Equal(t, s.Contents.Resources[0].Href, "http://server/api/14/system/acl/name.aclpolicy")
	assert.Equal(t, s.Contents.Resources[0].Name, "name.aclpolicy")
	assert.Equal(t, s.Contents.Resources[1].Name, "foo.aclpolicy")
}
