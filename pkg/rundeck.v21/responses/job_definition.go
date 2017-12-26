package responses

// JobYAMLResponse represents a rundeck job-yaml response
type JobYAMLResponse []*JobYAMLDetailResponse

// JobYAMLDetailResponse represents the details of a yaml job definition response
type JobYAMLDetailResponse struct {
	Description        string                   `yaml:"description"`
	ExecutionEnabled   string                   `yaml:"executionEnabled"`
	Group              string                   `yaml:"group"`
	ID                 string                   `yaml:"id"`
	MultipleExecutions bool                     `yaml:"multipleExecutions"`
	Name               string                   `yaml:"name"`
	NodeFilterEditable bool                     `yaml:"nodeFilterEditable"`
	Options            []*JobOptionYAMLResponse `yaml:"options,omitempty"`
	ScheduleEnabled    bool                     `yaml:"scheduleEnabled"`
	UUID               string                   `yaml:"uuid"`
	Sequence           *JobCommandsYAMLResponse `yaml:"sequence"`
}

// JobOptionYAMLResponse represents a jobs options in a yaml job definition response
type JobOptionYAMLResponse struct {
	Description string `yaml:"description,omitempty"`
	Name        string `yaml:"name"`
	Regex       string `yaml:"regex,omitempty"`
	Required    bool   `yaml:"required"`
	Value       string `yaml:"value,omitempty"`
}

// JobCommandsYAMLResponse represents a jobs commands in a yaml job definition response
type JobCommandsYAMLResponse struct {
	Commands map[string]interface{} `yaml:"commands,inline"`
}
