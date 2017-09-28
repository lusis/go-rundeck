package rundeck

import "encoding/xml"

type RundeckError struct {
	XMLName    xml.Name `xml:"result"`
	Error      bool     `xml:"error,attr"`
	ApiVersion string   `xml:"apiversion,attr"`
	Message    string   `xml:"error>message"`
}
