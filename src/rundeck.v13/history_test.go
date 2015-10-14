package rundeck

import (
	"encoding/xml"
	"io/ioutil"
	"os"
	"testing"
)

func TestHistory(t *testing.T) {
	xmlfile, err := os.Open("assets/test/history.xml")
	if err != nil {
		t.Fatalf(err.Error())
	}
	defer xmlfile.Close()
	xmlData, _ := ioutil.ReadAll(xmlfile)
	var s Events
	xml.Unmarshal(xmlData, &s)

	intexpects(s.Count, 2, t)
	intexpects(s.Total, 630, t)
	intexpects(s.Max, 20, t)
	intexpects(s.Offset, 0, t)
	intexpects(int64(len(s.Events)), s.Count, t)

	// Test first history entry
	strexpects(s.Events[0].Title, "admin_tasks/thing1", t)
	strexpects(s.Events[0].Status, "succeeded", t)
	strexpects(s.Events[0].User, "bob", t)
	strexpects(s.Events[0].AbortedBy, "", t)
	strexpects(s.Events[0].Summary, "sudo download-artifacts.sh username ${option.release_version} ('calls download-artifacts.sh with a release version')", t)
	intexpects(s.Events[0].NodeSummary.Succeeded, 1, t)
	intexpects(s.Events[0].NodeSummary.Failed, 0, t)
	intexpects(s.Events[0].NodeSummary.Total, 1, t)
	strexpects(s.Events[0].StartTime, "1425406161632", t)
	strexpects(s.Events[0].Project, "MYPROJECT", t)
	strexpects(s.Events[0].EndTime, "1425406188349", t)
	strexpects(s.Events[0].Job.ID, "c41857ff-576e-4dc5-8a8c-9870412f245d", t)
	strexpects(s.Events[0].DateStarted, "2015-03-03T18:09:21Z", t)
	strexpects(s.Events[0].DateEnded, "2015-03-03T18:09:48Z", t)
	intexpects(s.Events[0].Execution.ID, 630, t)

	// Test second history entry (no job id, has abortedby)
	strexpects(s.Events[1].Title, "adhoc", t)
	strexpects(s.Events[1].Status, "failed", t)
	strexpects(s.Events[1].Summary, "uptime", t)
	intexpects(s.Events[1].NodeSummary.Succeeded, 1, t)
	intexpects(s.Events[1].NodeSummary.Failed, 1, t)
	intexpects(s.Events[1].NodeSummary.Total, 2, t)
	strexpects(s.Events[1].StartTime, "1425073242106", t)
	strexpects(s.Events[1].Project, "MYPROJECT", t)
	strexpects(s.Events[1].EndTime, "1425073243318", t)
	if s.Events[1].Job != nil {
		t.Errorf("Expected Job to be nil but got %v", s.Events[1].Job)
	}
	strexpects(s.Events[1].DateStarted, "2015-02-27T21:40:42Z", t)
	strexpects(s.Events[1].DateEnded, "2015-02-27T21:40:43Z", t)
	intexpects(s.Events[1].Execution.ID, 629, t)
}
