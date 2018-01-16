package responses

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
)

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
	Integration string              `json:"integration"`
	Plugins     []SCMPluginResponse `json:"plugins"`
}

// ListSCMPluginsResponseImportTestFile is the test data for list scm plugins response for import scm
const ListSCMPluginsResponseImportTestFile = "list_scm_plugins_import.json"

// ListSCMPluginsResponseExportTestFile is the test data for list scm plugins response for export scm
const ListSCMPluginsResponseExportTestFile = "list_scm_plugins_export.json"

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

// GetSCMPluginInputFieldsResponseImportTestData is test data for GetSCMPluginInputFieldsResponse import plugins
const GetSCMPluginInputFieldsResponseImportTestData = "get_scm_input_plugin_fields_import.json"

// GetSCMPluginInputFieldsResponseExportTestData is test data for GetSCMPluginInputFieldsResponse export plugins
const GetSCMPluginInputFieldsResponseExportTestData = "get_scm_input_plugin_fields_export.json"

// SCMPluginForProjectResponse is the response for setting up, enabling or disabling an scm plugin for a project
// http://rundeck.org/docs/api/index.html#setup-scm-plugin-for-a-project
// http://rundeck.org/docs/api/index.html#enable-scm-plugin-for-a-project
// http://rundeck.org/docs/api/index.html#disable-scm-plugin-for-a-project
type SCMPluginForProjectResponse struct {
	SCMResponse
	Message          string            `json:"message"`
	NextAction       string            `json:"nextAction,omitempty"`
	Success          bool              `json:"success"`
	ValidationErrors map[string]string `json:"validationErrors,omitempty"`
}

// SCMPluginForProjectResponseEnableImportTestFile is test data for enabling an scm import plugin
const SCMPluginForProjectResponseEnableImportTestFile = "enable_scm_plugin_import.json"

// SCMPluginForProjectResponseDisableImportTestFile is test data for disabling an scm import plugin
const SCMPluginForProjectResponseDisableImportTestFile = "disable_scm_plugin_import.json"

// SCMPluginForProjectResponseEnableExportTestFile is test data for enabling an scm export plugin
const SCMPluginForProjectResponseEnableExportTestFile = "enable_scm_plugin_export.json"

// SCMPluginForProjectResponseDisableExportTestFile is test data for disabling an scm export plugin
const SCMPluginForProjectResponseDisableExportTestFile = "disable_scm_plugin_export.json"

// GetProjectSCMStatusResponse is the response for getting a project's scm status
// http://rundeck.org/docs/api/index.html#get-project-scm-status
type GetProjectSCMStatusResponse struct {
	SCMResponse
	Actions     []string `json:"actions"`
	Integration string   `json:"integration"`
	Message     string   `json:"message,omitempty"`
	Project     string   `json:"project"`
	SynchState  string   `json:"synchState"`
}

// GetProjectSCMStatusResponseImportTestFile is test data for a GetProjectSCMStatusResponse import plugin
const GetProjectSCMStatusResponseImportTestFile = "get_project_scm_status_import.json"

// GetProjectSCMStatusResponseExportTestFile is test data for a GetProjectSCMStatusResponse export plugin
const GetProjectSCMStatusResponseExportTestFile = "get_project_scm_status_export.json"

// GetProjectSCMConfigResponse is the response for getting a project's scm config
// http://rundeck.org/docs/api/index.html#get-project-scm-config
type GetProjectSCMConfigResponse struct {
	SCMResponse
	Config      *map[string]string `json:"config"`
	Enabled     bool               `json:"enabled"`
	Integration string             `json:"integration"`
	Project     string             `json:"project"`
	Type        string             `json:"type"`
}

// GetProjectSCMConfigResponseImportTestFile is testdata for GetProjectSCMConfigResponse for import plugins
const GetProjectSCMConfigResponseImportTestFile = "get_project_scm_config_import.json"

// GetProjectSCMConfigResponseExportTestFile is testdata for GetProjectSCMConfigResponse for export plugins
const GetProjectSCMConfigResponseExportTestFile = "get_project_scm_config_export.json"

// GetSCMActionInputFieldsResponse is the response for getting a project's scm action input fields
// http://rundeck.org/docs/api/index.html#get-project-scm-action-input-fields
// http://rundeck.org/docs/api/index.html#get-job-scm-action-input-fields
type GetSCMActionInputFieldsResponse struct {
	SCMResponse
	ActionID    string                   `json:"actionId"`
	Description string                   `json:"description"`
	Fields      []map[string]interface{} `json:"fields"`
	Integration string                   `json:"integration"`
	Title       string                   `json:"title"`
	ImportItems []struct {
		ItemID string `json:"itemId"`
		Job    struct {
			GroupPath string `json:"groupPath"`
			JobID     string `json:"jobId"`
			JobName   string `json:"jobName"`
		}
		Tracked bool `json:"tracked"`
	} `json:"importItems,omitempty"`
	ExportItems []struct {
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

// GetSCMActionInputFieldsResponseTestFileProjectImport is test data for project import plugins
const GetSCMActionInputFieldsResponseTestFileProjectImport = "get_project_scm_action_input_fields_import.json"

// GetSCMActionInputFieldsResponseTestFileProjectExport is test data for project export plugins
const GetSCMActionInputFieldsResponseTestFileProjectExport = "get_project_scm_action_input_fields_export.json"

// GetSCMActionInputFieldsResponseTestFileJobImport is test data for job import plugins
const GetSCMActionInputFieldsResponseTestFileJobImport = "get_job_scm_action_input_fields_import.json"

// GetSCMActionInputFieldsResponseTestFileJobExport is test data for job export plugins
const GetSCMActionInputFieldsResponseTestFileJobExport = "get_job_scm_action_input_fields_export.json"

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

// GetJobSCMStatusResponseTestFileImport is test data for job import plugins
const GetJobSCMStatusResponseTestFileImport = "get_job_scm_status_import.json"

// GetJobSCMStatusResponseTestFileExport is test data for job import plugins
const GetJobSCMStatusResponseTestFileExport = "get_job_scm_status_export.json"

// GetJobSCMDiffResponse is the response for a job scm diff
// http://rundeck.org/docs/api/index.html#get-job-scm-diff
// TODO: break Commit into its own type to avoid nested structs issues
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

// GetJobSCMDiffResponseTestFileImport is test data for job scm diff import plugins
const GetJobSCMDiffResponseTestFileImport = "get_job_scm_diff_import.json"

// GetJobSCMDiffResponseTestFileExport is test data for job scm diff import plugins
const GetJobSCMDiffResponseTestFileExport = "get_job_scm_diff_export.json"

// FromReader returns a GetJobSCMDiffResponse from an io.Reader
func (a *GetJobSCMDiffResponse) FromReader(i io.Reader) error {
	b, err := ioutil.ReadAll(i)
	if err != nil {
		return err
	}
	return json.Unmarshal(b, a)
}

// FromBytes returns a GetJobSCMDiffResponse from a byte slice
func (a *GetJobSCMDiffResponse) FromBytes(f []byte) error {
	file := bytes.NewReader(f)
	return a.FromReader(file)
}

// PerformJobSCMActionResponse is the response for performing a job scm action
// http://rundeck.org/docs/api/index.html#perform-job-scm-action
type PerformJobSCMActionResponse struct {
	SCMResponse
	Input struct {
		Message string `json:"message"`
	} `json:"input"`
}
