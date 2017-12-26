package rundeck

import (
	"encoding/json"

	responses "github.com/lusis/go-rundeck/pkg/rundeck.v21/responses"
)

// SystemInfo represents the rundeck server system info output
type SystemInfo responses.SystemInfoResponse

// GetSystemInfo gets system information from the rundeck server
func (c *Client) GetSystemInfo() (*SystemInfo, error) {
	ls := SystemInfo{}
	data, err := c.httpGet("system/info", requestJSON())
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(data, &ls); err != nil {
		return nil, err
	}
	return &ls, nil
}
