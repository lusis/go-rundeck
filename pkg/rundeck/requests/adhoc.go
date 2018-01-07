package requests

// AdHocCommandRequest represents an AdHocCommand request
/*
{
    "project":"[project]",
    "exec":"[exec]",
    "nodeThreadcount": #threadcount#,
    "nodeKeepgoing": true/false,
    "asUser": "[asUser]",
    "filter": "[node filter string]"
}
*/
type AdHocCommandRequest struct {
	Project         string `json:"project"`
	Exec            string `json:"exec"`
	NodeThreadCount int    `json:"nodeThreadcount,omitempty"`
	NodeKeepGoing   bool   `json:"nodeKeepGoing,omitempty"`
	AsUser          string `json:"asUser,omitempty"`
	Filter          string `json:"filter,omitempty"`
}
