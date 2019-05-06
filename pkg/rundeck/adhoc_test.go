package rundeck

import (
	"errors"
	"strings"
	"testing"

	requests "github.com/lusis/go-rundeck/pkg/rundeck/requests"
	"github.com/lusis/go-rundeck/pkg/rundeck/responses"

	"github.com/stretchr/testify/require"
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
	jsonfile, err := responses.GetTestData(responses.AdHocExecutionResponseTestFile)
	require.NoError(t, err)

	client, server, err := newTestRundeckClient(jsonfile, "application/json", 200)
	defer server.Close()
	require.NoError(t, err)

	res, err := client.RunAdHocCommand("testproject", "ps -ef")
	require.NoError(t, err)
	require.NotNil(t, res)
	require.Equal(t, "Immediate execution scheduled (X)", res.Message)
	require.Equal(t, 1, res.Execution.ID)
	require.Equal(t, "[API Href]", res.Execution.HRef)
	require.Equal(t, "[GUI Href]", res.Execution.Permalink)
}

func TestRunAdHocWithOptions(t *testing.T) {
	jsonfile, err := responses.GetTestData(responses.AdHocExecutionResponseTestFile)
	require.NoError(t, err)

	client, server, err := newTestRundeckClient(jsonfile, "application/json", 200)
	defer server.Close()
	require.NoError(t, err)

	options := []AdHocRunOption{
		CmdRunAs("auser"),
		CmdNodeFilters("*"),
		CmdThreadCount(2),
		CmdKeepGoing(true),
	}
	res, err := client.RunAdHocCommand("testproject", "ps -ef", options...)
	require.NoError(t, err)
	require.NotNil(t, res)
	require.Equal(t, "Immediate execution scheduled (X)", res.Message)
	require.Equal(t, 1, res.Execution.ID)
	require.Equal(t, "[API Href]", res.Execution.HRef)
	require.Equal(t, "[GUI Href]", res.Execution.Permalink)
}

func TestRunAdHocWithFailingOption(t *testing.T) {
	jsonfile, err := responses.GetTestData(responses.AdHocExecutionResponseTestFile)
	require.NoError(t, err)

	client, server, err := newTestRundeckClient(jsonfile, "application/json", 200)
	defer server.Close()
	require.NoError(t, err)

	res, err := client.RunAdHocCommand("testproject", "ps -ef", testFailedAdHocOption())
	require.Error(t, err)
	require.Nil(t, res)
	require.IsType(t, &OptionError{}, err)
}

func TestRunAdHocJSONError(t *testing.T) {
	client, server, err := newTestRundeckClient([]byte(""), "application/json", 200)
	defer server.Close()
	require.NoError(t, err)

	res, err := client.RunAdHocCommand("testproject", "ps -ef")
	require.Error(t, err)
	require.Nil(t, res)
	require.IsType(t, &UnmarshalError{}, err)
}

func TestRunAdHocHTTPError(t *testing.T) {
	client, server, err := newTestRundeckClient([]byte(""), "application/json", 500)
	defer server.Close()
	require.NoError(t, err)

	res, err := client.RunAdHocCommand("testproject", "ps -ef")
	require.Error(t, err)
	require.Nil(t, res)
}

func TestRunAdHocScript(t *testing.T) {
	jsonfile, err := responses.GetTestData(responses.AdHocExecutionResponseTestFile)
	require.NoError(t, err)

	client, server, err := newTestRundeckClient(jsonfile, "application/json", 200)
	defer server.Close()
	require.NoError(t, err)

	res, err := client.RunAdHocScript("testproject", strings.NewReader("ps -ef"))
	require.NoError(t, err)
	require.NotNil(t, res)
	require.Equal(t, "Immediate execution scheduled (X)", res.Message)
	require.Equal(t, 1, res.Execution.ID)
	require.Equal(t, "[API Href]", res.Execution.HRef)
	require.Equal(t, "[GUI Href]", res.Execution.Permalink)
}

func TestRunAdHocScriptOptions(t *testing.T) {
	jsonfile, err := responses.GetTestData(responses.AdHocExecutionResponseTestFile)
	require.NoError(t, err)

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
	client, server, err := newTestRundeckClient(jsonfile, "application/json", 200)
	defer server.Close()
	require.NoError(t, err)

	res, err := client.RunAdHocScript("testproject", strings.NewReader("ps -ef"), opts...)
	require.NoError(t, err)
	require.NotNil(t, res)
	require.Equal(t, "Immediate execution scheduled (X)", res.Message)
	require.Equal(t, 1, res.Execution.ID)
	require.Equal(t, "[API Href]", res.Execution.HRef)
	require.Equal(t, "[GUI Href]", res.Execution.Permalink)
}

func TestRunAdHocScriptFailedOption(t *testing.T) {
	jsonfile, err := responses.GetTestData(responses.AdHocExecutionResponseTestFile)
	require.NoError(t, err)

	client, server, err := newTestRundeckClient(jsonfile, "application/json", 200)
	defer server.Close()
	require.NoError(t, err)

	res, err := client.RunAdHocScript("testproject", strings.NewReader("ps -ef"), testFailedAdHocScriptOption("foo"))
	require.Error(t, err)
	require.Nil(t, res)
}

func TestRunAdHocScriptHTTPError(t *testing.T) {
	jsonfile, err := responses.GetTestData(responses.AdHocExecutionResponseTestFile)
	require.NoError(t, err)

	client, server, err := newTestRundeckClient(jsonfile, "application/json", 500)
	defer server.Close()
	require.NoError(t, err)

	res, err := client.RunAdHocScript("testproject", strings.NewReader("ps -ef"))
	require.Error(t, err)
	require.Nil(t, res)
}

func TestRunAdHocScriptJSONError(t *testing.T) {
	client, server, err := newTestRundeckClient([]byte(""), "application/json", 200)
	defer server.Close()
	require.NoError(t, err)

	res, err := client.RunAdHocScript("testproject", strings.NewReader("ps -ef"))
	require.Error(t, err)
	require.Nil(t, res)
}

func TestRunAdHocScriptURL(t *testing.T) {
	jsonfile, err := responses.GetTestData(responses.AdHocExecutionResponseTestFile)
	require.NoError(t, err)

	client, server, err := newTestRundeckClient(jsonfile, "application/json", 200)
	defer server.Close()
	require.NoError(t, err)

	res, err := client.RunAdHocScriptFromURL("testproject", "http://localhost/script.sh")
	require.NoError(t, err)
	require.NotNil(t, res)
	require.Equal(t, "Immediate execution scheduled (X)", res.Message)
	require.Equal(t, 1, res.Execution.ID)
	require.Equal(t, "[API Href]", res.Execution.HRef)
	require.Equal(t, "[GUI Href]", res.Execution.Permalink)
}

func TestRunAdHocScriptURLOptions(t *testing.T) {
	jsonfile, err := responses.GetTestData(responses.AdHocExecutionResponseTestFile)
	require.NoError(t, err)

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

	client, server, err := newTestRundeckClient(jsonfile, "application/json", 200)
	defer server.Close()
	require.NoError(t, err)
	
	res, err := client.RunAdHocScriptFromURL("testproject", "http://localhost/script.sh", opts...)
	require.NoError(t, err)
	require.NotNil(t, res)
	require.Equal(t, "Immediate execution scheduled (X)", res.Message)
	require.Equal(t, 1, res.Execution.ID)
	require.Equal(t, "[API Href]", res.Execution.HRef)
	require.Equal(t, "[GUI Href]", res.Execution.Permalink)
}

func TestRunAdHocScriptURLOptionError(t *testing.T) {
	jsonfile, err := responses.GetTestData(responses.AdHocExecutionResponseTestFile)
	require.NoError(t, err)

	client, server, err := newTestRundeckClient(jsonfile, "application/json", 200)
	defer server.Close()
	require.NoError(t, err)

	res, err := client.RunAdHocScriptFromURL("testproject", "http://localhost/script.sh", testFailedAdHocFromURLOption("foo"))
	require.Error(t, err)
	require.Nil(t, res)
}

func TestRunAdHocScriptFromURLHTTPError(t *testing.T) {
	jsonfile, err := responses.GetTestData(responses.AdHocExecutionResponseTestFile)
	require.NoError(t, err)

	client, server, err := newTestRundeckClient(jsonfile, "application/json", 500)
	defer server.Close()
	require.NoError(t, err)

	res, err := client.RunAdHocScriptFromURL("testproject", "http://localhost/script.sh")
	require.Error(t, err)
	require.Nil(t, res)
}

func TestRunAdHocScriptFromURLJSONError(t *testing.T) {
	client, server, err := newTestRundeckClient([]byte(""), "application/json", 200)
	defer server.Close()
	require.NoError(t, err)

	res, err := client.RunAdHocScriptFromURL("testproject", "http://localhost/script.sh")
	require.Error(t, err)
	require.Nil(t, res)
}
