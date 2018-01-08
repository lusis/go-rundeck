package responses

// PagingResponse represents paging data in a response
type PagingResponse struct {
	Offset int `json:"offset"`
	Max    int `json:"max"`
	Total  int `json:"total"`
	Count  int `json:"count"`
}

func (a PagingResponse) minVersion() int  { return AbsoluteMinimumVersion }
func (a PagingResponse) maxVersion() int  { return CurrentVersion }
func (a PagingResponse) deprecated() bool { return false }
