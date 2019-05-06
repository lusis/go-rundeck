package rundeck

import (
	"bufio"
	"bytes"
	"errors"
	"strings"
	"testing"

	"github.com/lusis/go-rundeck/pkg/rundeck/responses"

	"github.com/stretchr/testify/require"
)

func TestGetProject(t *testing.T) {
	jsonfile, err := responses.GetTestData(responses.ProjectInfoResponseTestFile)
	if err != nil {
		t.Fatalf(err.Error())
	}

	client, server, cErr := newTestRundeckClient(jsonfile, "application/json", 200)
	defer server.Close()
	if cErr != nil {
		t.Fatalf(cErr.Error())
	}
	obj, getErr := client.GetProjectInfo("testproject")
	require.NoError(t, getErr)
	require.NotNil(t, obj)
	require.NotEmpty(t, obj.Properties)
	require.Len(t, obj.Properties, 32)
	require.Equal(t, "[API Href]", obj.URL)
	require.Equal(t, "testproject", obj.Name)
	require.Equal(t, "test project", obj.Description)
}

func TestGetProjectHTTPError(t *testing.T) {
	jsonfile, err := responses.GetTestData(responses.ProjectInfoResponseTestFile)
	if err != nil {
		t.Fatalf(err.Error())
	}

	client, server, cErr := newTestRundeckClient(jsonfile, "application/json", 404)
	defer server.Close()
	if cErr != nil {
		t.Fatalf(cErr.Error())
	}
	obj, err := client.GetProjectInfo("testproject")
	require.Error(t, err)
	require.EqualError(t, ErrMissingResource, err.Error())
	require.Nil(t, obj)
}

func TestGetProjectDecodeError(t *testing.T) {
	client, server, cErr := newTestRundeckClient([]byte(""), "application/json", 200)
	defer server.Close()
	if cErr != nil {
		t.Fatalf(cErr.Error())
	}
	obj, oErr := client.GetProjectInfo("testproject")
	require.Error(t, oErr)
	require.Nil(t, obj)
}

func TestListProjects(t *testing.T) {
	jsonfile, err := responses.GetTestData(responses.ListProjectsResponseTestFile)
	if err != nil {
		t.Fatalf(err.Error())
	}

	client, server, cErr := newTestRundeckClient(jsonfile, "application/json", 200)
	defer server.Close()
	if cErr != nil {
		t.Fatalf(cErr.Error())
	}
	obj, cErr := client.ListProjects()
	require.NoError(t, cErr)
	require.NotNil(t, obj)
	require.Len(t, obj, 1)
	project := obj[0]
	require.Equal(t, "[API Href]", project.URL)
	require.Equal(t, "testproject", project.Name)
	require.Equal(t, "test project", project.Description)
}

func TestListProjectsDecodeError(t *testing.T) {
	client, server, cErr := newTestRundeckClient([]byte(""), "application/json", 200)
	defer server.Close()
	if cErr != nil {
		t.Fatalf(cErr.Error())
	}
	obj, oErr := client.ListProjects()
	require.Error(t, oErr)
	require.Nil(t, obj)
}

func TestListProjectsNotFound(t *testing.T) {
	client, server, cErr := newTestRundeckClient([]byte(""), "application/json", 404)
	defer server.Close()
	if cErr != nil {
		t.Fatalf(cErr.Error())
	}
	obj, oErr := client.ListProjects()
	require.Error(t, oErr)
	require.Nil(t, obj)
	require.EqualError(t, ErrMissingResource, oErr.Error())
}

func TestDeleteProject(t *testing.T) {
	client, server, cErr := newTestRundeckClient([]byte(""), "application/json", 204)
	defer server.Close()
	if cErr != nil {
		t.Fatalf(cErr.Error())
	}
	err := client.DeleteProject("abc123")
	require.NoError(t, err)
}

func TestDeleteProjectNotFound(t *testing.T) {
	client, server, cErr := newTestRundeckClient([]byte(""), "application/json", 404)
	defer server.Close()
	if cErr != nil {
		t.Fatalf(cErr.Error())
	}
	err := client.DeleteProject("abc123")
	require.EqualError(t, ErrMissingResource, err.Error())
}

func TestCreateProject(t *testing.T) {
	jsonfile, err := responses.GetTestData(responses.ProjectInfoResponseTestFile)
	if err != nil {
		t.Fatalf(err.Error())
	}

	client, server, _ := newTestRundeckClient(jsonfile, "application/json", 201)
	defer server.Close()
	obj, getErr := client.CreateProject("testproject", nil)
	require.NoError(t, getErr)
	require.NotNil(t, obj)
	require.NotEmpty(t, obj.Properties)
	require.Len(t, obj.Properties, 32)
	require.Equal(t, "[API Href]", obj.URL)
	require.Equal(t, "testproject", obj.Name)
	require.Equal(t, "test project", obj.Description)
}

func TestCreateProjectHTTPError(t *testing.T) {
	jsonfile, err := responses.GetTestData(responses.ProjectInfoResponseTestFile)
	if err != nil {
		t.Fatalf(err.Error())
	}

	client, server, _ := newTestRundeckClient(jsonfile, "application/json", 400)
	defer server.Close()
	obj, getErr := client.CreateProject("testproject", nil)
	require.Error(t, getErr)
	require.Nil(t, obj)
}

func TestCreateProjectJSONError(t *testing.T) {
	client, server, _ := newTestRundeckClient([]byte(""), "application/json", 201)
	defer server.Close()
	obj, getErr := client.CreateProject("testproject", nil)
	require.Error(t, getErr)
	require.Nil(t, obj)
}

func TestGetProjectConfiguration(t *testing.T) {
	jsonfile, err := responses.GetTestData(responses.ProjectConfigResponseTestFile)
	if err != nil {
		t.Fatalf(err.Error())
	}

	client, server, _ := newTestRundeckClient(jsonfile, "application/json", 200)
	defer server.Close()
	obj, cErr := client.GetProjectConfiguration("testproject")
	require.NoError(t, cErr)
	require.NotNil(t, obj)
	require.Len(t, obj, 33)
}

func TestGetProjectConfigurationJSONError(t *testing.T) {
	client, server, _ := newTestRundeckClient([]byte(""), "application/json", 200)
	defer server.Close()
	obj, cErr := client.GetProjectConfiguration("testproject")
	require.Error(t, cErr)
	require.Nil(t, obj)
}

func TestGetProjectConfigurationHTTPError(t *testing.T) {
	jsonfile, err := responses.GetTestData(responses.ProjectConfigResponseTestFile)
	if err != nil {
		t.Fatalf(err.Error())
	}

	client, server, _ := newTestRundeckClient(jsonfile, "application/json", 500)
	defer server.Close()
	obj, cErr := client.GetProjectConfiguration("testproject")
	require.Error(t, cErr)
	require.Nil(t, obj)
}

func TestGetProjectArchiveExport(t *testing.T) {
	client, server, _ := newTestRundeckClient([]byte("testdata"), "application/zip", 200)
	defer server.Close()
	var buf bytes.Buffer
	writer := bufio.NewWriter(&buf)
	cerr := client.GetProjectArchiveExport("testproject", writer)
	require.NoError(t, cerr)
}

func TestGetProjectArchiveExportOptions(t *testing.T) {
	client, server, _ := newTestRundeckClient([]byte("testdata"), "application/zip", 200)
	defer server.Close()
	var buf bytes.Buffer
	writer := bufio.NewWriter(&buf)
	opts := []ProjectExportOption{
		ProjectExportAcls(true),
		ProjectExportAll(true),
		ProjectExportConfigs(true),
		ProjectExportExecutionIDs("a", "b"),
		ProjectExportExecutions(true),
		ProjectExportJobs(true),
		ProjectExportReadmes(true),
	}
	cerr := client.GetProjectArchiveExport("testproject", writer, opts...)
	require.NoError(t, cerr)
}

func TestGetProjectArchiveHTTPError(t *testing.T) {
	client, server, _ := newTestRundeckClient([]byte("testdata"), "application/zip", 500)
	defer server.Close()
	var buf bytes.Buffer
	writer := bufio.NewWriter(&buf)
	cerr := client.GetProjectArchiveExport("testproject", writer)
	require.Error(t, cerr)
}

func TestGetProjectArchiveExportOptionError(t *testing.T) {
	client, server, _ := newTestRundeckClient([]byte("testdata"), "application/zip", 200)
	defer server.Close()
	var buf bytes.Buffer
	writer := bufio.NewWriter(&buf)
	opts := []ProjectExportOption{
		func() ProjectExportOption {
			return func(p *map[string]string) error { return errors.New("failed option") }
		}(),
	}
	cerr := client.GetProjectArchiveExport("testproject", writer, opts...)
	require.Error(t, cerr)
}

func TestGetProjectArchiveExportAsync(t *testing.T) {
	jsonfile, err := responses.GetTestData(responses.ProjectArchiveExportAsyncResponseTestFile)
	if err != nil {
		t.Fatalf(err.Error())
	}

	client, server, _ := newTestRundeckClient(jsonfile, "application/json", 200)
	defer server.Close()
	obj, cErr := client.GetProjectArchiveExportAsync("testproject")
	require.NoError(t, cErr)
	require.NotEmpty(t, obj)
}

func TestGetProjectArchiveExportAsyncOptions(t *testing.T) {
	jsonfile, err := responses.GetTestData(responses.ProjectArchiveExportAsyncResponseTestFile)
	if err != nil {
		t.Fatalf(err.Error())
	}

	client, server, _ := newTestRundeckClient(jsonfile, "application/zip", 200)
	defer server.Close()

	opts := []ProjectExportOption{
		func() ProjectExportOption {
			return func(p *map[string]string) error { return errors.New("failed option") }
		}(),
	}
	c, cerr := client.GetProjectArchiveExportAsync("testproject", opts...)
	require.Error(t, cerr)
	require.Empty(t, c)
}

func TestGetProjectArchiveExportAsyncOptionError(t *testing.T) {
	jsonfile, err := responses.GetTestData(responses.ProjectArchiveExportAsyncResponseTestFile)
	if err != nil {
		t.Fatalf(err.Error())
	}

	client, server, _ := newTestRundeckClient(jsonfile, "application/zip", 200)
	defer server.Close()

	opts := []ProjectExportOption{
		ProjectExportAcls(true),
		ProjectExportAll(true),
		ProjectExportConfigs(true),
		ProjectExportExecutionIDs("a", "b"),
		ProjectExportExecutions(true),
		ProjectExportJobs(true),
		ProjectExportReadmes(true),
	}
	c, cerr := client.GetProjectArchiveExportAsync("testproject", opts...)
	require.NoError(t, cerr)
	require.NotEmpty(t, c)
}

func TestGetProjectArchiveExportAsyncHTTPError(t *testing.T) {
	jsonfile, err := responses.GetTestData(responses.ProjectArchiveExportAsyncResponseTestFile)
	if err != nil {
		t.Fatalf(err.Error())
	}

	client, server, _ := newTestRundeckClient(jsonfile, "application/json", 500)
	defer server.Close()
	obj, cErr := client.GetProjectArchiveExportAsync("testproject")
	require.Error(t, cErr)
	require.Empty(t, obj)
}

func TestGetProjectArchiveExportAsyncJSONError(t *testing.T) {
	client, server, _ := newTestRundeckClient([]byte(""), "application/json", 200)
	defer server.Close()
	obj, cErr := client.GetProjectArchiveExportAsync("testproject")
	require.Error(t, cErr)
	require.Empty(t, obj)
}

func TestGetProjectArchiveExportAsyncStatus(t *testing.T) {
	jsonfile, err := responses.GetTestData(responses.ProjectArchiveExportAsyncResponseTestFile)
	if err != nil {
		t.Fatalf(err.Error())
	}

	client, server, _ := newTestRundeckClient(jsonfile, "application/json", 200)
	defer server.Close()
	obj, cErr := client.GetProjectArchiveExportAsyncStatus("testproject", "ABCDEFG")
	require.NoError(t, cErr)
	require.NotNil(t, obj)
}

func TestGetProjectArchiveExportAsyncStatusHTTPError(t *testing.T) {
	jsonfile, err := responses.GetTestData(responses.ProjectArchiveExportAsyncResponseTestFile)
	if err != nil {
		t.Fatalf(err.Error())
	}

	client, server, _ := newTestRundeckClient(jsonfile, "application/json", 500)
	defer server.Close()
	obj, cErr := client.GetProjectArchiveExportAsyncStatus("testproject", "abcdefg")
	require.Error(t, cErr)
	require.Nil(t, obj)
}

func TestGetProjectArchiveExportAsyncStatusJSONError(t *testing.T) {
	client, server, _ := newTestRundeckClient([]byte(""), "application/json", 200)
	defer server.Close()
	obj, cErr := client.GetProjectArchiveExportAsyncStatus("testproject", "abcdefg")
	require.Error(t, cErr)
	require.Nil(t, obj)
}

func TestGetProjectArchiveExportAsyncDownload(t *testing.T) {
	client, server, _ := newTestRundeckClient([]byte("testdata"), "application/zip", 200)
	defer server.Close()
	var buf bytes.Buffer
	writer := bufio.NewWriter(&buf)
	cerr := client.GetProjectArchiveExportAsyncDownload("testproject", "abcdefg", writer)
	require.NoError(t, cerr)
}

func TestGetProjectArchiveExportAsyncDownloadHTTPError(t *testing.T) {
	client, server, _ := newTestRundeckClient([]byte("testdata"), "application/zip", 500)
	defer server.Close()
	var buf bytes.Buffer
	writer := bufio.NewWriter(&buf)
	cerr := client.GetProjectArchiveExportAsyncDownload("testproject", "abcdefg", writer)
	require.Error(t, cerr)
}

func TestGetProjectArchiveImport(t *testing.T) {
	jsonfile, err := responses.GetTestData(responses.ProjectImportArchiveResponseTestFile)
	if err != nil {
		t.Fatalf(err.Error())
	}
	client, server, _ := newTestRundeckClient(jsonfile, "application/json", 200)
	defer server.Close()

	res, cerr := client.ProjectArchiveImport("testproject", strings.NewReader("hello"))
	require.NoError(t, cerr)
	require.NotNil(t, res)
}

func TestGetProjectArchiveImportOptions(t *testing.T) {
	jsonfile, err := responses.GetTestData(responses.ProjectImportArchiveResponseTestFile)
	if err != nil {
		t.Fatalf(err.Error())
	}
	client, server, _ := newTestRundeckClient(jsonfile, "application/json", 200)
	defer server.Close()

	opts := []ProjectImportOption{
		ProjectImportJobUUIDs("preserve"),
		ProjectImportExecutions(true),
		ProjectImportAcls(true),
		ProjectImportConfigs(true),
	}
	res, cerr := client.ProjectArchiveImport("testproject", strings.NewReader("hello"), opts...)
	require.NoError(t, cerr)
	require.NotNil(t, res)
}

func TestGetProjectArchiveImportOptionError(t *testing.T) {
	jsonfile, err := responses.GetTestData(responses.ProjectImportArchiveResponseTestFile)
	if err != nil {
		t.Fatalf(err.Error())
	}
	client, server, _ := newTestRundeckClient(jsonfile, "application/json", 200)
	defer server.Close()

	opts := []ProjectImportOption{
		func() ProjectImportOption {
			return func(p *map[string]string) error { return errors.New("failed option") }
		}(),
	}

	res, cerr := client.ProjectArchiveImport("testproject", strings.NewReader("hello"), opts...)
	require.Error(t, cerr)
	require.Nil(t, res)
}

func TestGetProjectArchiveImportJSONError(t *testing.T) {
	client, server, _ := newTestRundeckClient([]byte(""), "application/json", 200)
	defer server.Close()
	res, cerr := client.ProjectArchiveImport("testproject", strings.NewReader("hello"))
	require.Error(t, cerr)
	require.Nil(t, res)
}

func TestGetProjectArchiveImportHTTPError(t *testing.T) {
	client, server, _ := newTestRundeckClient([]byte(""), "application/json", 500)
	defer server.Close()
	res, cerr := client.ProjectArchiveImport("testproject", strings.NewReader("hello"))
	require.Error(t, cerr)
	require.Nil(t, res)
}

func TestPutProjectConfigurationKey(t *testing.T) {
	client, server, _ := newTestRundeckClient([]byte(""), "application/json", 200)
	defer server.Close()
	cerr := client.PutProjectConfigurationKey("testproject", "mykey", "myval")
	require.NoError(t, cerr)
}

func TestPutProjectConfigurationKeyHTTPError(t *testing.T) {
	client, server, _ := newTestRundeckClient([]byte(""), "application/json", 500)
	defer server.Close()
	cerr := client.PutProjectConfigurationKey("testproject", "mykey", "myval")
	require.Error(t, cerr)
}

func TestGetProjectConfigurationKey(t *testing.T) {
	client, server, _ := newTestRundeckClient([]byte("stub"), "text/plain", 200)
	defer server.Close()

	res, cerr := client.GetProjectConfigurationKey("testproject", "resources.source.1.type")
	require.NoError(t, cerr)
	require.Equal(t, "stub", res)
}

func TestGetProjectConfigurationKeyHTTPError(t *testing.T) {
	client, server, _ := newTestRundeckClient([]byte("stub"), "text/plain", 500)
	defer server.Close()

	res, cerr := client.GetProjectConfigurationKey("testproject", "resources.source.1.type")
	require.Error(t, cerr)
	require.Empty(t, res)
}

func TestDeleteProjectConfigurationKey(t *testing.T) {
	client, server, _ := newTestRundeckClient([]byte(""), "text/plain", 204)
	defer server.Close()

	cerr := client.DeleteProjectConfigurationKey("testproject", "resources.source.1.type")
	require.NoError(t, cerr)
}

func TestDeleteProjectConfigurationKeyHTTPError(t *testing.T) {
	client, server, _ := newTestRundeckClient([]byte(""), "text/plain", 500)
	defer server.Close()

	cerr := client.DeleteProjectConfigurationKey("testproject", "resources.source.1.type")
	require.Error(t, cerr)
}

func TestPutProjectConfiguration(t *testing.T) {
	conf := `{"foo":"bar","baz":"qux"}`
	data := map[string]string{
		"foo": "bar",
		"baz": "qux",
	}
	client, server, _ := newTestRundeckClient([]byte(conf), "application/json", 200)
	defer server.Close()
	res, cerr := client.PutProjectConfiguration("testproject", data)
	require.NoError(t, cerr)
	require.Equal(t, "bar", res["foo"])
	require.Equal(t, "qux", res["baz"])
}

func TestPutProjectConfigurationHTTPError(t *testing.T) {
	conf := `{"foo":"bar","baz":"qux"}`
	data := map[string]string{
		"foo": "bar",
		"baz": "qux",
	}
	client, server, _ := newTestRundeckClient([]byte(conf), "application/json", 500)
	defer server.Close()
	_, cerr := client.PutProjectConfiguration("testproject", data)
	require.Error(t, cerr)
}

func TestPutProjectConfigurationJSONError(t *testing.T) {
	data := map[string]string{
		"foo": "bar",
		"baz": "qux",
	}
	client, server, _ := newTestRundeckClient([]byte(""), "application/json", 200)
	defer server.Close()
	_, cerr := client.PutProjectConfiguration("testproject", data)
	require.Error(t, cerr)
}
