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

// GetLogstorage gets the logstorage
func (c *Client) GetLogstorage() (data LogStorage, err error) {
	u := make(map[string]string)
	var res []byte
	err = c.Get(&res, "system/logstorage", u)
	if err != nil {
		return data, err
	}
	xmlErr := xml.Unmarshal(res, &data)
	return data, xmlErr
}
