package rundeck

import (
	"encoding/xml"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"testing"
)

func TestJobOne(t *testing.T) {
	xmlfile, err := os.Open("assets/test/job1.xml")
	if err != nil {
		t.Fatal(err.Error())
	}
	defer xmlfile.Close()
	xmlData, _ := ioutil.ReadAll(xmlfile)
	var s JobList
	xml.Unmarshal(xmlData, &s)
	scope := s.Job
	options := *scope.Context.Options
	assert.Len(t, options, 4, "Should have 4 options")
	for _, o := range options {
		if o.Name == "password" {
			assert.True(t, o.Required, "Password should be required")
			assert.True(t, o.ValueExposed, "valueExposed should be true")
		}
		if o.Name == "lastname" {
			assert.False(t, o.Required, "Last name should not be required")
			assert.False(t, o.ValueExposed, "valueExposed should be false")
		}
	}
	assert.Equal(t, "node-first", scope.Sequence.Strategy, "Strategy should be node-first")
	assert.Len(t, scope.Sequence.Steps, 2, "Should have two steps")
}

func TestJobs(t *testing.T) {
	xmlfile, err := os.Open("assets/test/jobs.xml")
	if err != nil {
		t.Fatalf(err.Error())
	}
	defer xmlfile.Close()
	xmlData, _ := ioutil.ReadAll(xmlfile)
	var s Jobs
	xml.Unmarshal(xmlData, &s)

	intexpects(s.Count, 3, t)
	intexpects(int64(len(s.Jobs)), s.Count, t)

	// Test first Job entry
	strexpects(s.Jobs[0].ID, "00000000-0000-0000-0000-000000000000", t)
	strexpects(s.Jobs[0].Name, "job1", t)
	strexpects(s.Jobs[0].Group, "ad-hoc", t)
	strexpects(s.Jobs[0].Project, "MYPROJ", t)
	strexpects(s.Jobs[0].Description, "Run job1", t)

	// Test second Job entry (different group)
	strexpects(s.Jobs[1].ID, "11111111-1111-1111-1111-111111111111", t)
	strexpects(s.Jobs[1].Name, "job2", t)
	strexpects(s.Jobs[1].Group, "nested/group", t)
	strexpects(s.Jobs[1].Project, "MYPROJ", t)
	strexpects(s.Jobs[1].Description, "Run job2 with a long description", t)

	// Test the third Job entry (no description, no group)
	strexpects(s.Jobs[2].ID, "aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa", t)
	strexpects(s.Jobs[2].Name, "job-with-empty-fields", t)
	strexpects(s.Jobs[2].Group, "", t)
	strexpects(s.Jobs[2].Project, "MYPROJ", t)
	strexpects(s.Jobs[2].Description, "", t)
}
