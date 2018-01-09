package responses

// ListProjectsResponseTestFile is test data for a ListProjectsResponse
const ListProjectsResponseTestFile = "list_projects.json"

// ListProjectsResponse represents a list projects response
type ListProjectsResponse []*ListProjectsEntryResponse

func (a ListProjectsResponse) minVersion() int  { return AbsoluteMinimumVersion }
func (a ListProjectsResponse) maxVersion() int  { return CurrentVersion }
func (a ListProjectsResponse) deprecated() bool { return false }

// ListProjectsEntryResponse represents an item in a list projects response
type ListProjectsEntryResponse struct {
	URL         string `json:"url"`
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
}

func (a ListProjectsEntryResponse) minVersion() int  { return AbsoluteMinimumVersion }
func (a ListProjectsEntryResponse) maxVersion() int  { return CurrentVersion }
func (a ListProjectsEntryResponse) deprecated() bool { return false }

// ProjectInfoResponseTestFile is test data for a ProjectInfoResponse
const ProjectInfoResponseTestFile = "project_info.json"

// ProjectInfoResponse represents a project's details
type ProjectInfoResponse struct {
	URL         string                 `json:"url"`
	Name        string                 `json:"name"`
	Description string                 `json:"description,omitempty"`
	Config      *ProjectConfigResponse `json:"config"`
}

func (a ProjectInfoResponse) minVersion() int  { return AbsoluteMinimumVersion }
func (a ProjectInfoResponse) maxVersion() int  { return CurrentVersion }
func (a ProjectInfoResponse) deprecated() bool { return false }

// ProjectConfigResponse represents a projects configuration response
type ProjectConfigResponse map[string]string

func (a ProjectConfigResponse) minVersion() int  { return AbsoluteMinimumVersion }
func (a ProjectConfigResponse) maxVersion() int  { return CurrentVersion }
func (a ProjectConfigResponse) deprecated() bool { return false }

// ProjectConfigResponseTestFile is test data for a ProjectConfigResponse
const ProjectConfigResponseTestFile = "project_config.json"

// ProjectConfigItemResponseTestFile is test data for a ProjectConfigItemResponse
const ProjectConfigItemResponseTestFile = "config_item.json"

// ProjectConfigItemResponse represents the response from an individual key
type ProjectConfigItemResponse struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func (a ProjectConfigItemResponse) minVersion() int  { return AbsoluteMinimumVersion }
func (a ProjectConfigItemResponse) maxVersion() int  { return CurrentVersion }
func (a ProjectConfigItemResponse) deprecated() bool { return false }

// ProjectArchiveExportAsyncResponseTestFile is test data for a ProjectArchiveExportAsyncResponse
const ProjectArchiveExportAsyncResponseTestFile = "project_archive_export_async.json"

// ProjectArchiveExportAsyncResponse represents the response from an async project archive
type ProjectArchiveExportAsyncResponse struct {
	Token      string `json:"token"`
	Ready      bool   `json:"ready"`
	Percentage int    `json:"percentage"`
}

func (a ProjectArchiveExportAsyncResponse) minVersion() int  { return 19 }
func (a ProjectArchiveExportAsyncResponse) maxVersion() int  { return CurrentVersion }
func (a ProjectArchiveExportAsyncResponse) deprecated() bool { return false }

// ProjectImportArchiveResponseTestFile is test data for a ProjectImportArchiveResponse
const ProjectImportArchiveResponseTestFile = "project_archive_import.json"

// ProjectImportArchiveResponse represents the response from a project archive import
type ProjectImportArchiveResponse struct {
	ImportStatus    string    `json:"import_status"`
	Errors          *[]string `json:"errors,omitempty"`
	ExecutionErrors *[]string `json:"execution_errors,omitempty"`
	ACLErrors       *[]string `json:"acl_errors,omitempty"`
}

func (a ProjectImportArchiveResponse) minVersion() int  { return 19 }
func (a ProjectImportArchiveResponse) maxVersion() int  { return CurrentVersion }
func (a ProjectImportArchiveResponse) deprecated() bool { return false }
