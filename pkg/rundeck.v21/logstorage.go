package rundeck

import (
	"encoding/json"

	"github.com/lusis/go-rundeck/pkg/rundeck.v21/responses"
)

// LogStorage represents log storage
type LogStorage responses.LogStorageResponse

// GetLogStorage gets the logstorage
func (c *Client) GetLogStorage() (*LogStorage, error) {
	ls := &LogStorage{}
	data, err := c.httpGet("system/logstorage", requestJSON(), requestExpects(200))
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(data, &ls); err != nil {
		return nil, err
	}
	return ls, nil
}
