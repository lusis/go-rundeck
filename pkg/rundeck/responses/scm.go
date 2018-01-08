package responses

// SCMResponse is current a placeholder response for SCM
// TODO: implement
type SCMResponse struct{}

func (s SCMResponse) minVersion() int  { return 15 }
func (s SCMResponse) maxVersion() int  { return CurrentVersion }
func (s SCMResponse) deprecated() bool { return false }
