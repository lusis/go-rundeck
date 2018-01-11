package requests

// TokenRequest represents a user and token
// http://rundeck.org/docs/api/index.html#get-a-token
type TokenRequest struct {
	User     string `json:"user"`
	Roles    string `json:"roles"`
	Duration string `json:"duration,omitempty"`
}
