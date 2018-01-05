package rundeck

import (
	"testing"

	"github.com/lusis/go-rundeck/pkg/rundeck.v21/responses"
	"github.com/lusis/go-rundeck/pkg/rundeck.v21/responses/testdata"
	"github.com/stretchr/testify/assert"
)

func TestEnableSchedule(t *testing.T) {
	client, server, cErr := newTestRundeckClient([]byte(`{"success":true}`), "application/json", 200)
	defer server.Close()
	if cErr != nil {
		t.Fatalf(cErr.Error())
	}

	res, err := client.EnableSchedule("1")
	assert.NoError(t, err)
	assert.True(t, res)
}

func TestDisableSchedule(t *testing.T) {
	client, server, cErr := newTestRundeckClient([]byte(`{"success":true}`), "application/json", 200)
	defer server.Close()
	if cErr != nil {
		t.Fatalf(cErr.Error())
	}

	res, err := client.DisableSchedule("1")
	assert.NoError(t, err)
	assert.True(t, res)
}
func TestBulkEnableSchedule(t *testing.T) {
	jsonfile, err := testdata.GetBytes(responses.BulkToggleResponseTestFile)
	if err != nil {
		t.Fatalf(err.Error())
	}

	client, server, _ := newTestRundeckClient(jsonfile, "application/json", 200)
	defer server.Close()
	obj, cErr := client.BulkEnableSchedule("a", "b", "c")
	assert.NoError(t, cErr)
	assert.NotNil(t, obj)
}

func TestBulkDisableSchedule(t *testing.T) {
	jsonfile, err := testdata.GetBytes(responses.BulkToggleResponseTestFile)
	if err != nil {
		t.Fatalf(err.Error())
	}

	client, server, _ := newTestRundeckClient(jsonfile, "application/json", 200)
	defer server.Close()
	obj, cErr := client.BulkDisableSchedule("a", "b", "c")
	assert.NoError(t, cErr)
	assert.NotNil(t, obj)
}
