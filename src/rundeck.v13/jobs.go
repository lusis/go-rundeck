package rundeck

import (
	"encoding/xml"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"strings"
)

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
	Url   string `xml:"url,omitempty"`
}

type JobImportResultJob struct {
	XMLName xml.Name `xml:"job"`
	ID      string   `xml:"id,omitempty"`
	Name    string   `xml:"name"`
	Group   string   `xml:"group"`
	Project string   `xml:"project"`
	Index   int      `xml:"index,attr,omitempty"`
	Href    string   `xml:"href,attr,omitempty"`
	Error   string   `xml:"error,omitempty"`
	Url     string   `xml:"url,omitempty"`
}
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

type Options struct {
	XMLName xml.Name
	Options []Option `xml:"option"`
}

type Option struct {
	XMLName xml.Name `xml:"option"`
	Name    string   `xml:"name,attr"`
	Value   string   `xml:"value,attr,omitempty"`
}

type Jobs struct {
	XMLName xml.Name
	Count   int64 `xml:"count,attr"`
	Jobs    []Job `xml:"job"`
}

type RunOptions struct {
	Filter    string `qp:"filter,omitempty"`
	LogLevel  string `qp:"loglevel,omitempty"`
	RunAs     string `qp:"runAs,omitempty"`
	Arguments string `qp:"argString,omitempty"`
}

func (ro *RunOptions) toQueryParams() (u map[string]string) {
	q := make(map[string]string)
	f := reflect.TypeOf(ro).Elem()
	for i := 0; i < f.NumField(); i++ {
		field := f.Field(i)
		tag := field.Tag
		mytag := tag.Get("qp")
		tokens := strings.Split(mytag, ",")
		if len(tokens) == 1 {
			switch tokens[0] {
			case "-":
				//skip
			default:
				k := tokens[0]
				v := reflect.ValueOf(*ro).Field(i).String()
				q[k] = v
			}
		} else {
			switch tokens[1] {
			case "omitempty":
				if tokens[0] == "" {
					// skip
				} else {
					k := tokens[0]
					v := reflect.ValueOf(*ro).Field(i).String()
					q[k] = v
				}
			default:
				//skip
			}
		}
	}
	return q
}

type JobList struct {
	XMLName xml.Name   `xml:"joblist"`
	Job     JobDetails `xml:"job"`
}

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

type JobSequence struct {
	XMLName   xml.Name
	KeepGoing bool           `xml:"keepgoing,attr"`
	Strategy  string         `xml:"strategy,attr"`
	Steps     []SequenceStep `xml:"command"`
}

type SequenceStep struct {
	XMLName        xml.Name
	Description    string      `xml:"description,omitempty"`
	JobRef         *JobRefStep `xml:"jobref,omitempty"`
	NodeStepPlugin *PluginStep `xml:"node-step-plugin,omitempty"`
	StepPlugin     *PluginStep `xml:"step-plugin,omitempty"`
	Exec           *string     `xml:"exec,omitempty"`
	*ScriptStep    `xml:",omitempty"`
}

type ExecStep struct {
	XMLName xml.Name
	string  `xml:"exec,omitempty"`
}

type ScriptStep struct {
	XMLName           xml.Name
	Script            *string `xml:"script,omitempty"`
	ScriptArgs        *string `xml:"scriptargs,omitempty"`
	ScriptFile        *string `xml:"scriptfile,omitempty"`
	ScriptUrl         *string `xml:"scripturl,omitempty"`
	ScriptInterpreter *string `xml:"scriptinterpreter,omitempty"`
}

type PluginStep struct {
	XMLName       xml.Name
	Type          string `xml:"type,attr"`
	Configuration []struct {
		XMLName xml.Name `xml:"entry"`
		Key     string   `xml:"key,attr"`
		Value   string   `xml:"value,attr"`
	} `xml:"configuration>entry,omitempty"`
}

type JobRefStep struct {
	XMLName  xml.Name
	Name     string `xml:"name,attr,omitempty"`
	Group    string `xml:"group,attr,omitempty"`
	NodeStep bool   `xml:"nodeStep,attr,omitempty"`
}

type JobContext struct {
	XMLName xml.Name     `xml:"context"`
	Project string       `xml:"project"`
	Options *[]JobOption `xml:"options>option,omitempty"`
}

type JobOptions struct {
	XMLName xml.Name
	Options []JobOption `xml:"option"`
}

type JobOption struct {
	XMLName      xml.Name `xml:"option"`
	Name         string   `xml:"name,attr"`
	Required     bool     `xml:"required,attr,omitempty"`
	Secure       bool     `xml:"secure,attr,omitempty"`
	ValueExposed bool     `xml:"valueExposed,attr,omitempty"`
	DefaultValue string   `xml:"value,attr,omitempty"`
	Description  string   `xml:"description,omitempty"`
}

type JobNotifications struct {
	Notifications []JobNotification `xml:"notification,omitempty"`
}

type JobNotification struct {
	XMLName   xml.Name   `xml:"notification"`
	OnStart   JobPlugins `xml:"onstart,omitempty"`
	OnSuccess JobPlugins `xml:"onsuccess,omitempty"`
	OnFailure JobPlugins `xml:"onfailure,omitempty"`
}

type JobPlugins struct {
	Plugins []JobPlugin `xml:"plugin,omitempty"`
}

type JobPlugin struct {
	XMLName       xml.Name               `xml:"plugin"`
	PluginType    string                 `xml:"type,attr"`
	Configuration JobPluginConfiguration `xml:"configuration,omitempty"`
}

type JobPluginConfiguration struct {
	XMLName xml.Name                      `xml:"configuration"`
	Entries []JobPluginConfigurationEntry `xml:"entry,omitempty"`
}

type JobPluginConfigurationEntry struct {
	Key   string `xml:"key,attr"`
	Value string `xml:"value,attr,omitempty"`
}

type JobDispatch struct {
	XMLName           xml.Name `xml:"dispatch"`
	ThreadCount       int64    `xml:"threadcount"`
	KeepGoing         bool     `xml:"keepgoing"`
	ExcludePrecedence bool     `xml:"excludePrecendence"`
	RankOrder         string   `xml:"rankOrder"`
}

type ImportParams struct {
	Filename string
	Format   string
	Dupe     string
	Uuid     string
	Project  string
}

func (c *RundeckClient) GetJob(id string) (JobList, error) {
	u := make(map[string]string)
	var res []byte
	var data JobList
	err := c.Get(&res, "job/"+id, u)
	xml.Unmarshal(res, &data)
	return data, err
}

func (c *RundeckClient) DeleteJob(id string) error {
	return c.Delete("job/"+id, nil)
}
func (c *RundeckClient) ExportJob(id string, format string) (s string, e error) {
	if format != "xml" && format != "yaml" {
		errString := fmt.Sprintf("Unknown/unsupported format \"%s\"", format)
		return s, errors.New(errString)
	}
	var res []byte
	var opts map[string]string
	opts = make(map[string]string)
	opts["format"] = format
	err := c.Get(&res, "job/"+id, opts)
	if err != nil {
		e = err
	} else {
		s = string(res)

	}
	return s, e
}

func (c *RundeckClient) ImportJob(j ImportParams) (string, error) {
	var res []byte
	var jobRes JobImportResult
	var opts map[string]string
	opts = make(map[string]string)
	opts["project"] = j.Project
	opts["format"] = j.Format
	opts["dupeOption"] = j.Dupe
	opts["uuidOption"] = j.Uuid

	jobfile, err := os.Open(j.Filename)
	if err != nil {
		return "", err
	}
	defer jobfile.Close()
	data, _ := ioutil.ReadAll(jobfile)
	err = c.Post(&res, "jobs/import", data, opts)
	xml.Unmarshal(res, &jobRes)
	if err != nil {
		return "", err
	} else {
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
}

func (c *RundeckClient) GetRequiredOpts(j string) (map[string]string, error) {
	u := make(map[string]string)
	var res []byte
	var data JobList
	err := c.Get(&res, "job/"+j, u)
	if err != nil {
		return u, err
	} else {
		xml.Unmarshal(res, &data)
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
}

func (c *RundeckClient) RunJob(id string, options RunOptions) (Executions, error) {
	u := options.toQueryParams()
	var res []byte
	var data Executions

	err := c.Post(&res, "job/"+id+"/run", nil, u)
	if err != nil {
		return data, err
	} else {
		xml.Unmarshal(res, &data)
		return data, nil
	}
}

func (c *RundeckClient) FindJobByName(name string, project string) (*JobDetails, error) {
	var job *JobDetails
	var err error
	jobs, err := c.ListJobs(project)
	if err != nil {
		//
	} else {
		if len(jobs.Jobs) > 0 {
			for _, d := range jobs.Jobs {
				if d.Name == name {
					joblist, err := c.GetJob(d.ID)
					if err != nil {
						//
					} else {
						job = &joblist.Job
					}
				}
			}
		} else {
			err = errors.New("No matches found")
		}
	}
	return job, err
}

func (c *RundeckClient) ListJobs(projectId string) (Jobs, error) {
	options := make(map[string]string)
	options["project"] = projectId
	var res []byte
	var data Jobs
	err := c.Get(&res, "jobs", options)
	if err != nil {
		return data, err
	} else {
		xml.Unmarshal(res, &data)
		return data, nil
	}
}
