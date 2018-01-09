package responses

// ListKeysResponseTestFile is the test data for a ListKeysResponse
const ListKeysResponseTestFile = "list_keys.json"

// ListKeysResponse represents a list keys response
type ListKeysResponse struct {
	Resources []ListKeysResourceResponse `json:"resources"`
	URL       string                     `json:"url"`
	Type      string                     `json:"type"`
	Path      string                     `json:"path"`
}

func (a ListKeysResponse) minVersion() int  { return AbsoluteMinimumVersion }
func (a ListKeysResponse) maxVersion() int  { return CurrentVersion }
func (a ListKeysResponse) deprecated() bool { return false }

// ListKeysResourceResponse is an individual resource in a list keys response
type ListKeysResourceResponse struct {
	Meta KeyMetaResponse `json:"meta"`
	URL  string          `json:"url"`
	Name string          `json:"name"`
	Type string          `json:"type"`
	Path string          `json:"path"`
}

func (a ListKeysResourceResponse) minVersion() int  { return AbsoluteMinimumVersion }
func (a ListKeysResourceResponse) maxVersion() int  { return CurrentVersion }
func (a ListKeysResourceResponse) deprecated() bool { return false }

// ListKeysResourceResponseTestFile is the test data for a KeyMetaResponse
const ListKeysResourceResponseTestFile = "key_metadata.json"

// KeyMetaResponse is the metadata about an individual list keys resource
type KeyMetaResponse struct {
	RundeckKeyType     string `json:"Rundeck-key-type"`
	RundeckContentMask string `json:"Rundeck-content-mask"`
	RundeckContentSize string `json:"Rundeck-content-size"`
	RundeckContentType string `json:"Rundeck-content-type"`
}

func (a KeyMetaResponse) minVersion() int  { return AbsoluteMinimumVersion }
func (a KeyMetaResponse) maxVersion() int  { return CurrentVersion }
func (a KeyMetaResponse) deprecated() bool { return false }
