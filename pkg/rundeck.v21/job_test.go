package rundeck

import (
	"testing"

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
