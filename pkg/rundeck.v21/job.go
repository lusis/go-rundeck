package rundeck

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	multierror "github.com/hashicorp/go-multierror"
	httpclient "github.com/lusis/go-rundeck/pkg/httpclient"
	requests "github.com/lusis/go-rundeck/pkg/rundeck.v21/requests"
	responses "github.com/lusis/go-rundeck/pkg/rundeck.v21/responses"
)

// Job represents a rundeck job
type Job responses.JobResponse

// JobList is a list of rundeck jobs
type JobList responses.JobsResponse

// JobMetaData is the result of getting a job's metadata
type JobMetaData responses.JobMetaDataResponse

// JobOption represents a job option
type JobOption struct {
	Description string
	Name        string
	Regex       string
	Required    bool
	Value       string
}

// RunJobOption is a type for functional options
type RunJobOption func(*requests.RunJobRequest) error

// RunJobAs runs a job as a specific user
func RunJobAs(u string) RunJobOption {
	return func(r *requests.RunJobRequest) error {
		r.AsUser = u
		return nil
	}
}

// RunJobArgs runs the job with the specified arg string
func RunJobArgs(a string) RunJobOption {
	return func(r *requests.RunJobRequest) error {
		r.ArgString = a
		return nil
	}
}

// RunJobFilter runs the job with the specified filter
func RunJobFilter(a string) RunJobOption {
	return func(r *requests.RunJobRequest) error {
		r.Filter = a
		return nil
	}
}

// RunJobOpts runs the job with the specified filter
func RunJobOpts(a map[string]string) RunJobOption {
	return func(r *requests.RunJobRequest) error {
		r.Options = a
		return nil
	}
}

// RunJobLogLevel runs the job with the specified log level
func RunJobLogLevel(l string) RunJobOption {
	return func(r *requests.RunJobRequest) error {
		r.LogLevel = l
		return nil
	}
}

// RunJobRunAt runs the specified job at the specified time
func RunJobRunAt(t time.Time) RunJobOption {
	return func(r *requests.RunJobRequest) error {
		r.RunAtTime = &requests.JSONTime{t} // nolint: vet
		return nil
	}
}

// GetJobMetaData gets a job's metadata
// http://rundeck.org/docs/api/index.html#get-job-metadata
func (c *Client) GetJobMetaData(id string) (*JobMetaData, error) {
	if _, err := c.hasRequiredAPIVersion(18, maxRundeckVersionInt); err != nil {
		return nil, err
	}
	data := &JobMetaData{}
	res, err := c.httpGet("job/"+id+"/info", requestJSON(), requestExpects(200))
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(res, data); err != nil {
		return nil, &UnmarshalError{msg: multierror.Append(errDecoding, err).Error()}
	}
	return data, nil
}

// GetJobDefinition gets a job definition
// http://rundeck.org/docs/api/index.html#getting-a-job-definition
func (c *Client) GetJobDefinition(id string, format string) ([]byte, error) {
	if _, err := c.hasRequiredAPIVersion(1, maxRundeckVersionInt); err != nil {
		return nil, err
	}
	options := []httpclient.RequestOption{
		accept("application/" + format),
		contentType("application/x-www-form-urlencoded"),
		queryParams(map[string]string{"format": format}),
		requestExpects(200),
	}
	res, err := c.httpGet("job/"+id, options...)
	if err != nil {
		return nil, err
	}
	return res, nil

}

// GetJobInfo gets a job's details
func (c *Client) GetJobInfo(id string) (*JobMetaData, error) {
	return c.GetJobMetaData(id)
}

// DeleteJob deletes a job
// http://rundeck.org/docs/api/index.html#deleting-a-job-definition
func (c *Client) DeleteJob(id string) error {
	if _, err := c.hasRequiredAPIVersion(minJSONSupportedAPIVersion, maxRundeckVersionInt); err != nil {
		return err
	}
	return c.httpDelete("job/"+id, httpclient.ExpectStatus(204))

}

// ExportJob exports a job
// http://rundeck.org/docs/api/index.html#exporting-jobs
func (c *Client) ExportJob(id string, format string) ([]byte, error) {
	if _, err := c.hasRequiredAPIVersion(minJSONSupportedAPIVersion, maxRundeckVersionInt); err != nil {
		return nil, err
	}
	if format != "xml" && format != "yaml" {
		errString := fmt.Sprintf("Unknown/unsupported format \"%s\"", format)
		return nil, errors.New(errString)
	}
	opts := make(map[string]string)
	opts["format"] = format
	res, err := c.httpGet("job/"+id, queryParams(opts), requestExpects(200))
	if err != nil {
		return nil, err
	}
	return res, nil
}

// RunJob runs a job
// http://rundeck.org/docs/api/index.html#running-a-job
func (c *Client) RunJob(id string, opts ...RunJobOption) (*Execution, error) {
	if _, err := c.hasRequiredAPIVersion(18, maxRundeckVersionInt); err != nil {
		return nil, err
	}
	jobOpts := &requests.RunJobRequest{}
	data := &Execution{}
	for _, opt := range opts {
		if err := opt(jobOpts); err != nil {
			return nil, &OptionError{msg: multierror.Append(errOption, err).Error()}
		}
	}
	body := bytes.NewReader([]byte("{}"))
	if jobOpts != nil {
		req, _ := json.Marshal(jobOpts)
		body = bytes.NewReader(req)
	}
	res, pErr := c.httpPost("job/"+id+"/run", withBody(body), requestJSON(), requestExpects(200))
	if pErr != nil {
		return nil, pErr
	}
	if jsonErr := json.Unmarshal(res, data); jsonErr != nil {
		return nil, &UnmarshalError{msg: multierror.Append(errDecoding, jsonErr).Error()}
	}
	return data, nil
}

// ListJobs lists the jobs for a project
// http://rundeck.org/docs/api/index.html#listing-jobs
func (c *Client) ListJobs(projectID string) (*JobList, error) {
	if _, err := c.hasRequiredAPIVersion(17, maxRundeckVersionInt); err != nil {
		return nil, err
	}
	data := &JobList{}
	url := fmt.Sprintf("project/%s/jobs", projectID)
	res, err := c.httpGet(url, requestJSON(), requestExpects(200))
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(res, &data); err != nil {
		return nil, &UnmarshalError{msg: multierror.Append(errDecoding, err).Error()}
	}
	return data, nil
}

// BulkJobDelete deletes jobs in bulk
// http://rundeck.org/docs/api/index.html#bulk-job-delete
func (c *Client) BulkJobDelete(ids ...string) error {
	if _, err := c.hasRequiredAPIVersion(minJSONSupportedAPIVersion, maxRundeckVersionInt); err != nil {
		return err
	}
	return fmt.Errorf("not yet implemented")
}

// GetExecutionsForJob gets executions for a job
// http://rundeck.org/docs/api/index.html#getting-executions-for-a-job
func (c *Client) GetExecutionsForJob(jobid string) error {
	if _, err := c.hasRequiredAPIVersion(minJSONSupportedAPIVersion, maxRundeckVersionInt); err != nil {
		return err
	}
	return fmt.Errorf("not yet implemented")
}

// DeleteAllExecutionsForJob deletes all executions for a job
// http://rundeck.org/docs/api/index.html#delete-all-executions-for-a-job
func (c *Client) DeleteAllExecutionsForJob(jobid string) error {
	if _, err := c.hasRequiredAPIVersion(minJSONSupportedAPIVersion, maxRundeckVersionInt); err != nil {
		return err
	}
	return fmt.Errorf("not yet implemented")
}
