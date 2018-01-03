package rundeck

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	httpclient "github.com/lusis/go-rundeck/pkg/httpclient"
	requests "github.com/lusis/go-rundeck/pkg/rundeck.v21/requests"
	responses "github.com/lusis/go-rundeck/pkg/rundeck.v21/responses"
	yaml "gopkg.in/yaml.v2"
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

// RunJobAt runs the job at the specified time
func RunJobAt(t time.Time) RunJobOption {
	return func(r *requests.RunJobRequest) error {
		r.RunAtTime = &requests.JSONTime{t} // nolint: vet
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
		r.RunAtTime = &requests.JSONTime{t}
		return nil
	}
}

// GetJobMetaData gets a job's metadata
func (c *Client) GetJobMetaData(id string) (*JobMetaData, error) {
	data := &JobMetaData{}
	res, err := c.httpGet("job/"+id+"/info", requestJSON())
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(res, data); err != nil {
		return nil, err
	}
	return data, nil
}

// GetJobDefinition gets a job definition
func (c *Client) GetJobDefinition(id string, format string) ([]byte, error) {
	options := []httpclient.RequestOption{
		accept("application/" + format),
		contentType("application/x-www-form-urlencoded"),
		queryParams(map[string]string{"format": format}),
	}
	return c.httpGet("job/"+id, options...)

}

// GetJobInfo gets a job's details
func (c *Client) GetJobInfo(id string) (*JobMetaData, error) {
	return c.GetJobMetaData(id)
}

// DeleteJob deletes a job
func (c *Client) DeleteJob(id string) error {
	return c.httpDelete("job/"+id, httpclient.ExpectStatus(404))

}

// ExportJob exports a job
func (c *Client) ExportJob(id string, format string) (string, error) {
	if format != "xml" && format != "yaml" {
		errString := fmt.Sprintf("Unknown/unsupported format \"%s\"", format)
		return "", errors.New(errString)
	}
	opts := make(map[string]string)
	opts["format"] = format
	res, err := c.httpGet("job/"+id, queryParams(opts))
	if err != nil {
		return "", err
	}
	return string(res), nil
}

// GetJobOpts returns the required options for a job
func (c *Client) GetJobOpts(j string) ([]*JobOption, error) {
	options := make([]*JobOption, 0)
	data := &responses.JobYAMLResponse{}
	res, err := c.httpGet("job/"+j, accept("application/yaml"))
	if err != nil {
		return nil, err
	}
	if err := yaml.Unmarshal(res, &data); err != nil {
		return nil, err
	}
	if data != nil {
		for _, d := range *data {
			for _, o := range d.Options {
				options = append(options, &JobOption{
					Description: o.Description,
					Required:    o.Required,
					Regex:       o.Regex,
					Name:        o.Name,
					Value:       o.Value,
				})
			}
		}
	}
	return options, nil
}

// GetRequiredOpts returns the required options for a job
func (c *Client) GetRequiredOpts(j string) (map[string]string, error) {
	u := make(map[string]string)
	data := &responses.JobYAMLResponse{}
	res, err := c.httpGet("job/"+j, accept("application/yaml"))
	if err != nil {
		return u, err
	}
	if err := yaml.Unmarshal(res, &data); err != nil {
		return nil, err
	}
	if data != nil {
		for _, d := range *data {
			for _, o := range d.Options {
				if o.Required {
					if o.Value == "" {
						u[o.Name] = "<no default>"
					} else {
						u[o.Name] = o.Value
					}
				}
			}
		}
	}
	return u, nil
}

// RunJob runs a job
func (c *Client) RunJob(id string, opts ...RunJobOption) (*Execution, error) {
	jobOpts := &requests.RunJobRequest{}
	data := &Execution{}
	for _, opt := range opts {
		if err := opt(jobOpts); err != nil {
			return nil, err
		}
	}
	body := bytes.NewReader([]byte("{}"))
	if jobOpts != nil {
		req, err := json.Marshal(jobOpts)
		if err != nil {
			return nil, err
		}
		body = bytes.NewReader(req)
	}
	res, pErr := c.httpPost("job/"+id+"/run", withBody(body), requestJSON())
	if pErr != nil {
		return nil, pErr
	}
	jsonErr := json.Unmarshal(res, data)
	return data, jsonErr
}

// FindJobByName runs a job by name
func (c *Client) FindJobByName(name string, project string) (*JobMetaData, error) {
	jobs, err := c.ListJobs(project)
	if err != nil {
		return nil, err
	}
	if len(*jobs) > 0 {
		for _, d := range *jobs {
			if d.Name == name {
				job, joblistErr := c.GetJobInfo(d.ID)
				if err != nil {
					return nil, joblistErr
				}
				return job, nil
			}
		}
	}
	return nil, errors.New("No matches found")
}

// ListJobs lists the jobs for a project
func (c *Client) ListJobs(projectID string) (*JobList, error) {
	data := &JobList{}
	url := fmt.Sprintf("project/%s/jobs", projectID)
	res, err := c.httpGet(url, requestJSON())
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(res, &data); err != nil {
		return nil, err
	}
	return data, nil
}
