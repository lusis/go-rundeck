package rundeck

import "encoding/xml"

// Result represents an XML result
type Result struct {
	XMLName         xml.Name `xml:"result"`
	Succeeded       bool     `xml:"success,attr,omitempty"`
	Errored         bool     `xml:"error,attr,omitempty"`
	APIVersion      string   `xml:"apiversion,attr,omitempty"`
	SuccessMessages []string `xml:"success>message,omitempty"`
	ErrorMessages   []string `xml:"error>message,omitempty"`
}
