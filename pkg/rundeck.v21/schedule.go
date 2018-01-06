package rundeck

import (
	"bytes"
	"encoding/json"

	multierror "github.com/hashicorp/go-multierror"
	requests "github.com/lusis/go-rundeck/pkg/rundeck.v21/requests"
	responses "github.com/lusis/go-rundeck/pkg/rundeck.v21/responses"
)

// DisableSchedule disables a scheduled job
func (c *Client) DisableSchedule(id string) (bool, error) {
	t := &responses.ToggleResponse{}
	res, err := c.httpPost("job/"+id+"/schedule/disable", requestJSON(), requestExpects(200))
	if err != nil {
		return false, err
	}
	if jsonErr := json.Unmarshal(res, t); jsonErr != nil {
		return false, &UnmarshalError{msg: multierror.Append(errDecoding, jsonErr).Error()}
	}
	return t.Success, nil
}

// EnableSchedule enables a scheduled job
func (c *Client) EnableSchedule(id string) (bool, error) {
	t := &responses.ToggleResponse{}
	res, err := c.httpPost("job/"+id+"/schedule/enable", requestExpects(200), requestJSON())
	if err != nil {
		return false, err
	}
	if jsonErr := json.Unmarshal(res, t); jsonErr != nil {
		return false, &UnmarshalError{msg: multierror.Append(errDecoding, jsonErr).Error()}
	}
	return t.Success, nil
}

// BulkEnableSchedule enables scheduled jobs in bulk
func (c *Client) BulkEnableSchedule(ids ...string) (*responses.BulkToggleResponse, error) {
	req := &requests.BulkToggleRequest{
		IDs: ids,
	}
	results := &responses.BulkToggleResponse{}
	data, _ := json.Marshal(req)
	res, err := c.httpPost("jobs/schedule/enable",
		withBody(bytes.NewReader(data)),
		requestJSON(),
		requestExpects(200))
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(res, results); err != nil {
		return nil, &UnmarshalError{msg: multierror.Append(errEncoding, err).Error()}
	}
	return results, nil
}

// BulkDisableSchedule enables scheduled jobs in bulk
func (c *Client) BulkDisableSchedule(ids ...string) (*responses.BulkToggleResponse, error) {

	req := &requests.BulkToggleRequest{
		IDs: ids,
	}
	results := &responses.BulkToggleResponse{}
	data, _ := json.Marshal(req)
	res, err := c.httpPost("jobs/schedule/disable",
		withBody(bytes.NewReader(data)),
		requestJSON(),
		requestExpects(200))
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(res, results); err != nil {
		return nil, &UnmarshalError{msg: multierror.Append(errEncoding, err).Error()}
	}
	return results, nil
}
