package rundeck

import (
	"encoding/xml"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"testing"
)

func TestErrorMessage(t *testing.T) {
	xmlfile, err := os.Open("assets/test/errormessage.xml")
	if err != nil {
		t.Fatal(err.Error())
	}
	defer xmlfile.Close()
	xmlData, _ := ioutil.ReadAll(xmlfile)
	var s RundeckResult
	xml.Unmarshal(xmlData, &s)
	assert.True(t, s.Errored, "It should be an error")
	assert.False(t, s.Succeeded, "It should not be successful")
	assert.Len(t, s.ErrorMessages, 2, "Should have two messages")
	assert.Len(t, s.SuccessMessages, 0, "Should not have success messages")
}

func TestSuccessMessage(t *testing.T) {
	xmlfile, err := os.Open("assets/test/successmessage.xml")
	if err != nil {
		t.Fatal(err.Error())
	}
	defer xmlfile.Close()
	xmlData, _ := ioutil.ReadAll(xmlfile)
	var s RundeckResult
	xml.Unmarshal(xmlData, &s)
	assert.False(t, s.Errored, "It should not be an error")
	assert.True(t, s.Succeeded, "It should be successful")
	assert.Len(t, s.SuccessMessages, 2, "Should have two messages")
	assert.Len(t, s.ErrorMessages, 0, "Should not have error messages")
}
