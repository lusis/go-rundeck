package rundeck

/*
Functions in this package are not part of the standard rundeck API but are useful from an enduser perspective

*/
import "errors"

// FindJobByName runs a job by name
func (c *Client) FindJobByName(name string) ([]*JobMetaData, error) {
	projects, pErr := c.ListProjects()
	if pErr != nil {
		return nil, pErr
	}

	var results []*JobMetaData
	for _, project := range *projects {
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
