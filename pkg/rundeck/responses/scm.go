package responses

// SCMResponse is current a placeholder response for SCM
// http://rundeck.org/docs/api/index.html#scm
type SCMResponse struct{}

func (s SCMResponse) minVersion() int  { return 15 }
func (s SCMResponse) maxVersion() int  { return CurrentVersion }
func (s SCMResponse) deprecated() bool { return false }

// ListSCMPluginsResponse is the response listing Scm plugins
// http://rundeck.org/docs/api/index.html#list-scm-plugins
type ListSCMPluginsResponse struct {
	SCMResponse
	Integration string               `json:"integration"`
	Plugins     []*SCMPluginResponse `json:"plugins"`
}

// SCMPluginResponse is an individual Plugin entry in ListSCMPluginsResponse
type SCMPluginResponse struct {
	SCMResponse
	Configured  bool   `json:"configured"`
	Description string `json:"description"`
	Enabled     bool   `json:"enabled"`
	Title       string `json:"title"`
	Type        string `json:"type"`
}

// GetSCMPluginInputFieldsResponse is the response listing scm plugin input fields
// http://rundeck.org/docs/api/index.html#get-scm-plugin-input-fields
type GetSCMPluginInputFieldsResponse struct {
	SCMResponse
	Fields []struct {
		DefaultValue     string            `json:"defaultValue"`
		Description      string            `json:"description"`
		Name             string            `json:"name"`
		RenderingOptions map[string]string `json:"renderingOptions"`
		Required         bool              `json:"required"`
		Scope            string            `json:"scope"`
		Title            string            `json:"title"`
		Type             string            `json:"type"`
		Values           []string          `json:"values,omitempty"`
	} `json:"fields"`
	Integration string `json:"integration"`
	Type        string `json:"type"`
}

// SCMPluginForProjectResponse is the response for setting up, enabling or disabling an scm plugin for a project
// http://rundeck.org/docs/api/index.html#setup-scm-plugin-for-a-project
// http://rundeck.org/docs/api/index.html#enable-scm-plugin-for-a-project
// http://rundeck.org/docs/api/index.html#disable-scm-plugin-for-a-project
/*
failed
{
  "message": "Some input values were not valid.",
  "nextAction": null,
  "success": false,
  "validationErrors": {
    "dir": "required",
    "url": "required"
  }
}
succeeded
{
  "message": "$string",
  "nextAction": null,
  "success": true,
  "validationErrors": null
}
*/
type SCMPluginForProjectResponse struct {
	SCMResponse
	Message          string            `json:"message"`
	NextAction       string            `json:"nextAction,omitempty"`
	Success          bool              `json:"success"`
	ValidationErrors map[string]string `json:"validationErrors,omitempty"`
}

// GetProjectSCMStatusResponse is the response for getting a project's scm status
// http://rundeck.org/docs/api/index.html#get-project-scm-status
/*
{
  "actions": ['action1','action2',..],
  "integration": "$integration",
  "message": null,
  "project": "$project",
  "synchState": "$state"
}
*/
type GetProjectSCMStatusResponse struct {
	SCMResponse
	Actions     []string `json:"actions"`
	Integration string   `json:"integration"`
	Message     string   `json:"message,omitempty"`
	Project     string   `json:"project"`
	SynchState  string   `json:"synchState"`
}

// GetProjectSCMConfigResponse is the response for getting a project's scm config
// http://rundeck.org/docs/api/index.html#get-project-scm-config
/*
{
  "config": {
    "key": "$string",
    "key2": "$string"
  },
  "enabled": $boolean,
  "integration": "$integration",
  "project": "$project",
  "type": "$type"
}
*/
type GetProjectSCMConfigResponse struct {
	SCMResponse
	Config      map[string]string `json:"config"`
	Enabled     bool              `json:"enabled"`
	Integration string            `json:"integration"`
	Project     string            `json:"project"`
	Type        string            `json:"type"`
}

// GetSCMActionInputFieldsResponse is the response for getting a project's scm action input fields
// http://rundeck.org/docs/api/index.html#get-project-scm-action-input-fields
// http://rundeck.org/docs/api/index.html#get-job-scm-action-input-fields
type GetSCMActionInputFieldsResponse struct {
	SCMResponse
	ActionID    string              `json:"actionId"`
	Description string              `json:"description"`
	Fields      []map[string]string `json:"fields"`
	Integration string              `json:"integration"`
	Title       string              `json:"title"`
	ImportItems *[]struct {
		ItemID string `json:"itemId"`
		Job    struct {
			GroupPath string `json:"groupPath"`
			JobID     string `json:"jobId"`
			JobName   string `json:"jobName"`
		}
		Tracked bool `json:"tracked"`
	} `json:"importItems,omitempty"`
	ExportItems *[]struct {
		Deleted bool   `json:"deleted"`
		ItemID  string `json:"itemId"`
		Job     struct {
			GroupPath string `json:"groupPath"`
			JobID     string `json:"jobId"`
			JobName   string `json:"jobName"`
		}
		OriginalID string `json:"originalId"`
		Renamed    bool   `json:"renamed"`
	} `json:"exportItems,omitempty"`
}

// GetJobSCMStatusResponse is the response for getting a job's scm status
// http://rundeck.org/docs/api/index.html#get-job-scm-status
// Note: import status will not include any actions for the job, refer to the Project status to list import actions.
type GetJobSCMStatusResponse struct {
	SCMResponse
	Actions *[]string `json:"actions,omitempty"`
	Commit  struct {
		Author   string            `json:"author"`
		CommitID string            `json:"commitId"`
		Date     string            `json:"date"`
		Info     map[string]string `json:"info"`
		Message  string            `json:"message"`
	}
	ID          string `json:"id"`
	Integration string `json:"integration"`
	Message     string `json:"message"`
	Project     string `json:"project"`
	SynchState  string `json:"synchState"`
}

// GetJobSCMDiffResponse is the response for a job scm diff
// http://rundeck.org/docs/api/index.html#get-job-scm-diff
type GetJobSCMDiffResponse struct {
	SCMResponse
	Commit struct {
		Author   string            `json:"author"`
		CommitID string            `json:"commitId"`
		Date     string            `json:"date"`
		Info     map[string]string `json:"info"`
		Message  string            `json:"message"`
	} `json:"commit,omitempty"`
	DiffContent    string `json:"diffContent"`
	ID             string `json:"id"`
	IncomingCommit struct {
		Author   string            `json:"author"`
		CommitID string            `json:"commitId"`
		Date     string            `json:"date"`
		Info     map[string]string `json:"info"`
		Message  string            `json:"message"`
	} `json:"incomingCommit,omitempty"`
	Integration string `json:"integration"`
	Project     string `json:"project"`
}

// PerformJobSCMActionResponse is the response for performing a job scm action
// http://rundeck.org/docs/api/index.html#perform-job-scm-action
type PerformJobSCMActionResponse struct {
	SCMResponse
	Input struct {
		Message string `json:"message"`
	} `json:"input"`
}
