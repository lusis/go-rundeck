package rundeck

import (
	"testing"
	"time"

	requests "github.com/lusis/go-rundeck/pkg/rundeck.v21/requests"
	"github.com/lusis/go-rundeck/pkg/rundeck.v21/responses"
	"github.com/lusis/go-rundeck/pkg/rundeck.v21/responses/testdata"
	yaml "gopkg.in/yaml.v2"

	"github.com/stretchr/testify/assert"
)

func TestGetJobMetaData(t *testing.T) {
	jsonfile, err := testdata.GetBytes(responses.JobMetaDataResponseTestFile)
	if err != nil {
		t.Fatalf(err.Error())
	}

	client, server, cErr := newTestRundeckClient(jsonfile, "application/json", 200)
	defer server.Close()
	if cErr != nil {
		t.Fatalf(cErr.Error())
	}
	obj, cErr := client.GetJobMetaData("1")
	assert.NoError(t, cErr)
	assert.NotNil(t, obj)
}

func TestGetJobDefinition(t *testing.T) {
	jsonfile, err := testdata.GetBytes("job_definition.yaml")
	if err != nil {
		t.Fatalf(err.Error())
	}

	client, server, cErr := newTestRundeckClient(jsonfile, "application/yaml", 200)
	defer server.Close()
	if cErr != nil {
		t.Fatalf(cErr.Error())
	}
	obj, cErr := client.GetJobDefinition("1", "yaml")
	assert.NoError(t, cErr)
	assert.NotEmpty(t, obj)
	data := &responses.JobYAMLResponse{}
	yErr := yaml.Unmarshal(obj, &data)
	assert.NoError(t, yErr)
	assert.NotNil(t, data)
}

func TestGetJobInfo(t *testing.T) {
	jsonfile, err := testdata.GetBytes(responses.JobMetaDataResponseTestFile)
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
	jsonfile, err := testdata.GetBytes(responses.JobsResponseTestFile)
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

func TestDeleteJobFound(t *testing.T) {
	client, server, cErr := newTestRundeckClient([]byte(""), "application/json", 201)
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
