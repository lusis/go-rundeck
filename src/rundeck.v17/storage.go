package rundeck

import (
	"encoding/xml"
)

type Resources struct {
	Resources []Resource `xml:"resource"`
}

type Resource struct {
	XMLName      xml.Name     `xml:"resource"`
	Path         string       `xml:"path,attr"`
	ResourceType string       `xml:"type,attr"`
	URL          string       `xml:"url,attr"`
	Name         string       `xml:"name,omitempty"`
	MetaData     ResourceMeta `xml:"resource-meta,omitempty"`
	Contents     Contents     `xml:"contents,omitempty"`
}

type Contents struct {
	XMLName   xml.Name   `xml:"contents"`
	Count     int64      `xml:"count"`
	Resources []Resource `xml:"resource"`
}

type ResourceMeta struct {
	XMLName             xml.Name `xml:"resource-meta"`
	ContentType         string   `xml:"Rundeck-content-type"`
	Size                int64    `xml:"Rundeck-content-size"`
	CreationTime        string   `xml:"Rundeck-content-creation-time"`
	ModifyTime          string   `xml:"Rundeck-content-modify-time"`
	AuthCreatedUsername string   `xml:"Rundeck-auth-created-username"`
	KeyType             string   `xml:"Rundeck-key-type"`
}

// returns the raw data associated with a storage item
//func (r *Resource) GetRaw() (data string, err error) {
//
//}
