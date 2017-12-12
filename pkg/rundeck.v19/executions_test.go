package rundeck

import (
	"encoding/xml"
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeletedExecutionsOutput(t *testing.T) {
	xmlfile, err := os.Open("assets/test/delete_executions.xml")
	if err != nil {
		t.Fatalf(err.Error())
	}
	defer func() { _ = xmlfile.Close() }()
	xmlData, _ := ioutil.ReadAll(xmlfile)
	var s ExecutionsDeleted
	err = xml.Unmarshal(xmlData, &s)
	if err != nil {
		t.Fatalf(err.Error())
	}
	assert.EqualValues(t, s.RequestCount, 4)
	assert.Len(t, s.Failed.Failures, 4)
}
