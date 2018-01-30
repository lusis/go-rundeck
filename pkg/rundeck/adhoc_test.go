package rundeck

import (
	"errors"
	"strings"
	"testing"

	requests "github.com/lusis/go-rundeck/pkg/rundeck/requests"
	"github.com/lusis/go-rundeck/pkg/rundeck/responses"
	"github.com/lusis/go-rundeck/pkg/rundeck/responses/testdata"

	"github.com/stretchr/testify/assert"
)

func testFailedAdHocOption() AdHocRunOption {
	return func(r *requests.AdHocCommandRequest) error {
		return errors.New("option setting failed")
	}
}

func testFailedAdHocFromURLOption(e string) AdHocScriptURLOption {
	return func(c *map[string]string) error {
		return errors.New("option setting failed")
	}
}

func testFailedAdHocScriptOption(e string) AdHocScriptOption {
	return func(c *map[string]string) error {
		return errors.New("option setting failed")
	}
}

func TestRunAdHoc(t *testing.T) {
	jsonfile, err := testdata.GetBytes(responses.AdHocExecutionResponseTestFile)
	if err != nil {
		t.Fatalf(err.Error())
	}

	client, server, cErr := newTestRundeckClient(jsonfile, "application/json", 200)
	defer server.Close()
	if cErr != nil {
		t.Fatalf(cErr.Error())
	}
	res, resErr := client.RunAdHocCommand("testproject", "ps -ef")
	assert.NoError(t, resErr)
	assert.NotNil(t, res)
	assert.Equal(t, "Immediate execution scheduled (X)", res.Message)
	assert.Equal(t, 1, res.Execution.ID)
	assert.Equal(t, "[API Href]", res.Execution.HRef)
	assert.Equal(t, "[GUI Href]", res.Execution.Permalink)
}

func TestRunAdHocWithOptions(t *testing.T) {
	jsonfile, err := testdata.GetBytes(responses.AdHocExecutionResponseTestFile)
	if err != nil {
		t.Fatalf(err.Error())
	}

	client, server, cErr := newTestRundeckClient(jsonfile, "application/json", 200)
	defer server.Close()
	if cErr != nil {
		t.Fatalf(cErr.Error())
	}
	options := []AdHocRunOption{
		CmdRunAs("auser"),
		CmdNodeFilters("*"),
		CmdThreadCount(2),
		CmdKeepGoing(true),
	}
	res, resErr := client.RunAdHocCommand("testproject", "ps -ef", options...)
	assert.NoError(t, resErr)
	assert.NotNil(t, res)
	assert.Equal(t, "Immediate execution scheduled (X)", res.Message)
	assert.Equal(t, 1, res.Execution.ID)
	assert.Equal(t, "[API Href]", res.Execution.HRef)
	assert.Equal(t, "[GUI Href]", res.Execution.Permalink)
}

func TestRunAdHocWithFailingOption(t *testing.T) {
	jsonfile, err := testdata.GetBytes(responses.AdHocExecutionResponseTestFile)
	if err != nil {
		t.Fatalf(err.Error())
	}

	client, server, cErr := newTestRundeckClient(jsonfile, "application/json", 200)
	defer server.Close()
	if cErr != nil {
		t.Fatalf(cErr.Error())
	}
	res, resErr := client.RunAdHocCommand("testproject", "ps -ef", testFailedAdHocOption())
	assert.Error(t, resErr)
	assert.Nil(t, res)
	assert.IsType(t, &OptionError{}, resErr)
}

func TestRunAdHocJSONError(t *testing.T) {
	client, server, cErr := newTestRundeckClient([]byte(""), "application/json", 200)
	defer server.Close()
	if cErr != nil {
		t.Fatalf(cErr.Error())
	}
	res, resErr := client.RunAdHocCommand("testproject", "ps -ef")
	assert.Error(t, resErr)
	assert.Nil(t, res)
	assert.IsType(t, &UnmarshalError{}, resErr)
}

func TestRunAdHocHTTPError(t *testing.T) {
	client, server, cErr := newTestRundeckClient([]byte(""), "application/json", 500)
	defer server.Close()
	if cErr != nil {
		t.Fatalf(cErr.Error())
	}
	res, resErr := client.RunAdHocCommand("testproject", "ps -ef")
	assert.Error(t, resErr)
	assert.Nil(t, res)
}

func TestRunAdHocScript(t *testing.T) {
	jsonfile, err := testdata.GetBytes(responses.AdHocExecutionResponseTestFile)
	if err != nil {
		t.Fatalf(err.Error())
	}

	client, server, cErr := newTestRundeckClient(jsonfile, "application/json", 200)
	defer server.Close()
	if cErr != nil {
		t.Fatalf(cErr.Error())
	}
	res, resErr := client.RunAdHocScript("testproject", strings.NewReader("ps -ef"))
	assert.NoError(t, resErr)
	assert.NotNil(t, res)
	assert.Equal(t, "Immediate execution scheduled (X)", res.Message)
	assert.Equal(t, 1, res.Execution.ID)
	assert.Equal(t, "[API Href]", res.Execution.HRef)
	assert.Equal(t, "[GUI Href]", res.Execution.Permalink)
}

func TestRunAdHocScriptOptions(t *testing.T) {
	jsonfile, err := testdata.GetBytes(responses.AdHocExecutionResponseTestFile)
	if err != nil {
		t.Fatalf(err.Error())
	}

	opts := []AdHocScriptOption{
		ScriptArgsQuoted(true),
		ScriptArgString("-i"),
		ScriptFileExtension(".zz"),
		ScriptInterpreter("/usr/bin/python"),
		ScriptKeepGoing(true),
		ScriptThreadCount(2),
		ScriptNodeFilters(".*"),
		ScriptRunAs("auser"),
	}
	client, server, cErr := newTestRundeckClient(jsonfile, "application/json", 200)
	defer server.Close()
	if cErr != nil {
		t.Fatalf(cErr.Error())
	}
	res, resErr := client.RunAdHocScript("testproject", strings.NewReader("ps -ef"), opts...)
	assert.NoError(t, resErr)
	assert.NotNil(t, res)
	assert.Equal(t, "Immediate execution scheduled (X)", res.Message)
	assert.Equal(t, 1, res.Execution.ID)
	assert.Equal(t, "[API Href]", res.Execution.HRef)
	assert.Equal(t, "[GUI Href]", res.Execution.Permalink)
}

func TestRunAdHocScriptFailedOption(t *testing.T) {
	jsonfile, err := testdata.GetBytes(responses.AdHocExecutionResponseTestFile)
	if err != nil {
		t.Fatalf(err.Error())
	}

	client, server, cErr := newTestRundeckClient(jsonfile, "application/json", 200)
	defer server.Close()
	if cErr != nil {
		t.Fatalf(cErr.Error())
	}
	res, resErr := client.RunAdHocScript("testproject", strings.NewReader("ps -ef"), testFailedAdHocScriptOption("foo"))
	assert.Error(t, resErr)
	assert.Nil(t, res)
}

func TestRunAdHocScriptHTTPError(t *testing.T) {
	jsonfile, err := testdata.GetBytes(responses.AdHocExecutionResponseTestFile)
	if err != nil {
		t.Fatalf(err.Error())
	}

	client, server, cErr := newTestRundeckClient(jsonfile, "application/json", 500)
	defer server.Close()
	if cErr != nil {
		t.Fatalf(cErr.Error())
	}
	res, resErr := client.RunAdHocScript("testproject", strings.NewReader("ps -ef"))
	assert.Error(t, resErr)
	assert.Nil(t, res)
}

func TestRunAdHocScriptJSONError(t *testing.T) {
	client, server, cErr := newTestRundeckClient([]byte(""), "application/json", 200)
	defer server.Close()
	if cErr != nil {
		t.Fatalf(cErr.Error())
	}

	res, resErr := client.RunAdHocScript("testproject", strings.NewReader("ps -ef"))
	assert.Error(t, resErr)
	assert.Nil(t, res)
}

func TestRunAdHocScriptURL(t *testing.T) {
	jsonfile, err := testdata.GetBytes(responses.AdHocExecutionResponseTestFile)
	if err != nil {
		t.Fatalf(err.Error())
	}

	client, server, cErr := newTestRundeckClient(jsonfile, "application/json", 200)
	defer server.Close()
	if cErr != nil {
		t.Fatalf(cErr.Error())
	}
	res, resErr := client.RunAdHocScriptFromURL("testproject", "http://localhost/script.sh")
	assert.NoError(t, resErr)
	assert.NotNil(t, res)
	assert.Equal(t, "Immediate execution scheduled (X)", res.Message)
	assert.Equal(t, 1, res.Execution.ID)
	assert.Equal(t, "[API Href]", res.Execution.HRef)
	assert.Equal(t, "[GUI Href]", res.Execution.Permalink)
}

func TestRunAdHocScriptURLOptions(t *testing.T) {
	jsonfile, err := testdata.GetBytes(responses.AdHocExecutionResponseTestFile)
	if err != nil {
		t.Fatalf(err.Error())
	}

	opts := []AdHocScriptURLOption{
		ScriptURLArgsQuoted(true),
		ScriptURLArgString("-i"),
		ScriptURLFileExtension(".zz"),
		ScriptURLInterpreter("/usr/bin/python"),
		ScriptURLKeepGoing(true),
		ScriptURLThreadCount(2),
		ScriptURLNodeFilters(".*"),
		ScriptURLRunAs("auser"),
	}

	client, server, cErr := newTestRundeckClient(jsonfile, "application/json", 200)
	defer server.Close()
	if cErr != nil {
		t.Fatalf(cErr.Error())
	}
	res, resErr := client.RunAdHocScriptFromURL("testproject", "http://localhost/script.sh", opts...)
	assert.NoError(t, resErr)
	assert.NotNil(t, res)
	assert.Equal(t, "Immediate execution scheduled (X)", res.Message)
	assert.Equal(t, 1, res.Execution.ID)
	assert.Equal(t, "[API Href]", res.Execution.HRef)
	assert.Equal(t, "[GUI Href]", res.Execution.Permalink)
}

func TestRunAdHocScriptURLOptionError(t *testing.T) {
	jsonfile, err := testdata.GetBytes(responses.AdHocExecutionResponseTestFile)
	if err != nil {
		t.Fatalf(err.Error())
	}

	client, server, cErr := newTestRundeckClient(jsonfile, "application/json", 200)
	defer server.Close()
	if cErr != nil {
		t.Fatalf(cErr.Error())
	}
	res, resErr := client.RunAdHocScriptFromURL("testproject", "http://localhost/script.sh", testFailedAdHocFromURLOption("foo"))
	assert.Error(t, resErr)
	assert.Nil(t, res)
}

func TestRunAdHocScriptFromURLHTTPError(t *testing.T) {
	jsonfile, err := testdata.GetBytes(responses.AdHocExecutionResponseTestFile)
	if err != nil {
		t.Fatalf(err.Error())
	}

	client, server, cErr := newTestRundeckClient(jsonfile, "application/json", 500)
	defer server.Close()
	if cErr != nil {
		t.Fatalf(cErr.Error())
	}
	res, resErr := client.RunAdHocScriptFromURL("testproject", "http://localhost/script.sh")
	assert.Error(t, resErr)
	assert.Nil(t, res)
}

func TestRunAdHocScriptFromURLJSONError(t *testing.T) {
	client, server, cErr := newTestRundeckClient([]byte(""), "application/json", 200)
	defer server.Close()
	if cErr != nil {
		t.Fatalf(cErr.Error())
	}
	res, resErr := client.RunAdHocScriptFromURL("testproject", "http://localhost/script.sh")
	assert.Error(t, resErr)
	assert.Nil(t, res)
}
