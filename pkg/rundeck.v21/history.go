package rundeck

import (
	"encoding/json"

	multierror "github.com/hashicorp/go-multierror"
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
		requestExpects(200),
	}
	res, err := c.httpGet("project/"+project+"/history", options...)
	if err != nil {
		return nil, err
	}
	if jsonErr := json.Unmarshal(res, data); jsonErr != nil {
		return nil, &UnmarshalError{msg: multierror.Append(errDecoding, jsonErr).Error()}
	}
	return data, nil
}
