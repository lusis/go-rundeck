package rundeck

import (
	"errors"
	"testing"
	"time"

	requests "github.com/lusis/go-rundeck/pkg/rundeck/requests"
	"github.com/lusis/go-rundeck/pkg/rundeck/responses"
	yaml "gopkg.in/yaml.v2"

	"github.com/stretchr/testify/require"
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
	require.NoError(t, oErr)
	require.NotNil(t, obj)
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
	require.Error(t, oErr)
	require.Nil(t, obj)
}

func TestGetJobMetaDataJSONError(t *testing.T) {
	client, server, cErr := newTestRundeckClient([]byte(""), "application/json", 200)
	defer server.Close()
	if cErr != nil {
		t.Fatalf(cErr.Error())
	}
	obj, oErr := client.GetJobMetaData("1")
	require.Error(t, oErr)
	require.Nil(t, obj)
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
	require.NoError(t, oErr)
	require.NotEmpty(t, obj)
	data := &responses.JobYAMLResponse{}
	yErr := yaml.Unmarshal(obj, &data)
	require.NoError(t, yErr)
	require.NotNil(t, data)
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
	require.Error(t, oErr)
	require.Nil(t, obj)
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
	require.NoError(t, cErr)
	require.NotNil(t, obj)
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
	require.NoError(t, cErr)
	require.NotNil(t, obj)
}

func TestListJobsJSONError(t *testing.T) {
	client, server, cErr := newTestRundeckClient([]byte(""), "application/json", 200)
	defer server.Close()
	if cErr != nil {
		t.Fatalf(cErr.Error())
	}
	obj, cErr := client.ListJobs("testproject")
	require.Error(t, cErr)
	require.Empty(t, obj)
}

func TestListJobsHTTPError(t *testing.T) {
	client, server, cErr := newTestRundeckClient([]byte(""), "application/json", 500)
	defer server.Close()
	if cErr != nil {
		t.Fatalf(cErr.Error())
	}
	obj, cErr := client.ListJobs("testproject")
	require.Error(t, cErr)
	require.Empty(t, obj)
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
			require.NoError(t, err)
		}
	}
	require.Equal(t, "-foo bar", jobOpts.ArgString)
	require.Equal(t, "auser", jobOpts.AsUser)
	require.Equal(t, "bar", jobOpts.Options["foo"])
	require.Equal(t, "DEBUG", jobOpts.LogLevel)
	require.NotNil(t, jobOpts.RunAtTime)
	require.Equal(t, ".*", jobOpts.Filter)
}

func TestRunJobOptionError(t *testing.T) {
	client, server, _ := newTestRundeckClient([]byte(""), "application/json", 200)
	defer server.Close()
	res, err := client.RunJob("abcdefg", testFailedJobOption())
	require.Nil(t, res)
	require.Error(t, err)
}

func TestRunJobHTTPError(t *testing.T) {
	client, server, _ := newTestRundeckClient([]byte(""), "application/json", 500)
	defer server.Close()
	res, err := client.RunJob("abcdefg")
	require.Nil(t, res)
	require.Error(t, err)
}

func TestRunJobJSONError(t *testing.T) {
	client, server, _ := newTestRundeckClient([]byte(""), "application/json", 200)
	defer server.Close()
	res, err := client.RunJob("abcdefg")
	require.Nil(t, res)
	require.Error(t, err)
}

func TestDeleteJobFound(t *testing.T) {
	client, server, cErr := newTestRundeckClient([]byte(""), "application/json", 204)
	defer server.Close()
	if cErr != nil {
		t.Fatalf(cErr.Error())
	}
	err := client.DeleteJob("testproject")
	require.NoError(t, err)
}

func TestDeleteJobNotFound(t *testing.T) {
	client, server, cErr := newTestRundeckClient([]byte(""), "application/json", 404)
	defer server.Close()
	if cErr != nil {
		t.Fatalf(cErr.Error())
	}
	err := client.DeleteJob("testproject")
	require.EqualError(t, ErrMissingResource, err.Error())
}

func TestGetRequiredOptsHTTPError(t *testing.T) {
	client, server, _ := newTestRundeckClient([]byte(""), "application/json", 500)
	defer server.Close()
	res, err := client.GetRequiredOpts("abcdefg")
	require.Nil(t, res)
	require.Error(t, err)
}

func TestGetRequiredOptsYAMLError(t *testing.T) {
	client, server, _ := newTestRundeckClient([]byte("1234"), "application/json", 200)
	defer server.Close()
	res, err := client.GetRequiredOpts("abcdefg")
	require.Nil(t, res)
	require.Error(t, err)
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
	require.NoError(t, cErr)
	require.NotNil(t, obj)
}

func TestGetJobOptsHTTPError(t *testing.T) {
	client, server, _ := newTestRundeckClient([]byte(""), "application/json", 500)
	defer server.Close()
	res, err := client.GetJobOpts("abcdefg")
	require.Nil(t, res)
	require.Error(t, err)
}

func TestGetJobOptsYAMLError(t *testing.T) {
	client, server, _ := newTestRundeckClient([]byte("1234"), "application/json", 200)
	defer server.Close()
	res, err := client.GetJobOpts("abcdefg")
	require.Nil(t, res)
	require.Error(t, err)
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
	require.NoError(t, cErr)
	require.NotNil(t, obj)
}

func TestExportJobInvalidFormat(t *testing.T) {
	client, server, cErr := newTestRundeckClient([]byte(""), "application/yaml", 200)
	defer server.Close()
	if cErr != nil {
		t.Fatalf(cErr.Error())
	}
	res, err := client.ExportJob("abcdefg", "json")
	require.Nil(t, res)
	require.Error(t, err)
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
	require.NoError(t, cErr)
	require.NotEmpty(t, obj)
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
	require.Error(t, oErr)
	require.Nil(t, obj)
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
	require.NoError(t, oErr)
	require.NotNil(t, obj)
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
	require.Error(t, oErr)
	require.Nil(t, obj)
}

func TestDeleteAllExecutionsForJobJSONError(t *testing.T) {
	client, server, cErr := newTestRundeckClient([]byte(""), "application/json", 200)
	defer server.Close()
	if cErr != nil {
		t.Fatalf(cErr.Error())
	}
	obj, oErr := client.DeleteAllExecutionsForJob("abcdefg")
	require.Error(t, oErr)
	require.Nil(t, obj)
}
