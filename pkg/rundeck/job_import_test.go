package rundeck

import (
	"bytes"
	"errors"
	"testing"

	"github.com/lusis/go-rundeck/pkg/rundeck/responses"
	"github.com/lusis/go-rundeck/pkg/rundeck/responses/testdata"
	"github.com/stretchr/testify/assert"
)

func testFailedImportOption() JobImportOption {
	return func(r *JobImportDefinition) error {
		return errors.New("option setting failed")
	}
}

func TestJobImport(t *testing.T) {
	jsonfile, err := testdata.GetBytes(responses.ImportedJobResponseTestFile)
	if err != nil {
		t.Fatalf(err.Error())
	}

	client, server, _ := newTestRundeckClient(jsonfile, "application/json", 200)
	defer server.Close()
	res, resErr := client.ImportJob("testproject", bytes.NewReader(jsonfile), ImportFormat("yaml"))
	assert.NoError(t, resErr)
	assert.NotNil(t, res)
}

func TestJobImportOptions(t *testing.T) {
	jsonfile, err := testdata.GetBytes(responses.ImportedJobResponseTestFile)
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
	assert.NoError(t, resErr)
	assert.NotNil(t, res)
}

func TestJobImportInvalidFormat(t *testing.T) {
	jsonfile, err := testdata.GetBytes(responses.ImportedJobResponseTestFile)
	if err != nil {
		t.Fatalf(err.Error())
	}

	client, server, _ := newTestRundeckClient(jsonfile, "application/json", 200)
	defer server.Close()
	res, resErr := client.ImportJob("testproject", bytes.NewReader(jsonfile), ImportFormat("text"))
	assert.Error(t, resErr)
	assert.Nil(t, res)
}

func TestJobImportHTTPError(t *testing.T) {
	jsonfile, err := testdata.GetBytes(responses.ImportedJobResponseTestFile)
	if err != nil {
		t.Fatalf(err.Error())
	}

	client, server, _ := newTestRundeckClient(jsonfile, "application/json", 500)

	defer server.Close()
	res, resErr := client.ImportJob("testproject", bytes.NewReader(jsonfile), ImportFormat("yaml"))
	assert.Error(t, resErr)
	assert.Nil(t, res)
}

func TestJobImportJSONError(t *testing.T) {
	jsonfile, err := testdata.GetBytes(responses.ImportedJobResponseTestFile)
	if err != nil {
		t.Fatalf(err.Error())
	}

	client, server, _ := newTestRundeckClient([]byte(""), "application/json", 200)

	defer server.Close()
	res, resErr := client.ImportJob("testproject", bytes.NewReader(jsonfile), ImportFormat("yaml"))
	assert.Error(t, resErr)
	assert.Nil(t, res)
}

func TestJobImportOptionError(t *testing.T) {
	jsonfile, err := testdata.GetBytes(responses.ImportedJobResponseTestFile)
	if err != nil {
		t.Fatalf(err.Error())
	}

	client, server, _ := newTestRundeckClient([]byte(""), "application/json", 200)

	defer server.Close()
	res, resErr := client.ImportJob("testproject", bytes.NewReader(jsonfile), testFailedImportOption())
	assert.Error(t, resErr)
	assert.Nil(t, res)
}
