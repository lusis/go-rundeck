package rundeck

import "encoding/xml"

// RundeckError represents a rundeck xml error
type RundeckError struct {
	XMLName    xml.Name `xml:"result"`
	Error      bool     `xml:"error,attr"`
	APIVersion string   `xml:"apiversion,attr"`
	Message    string   `xml:"error>message"`
}
