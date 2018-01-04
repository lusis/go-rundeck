package rundeck

import (
	"encoding/json"

	multierror "github.com/hashicorp/go-multierror"
	responses "github.com/lusis/go-rundeck/pkg/rundeck.v21/responses"
)

// SystemInfo represents the rundeck server system info output
type SystemInfo responses.SystemInfoResponse

// GetSystemInfo gets system information from the rundeck server
func (c *Client) GetSystemInfo() (*SystemInfo, error) {
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
