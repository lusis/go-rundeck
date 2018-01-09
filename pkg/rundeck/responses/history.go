package responses

// HistoryResponseTestFile is the test data for a HistoryResponse
const HistoryResponseTestFile = "history.json"

// HistoryResponse represents a project history response
// http://rundeck.org/docs/api/index.html#listing-history
type HistoryResponse struct {
	Paging *PagingResponse         `json:"paging"`
	Events []*HistoryEventResponse `json:"events"`
}

func (a HistoryResponse) minVersion() int  { return AbsoluteMinimumVersion }
func (a HistoryResponse) maxVersion() int  { return CurrentVersion }
func (a HistoryResponse) deprecated() bool { return false }

// HistoryEventResponse represents an individual event in a history response
type HistoryEventResponse struct {
	StartTime    int       `json:"starttime"`
	EndTime      int       `json:"endtime"`
	DateStarted  *JSONTime `json:"date-started"`
	DateEnded    *JSONTime `json:"date-ended"`
	Title        string    `json:"title"`
	Status       string    `json:"status"`
	StatusString string    `json:"statusString"`
	Job          *struct {
		ID        string `json:"id"`
		HRef      string `json:"href"`
		Permalink string `json:"permalink"`
	} `json:"job,omitempty"`
	Summary     string `json:"summary"`
	NodeSummary *struct {
		Succeeded int `json:"succeeded"`
		Failed    int `json:"failed"`
		Total     int `json:"total"`
	} `json:"node-summary"`
	User      string `json:"user"`
	Project   string `json:"project"`
	Execution *struct {
		ID        string `json:"id"`
		HRef      string `json:"href"`
		Permalink string `json:"permalink"`
	} `json:"execution"`
}

func (a HistoryEventResponse) minVersion() int  { return AbsoluteMinimumVersion }
func (a HistoryEventResponse) maxVersion() int  { return CurrentVersion }
func (a HistoryEventResponse) deprecated() bool { return false }
