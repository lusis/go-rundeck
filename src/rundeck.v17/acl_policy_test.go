package rundeck

import (
	"encoding/xml"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
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
	var s AclPolicies
	err = xml.Unmarshal(xmlData, &s)
	if err != nil {
		t.Fatalf(err.Error())
	}
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
