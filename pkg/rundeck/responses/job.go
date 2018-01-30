package responses

// JobsResponse is a collection of JobResponse
type JobsResponse []JobResponse

func (a JobsResponse) minVersion() int  { return 17 }
func (a JobsResponse) maxVersion() int  { return CurrentVersion }
func (a JobsResponse) deprecated() bool { return false }

// JobsResponseTestFile is the test data for JobsResponse
const JobsResponseTestFile = "jobs.json"

// JobResponse represents a job response
type JobResponse struct {
	ID              string `json:"id"`
	Name            string `json:"name"`
	Group           string `json:"group"`
	Project         string `json:"project"`
	Description     string `json:"description"`
	HRef            string `json:"href"`
	Permalink       string `json:"permalink"`
	Scheduled       bool   `json:"scheduled"`
	ScheduleEnabled bool   `json:"scheduleEnabled"`
	Enabled         bool   `json:"enabled"`
	// The following are only visible in cluster mode
	ServerNodeUUID string `json:"serverNodeUUID"`
	ServerOwned    bool   `json:"serverOwned"`
}

func (a JobResponse) minVersion() int  { return 17 }
func (a JobResponse) maxVersion() int  { return CurrentVersion }
func (a JobResponse) deprecated() bool { return false }

// JobMetaDataResponseTestFile is the test data for a JobMetaDataResponse
const JobMetaDataResponseTestFile = "job_metadata.json"

// JobMetaDataResponse represents a job metadata response
type JobMetaDataResponse struct {
	ID              string `json:"id"`
	Name            string `json:"name"`
	Group           string `json:"group"`
	Project         string `json:"project"`
	Description     string `json:"description"`
	HRef            string `json:"href"`
	Permalink       string `json:"permalink"`
	Scheduled       bool   `json:"scheduled"`
	ScheduleEnabled bool   `json:"scheduleEnabled"`
	Enabled         bool   `json:"enabled"`
	AverageDuration int64  `json:"averageDuration"`
}

func (a JobMetaDataResponse) minVersion() int  { return 18 }
func (a JobMetaDataResponse) maxVersion() int  { return CurrentVersion }
func (a JobMetaDataResponse) deprecated() bool { return false }

// ImportedJobEntryResponse is an imported Job response
type ImportedJobEntryResponse struct {
	Index     int    `json:"index"`
	ID        string `json:"id"`
	Name      string `json:"name"`
	Group     string `json:"group"`
	Project   string `json:"project"`
	HRef      string `json:"href"`
	Permalink string `json:"permalink"`
	Messages  string `json:"error,omitempty"`
}

func (a ImportedJobEntryResponse) minVersion() int  { return AbsoluteMinimumVersion }
func (a ImportedJobEntryResponse) maxVersion() int  { return CurrentVersion }
func (a ImportedJobEntryResponse) deprecated() bool { return false }

// ImportedJobResponseTestFile is the test data for an ImportedJobResponse
const ImportedJobResponseTestFile = "imported_job.json"

// ImportedJobResponse is an imported jobs response
type ImportedJobResponse struct {
	Succeeded []ImportedJobEntryResponse `json:"succeeded"`
	Failed    []ImportedJobEntryResponse `json:"failed"`
	Skipped   []ImportedJobEntryResponse `json:"skipped"`
}

func (a ImportedJobResponse) minVersion() int  { return AbsoluteMinimumVersion }
func (a ImportedJobResponse) maxVersion() int  { return CurrentVersion }
func (a ImportedJobResponse) deprecated() bool { return false }

// BulkJobEntryResponse represents a bulk job entry response
type BulkJobEntryResponse struct {
	ID        string `json:"id"`
	Message   string `json:"message"`
	ErrorCode string `json:"errorCode,omitempty"`
}

func (a BulkJobEntryResponse) minVersion() int  { return AbsoluteMinimumVersion }
func (a BulkJobEntryResponse) maxVersion() int  { return CurrentVersion }
func (a BulkJobEntryResponse) deprecated() bool { return false }

// BulkDeleteJobResponseTestFile is the test data for BulkDeleteJobResponse
const BulkDeleteJobResponseTestFile = "bulk_job_delete.json"

// BulkDeleteJobResponse represents a bulk job delete response
type BulkDeleteJobResponse struct {
	RequestCount  int                    `json:"requestCount"`
	AllSuccessful bool                   `json:"allSuccessful"`
	Succeeded     []BulkJobEntryResponse `json:"succeeded"`
	Failed        []BulkJobEntryResponse `json:"failed"`
}

func (a BulkDeleteJobResponse) minVersion() int  { return AbsoluteMinimumVersion }
func (a BulkDeleteJobResponse) maxVersion() int  { return CurrentVersion }
func (a BulkDeleteJobResponse) deprecated() bool { return false }

// JobOptionFileUploadResponseTestFile is the test data for a JobOptionFileUploadResponse
const JobOptionFileUploadResponseTestFile = "job_option_upload.json"

// JobOptionFileUploadResponse represents a job option file upload response
type JobOptionFileUploadResponse struct {
	Total   int               `json:"total"`
	Options map[string]string `json:"options"`
}

func (a JobOptionFileUploadResponse) minVersion() int  { return 19 }
func (a JobOptionFileUploadResponse) maxVersion() int  { return CurrentVersion }
func (a JobOptionFileUploadResponse) deprecated() bool { return false }

// UploadedJobInputFileResponse represents an entry in an UploadedJobInputFilesResponse
type UploadedJobInputFileResponse struct {
	ID             string    `json:"id"`
	User           string    `json:"user"`
	FileState      string    `json:"fileState"`
	SHA            string    `json:"sha"`
	JobID          string    `json:"jobId"`
	DateCreated    *JSONTime `json:"dateCreated"`
	ServerNodeUUID string    `json:"serverNodeUUID"`
	FileName       string    `json:"fileName"`
	Size           int64     `json:"size"`
	ExpirationDate *JSONTime `json:"expirationDate"`
	ExecID         int       `json:"execId"`
}

func (a UploadedJobInputFileResponse) minVersion() int  { return 19 }
func (a UploadedJobInputFileResponse) maxVersion() int  { return CurrentVersion }
func (a UploadedJobInputFileResponse) deprecated() bool { return false }

// UploadedJobInputFilesResponseTestFile is the test data for a UploadedJobInputFileResponse
const UploadedJobInputFilesResponseTestFile = "uploaded_job_input_files.json"

// UploadedJobInputFilesResponse is a response to an uploaded job input file list request
type UploadedJobInputFilesResponse struct {
	Paging PagingResponse                 `json:"paging"`
	Files  []UploadedJobInputFileResponse `json:"files"`
}

func (a UploadedJobInputFilesResponse) minVersion() int  { return 19 }
func (a UploadedJobInputFilesResponse) maxVersion() int  { return CurrentVersion }
func (a UploadedJobInputFilesResponse) deprecated() bool { return false }
