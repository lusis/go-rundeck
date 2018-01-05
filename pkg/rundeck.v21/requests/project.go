package requests

// ProjectCreationRequest is the payload for creating a new project
/*
{
	"name":"fooproject3",
	"config":{
		"somekey":"somevar",
		"anotherkey":"anothervar"
	}
}
*/
type ProjectCreationRequest struct {
	Name   string             `json:"name"`
	Config *map[string]string `json:"config,omitempty"`
}
