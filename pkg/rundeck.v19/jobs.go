package rundeck

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

// Job represents a rundeck job
type Job struct {
	XMLName     xml.Name `xml:"job"`
	ID          string   `xml:"id,attr"`
	Name        string   `xml:"name"`
	Group       string   `xml:"group"`
	Project     string   `xml:"project"`
	Description string   `xml:"description,omitempty"`
	// These two come from Execution output
	AverageDuration int64   `xml:"averageDuration,attr,omitempty"`
	Options         Options `xml:"options,omitempty"`
	// These four come from Import output (depending on success,error,skipped)
	Index int    `xml:"index,attr,omitempty"`
	Href  string `xml:"href,attr,omitempty"`
	Error string `xml:"error,omitempty"`
	URL   string `xml:"url,omitempty"`
}

// JobInfo represents a rundeck jobinfo
type JobInfo struct {
	XMLName         xml.Name `xml:"job"`
	ID              string   `xml:"id,attr"`
	Href            string   `xml:"href,attr,omitempty"`
	Permalink       string   `xml:"permalink,attr,omitempty"`
	Scheduled       bool     `xml:"scheduled,attr,omitempty"`
	ScheduleEnabled bool     `xml:"scheduleEnabled,attr,omitempty"`
	Enabled         bool     `xml:"enabled,attr,omitempty"`
	AverageDuration int64    `xml:"averageDuration,attr,omitempty"`
	Name            string   `xml:"name,omitempty"`
	Group           string   `xml:"group,omitempty"`
	Project         string   `xml:"project,omitempty"`
	Description     string   `xml:"description,omitempty"`
}

// JobImportResultJob represents an imported job
type JobImportResultJob struct {
	XMLName xml.Name `xml:"job"`
	ID      string   `xml:"id,omitempty"`
	Name    string   `xml:"name"`
	Group   string   `xml:"group"`
	Project string   `xml:"project"`
	Index   int      `xml:"index,attr,omitempty"`
	Href    string   `xml:"href,attr,omitempty"`
	Error   string   `xml:"error,omitempty"`
	URL     string   `xml:"url,omitempty"`
}

// JobImportResult represents an imported job result
type JobImportResult struct {
	XMLName    xml.Name `xml:"result"`
	Success    bool     `xml:"success,attr,omitempty"`
	Error      bool     `xml:"error,attr,omitempty"`
	APIVersion int64    `xml:"apiversion,attr"`
	Succeeded  struct {
		XMLName xml.Name             `xml:"succeeded"`
		Count   int64                `xml:"count,attr"`
		Jobs    []JobImportResultJob `xml:"job,omitempty"`
	} `xml:"succeeded,omitempty"`
	Failed struct {
		XMLName xml.Name             `xml:"failed"`
		Count   int64                `xml:"count,attr"`
		Jobs    []JobImportResultJob `xml:"job,omitempty"`
	} `xml:"failed,omitempty"`
	Skipped struct {
		XMLName xml.Name             `xml:"skipped"`
		Count   int64                `xml:"count,attr"`
		Jobs    []JobImportResultJob `xml:"job,omitempty"`
	} `xml:"skipped,omitempty"`
}

// Options represents a group of xml `Option`
type Options struct {
	XMLName xml.Name
	Options []Option `xml:"option"`
}

// Option represents an xml option
type Option struct {
	XMLName xml.Name `xml:"option"`
	Name    string   `xml:"name,attr"`
	Value   string   `xml:"value,attr,omitempty"`
}

// Jobs is a collection of `Job`
type Jobs struct {
	XMLName xml.Name
	Count   int64 `xml:"count,attr"`
	Jobs    []Job `xml:"job"`
}

// RunOptions represents the options for a job
type RunOptions struct {
	LogLevel  string            `json:"loglevel,omitempty"`
	AsUser    string            `json:"asUser,omitempty"`
	Filter    string            `json:"filter,omitempty"`
	runAtTime string            `json:"runAtTime,omitempty"` // nolint: vet
	RunAtTime time.Time         `json:"-"`
	Options   map[string]string `json:"options,omitempty"`
	// backwards compatibility for now
	Arguments string `json:"argString,omitempty"`
}

// JobList represents a list of `Job`
type JobList struct {
	XMLName xml.Name   `xml:"joblist"`
	Job     JobDetails `xml:"job"`
}

// JobDetails represents the details of a `Job`
type JobDetails struct {
	ID                string          `xml:"id"`
	Name              string          `xml:"name"`
	LogLevel          string          `xml:"loglevel"`
	Description       string          `xml:"description,omitempty"`
	UUID              string          `xml:"uuid"`
	Group             string          `xml:"group"`
	Context           JobContext      `xml:"context"`
	Notification      JobNotification `xml:"notification"`
	MultipleExections bool            `xml:"multipleExecutions"`
	Dispatch          JobDispatch     `xml:"dispatch"`
	NodeFilters       struct {
		Filter []string `xml:"filter"`
	} `xml:"nodefilters"`
	Sequence JobSequence `xml:"sequence"`
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

// JobContext represents the context of a job
type JobContext struct {
	XMLName xml.Name     `xml:"context"`
	Project string       `xml:"project"`
	Options *[]JobOption `xml:"options>option,omitempty"`
}

// JobOptions is a collection of `JobOption`
type JobOptions struct {
	XMLName xml.Name
	Options []JobOption `xml:"option"`
}

// JobOption is a single option for a job
type JobOption struct {
	XMLName      xml.Name `xml:"option"`
	Name         string   `xml:"name,attr"`
	Required     bool     `xml:"required,attr,omitempty"`
	Secure       bool     `xml:"secure,attr,omitempty"`
	ValueExposed bool     `xml:"valueExposed,attr,omitempty"`
	DefaultValue string   `xml:"value,attr,omitempty"`
	Description  string   `xml:"description,omitempty"`
}

// JobNotifications is a collection of `JobNotification`
type JobNotifications struct {
	Notifications []JobNotification `xml:"notification,omitempty"`
}

// JobNotification represents a job notification
type JobNotification struct {
	XMLName   xml.Name   `xml:"notification"`
	OnStart   JobPlugins `xml:"onstart,omitempty"`
	OnSuccess JobPlugins `xml:"onsuccess,omitempty"`
	OnFailure JobPlugins `xml:"onfailure,omitempty"`
}

// JobPlugins is a collection on `JobPlugin`
type JobPlugins struct {
	Plugins []JobPlugin `xml:"plugin,omitempty"`
}

// JobPlugin represents a job plugin
type JobPlugin struct {
	XMLName       xml.Name               `xml:"plugin"`
	PluginType    string                 `xml:"type,attr"`
	Configuration JobPluginConfiguration `xml:"configuration,omitempty"`
}

// JobPluginConfiguration represents the configuration for a job plugin
type JobPluginConfiguration struct {
	XMLName xml.Name                      `xml:"configuration"`
	Entries []JobPluginConfigurationEntry `xml:"entry,omitempty"`
}

// JobPluginConfigurationEntry is an entry for a job plugin configuration
type JobPluginConfigurationEntry struct {
	Key   string `xml:"key,attr"`
	Value string `xml:"value,attr,omitempty"`
}

// JobDispatch represents a job dispatch
type JobDispatch struct {
	XMLName           xml.Name `xml:"dispatch"`
	ThreadCount       int64    `xml:"threadcount"`
	KeepGoing         bool     `xml:"keepgoing"`
	ExcludePrecedence bool     `xml:"excludePrecendence"`
	RankOrder         string   `xml:"rankOrder"`
}

// ImportParams represents the params for importing a job
type ImportParams struct {
	Filename string
	Format   string
	Dupe     string
	UUID     string
	Project  string
}

// GetJob gets a job
func (c *Client) GetJob(id string) (JobList, error) {
	u := make(map[string]string)
	var res []byte
	var data JobList
	if err := c.Get(&res, "job/"+id, u); err != nil {
		return data, err
	}
	xmlErr := xml.Unmarshal(res, &data)
	return data, xmlErr
}

// GetJobinfo gets a jobs detail
func (c *Client) GetJobinfo(id string) (JobInfo, error) {
	u := make(map[string]string)
	var res []byte
	var data JobInfo
	err := c.Get(&res, "job/"+id+"/info", u)
	if err != nil {
		return data, err
	}
	xmlErr := xml.Unmarshal(res, &data)
	return data, xmlErr
}

// DeleteJob deletes a job
func (c *Client) DeleteJob(id string) error {
	return c.Delete("job/"+id, nil)
}

// ExportJob exports a job
func (c *Client) ExportJob(id string, format string) (s string, e error) {
	if format != "xml" && format != "yaml" {
		errString := fmt.Sprintf("Unknown/unsupported format \"%s\"", format)
		return s, errors.New(errString)
	}
	var res []byte
	opts := make(map[string]string)
	opts["format"] = format
	err := c.Get(&res, "job/"+id, opts)
	if err != nil {
		e = err
	} else {
		s = string(res)

	}
	return s, e
}

// ImportJob imports a job
func (c *Client) ImportJob(j ImportParams) (string, error) {
	var res []byte
	var jobRes JobImportResult
	opts := make(map[string]string)
	opts["project"] = j.Project
	opts["format"] = j.Format
	opts["dupeOption"] = j.Dupe
	opts["uuidOption"] = j.UUID

	jobfile, err := os.Open(j.Filename)
	if err != nil {
		return "", err
	}
	defer func() { _ = jobfile.Close() }()
	data, readErr := ioutil.ReadAll(jobfile)
	if readErr != nil {
		return "", readErr
	}
	if postErr := c.Post(&res, "jobs/import", data, opts); postErr != nil {
		return "", postErr
	}
	xmlErr := xml.Unmarshal(res, &jobRes)
	if xmlErr != nil {
		return "", err
	}
	if jobRes.Skipped.Count > 0 {
		var errString []string
		for _, e := range jobRes.Skipped.Jobs {
			errString = append(errString, e.Error)
		}
		return "", errors.New(strings.Join(errString, "\n"))
	}
	if jobRes.Failed.Count > 0 {
		var errString []string
		for _, e := range jobRes.Failed.Jobs {
			errString = append(errString, e.Error)
		}
		return "", errors.New(strings.Join(errString, "\n"))
	}
	retStr := fmt.Sprintf("%s (ID: %s)", jobRes.Succeeded.Jobs[0].Name, jobRes.Succeeded.Jobs[0].ID)
	return retStr, nil
}

// GetRequiredOpts returns the required options for a job
func (c *Client) GetRequiredOpts(j string) (map[string]string, error) {
	u := make(map[string]string)
	var res []byte
	var data JobList
	err := c.Get(&res, "job/"+j, u)
	if err != nil {
		return u, err
	}
	xmlErr := xml.Unmarshal(res, &data)
	if xmlErr != nil {
		return u, xmlErr
	}
	if data.Job.Context.Options != nil {
		for _, o := range *data.Job.Context.Options {
			if o.Required {
				if o.DefaultValue == "" {
					u[o.Name] = "<no default>"
				} else {
					u[o.Name] = o.DefaultValue
				}
			}
		}
	}
	return u, nil
}

// RunJob runs a job
func (c *Client) RunJob(id string, options RunOptions) (Executions, error) {
	u := make(map[string]string)
	u["content_type"] = "application/json"
	var res []byte
	var data Executions
	options.runAtTime = strings.Replace(options.RunAtTime.Format(time.RFC3339), "Z", "-", -1)
	opts, err := json.Marshal(options)
	if err != nil {
		return data, err
	}
	err = c.Post(&res, "job/"+id+"/run", opts, u)
	if err != nil {
		return data, err
	}
	xmlErr := xml.Unmarshal(res, &data)
	return data, xmlErr
}

// FindJobByName runs a job by name
func (c *Client) FindJobByName(name string, project string) (*JobDetails, error) {
	var job *JobDetails
	jobs, err := c.ListJobs(project)
	if err != nil {
		return job, err
	}
	if len(jobs.Jobs) > 0 {
		for _, d := range jobs.Jobs {
			if d.Name == name {
				joblist, err := c.GetJob(d.ID)
				if err != nil {
					return job, err
				}
				job = &joblist.Job
			}
		}
	} else {
		err := errors.New("No matches found")
		return job, err
	}
	return job, nil
}

// ListJobs lists the jobs for a project
func (c *Client) ListJobs(projectID string) (Jobs, error) {
	var res []byte
	var data Jobs
	url := fmt.Sprintf("project/%s/jobs", projectID)
	err := c.Get(&res, url, nil)
	if err != nil {
		return data, err
	}
	xmlErr := xml.Unmarshal(res, &data)
	return data, xmlErr
}
