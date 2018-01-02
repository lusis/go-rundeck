package rundeck

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"

	httpclient "github.com/lusis/go-rundeck/pkg/httpclient"
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

// JobSequence represents a job sequence
type JobSequence struct {
	XMLName   xml.Name
	KeepGoing bool           `xml:"keepgoing,attr"`
	Strategy  string         `xml:"strategy,attr"`
	Steps     []SequenceStep `xml:"command"`
}

// SequenceStep represents a sequence step
type SequenceStep struct {
	XMLName        xml.Name
	Description    string      `xml:"description,omitempty"`
	JobRef         *JobRefStep `xml:"jobref,omitempty"`
	NodeStepPlugin *PluginStep `xml:"node-step-plugin,omitempty"`
	StepPlugin     *PluginStep `xml:"step-plugin,omitempty"`
	Exec           *string     `xml:"exec,omitempty"`
	*ScriptStep    `xml:",omitempty"`
}

// ExecStep represents an exec step
type ExecStep struct {
	XMLName xml.Name
	string  `xml:"exec,omitempty"`
}

// ScriptStep represents a script step
type ScriptStep struct {
	XMLName           xml.Name
	Script            *string `xml:"script,omitempty"`
	ScriptArgs        *string `xml:"scriptargs,omitempty"`
	ScriptFile        *string `xml:"scriptfile,omitempty"`
	ScriptURL         *string `xml:"scripturl,omitempty"`
	ScriptInterpreter *string `xml:"scriptinterpreter,omitempty"`
}

// PluginStep represents a plugin step
type PluginStep struct {
	XMLName       xml.Name
	Type          string `xml:"type,attr"`
	Configuration []struct {
		XMLName xml.Name `xml:"entry"`
		Key     string   `xml:"key,attr"`
		Value   string   `xml:"value,attr"`
	} `xml:"configuration>entry,omitempty"`
}

// JobRefStep represents a job reference step
type JobRefStep struct {
	XMLName  xml.Name
	Name     string `xml:"name,attr,omitempty"`
	Group    string `xml:"group,attr,omitempty"`
	NodeStep bool   `xml:"nodeStep,attr,omitempty"`
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
	return c.httpDelete("job/" + id)
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
/*
func (c *Client) RunJob(id string, options RunOptions) (*Executions, error) {
	var res []byte
	data := &Executions{}
	options.runAtTime = strings.Replace(options.RunAtTime.Format(time.RFC3339), "Z", "-", -1)
	opts, err := json.Marshal(options)
	if err != nil {
		return data, err
	}
	res, pErr := c.httpPost("job/"+id+"/run", withBody(bytes.NewReader(opts)), requestJSON())
	if pErr != nil {
		return data, pErr
	}
	xmlErr := xml.Unmarshal(res, &data)
	return data, xmlErr
}
*/

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
