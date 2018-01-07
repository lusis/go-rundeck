package requests

// BulkToggleRequest represents a bulk toggle request body
type BulkToggleRequest struct {
	IDs    []string `json:"ids,omitempty"`
	IDList string   `json:"idlist,omitempty"`
}
