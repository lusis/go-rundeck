package rundeck

import (
	"encoding/json"

	multierror "github.com/hashicorp/go-multierror"
	httpclient "github.com/lusis/go-rundeck/pkg/httpclient"
	responses "github.com/lusis/go-rundeck/pkg/rundeck/responses"
)

// History represents a project history
type History struct {
	responses.HistoryResponse
}

// ListHistory returns the history for a project
// http://rundeck.org/docs/api/index.html#listing-history
func (c *Client) ListHistory(project string, opts ...map[string]string) (*History, error) {
	if err := c.checkRequiredAPIVersion(responses.HistoryResponse{}); err != nil {
		return nil, err
	}
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
