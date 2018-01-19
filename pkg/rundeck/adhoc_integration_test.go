package rundeck_test

import (
	"bytes"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/lusis/go-rundeck/pkg/rundeck"
	"github.com/stretchr/testify/suite"
)

type AdHocIntegrationTestSuite struct {
	suite.Suite
	CreatedProjects []rundeck.Project
	TestClient      *rundeck.Client
	sync.Mutex
}

func (s *AdHocIntegrationTestSuite) testCreateProject(slow bool) (rundeck.Project, rundeck.JobMetaData) {
	projectName := testGenerateRandomName(s.T().Name())
	props := map[string]string{
		"project.description": projectName,
	}
	for k, v := range testDefaultProjectProperties {
		props[k] = v
	}
	if slow {
		// make these slow nodes
		props["resources.source.1.config.delay"] = "10"
		props["resources.source.1.config.count"] = "100"
	}
	project, createErr := s.TestClient.CreateProject(projectName, props)
	if createErr != nil {
		s.T().Fatalf("Unable to create test project: %s", createErr.Error())
	}
	s.Lock()
	s.CreatedProjects = append(s.CreatedProjects, *project)
	s.Unlock()
	jobbytes, joberr := testJobFromTemplate(projectName+"-job", "job for "+projectName)
	if joberr != nil {
		s.T().Fatalf("cannot create a job import from template. cannot continue: %s", joberr.Error())
	}
	importJob, importErr := s.TestClient.ImportJob(project.Name,
		bytes.NewReader(jobbytes),
		rundeck.ImportFormat("yaml"),
		rundeck.ImportUUID("remove"))
	if importErr != nil {
		s.T().Fatalf("job did not import. cannot continue: %s", importErr.Error())
	}
	j, jerr := s.TestClient.GetJobMetaData(importJob.Succeeded[0].ID)
	if jerr != nil {
		s.T().Fatalf("unable to get job meta data for imported job. cannot continue: %s", jerr.Error())
	}
	return *project, *j
}

func (s *AdHocIntegrationTestSuite) SetupSuite() {
	client := testNewTokenAuthClient()
	s.CreatedProjects = []rundeck.Project{}
	s.TestClient = client
}

func (s *AdHocIntegrationTestSuite) TearDownSuite() {
	for _, p := range s.CreatedProjects {
		e := s.TestClient.DeleteProject(p.Name)
		if e != nil {
			s.T().Logf("unable to clean up test project: %s", e.Error())
		}
	}
}

func (s *AdHocIntegrationTestSuite) TestAdHocCommand() {
	project, _ := s.testCreateProject(false)
	ahe, aheErr := s.TestClient.RunAdHocCommand(project.Name, "ps -ef", rundeck.CmdThreadCount(3))
	if aheErr != nil {
		s.T().Fatalf("unable to run adhoc command. cannot continue: %s", aheErr.Error())
	}
	doneFunc := func() (bool, error) {
		time.Sleep(500 * time.Millisecond)
		info, infoErr := s.TestClient.GetExecutionState(ahe.Execution.ID)
		if infoErr != nil {
			return false, infoErr
		}
		return info.Completed, nil
	}
	done, doneErr := s.TestClient.WaitFor(doneFunc, 5*time.Second)
	s.NoError(doneErr)
	s.True(done)
	info, _ := s.TestClient.GetExecutionState(ahe.Execution.ID)
	s.Equal("SUCCEEDED", info.ExecutionState)
}

func (s *AdHocIntegrationTestSuite) TestAdHocScript() {
	project, _ := s.testCreateProject(false)
	testScript := `#!/bin/bash
echo "hello"
	`
	ahe, aheErr := s.TestClient.RunAdHocScript(project.Name, strings.NewReader(testScript))
	if aheErr != nil {
		s.T().Fatalf("unable to run adhoc script. cannot continue: %s", aheErr.Error())
	}
	doneFunc := func() (bool, error) {
		time.Sleep(500 * time.Millisecond)
		info, infoErr := s.TestClient.GetExecutionState(ahe.Execution.ID)
		if infoErr != nil {
			return false, infoErr
		}
		return info.Completed, nil
	}
	done, doneErr := s.TestClient.WaitFor(doneFunc, 5*time.Second)
	s.NoError(doneErr)
	s.True(done)
	info, _ := s.TestClient.GetExecutionState(ahe.Execution.ID)
	s.Equal("SUCCEEDED", info.ExecutionState)
}

func (s *AdHocIntegrationTestSuite) TestAdHocScriptURL() {
	project, _ := s.testCreateProject(false)
	ahe, aheErr := s.TestClient.RunAdHocScriptFromURL(project.Name, testAdHocScriptURL)
	if aheErr != nil {
		s.T().Fatalf("unable to run adhoc script. cannot continue: %s", aheErr.Error())
	}
	doneFunc := func() (bool, error) {
		time.Sleep(500 * time.Millisecond)
		info, infoErr := s.TestClient.GetExecutionState(ahe.Execution.ID)
		if infoErr != nil {
			return false, infoErr
		}
		return info.Completed, nil
	}
	done, doneErr := s.TestClient.WaitFor(doneFunc, 5*time.Second)
	s.NoError(doneErr)
	s.True(done)
	info, _ := s.TestClient.GetExecutionState(ahe.Execution.ID)
	s.Equal("SUCCEEDED", info.ExecutionState)
}

func TestIntegrationAdHocSuite(t *testing.T) {
	if testRundeckRunning() {
		suite.Run(t, &AdHocIntegrationTestSuite{})
	} else {
		t.Skip("rundeck isn't running for integration testing")
	}
}
