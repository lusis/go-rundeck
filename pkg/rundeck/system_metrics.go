package rundeck

import (
	"encoding/json"

	multierror "github.com/hashicorp/go-multierror"
	responses "github.com/lusis/go-rundeck/pkg/rundeck/responses"
)

// SystemMetrics represents the rundeck server system info output
type SystemMetrics struct {
	Version    string                 `json:"version"`
	Gauges     map[string]interface{} `json:"gauges"`
	Counters   map[string]interface{} `json:"counters"`
	Histograms map[string]interface{} `json:"histograms"`
	Meters     map[string]interface{} `json:"meters"`
	Timers     map[string]interface{} `json:"timers"`
}

// type SystemMetrics []string

// GetMetrics gets system information from the rundeck server
// http://rundeck.org/docs/api/index.html#system-info
func (c *Client) GetMetrics() (*SystemMetrics, error) {
	if err := c.checkRequiredAPIVersion(responses.SystemInfoResponse{}); err != nil {
		return nil, err
	}
	ls := SystemMetrics{}
	// data, err := c.metricsHttpGet("/metrics/metrics?pretty=true", requestJSON(), requestExpects(200))
	data, err := c.httpGet("/metrics/metrics?pretty=true", requestJSON(), requestExpects(200))
	if err != nil {
		return nil, err
	}
	if jsonErr := json.Unmarshal(data, &ls); jsonErr != nil {
		return nil, &UnmarshalError{msg: multierror.Append(errDecoding, jsonErr).Error()}
	}
	return &ls, nil
}
