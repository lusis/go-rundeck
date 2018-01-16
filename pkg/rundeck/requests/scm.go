package requests

// SetupSCMPluginRequest is the request body for calling setup scm plugin
type SetupSCMPluginRequest struct {
	Config map[string]string `json:"config"`
}

// PerformSCMActionRequest is the request body for performing an scm action
/*
{
    "input":{
        "message":"$commitMessage"
    },
    "jobs":[
        "$jobId"
    ],
    "items":[
        "$itemId"
    ],
    "deleted":null
}
*/
type PerformSCMActionRequest struct {
	Input   map[string]string `json:"input"`
	Jobs    []string          `json:"jobs,omitempty"`
	Items   []string          `json:"items,omitempty"`
	Deleted []string          `json:"deleted,omitempty"`
}
