package rundeck

import (
	"errors"
	"testing"
	"time"

	requests "github.com/lusis/go-rundeck/pkg/rundeck/requests"
	"github.com/lusis/go-rundeck/pkg/rundeck/responses"
	yaml "gopkg.in/yaml.v2"

	"github.com/stretchr/testify/assert"
)

func testFailedJobOption() RunJobOption {
	return func(r *requests.RunJobRequest) error {
		return errors.New("option setting failed")
	}
}

func TestGetJobMetaData(t *testing.T) {
	jsonfile, err := responses.GetTestData(responses.JobMetaDataResponseTestFile)
	if err != nil {
		t.Fatalf(err.Error())
	}

	client, server, cErr := newTestRundeckClient(jsonfile, "application/json", 200)
	defer server.Close()
	if cErr != nil {
		t.Fatalf(cErr.Error())
	}
	obj, oErr := client.GetJobMetaData("1")
	assert.NoError(t, oErr)
	assert.NotNil(t, obj)
}

func TestGetJobMetaDataHTTPError(t *testing.T) {
	jsonfile, err := responses.GetTestData(responses.JobMetaDataResponseTestFile)
	if err != nil {
		t.Fatalf(err.Error())
	}

	client, server, cErr := newTestRundeckClient(jsonfile, "application/json", 500)
	defer server.Close()
	if cErr != nil {
		t.Fatalf(cErr.Error())
	}
	obj, oErr := client.GetJobMetaData("1")
	assert.Error(t, oErr)
	assert.Nil(t, obj)
}

func TestGetJobMetaDataJSONError(t *testing.T) {
	client, server, cErr := newTestRundeckClient([]byte(""), "application/json", 200)
	defer server.Close()
	if cErr != nil {
		t.Fatalf(cErr.Error())
	}
	obj, oErr := client.GetJobMetaData("1")
	assert.Error(t, oErr)
	assert.Nil(t, obj)
}

func TestGetJobDefinition(t *testing.T) {
	jsonfile, err := responses.GetTestData("job_definition.yaml")
	if err != nil {
		t.Fatalf(err.Error())
	}

	client, server, cErr := newTestRundeckClient(jsonfile, "application/yaml", 200)
	defer server.Close()
	if cErr != nil {
		t.Fatalf(cErr.Error())
	}
	obj, oErr := client.GetJobDefinition("1", "yaml")
	assert.NoError(t, oErr)
	assert.NotEmpty(t, obj)
	data := &responses.JobYAMLResponse{}
	yErr := yaml.Unmarshal(obj, &data)
	assert.NoError(t, yErr)
	assert.NotNil(t, data)
}

func TestGetJobDefinitionHTTPError(t *testing.T) {
	jsonfile, err := responses.GetTestData("job_definition.yaml")
	if err != nil {
		t.Fatalf(err.Error())
	}

	client, server, cErr := newTestRundeckClient(jsonfile, "application/yaml", 500)
	defer server.Close()
	if cErr != nil {
		t.Fatalf(cErr.Error())
	}
	obj, oErr := client.GetJobDefinition("1", "yaml")
	assert.Error(t, oErr)
	assert.Nil(t, obj)
}

func TestGetJobInfo(t *testing.T) {
	jsonfile, err := responses.GetTestData(responses.JobMetaDataResponseTestFile)
	if err != nil {
		t.Fatalf(err.Error())
	}

	client, server, cErr := newTestRundeckClient(jsonfile, "application/json", 200)
	defer server.Close()
	if cErr != nil {
		t.Fatalf(cErr.Error())
	}
	obj, cErr := client.GetJobInfo("1")
	assert.NoError(t, cErr)
	assert.NotNil(t, obj)
}

func TestListJobs(t *testing.T) {
	jsonfile, err := responses.GetTestData(responses.JobsResponseTestFile)
	if err != nil {
		t.Fatalf(err.Error())
	}

	client, server, cErr := newTestRundeckClient(jsonfile, "application/json", 200)
	defer server.Close()
	if cErr != nil {
		t.Fatalf(cErr.Error())
	}
	obj, cErr := client.ListJobs("testproject")
	assert.NoError(t, cErr)
	assert.NotNil(t, obj)
}

func TestListJobsJSONError(t *testing.T) {
	client, server, cErr := newTestRundeckClient([]byte(""), "application/json", 200)
	defer server.Close()
	if cErr != nil {
		t.Fatalf(cErr.Error())
	}
	obj, cErr := client.ListJobs("testproject")
	assert.Error(t, cErr)
	assert.Empty(t, obj)
}

func TestListJobsHTTPError(t *testing.T) {
	client, server, cErr := newTestRundeckClient([]byte(""), "application/json", 500)
	defer server.Close()
	if cErr != nil {
		t.Fatalf(cErr.Error())
	}
	obj, cErr := client.ListJobs("testproject")
	assert.Error(t, cErr)
	assert.Empty(t, obj)
}

func TestRunJobOption(t *testing.T) {
	curTime := time.Now().UTC()
	jobOpts := &requests.RunJobRequest{}
	opts := []RunJobOption{
		RunJobArgs("-foo bar"),
		RunJobAs("auser"),
		RunJobOpts(map[string]string{"foo": "bar"}),
		RunJobFilter(".*"),
		RunJobLogLevel("DEBUG"),
		RunJobRunAt(curTime),
	}
	for _, opt := range opts {
		if err := opt(jobOpts); err != nil {
			assert.NoError(t, err)
		}
	}
	assert.Equal(t, "-foo bar", jobOpts.ArgString)
	assert.Equal(t, "auser", jobOpts.AsUser)
	assert.Equal(t, "bar", jobOpts.Options["foo"])
	assert.Equal(t, "DEBUG", jobOpts.LogLevel)
	assert.NotNil(t, jobOpts.RunAtTime)
	assert.Equal(t, ".*", jobOpts.Filter)
}

func TestRunJobOptionError(t *testing.T) {
	client, server, _ := newTestRundeckClient([]byte(""), "application/json", 200)
	defer server.Close()
	res, err := client.RunJob("abcdefg", testFailedJobOption())
	assert.Nil(t, res)
	assert.Error(t, err)
}

func TestRunJobHTTPError(t *testing.T) {
	client, server, _ := newTestRundeckClient([]byte(""), "application/json", 500)
	defer server.Close()
	res, err := client.RunJob("abcdefg")
	assert.Nil(t, res)
	assert.Error(t, err)
}

func TestRunJobJSONError(t *testing.T) {
	client, server, _ := newTestRundeckClient([]byte(""), "application/json", 200)
	defer server.Close()
	res, err := client.RunJob("abcdefg")
	assert.Nil(t, res)
	assert.Error(t, err)
}

func TestDeleteJobFound(t *testing.T) {
	client, server, cErr := newTestRundeckClient([]byte(""), "application/json", 204)
	defer server.Close()
	if cErr != nil {
		t.Fatalf(cErr.Error())
	}
	err := client.DeleteJob("testproject")
	assert.NoError(t, err)
}

func TestDeleteJobNotFound(t *testing.T) {
	client, server, cErr := newTestRundeckClient([]byte(""), "application/json", 404)
	defer server.Close()
	if cErr != nil {
		t.Fatalf(cErr.Error())
	}
	err := client.DeleteJob("testproject")
	assert.EqualError(t, ErrMissingResource, err.Error())
}

func TestGetRequiredOptsHTTPError(t *testing.T) {
	client, server, _ := newTestRundeckClient([]byte(""), "application/json", 500)
	defer server.Close()
	res, err := client.GetRequiredOpts("abcdefg")
	assert.Nil(t, res)
	assert.Error(t, err)
}

func TestGetRequiredOptsYAMLError(t *testing.T) {
	client, server, _ := newTestRundeckClient([]byte("1234"), "application/json", 200)
	defer server.Close()
	res, err := client.GetRequiredOpts("abcdefg")
	assert.Nil(t, res)
	assert.Error(t, err)
}

func TestGetRequiredOpts(t *testing.T) {
	jsonfile, err := responses.GetTestData("job_definition.yaml")
	if err != nil {
		t.Fatalf(err.Error())
	}

	client, server, cErr := newTestRundeckClient(jsonfile, "application/json", 200)
	defer server.Close()
	if cErr != nil {
		t.Fatalf(cErr.Error())
	}
	obj, cErr := client.GetRequiredOpts("abcdefg")
	assert.NoError(t, cErr)
	assert.NotNil(t, obj)
}

func TestGetJobOptsHTTPError(t *testing.T) {
	client, server, _ := newTestRundeckClient([]byte(""), "application/json", 500)
	defer server.Close()
	res, err := client.GetJobOpts("abcdefg")
	assert.Nil(t, res)
	assert.Error(t, err)
}

func TestGetJobOptsYAMLError(t *testing.T) {
	client, server, _ := newTestRundeckClient([]byte("1234"), "application/json", 200)
	defer server.Close()
	res, err := client.GetJobOpts("abcdefg")
	assert.Nil(t, res)
	assert.Error(t, err)
}

func TestGetJobOpts(t *testing.T) {
	jsonfile, err := responses.GetTestData("job_definition.yaml")
	if err != nil {
		t.Fatalf(err.Error())
	}

	client, server, cErr := newTestRundeckClient(jsonfile, "application/json", 200)
	defer server.Close()
	if cErr != nil {
		t.Fatalf(cErr.Error())
	}
	obj, cErr := client.GetJobOpts("abcdefg")
	assert.NoError(t, cErr)
	assert.NotNil(t, obj)
}

func TestExportJobInvalidFormat(t *testing.T) {
	client, server, cErr := newTestRundeckClient([]byte(""), "application/yaml", 200)
	defer server.Close()
	if cErr != nil {
		t.Fatalf(cErr.Error())
	}
	res, err := client.ExportJob("abcdefg", "json")
	assert.Nil(t, res)
	assert.Error(t, err)
}

func TestExportJob(t *testing.T) {
	jsonfile, err := responses.GetTestData("job_definition.yaml")
	if err != nil {
		t.Fatalf(err.Error())
	}

	client, server, cErr := newTestRundeckClient(jsonfile, "application/yaml", 200)
	defer server.Close()
	if cErr != nil {
		t.Fatalf(cErr.Error())
	}
	obj, cErr := client.ExportJob("abcdefg", "yaml")
	assert.NoError(t, cErr)
	assert.NotEmpty(t, obj)
}

func TestExportJobHTTPError(t *testing.T) {
	jsonfile, err := responses.GetTestData("job_definition.yaml")
	if err != nil {
		t.Fatalf(err.Error())
	}

	client, server, cErr := newTestRundeckClient(jsonfile, "application/yaml", 500)
	defer server.Close()
	if cErr != nil {
		t.Fatalf(cErr.Error())
	}
	obj, oErr := client.ExportJob("abcdefg", "yaml")
	assert.Error(t, oErr)
	assert.Nil(t, obj)
}

func TestDeleteAllExecutionsForJob(t *testing.T) {
	jsonfile, err := responses.GetTestData(responses.BulkDeleteExecutionsResponseTestFile)
	if err != nil {
		t.Fatalf(err.Error())
	}

	client, server, cErr := newTestRundeckClient(jsonfile, "application/json", 200)
	defer server.Close()
	if cErr != nil {
		t.Fatalf(cErr.Error())
	}
	obj, oErr := client.DeleteAllExecutionsForJob("abcdefg")
	assert.NoError(t, oErr)
	assert.NotNil(t, obj)
}

func TestDeleteAllExecutionsForJobHTTPError(t *testing.T) {
	jsonfile, err := responses.GetTestData(responses.BulkDeleteExecutionsResponseTestFile)
	if err != nil {
		t.Fatalf(err.Error())
	}

	client, server, cErr := newTestRundeckClient(jsonfile, "application/json", 500)
	defer server.Close()
	if cErr != nil {
		t.Fatalf(cErr.Error())
	}
	obj, oErr := client.DeleteAllExecutionsForJob("abcdefg")
	assert.Error(t, oErr)
	assert.Nil(t, obj)
}

func TestDeleteAllExecutionsForJobJSONError(t *testing.T) {
	client, server, cErr := newTestRundeckClient([]byte(""), "application/json", 200)
	defer server.Close()
	if cErr != nil {
		t.Fatalf(cErr.Error())
	}
	obj, oErr := client.DeleteAllExecutionsForJob("abcdefg")
	assert.Error(t, oErr)
	assert.Nil(t, obj)
}
