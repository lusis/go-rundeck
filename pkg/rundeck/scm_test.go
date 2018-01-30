package rundeck

import (
	"errors"
	"testing"

	"github.com/lusis/go-rundeck/pkg/rundeck/responses"
	"github.com/lusis/go-rundeck/pkg/rundeck/responses/testdata"
	"github.com/stretchr/testify/assert"
)

func TestListSCMPlugins(t *testing.T) {
	// note that our test server doesn't do any routing/muxing
	// for this test the Import and Export sections will be the same due
	// to the fact that both requests serve the same data
	// however this is fine for this test
	jsonfile, err := testdata.GetBytes(responses.ListSCMPluginsResponseImportTestFile)
	if err != nil {
		t.Fatalf(err.Error())
	}

	client, server, cErr := newTestRundeckClient(jsonfile, "application/json", 200)
	defer server.Close()
	if cErr != nil {
		t.Fatalf(cErr.Error())
	}

	s, serr := client.ListSCMPlugins("testproject")
	assert.NoError(t, serr)
	assert.NotNil(t, s)
	assert.Len(t, s.Import, 1)
	assert.Len(t, s.Export, 1)
}

func TestListSCMPluginsHTTPError(t *testing.T) {
	jsonfile, err := testdata.GetBytes(responses.ListSCMPluginsResponseImportTestFile)
	if err != nil {
		t.Fatalf(err.Error())
	}

	client, server, cErr := newTestRundeckClient(jsonfile, "application/json", 500)
	defer server.Close()
	if cErr != nil {
		t.Fatalf(cErr.Error())
	}

	s, serr := client.ListSCMPlugins("testproject")
	assert.Error(t, serr)
	assert.Nil(t, s)
}

func TestListSCMPluginsJSONError(t *testing.T) {
	client, server, cErr := newTestRundeckClient([]byte(""), "application/json", 200)
	defer server.Close()
	if cErr != nil {
		t.Fatalf(cErr.Error())
	}

	s, serr := client.ListSCMPlugins("testproject")
	assert.Error(t, serr)
	assert.Nil(t, s)
}

func TestGetSCMPluginInputFields(t *testing.T) {
	jsonfile, err := testdata.GetBytes(responses.GetSCMActionInputFieldsResponseTestFileJobExport)
	if err != nil {
		t.Fatalf(err.Error())
	}

	client, server, cErr := newTestRundeckClient(jsonfile, "application/json", 200)
	defer server.Close()
	if cErr != nil {
		t.Fatalf(cErr.Error())
	}

	s, serr := client.GetSCMPluginInputFields("testproject", "export", "git-export")
	assert.NoError(t, serr)
	assert.NotNil(t, s)
}

func TestGetSCMPluginInputFieldsHTTPError(t *testing.T) {
	jsonfile, err := testdata.GetBytes(responses.GetSCMActionInputFieldsResponseTestFileJobExport)
	if err != nil {
		t.Fatalf(err.Error())
	}

	client, server, cErr := newTestRundeckClient(jsonfile, "application/json", 500)
	defer server.Close()
	if cErr != nil {
		t.Fatalf(cErr.Error())
	}

	s, serr := client.GetSCMPluginInputFields("testproject", "export", "git-export")
	assert.Error(t, serr)
	assert.Nil(t, s)
}

func TestGetSCMPluginInputFieldsJSONError(t *testing.T) {

	client, server, cErr := newTestRundeckClient([]byte(""), "application/json", 200)
	defer server.Close()
	if cErr != nil {
		t.Fatalf(cErr.Error())
	}

	s, serr := client.GetSCMPluginInputFields("testproject", "export", "git-export")
	assert.Error(t, serr)
	assert.Nil(t, s)
}

func TestSetupSCMPluginForProject(t *testing.T) {
	jsonfile, err := testdata.GetBytes(responses.SCMPluginForProjectResponseEnableExportTestFile)
	if err != nil {
		t.Fatalf(err.Error())
	}

	client, server, cErr := newTestRundeckClient(jsonfile, "application/json", 200)
	defer server.Close()
	if cErr != nil {
		t.Fatalf(cErr.Error())
	}

	s, serr := client.SetupSCMPluginForProject("testproject", "export", "git-export", make(map[string]string))
	assert.NoError(t, serr)
	assert.NotNil(t, s)
}

func TestSetupSCMPluginForProjectHTTPError(t *testing.T) {
	jsonfile, err := testdata.GetBytes(responses.SCMPluginForProjectResponseEnableExportTestFile)
	if err != nil {
		t.Fatalf(err.Error())
	}

	client, server, cErr := newTestRundeckClient(jsonfile, "application/json", 500)
	defer server.Close()
	if cErr != nil {
		t.Fatalf(cErr.Error())
	}

	s, serr := client.SetupSCMPluginForProject("testproject", "export", "git-export", make(map[string]string))
	assert.Error(t, serr)
	assert.Nil(t, s)
}

func TestSetupSCMPluginForProjectJSONError(t *testing.T) {
	client, server, cErr := newTestRundeckClient([]byte(""), "application/json", 200)
	defer server.Close()
	if cErr != nil {
		t.Fatalf(cErr.Error())
	}

	s, serr := client.SetupSCMPluginForProject("testproject", "export", "git-export", make(map[string]string))
	assert.Error(t, serr)
	assert.Nil(t, s)
}

func TestGetProjectSCMStatus(t *testing.T) {
	jsonfile, err := testdata.GetBytes(responses.GetProjectSCMStatusResponseExportTestFile)
	if err != nil {
		t.Fatalf(err.Error())
	}

	client, server, cErr := newTestRundeckClient(jsonfile, "application/json", 200)
	defer server.Close()
	if cErr != nil {
		t.Fatalf(cErr.Error())
	}

	s, serr := client.GetProjectSCMStatus("testproject", "export")
	assert.NoError(t, serr)
	assert.NotNil(t, s)
}

func TestGetProjectSCMStatusHTTPError(t *testing.T) {
	jsonfile, err := testdata.GetBytes(responses.GetProjectSCMStatusResponseExportTestFile)
	if err != nil {
		t.Fatalf(err.Error())
	}

	client, server, cErr := newTestRundeckClient(jsonfile, "application/json", 500)
	defer server.Close()
	if cErr != nil {
		t.Fatalf(cErr.Error())
	}

	s, serr := client.GetProjectSCMStatus("testproject", "export")
	assert.Error(t, serr)
	assert.Nil(t, s)
}

func TestGetProjectSCMStatusJSONError(t *testing.T) {
	client, server, cErr := newTestRundeckClient([]byte(""), "application/json", 200)
	defer server.Close()
	if cErr != nil {
		t.Fatalf(cErr.Error())
	}

	s, serr := client.GetProjectSCMStatus("testproject", "export")
	assert.Error(t, serr)
	assert.Nil(t, s)
}

func TestGetProjectSCMActionInputFields(t *testing.T) {
	jsonfile, err := testdata.GetBytes(responses.GetSCMActionInputFieldsResponseTestFileJobExport)
	if err != nil {
		t.Fatalf(err.Error())
	}

	client, server, cErr := newTestRundeckClient(jsonfile, "application/json", 200)
	defer server.Close()
	if cErr != nil {
		t.Fatalf(cErr.Error())
	}

	s, serr := client.GetProjectSCMActionInputFields("testproject", "export", "foo")
	assert.NoError(t, serr)
	assert.NotNil(t, s)
}

func TestGetProjectSCMActionInputFieldsHTTPError(t *testing.T) {
	jsonfile, err := testdata.GetBytes(responses.GetSCMActionInputFieldsResponseTestFileJobExport)
	if err != nil {
		t.Fatalf(err.Error())
	}

	client, server, cErr := newTestRundeckClient(jsonfile, "application/json", 500)
	defer server.Close()
	if cErr != nil {
		t.Fatalf(cErr.Error())
	}

	s, serr := client.GetProjectSCMActionInputFields("testproject", "export", "foo")
	assert.Error(t, serr)
	assert.Nil(t, s)
}

func TestGetProjectSCMActionInputFieldsJSONError(t *testing.T) {
	client, server, cErr := newTestRundeckClient([]byte("jsonfile"), "application/json", 200)
	defer server.Close()
	if cErr != nil {
		t.Fatalf(cErr.Error())
	}

	s, serr := client.GetProjectSCMActionInputFields("testproject", "export", "foo")
	assert.Error(t, serr)
	assert.Nil(t, s)
}

func TestPerformProjectSCMAction(t *testing.T) {
	jsonfile, err := testdata.GetBytes(responses.SCMPluginForProjectResponseEnableExportTestFile)
	if err != nil {
		t.Fatalf(err.Error())
	}

	client, server, cErr := newTestRundeckClient(jsonfile, "application/json", 200)
	defer server.Close()
	if cErr != nil {
		t.Fatalf(cErr.Error())
	}

	s, serr := client.PerformProjectSCMAction("testproject", "export", "foo")
	assert.NoError(t, serr)
	assert.NotNil(t, s)
}

func TestPerformProjectSCMActionOptions(t *testing.T) {
	jsonfile, err := testdata.GetBytes(responses.SCMPluginForProjectResponseEnableExportTestFile)
	if err != nil {
		t.Fatalf(err.Error())
	}

	opts := []SCMActionOption{
		SCMActionDeleted("a", "b"),
		SCMActionInput(map[string]string{"foo": "bar"}),
		SCMActionItems("a", "b"),
		SCMActionJobs("a", "b"),
	}
	client, server, cErr := newTestRundeckClient(jsonfile, "application/json", 200)
	defer server.Close()
	if cErr != nil {
		t.Fatalf(cErr.Error())
	}

	s, serr := client.PerformProjectSCMAction("testproject", "export", "foo", opts...)
	assert.NoError(t, serr)
	assert.NotNil(t, s)
}

func TestPerformProjectSCMActionOptionError(t *testing.T) {
	jsonfile, err := testdata.GetBytes(responses.SCMPluginForProjectResponseEnableExportTestFile)
	if err != nil {
		t.Fatalf(err.Error())
	}

	myopt := func() SCMActionOption {
		return func(a *SCMAction) error {
			return errors.New("option error")
		}
	}
	client, server, cErr := newTestRundeckClient(jsonfile, "application/json", 200)
	defer server.Close()
	if cErr != nil {
		t.Fatalf(cErr.Error())
	}

	s, serr := client.PerformProjectSCMAction("testproject", "export", "foo", myopt())
	assert.Error(t, serr)
	assert.Nil(t, s)
}

func TestPerformProjectSCMActionHTTPError(t *testing.T) {
	jsonfile, err := testdata.GetBytes(responses.SCMPluginForProjectResponseEnableExportTestFile)
	if err != nil {
		t.Fatalf(err.Error())
	}

	client, server, cErr := newTestRundeckClient(jsonfile, "application/json", 500)
	defer server.Close()
	if cErr != nil {
		t.Fatalf(cErr.Error())
	}

	s, serr := client.PerformProjectSCMAction("testproject", "export", "foo")
	assert.Error(t, serr)
	assert.Nil(t, s)
}

func TestPerformProjectSCMActionJSONError(t *testing.T) {

	client, server, _ := newTestRundeckClient([]byte(""), "application/json", 200)
	defer server.Close()

	s, serr := client.PerformProjectSCMAction("testproject", "export", "foo")
	assert.Error(t, serr)
	assert.Nil(t, s)
}

func TestGetProjectSCMConfig(t *testing.T) {
	jsonfile, err := testdata.GetBytes(responses.GetProjectSCMConfigResponseImportTestFile)
	if err != nil {
		t.Fatalf(err.Error())
	}
	client, server, _ := newTestRundeckClient(jsonfile, "application/json", 200)
	defer server.Close()

	s, serr := client.GetProjectSCMConfig("testproject", "export")
	assert.NoError(t, serr)
	assert.NotNil(t, s)
}
