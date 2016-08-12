package rundeck

import (
	"encoding/xml"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"testing"
)

func TestDeletedExecutionsOutput(t *testing.T) {
	xmlfile, err := os.Open("assets/test/delete_executions.xml")
	if err != nil {
		t.Fatalf(err.Error())
	}
	defer xmlfile.Close()
	xmlData, _ := ioutil.ReadAll(xmlfile)
	var s ExecutionsDeleted
	err = xml.Unmarshal(xmlData, &s)
	if err != nil {
		t.Fatalf(err.Error())
	}
	assert.EqualValues(t, s.RequestCount, 4)
	assert.Len(t, s.Failed.Failures, 4)
}
