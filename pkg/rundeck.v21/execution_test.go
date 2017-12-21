package rundeck

import (
	"encoding/xml"
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExecutionOutput(t *testing.T) {
	xmlfile, err := os.Open("assets/test/execution.xml")
	if err != nil {
		t.Fatalf(err.Error())
	}
	defer func() { _ = xmlfile.Close() }()
	xmlData, _ := ioutil.ReadAll(xmlfile)
	var s ExecutionOutput
	_ = xml.Unmarshal(xmlData, &s)

	assert.Equal(t, int64(23), s.ID)
	assert.Equal(t, int64(424), s.Offset)
	assert.Equal(t, "succeeded", s.ExecState)
	assert.Equal(t, int64(409), s.ExecDuration)
	assert.Equal(t, int64(430), s.TotalSize)
	assert.Equal(t, "hello", s.Entries.Entry[0].Log)
	assert.Equal(t, "bye", s.Entries.Entry[1].Log)
}
