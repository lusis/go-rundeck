package rundeck

import (
	"encoding/xml"
)

type SystemInfo struct {
	XMLName    xml.Name   `xml:"system"`
	Timestamp  TS         `xml:"timestamp"`
	Rundeck    Rundeck    `xml:"rundeck"`
	OS         OS         `xml:"os"`
	JVM        JVM        `xml:"jvm"`
	Stats      Stats      `xml:"stats"`
	Metrics    Metrics    `xml:"metrics"`
	ThreadDump ThreadDump `xml:"threadDump"`
}

type Metrics struct {
	XMLName     xml.Name `xml:"metrics"`
	Href        string   `xml:"href,attr"`
	ContentType string   `xml:"contentType,attr"`
}

type ThreadDump struct {
	XMLName     xml.Name `xml:"threadDump"`
	Href        string   `xml:"href,attr"`
	ContentType string   `xml:"contentType,attr"`
}

type Rundeck struct {
	XMLName    xml.Name `xml:"rundeck"`
	Version    string   `xml:"version"`
	ApiVersion int64    `xml:"apiversion"`
	Build      string   `xml:"build"`
	Node       string   `xml:"node"`
	Base       string   `xml:"base"`
	ServerUUID string   `xml:"serverUUID,omitempty"`
}

type TS struct {
	Epoch    string `xml:"epoch,attr"`
	Unit     string `xml:"unit,attr"`
	DateTime string `xml:"datetime"`
}

type OS struct {
	Arch    string `xml:"arch"`
	Name    string `xml:"name"`
	Version string `xml:"version"`
}

type JVM struct {
	Name                  string `xml:"name"`
	Vendor                string `xml:"vendor"`
	Version               string `xml:"version"`
	ImplementationVersion string `xml:"implementationVersion"`
}

type Stats struct {
	XMLName   xml.Name  `xml:"stats"`
	Uptime    Uptime    `xml:"uptime"`
	CPU       CPU       `xml:"cpu"`
	Memory    Memory    `xml:"memory"`
	Scheduler Scheduler `xml:"scheduler"`
	Threads   Threads   `xml:"threads"`
}

type Uptime struct {
	XMLName  xml.Name `xml:"uptime"`
	Duration string   `xml:"duration,attr"`
	Unit     string   `xml:"unit,attr"`
	Since    struct {
		XMLName  xml.Name
		TS       `xml:"since"`
		DateTime string `xml:"datetime"`
	} `xml:"since"`
}

type CPU struct {
	XMLName     xml.Name `xml:"cpu"`
	LoadAverage struct {
		XMLName xml.Name
		Unit    string  `xml:"unit,attr"`
		Value   float64 `xml:",chardata"`
	} `xml:"loadAverage"`
	Processors int64 `xml:"processors"`
}

type Memory struct {
	XMLName xml.Name `xml:"memory"`
	Unit    string   `xml:"unit,attr"`
	Max     int64    `xml:"max"`
	Free    int64    `xml:"free"`
	Total   int64    `xml:"total"`
}

type Scheduler struct {
	Running int64 `xml:"running"`
}

type Threads struct {
	Active int64 `xml:"active"`
}
