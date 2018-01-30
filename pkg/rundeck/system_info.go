package rundeck

import (
	"encoding/json"

	multierror "github.com/hashicorp/go-multierror"
	responses "github.com/lusis/go-rundeck/pkg/rundeck/responses"
)

// SystemInfo represents the rundeck server system info output
type SystemInfo struct {
	responses.SystemInfoResponse
}

// GetSystemInfo gets system information from the rundeck server
// http://rundeck.org/docs/api/index.html#system-info
func (c *Client) GetSystemInfo() (*SystemInfo, error) {
	if err := c.checkRequiredAPIVersion(responses.SystemInfoResponse{}); err != nil {
		return nil, err
	}
	ls := SystemInfo{}
	data, err := c.httpGet("system/info", requestJSON(), requestExpects(200))
	if err != nil {
		return nil, err
	}
	if jsonErr := json.Unmarshal(data, &ls); jsonErr != nil {
		return nil, &UnmarshalError{msg: multierror.Append(errDecoding, jsonErr).Error()}
	}
	return &ls, nil
}
