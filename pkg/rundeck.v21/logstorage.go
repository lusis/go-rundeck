package rundeck

import (
	"encoding/xml"
)

// LogStorage represents log storage
type LogStorage struct {
	XMLName         xml.Name `xml:"logStorage"`
	Enabled         bool     `xml:"enabled,attr"`
	PluginName      string   `xml:"pluginName,attr"`
	SucceededCount  int64    `xml:"succeededCount"`
	FailedCount     int64    `xml:"failedCount"`
	QueuedCount     int64    `xml:"queuedCount"`
	TotalCount      int64    `xml:"TotalCount"`
	IncompleteCount int64    `xml:"incompleteCount"`
	MissingCount    int64    `xml:"missingCount"`
}

// GetLogStorage gets the logstorage
func (c *Client) GetLogStorage() (*LogStorage, error) {
	ls := &LogStorage{}
	data, err := c.httpGet("system/logstorage", requestXML())
	if err != nil {
		return nil, err
	}
	xmlErr := xml.Unmarshal(data, &ls)
	return ls, xmlErr
}
