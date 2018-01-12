package rundeck

/*
Functions in this package are not part of the standard rundeck API but are useful from an enduser perspective

*/
import (
	"errors"

	multierror "github.com/hashicorp/go-multierror"
	responses "github.com/lusis/go-rundeck/pkg/rundeck/responses"
	yaml "gopkg.in/yaml.v2"
)

// FindJobByName runs a job by name
func (c *Client) FindJobByName(name string) ([]*JobMetaData, error) {
	projects, pErr := c.ListProjects()
	if pErr != nil {
		return nil, pErr
	}

	var results []*JobMetaData
	for _, project := range projects {
		jobs, err := c.ListJobs(project.Name)
		if err != nil {
			return nil, err
		}
		for _, d := range *jobs {
			if d.Name == name {
				job, joblistErr := c.GetJobInfo(d.ID)
				if joblistErr != nil {
					return nil, joblistErr
				}
				results = append(results, job)
			}
		}
	}
	if len(results) == 0 {
		return nil, errors.New("No matches found")
	}
	return results, nil
}

// GetJobOpts returns the required options for a job
func (c *Client) GetJobOpts(j string) ([]*JobOption, error) {
	options := make([]*JobOption, 0)
	data := &responses.JobYAMLResponse{}
	res, err := c.httpGet("job/"+j, accept("application/yaml"), requestExpects(200))
	if err != nil {
		return nil, err
	}
	if err := yaml.Unmarshal(res, &data); err != nil {
		return nil, &UnmarshalError{msg: multierror.Append(errDecoding, err).Error()}
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
	res, err := c.httpGet("job/"+j, accept("application/yaml"), requestExpects(200))

	if err != nil {
		return nil, err
	}
	if err := yaml.Unmarshal(res, &data); err != nil {
		return nil, &UnmarshalError{msg: multierror.Append(errDecoding, err).Error()}
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
