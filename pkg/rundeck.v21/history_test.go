package rundeck

import (
	"encoding/xml"
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHistory(t *testing.T) {
	xmlfile, err := os.Open("assets/test/history.xml")
	if err != nil {
		t.Fatalf(err.Error())
	}
	defer func() { _ = xmlfile.Close() }()
	xmlData, _ := ioutil.ReadAll(xmlfile)
	var s Events
	_ = xml.Unmarshal(xmlData, &s)

	assert.Equal(t, int64(2), s.Count)
	assert.Equal(t, int64(630), s.Total)
	assert.Equal(t, int64(20), s.Max)
	assert.Equal(t, int64(0), s.Offset)
	assert.Len(t, s.Events, int(s.Count))

	// Test first history entry
	assert.Equal(t, "admin_tasks/thing1", s.Events[0].Title)
	assert.Equal(t, "succeeded", s.Events[0].Status)
	assert.Equal(t, "bob", s.Events[0].User)
	assert.Equal(t, "", s.Events[0].AbortedBy)
	assert.Equal(t, "sudo download-artifacts.sh username ${option.release_version} ('calls download-artifacts.sh with a release version')", s.Events[0].Summary)
	assert.Equal(t, 1, s.Events[0].NodeSummary.Succeeded)
	assert.Equal(t, 0, s.Events[0].NodeSummary.Failed)
	assert.Equal(t, 1, s.Events[0].NodeSummary.Total)
	assert.Equal(t, "1425406161632", s.Events[0].StartTime)
	assert.Equal(t, "MYPROJECT", s.Events[0].Project)
	assert.Equal(t, "1425406188349", s.Events[0].EndTime)
	assert.Equal(t, "c41857ff-576e-4dc5-8a8c-9870412f245d", s.Events[0].Job.ID)
	assert.Equal(t, "2015-03-03T18:09:21Z", s.Events[0].DateStarted)
	assert.Equal(t, "2015-03-03T18:09:48Z", s.Events[0].DateEnded)
	assert.Equal(t, int64(630), s.Events[0].Execution.ID)

	// Test second history entry (no job id, has abortedby)
	assert.Equal(t, "adhoc", s.Events[1].Title)
	assert.Equal(t, "failed", s.Events[1].Status)
	assert.Equal(t, "uptime", s.Events[1].Summary)
	assert.Equal(t, 1, s.Events[1].NodeSummary.Succeeded)
	assert.Equal(t, 1, s.Events[1].NodeSummary.Failed)
	assert.Equal(t, 2, s.Events[1].NodeSummary.Total)
	assert.Equal(t, "1425073242106", s.Events[1].StartTime)
	assert.Equal(t, "MYPROJECT", s.Events[1].Project)
	assert.Equal(t, "1425073243318", s.Events[1].EndTime)
	if s.Events[1].Job != nil {
		t.Errorf("Expected Job to be nil but got %v", s.Events[1].Job)
	}
	assert.Equal(t, "2015-02-27T21:40:42Z", s.Events[1].DateStarted)
	assert.Equal(t, "2015-02-27T21:40:43Z", s.Events[1].DateEnded)
	assert.Equal(t, int64(629), s.Events[1].Execution.ID)
}
