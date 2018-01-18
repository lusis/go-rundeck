package rundeck

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	multierror "github.com/hashicorp/go-multierror"
	httpclient "github.com/lusis/go-rundeck/pkg/httpclient"
	requests "github.com/lusis/go-rundeck/pkg/rundeck/requests"
	responses "github.com/lusis/go-rundeck/pkg/rundeck/responses"
)

// Job represents a rundeck job
type Job struct {
	responses.JobResponse
}

// JobList is a list of rundeck jobs
type JobList []Job

// JobMetaData is the result of getting a job's metadata
type JobMetaData struct {
	responses.JobMetaDataResponse
}

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
	if err := c.checkRequiredAPIVersion(responses.JobMetaDataResponse{}); err != nil {
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
	if err := c.checkRequiredAPIVersion(responses.JobYAMLResponse{}); err != nil {
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
// http://rundeck.org/docs/api/index.html#get-job-metadata
func (c *Client) GetJobInfo(id string) (*JobMetaData, error) {
	return c.GetJobMetaData(id)
}

// DeleteJob deletes a job
// http://rundeck.org/docs/api/index.html#deleting-a-job-definition
func (c *Client) DeleteJob(id string) error {
	if err := c.checkRequiredAPIVersion(responses.GenericVersionedResponse{}); err != nil {
		return err
	}
	_, err := c.httpDelete("job/"+id, httpclient.ExpectStatus(204))
	return err

}

// ExportJob exports a job
// http://rundeck.org/docs/api/index.html#exporting-jobs
func (c *Client) ExportJob(id string, format string) ([]byte, error) {
	if err := c.checkRequiredAPIVersion(responses.JobYAMLResponse{}); err != nil {
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
	if err := c.checkRequiredAPIVersion(responses.ExecutionResponse{}); err != nil {
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
func (c *Client) ListJobs(projectID string) (JobList, error) {
	data := JobList{}
	if err := c.checkRequiredAPIVersion(responses.JobsResponse{}); err != nil {
		return data, err
	}
	url := fmt.Sprintf("project/%s/jobs", projectID)
	res, err := c.httpGet(url, requestJSON(), requestExpects(200))
	if err != nil {
		return data, err
	}
	if err := json.Unmarshal(res, &data); err != nil {
		return data, &UnmarshalError{msg: multierror.Append(errDecoding, err).Error()}
	}
	return data, nil
}

// BulkJobDelete deletes jobs in bulk
// http://rundeck.org/docs/api/index.html#bulk-job-delete
func (c *Client) BulkJobDelete(ids ...string) error {
	if err := c.checkRequiredAPIVersion(responses.BulkDeleteJobResponse{}); err != nil {
		return err
	}
	return fmt.Errorf("not yet implemented")
}

// GetExecutionsForJob gets executions for a job
// http://rundeck.org/docs/api/index.html#getting-executions-for-a-job
func (c *Client) GetExecutionsForJob(jobid string) error {
	if err := c.checkRequiredAPIVersion(responses.JobExecutionsResponse{}); err != nil {
		return err
	}
	return fmt.Errorf("not yet implemented")
}

// DeleteAllExecutionsForJob deletes all executions for a job
// http://rundeck.org/docs/api/index.html#delete-all-executions-for-a-job
func (c *Client) DeleteAllExecutionsForJob(jobid string) (*DeletedExecutions, error) {
	if err := c.checkRequiredAPIVersion(responses.BulkDeleteExecutionsResponse{}); err != nil {
		return nil, err
	}
	data := &DeletedExecutions{}

	u := fmt.Sprintf("job/%s/executions", jobid)
	res, err := c.httpDelete(u,
		accept("application/json"),
		requestExpects(200))
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(res, data); err != nil {
		return nil, &UnmarshalError{msg: multierror.Append(errEncoding, err).Error()}
	}
	return data, nil
}

// UploadFileForJobOption uploads a file for a job 'file' option type
// http://rundeck.org/docs/api/index.html#upload-a-file-for-a-job-option
func (c *Client) UploadFileForJobOption(ids ...string) error {
	if err := c.checkRequiredAPIVersion(responses.JobOptionFileUploadResponse{}); err != nil {
		return err
	}
	return fmt.Errorf("not yet implemented")
}

// ListFilesUploadedForJob lists files that have been uploaded for a job
// http://rundeck.org/docs/api/index.html#list-files-uploaded-for-a-job
func (c *Client) ListFilesUploadedForJob(ids ...string) error {
	if err := c.checkRequiredAPIVersion(responses.UploadedJobInputFilesResponse{}); err != nil {
		return err
	}
	return fmt.Errorf("not yet implemented")
}

// GetUploadedFileInfo gets info about an uploaded file
// http://rundeck.org/docs/api/index.html#get-info-about-an-uploaded-file
func (c *Client) GetUploadedFileInfo(ids ...string) error {
	if err := c.checkRequiredAPIVersion(responses.UploadedJobInputFileResponse{}); err != nil {
		return err
	}
	return fmt.Errorf("not yet implemented")
}
