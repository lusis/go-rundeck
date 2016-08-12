package rundeck

import (
	"encoding/xml"
	"io/ioutil"
	"os"
	"testing"
)

func TestExecutionOutput(t *testing.T) {
	xmlfile, err := os.Open("assets/test/execution.xml")
	if err != nil {
		t.Fatalf(err.Error())
	}
	defer xmlfile.Close()
	xmlData, _ := ioutil.ReadAll(xmlfile)
	var s ExecutionOutput
	xml.Unmarshal(xmlData, &s)

	intexpects(s.ID, 23, t)
	intexpects(s.Offset, 424, t)
	strexpects(s.ExecState, "succeeded", t)
	intexpects(s.ExecDuration, 409, t)
	intexpects(s.TotalSize, 430, t)
	strexpects(s.Entries.Entry[0].Log, "hello", t)
	strexpects(s.Entries.Entry[1].Log, "bye", t)
}
