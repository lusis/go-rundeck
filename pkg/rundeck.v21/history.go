package rundeck

import (
	"encoding/json"

	httpclient "github.com/lusis/go-rundeck/pkg/httpclient"
	responses "github.com/lusis/go-rundeck/pkg/rundeck.v21/responses"
)

// History represents a project history
type History responses.HistoryResponse

// GetHistory returns the history for a project
func (c *Client) GetHistory(project string, opts ...map[string]string) (*History, error) {
	u := make(map[string]string)
	for _, opt := range opts {
		for k, v := range opt {
			u[k] = v
		}
	}
	data := &History{}
	options := []httpclient.RequestOption{
		accept("application/json"),
		contentType("application/x-www-form-urlencoded"),
		queryParams(u),
	}
	res, err := c.httpGet("project/"+project+"/history", options...)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(res, data); err != nil {
		return nil, err
	}
	return data, nil
}
