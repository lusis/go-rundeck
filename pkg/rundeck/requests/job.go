package requests

/*
{
    "argString":"...",
    "loglevel":"...",
    "asUser":"...",
    "filter":"...",
    "runAtTime":"...",
    "options": {
        "myopt1":"value",
        ...
    }
}
*/

// RunJobRequest is the payload for running a job
type RunJobRequest struct {
	ArgString string            `json:"argString,omitempty"`
	LogLevel  string            `json:"loglevel,omitempty"`
	AsUser    string            `json:"asUser,omitempty"`
	Filter    string            `json:"filter,omitempty"`
	RunAtTime *JSONTime         `json:"runAtTime,omitempty"`
	Options   map[string]string `json:"options,omitempty"`
}
