package rundeck

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"

	responses "github.com/lusis/go-rundeck/pkg/rundeck.v21/responses"
)

// JobImportResult is the result of a job import
type JobImportResult responses.ImportedJobResponse

// JobImportOption is a functional option for importing jobs
type JobImportOption func(j *JobImportDefinition) error

// JobImportDefinition is a type for importing a job
type JobImportDefinition struct {
	File       io.Reader
	Format     string
	Project    string
	DupeOption string
	UUIDOption string
}

// ImportProject sets the project for the job import
func ImportProject(p string) JobImportOption {
	return func(j *JobImportDefinition) error {
		j.Project = p
		return nil
	}
}

// ImportFormat sets the format of the job import
func ImportFormat(f string) JobImportOption {
	return func(j *JobImportDefinition) error {
		j.Format = f
		return nil
	}
}

// ImportData sets the job import file source
func ImportData(f io.Reader) JobImportOption {
	return func(j *JobImportDefinition) error {
		j.File = f
		return nil
	}
}

// ImportDupe sets dupe handling for the job import
func ImportDupe(f string) JobImportOption {
	return func(j *JobImportDefinition) error {
		j.DupeOption = f
		return nil
	}
}

// ImportUUID sets uuid handling for the job import
func ImportUUID(f string) JobImportOption {
	return func(j *JobImportDefinition) error {
		j.UUIDOption = f
		return nil
	}
}

// ImportJob imports a job
func (c *Client) ImportJob(opt ...JobImportOption) (*JobImportResult, error) {
	jobRes := &JobImportResult{}
	importDef := &JobImportDefinition{}
	for _, o := range opt {
		if err := o(importDef); err != nil {
			return nil, err
		}
	}
	if &importDef.Project == nil {
		return nil, errors.New("project is required")
	}
	if importDef.Format != "xml" && importDef.Format != "yaml" {
		return nil, errors.New("unsupported import format")
	}

	opts := make(map[string]string)
	opts["format"] = importDef.Format
	if &importDef.DupeOption != nil {
		opts["dupeOption"] = importDef.DupeOption
	}
	if &importDef.UUIDOption != nil {
		opts["uuidOption"] = importDef.UUIDOption
	}

	data, readErr := ioutil.ReadAll(importDef.File)
	if readErr != nil {
		return nil, readErr
	}
	res, postErr := c.httpPost("project/"+importDef.Project+"/jobs/import",
		withBody(bytes.NewReader(data)),
		contentType("application/"+importDef.Format),
		queryParams(opts),
		accept("application/json"))
	if postErr != nil {
		return nil, postErr
	}
	if err := json.Unmarshal(res, &jobRes); err != nil {
		return nil, err
	}
	return jobRes, nil
}
