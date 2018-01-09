package responses

// JobYAMLResponse represents a rundeck job-yaml response
type JobYAMLResponse []*JobYAMLDetailResponse

// JobYAMLResponseTestFile is test data for a job definition
const JobYAMLResponseTestFile = "job_definition.yaml"

func (a JobYAMLResponse) minVersion() int  { return AbsoluteMinimumVersion }
func (a JobYAMLResponse) maxVersion() int  { return CurrentVersion }
func (a JobYAMLResponse) deprecated() bool { return false }

// JobYAMLDetailResponse represents the details of a yaml job definition response
type JobYAMLDetailResponse struct {
	Description        string                   `yaml:"description"`
	LogLevel           string                   `yaml:"loglevel"`
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

func (a JobYAMLDetailResponse) minVersion() int  { return AbsoluteMinimumVersion }
func (a JobYAMLDetailResponse) maxVersion() int  { return CurrentVersion }
func (a JobYAMLDetailResponse) deprecated() bool { return false }

// JobOptionYAMLResponse represents a jobs options in a yaml job definition response
type JobOptionYAMLResponse struct {
	Description string `yaml:"description,omitempty"`
	Name        string `yaml:"name"`
	Regex       string `yaml:"regex,omitempty"`
	Required    bool   `yaml:"required"`
	Value       string `yaml:"value,omitempty"`
}

func (a JobOptionYAMLResponse) minVersion() int  { return AbsoluteMinimumVersion }
func (a JobOptionYAMLResponse) maxVersion() int  { return CurrentVersion }
func (a JobOptionYAMLResponse) deprecated() bool { return false }

// JobCommandsYAMLResponse represents a jobs commands in a yaml job definition response
type JobCommandsYAMLResponse struct {
	Commands map[string]interface{} `yaml:"commands,inline"`
}

func (a JobCommandsYAMLResponse) minVersion() int  { return AbsoluteMinimumVersion }
func (a JobCommandsYAMLResponse) maxVersion() int  { return CurrentVersion }
func (a JobCommandsYAMLResponse) deprecated() bool { return false }
