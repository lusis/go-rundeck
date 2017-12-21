package rundeck

import "encoding/xml"

// Error represents a rundeck xml error
type Error struct {
	XMLName    xml.Name `xml:"result"`
	Error      bool     `xml:"error,attr"`
	APIVersion int      `xml:"apiversion,attr"`
	Message    string   `xml:"error>message"`
}

// JSONError represents a rundeck json error
type JSONError struct {
	IsError    bool   `json:"error"`
	APIVersion int    `json:"apiVersion"`
	ErrorCode  string `json:"errorCode"`
	Message    string `json:"message"`
}

//func GetRundeckError(data []byte)
