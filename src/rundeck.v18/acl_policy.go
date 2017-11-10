package rundeck

import (
	"encoding/xml"
	"fmt"
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

func (c *RundeckClient) CreateSystemAclPolicy(name string, contents []byte) error {
	var res []byte
	u := make(map[string]string)
	u["content_type"] = "application/yaml"
	url := fmt.Sprintf("system/acl/%s.aclpolicy", name)
	payload := fmt.Sprintf("<contents><![CDATA[%s]]></contents>", string(contents))
	fmt.Printf("%s\n", payload)
	err := c.Post(&res, url, []byte(payload), u)
	if err != nil {
		fmt.Printf("%#v\n", string(res))
		return err
	}
	return nil
}
