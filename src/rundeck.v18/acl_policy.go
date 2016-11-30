package rundeck

import (
	"encoding/xml"
)

type AclPolicies struct {
	XMLName  xml.Name        `xml:"resource"`
	Path     string          `xml:"path,attr"`
	Type     string          `xml:"type,attr"`
	Href     string          `xml:"href,attr"`
	Contents AclListContents `xml:"contents,omitempty"`
}

type AclListContents struct {
	XMLName   xml.Name `xml:"contents"`
	Resources []struct {
		XMLName xml.Name `xml:"resource"`
		Path    string   `xml:"path,attr"`
		Type    string   `xml:"type,attr"`
		Href    string   `xml:"href,attr"`
		Name    string   `xml:"name,attr"`
	} `xml:"resource,omitempty"`
}

func (c *RundeckClient) GetSystemAclPolicies() (data AclPolicies, err error) {
	u := make(map[string]string)
	var res []byte
	err = c.Get(&res, "system/acl/", u)
	if err != nil {
		return data, err
	} else {
		xml.Unmarshal(res, &data)
		return data, nil
	}
}
