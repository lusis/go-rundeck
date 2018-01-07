package responses

// PagingResponse represents paging data in a response
type PagingResponse struct {
	Offset int `json:"offset"`
	Max    int `json:"max"`
	Total  int `json:"total"`
	Count  int `json:"count"`
}
