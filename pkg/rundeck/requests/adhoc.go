package requests

// AdHocCommonRequest represents fields that are common in all adhoc requests
type AdHocCommonRequest struct {
	Project         string `json:"project"`
	NodeThreadCount int    `json:"nodeThreadcount,omitempty"`
	NodeKeepGoing   bool   `json:"nodeKeepgoing,omitempty"`
	AsUser          string `json:"asUser,omitempty"`
	Filter          string `json:"filter,omitempty"`
}

// AdHocCommandRequest represents an AdHocCommand request
// http://rundeck.org/docs/api/index.html#running-adhoc-commands
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
	Exec string `json:"exec"`
	AdHocCommonRequest
}

// AdHocScriptRequest represents an AdHocScript request
// http://rundeck.org/docs/api/index.html#running-adhoc-scripts
/*
{
    "project":"[project]",
    "script":"[script]",
    "nodeThreadcount": #threadcount#,
    "nodeKeepgoing": true/false,
    "asUser": "[asUser]",
    "argString": "[argString]",
    "scriptInterpreter": "[scriptInterpreter]",
    "interpreterArgsQuoted": true/false,
    "fileExtension": "[fileExtension]",
    "filter": "[node filter string]"
}
*/
type AdHocScriptRequest struct {
	Script                string `json:"script"`
	ArgString             string `json:"argString,omitempty"`
	ScriptInterpreter     string `json:"scriptInterpreter,omitempty"`
	InterpreterArgsQuoted bool   `json:"interpreterArgsQuoted,omitempty"`
	FileExtension         string `json:"fileExtension,omitempty"`
	AdHocCommonRequest
}

// AdHocScriptURLRequest represents an adhoc script request
// http://rundeck.org/docs/api/index.html#running-adhoc-script-urls
/*
{
    "project":"[project]",
    "url":"[scriptURL]",
    "nodeThreadcount": #threadcount#,
    "nodeKeepgoing": true/false,
    "asUser": "[asUser]",
    "argString": "[argString]",
    "scriptInterpreter": "[scriptInterpreter]",
    "interpreterArgsQuoted": true/false,
    "fileExtension": "[fileExtension]",
    "filter": "[node filter string]"
}
*/
type AdHocScriptURLRequest struct {
	URL                   string `json:"url"`
	ArgString             string `json:"argString,omitempty"`
	ScriptInterpreter     string `json:"scriptInterpreter,omitempty"`
	InterpreterArgsQuoted bool   `json:"interpreterArgsQuoted,omitempty"`
	FileExtension         string `json:"fileExtension,omitempty"`
	AdHocCommonRequest
}
