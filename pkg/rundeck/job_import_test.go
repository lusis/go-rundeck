package rundeck

import (
	"bytes"
	"errors"
	"testing"

	"github.com/lusis/go-rundeck/pkg/rundeck/responses"
	"github.com/stretchr/testify/require"
)

func testFailedImportOption() JobImportOption {
	return func(r *JobImportDefinition) error {
		return errors.New("option setting failed")
	}
}

func TestJobImport(t *testing.T) {
	jsonfile, err := responses.GetTestData(responses.ImportedJobResponseTestFile)
	if err != nil {
		t.Fatalf(err.Error())
	}

	client, server, _ := newTestRundeckClient(jsonfile, "application/json", 200)
	defer server.Close()
	res, resErr := client.ImportJob("testproject", bytes.NewReader(jsonfile), ImportFormat("yaml"))
	require.NoError(t, resErr)
	require.NotNil(t, res)
}

func TestJobImportOptions(t *testing.T) {
	jsonfile, err := responses.GetTestData(responses.ImportedJobResponseTestFile)
	if err != nil {
		t.Fatalf(err.Error())
	}

	client, server, _ := newTestRundeckClient(jsonfile, "application/json", 200)
	defer server.Close()
	res, resErr := client.ImportJob("testproject",
		bytes.NewReader(jsonfile),
		ImportFormat("yaml"),
		ImportDupe("foo"),
		ImportUUID("foo"))
	require.NoError(t, resErr)
	require.NotNil(t, res)
}

func TestJobImportInvalidFormat(t *testing.T) {
	jsonfile, err := responses.GetTestData(responses.ImportedJobResponseTestFile)
	if err != nil {
		t.Fatalf(err.Error())
	}

	client, server, _ := newTestRundeckClient(jsonfile, "application/json", 200)
	defer server.Close()
	res, resErr := client.ImportJob("testproject", bytes.NewReader(jsonfile), ImportFormat("text"))
	require.Error(t, resErr)
	require.Nil(t, res)
}

func TestJobImportHTTPError(t *testing.T) {
	jsonfile, err := responses.GetTestData(responses.ImportedJobResponseTestFile)
	if err != nil {
		t.Fatalf(err.Error())
	}

	client, server, _ := newTestRundeckClient(jsonfile, "application/json", 500)

	defer server.Close()
	res, resErr := client.ImportJob("testproject", bytes.NewReader(jsonfile), ImportFormat("yaml"))
	require.Error(t, resErr)
	require.Nil(t, res)
}

func TestJobImportJSONError(t *testing.T) {
	jsonfile, err := responses.GetTestData(responses.ImportedJobResponseTestFile)
	if err != nil {
		t.Fatalf(err.Error())
	}

	client, server, _ := newTestRundeckClient([]byte(""), "application/json", 200)

	defer server.Close()
	res, resErr := client.ImportJob("testproject", bytes.NewReader(jsonfile), ImportFormat("yaml"))
	require.Error(t, resErr)
	require.Nil(t, res)
}

func TestJobImportVersionError(t *testing.T) {
	jsonfile, err := responses.GetTestData(responses.ImportedJobResponseTestFile)
	if err != nil {
		t.Fatalf(err.Error())
	}

	client, server, _ := newTestRundeckClient([]byte(""), "application/json", 200)
	client.Config.APIVersion = "1"

	defer server.Close()
	res, resErr := client.ImportJob("testproject", bytes.NewReader(jsonfile), ImportFormat("yaml"))
	require.Error(t, resErr)
	require.Nil(t, res)
}

func TestJobImportOptionError(t *testing.T) {
	jsonfile, err := responses.GetTestData(responses.ImportedJobResponseTestFile)
	if err != nil {
		t.Fatalf(err.Error())
	}

	client, server, _ := newTestRundeckClient([]byte(""), "application/json", 200)

	defer server.Close()
	res, resErr := client.ImportJob("testproject", bytes.NewReader(jsonfile), testFailedImportOption())
	require.Error(t, resErr)
	require.Nil(t, res)
}
