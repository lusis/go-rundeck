package rundeck

import (
	"bytes"
	"encoding/json"

	multierror "github.com/hashicorp/go-multierror"
	requests "github.com/lusis/go-rundeck/pkg/rundeck/requests"
	responses "github.com/lusis/go-rundeck/pkg/rundeck/responses"
)

// DisableSchedule disables a scheduled job
// http://rundeck.org/docs/api/index.html#disable-scheduling-for-a-job
func (c *Client) DisableSchedule(id string) (bool, error) {
	if err := c.checkRequiredAPIVersion(responses.ToggleResponse{}); err != nil {
		return false, err
	}
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
// http://rundeck.org/docs/api/index.html#enable-scheduling-for-a-job
func (c *Client) EnableSchedule(id string) (bool, error) {
	if err := c.checkRequiredAPIVersion(responses.ToggleResponse{}); err != nil {
		return false, err
	}
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
// http://rundeck.org/docs/api/index.html#bulk-toggle-job-schedules
func (c *Client) BulkEnableSchedule(ids ...string) (*responses.BulkToggleResponse, error) {
	if err := c.checkRequiredAPIVersion(responses.BulkToggleResponse{}); err != nil {
		return nil, err
	}
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
// http://rundeck.org/docs/api/index.html#bulk-toggle-job-schedules
func (c *Client) BulkDisableSchedule(ids ...string) (*responses.BulkToggleResponse, error) {
	if err := c.checkRequiredAPIVersion(responses.BulkToggleResponse{}); err != nil {
		return nil, err
	}
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
