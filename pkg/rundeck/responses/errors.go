package responses

// ErrorResponseTestFile is the test data for an ErrorResponse
const ErrorResponseTestFile = "error.json"

// ErrorResponse is the response for an api error
type ErrorResponse struct {
	IsError    bool   `json:"error"`
	APIVersion int    `json:"apiVersion"`
	ErrorCode  string `json:"errorCode"`
	Message    string `json:"message"`
}

func (a ErrorResponse) minVersion() int  { return AbsoluteMinimumVersion }
func (a ErrorResponse) maxVersion() int  { return CurrentVersion }
func (a ErrorResponse) deprecated() bool { return false }
