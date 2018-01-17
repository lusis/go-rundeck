package rundeck

import (
	"encoding/json"
	"errors"
	"io"

	responses "github.com/lusis/go-rundeck/pkg/rundeck/responses"
)

// JobImportResult is the result of a job import
type JobImportResult struct {
	responses.ImportedJobResponse
}

// JobImportOption is a functional option for importing jobs
type JobImportOption func(j *JobImportDefinition) error

// JobImportDefinition is a type for importing a job
type JobImportDefinition struct {
	Format     string
	DupeOption string
	UUIDOption string
}

// ImportFormat sets the format of the job import
func ImportFormat(f string) JobImportOption {
	return func(j *JobImportDefinition) error {
		j.Format = f
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
// http://rundeck.org/docs/api/index.html#importing-jobs
func (c *Client) ImportJob(project string, data io.Reader, opt ...JobImportOption) (*JobImportResult, error) {
	if err := c.checkRequiredAPIVersion(responses.ImportedJobResponse{}); err != nil {
		return nil, err
	}
	jobRes := &JobImportResult{}
	importDef := &JobImportDefinition{}
	for _, o := range opt {
		if err := o(importDef); err != nil {
			return nil, err
		}
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

	res, postErr := c.httpPost("project/"+project+"/jobs/import",
		withBody(data),
		contentType("application/"+importDef.Format),
		queryParams(opts),
		accept("application/json"),
		requestExpects(200))
	if postErr != nil {
		return nil, postErr
	}
	if err := json.Unmarshal(res, &jobRes); err != nil {
		return nil, err
	}
	return jobRes, nil
}
