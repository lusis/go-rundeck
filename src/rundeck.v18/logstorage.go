package rundeck

import (
	"encoding/xml"
)

type LogStorage struct {
	XMLName         xml.Name `xml:"logStorage"`
	Enabled         bool     `xml:"enabled,attr"`
	PluginName      string   `xml:"enabled,attr"`
	SucceededCount  int64    `xml:"succeededCount"`
	FailedCount     int64    `xml:"failedCount"`
	QueuedCount     int64    `xml:"queuedCount"`
	TotalCount      int64    `xml:"TotalCount"`
	IncompleteCount int64    `xml:"incompleteCount"`
	MissingCount    int64    `xml:"missingCount"`
}

func (c *RundeckClient) GetLogstorage() (data LogStorage, err error) {
	u := make(map[string]string)
	var res []byte
	err = c.Get(&res, "system/logstorage", u)
	if err != nil {
		return data, err
	} else {
		xml.Unmarshal(res, &data)
		return data, nil
	}
}
