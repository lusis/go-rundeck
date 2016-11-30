package rundeck

import "encoding/xml"

type RundeckResult struct {
	XMLName         xml.Name `xml:"result"`
	Succeeded       bool     `xml:"success,attr,omitempty"`
	Errored         bool     `xml:"error,attr,omitempty"`
	ApiVersion      string   `xml:"apiversion,attr,omitempty"`
	SuccessMessages []string `xml:"success>message,omitempty"`
	ErrorMessages   []string `xml:"error>message,omitempty"`
}
