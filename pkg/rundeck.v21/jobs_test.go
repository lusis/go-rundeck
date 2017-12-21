package rundeck

import (
	"encoding/xml"
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJobOne(t *testing.T) {
	xmlfile, err := os.Open("assets/test/job1.xml")
	if err != nil {
		t.Fatal(err.Error())
	}
	defer func() { _ = xmlfile.Close() }()
	xmlData, _ := ioutil.ReadAll(xmlfile)
	var s JobList
	_ = xml.Unmarshal(xmlData, &s)
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
	defer func() { _ = xmlfile.Close() }()
	client, server, err := newTestRundeckClient(xmlfile, "application/xml", 200)
	defer server.Close()
	if err != nil {
		t.FailNow()
	}
	s, jobsErr := client.ListJobs("MYPROJ")
	if jobsErr != nil {
		t.FailNow()
	}
	assert.Equal(t, int64(3), s.Count)
	assert.Len(t, s.Jobs, int(s.Count))

	// Test first Job entry
	assert.Equal(t, "00000000-0000-0000-0000-000000000000", s.Jobs[0].ID)
	assert.Equal(t, "job1", s.Jobs[0].Name)
	assert.Equal(t, "ad-hoc", s.Jobs[0].Group)
	assert.Equal(t, "MYPROJ", s.Jobs[0].Project)
	assert.Equal(t, "Run job1", s.Jobs[0].Description)

	// Test second Job entry (different group)
	assert.Equal(t, "11111111-1111-1111-1111-111111111111", s.Jobs[1].ID)
	assert.Equal(t, "job2", s.Jobs[1].Name)
	assert.Equal(t, "nested/group", s.Jobs[1].Group)
	assert.Equal(t, "MYPROJ", s.Jobs[1].Project)
	assert.Equal(t, "Run job2 with a long description", s.Jobs[1].Description)

	// Test the third Job entry (no description, no group)
	assert.Equal(t, "aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa", s.Jobs[2].ID)
	assert.Equal(t, "job-with-empty-fields", s.Jobs[2].Name)
	assert.Equal(t, "", s.Jobs[2].Group)
	assert.Equal(t, "MYPROJ", s.Jobs[2].Project)
	assert.Equal(t, "", s.Jobs[2].Description)
}
