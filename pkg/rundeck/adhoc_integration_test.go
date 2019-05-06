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
	project, err := s.TestClient.CreateProject(projectName, props)
	s.Require().NoError(err)
	s.Lock()
	s.CreatedProjects = append(s.CreatedProjects, *project)
	s.Unlock()
	jobbytes, err := testJobFromTemplate(projectName+"-job", "job for "+projectName)
	s.Require().NoError(err)
	importJob, err := s.TestClient.ImportJob(project.Name,
		bytes.NewReader(jobbytes),
		rundeck.ImportFormat("yaml"),
		rundeck.ImportUUID("remove"))
	s.Require().NoError(err)
	j, err := s.TestClient.GetJobMetaData(importJob.Succeeded[0].ID)
	s.Require().NoError(err)
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
	ahe, err := s.TestClient.RunAdHocCommand(project.Name, "ps -ef", rundeck.CmdThreadCount(3))
	s.Require().NoError(err)

	doneFunc := func() (bool, error) {
		time.Sleep(500 * time.Millisecond)
		info, infoErr := s.TestClient.GetExecutionState(ahe.Execution.ID)
		if infoErr != nil {
			return false, infoErr
		}
		return info.Completed, nil
	}
	done, err := s.TestClient.WaitFor(doneFunc, 5*time.Second)
	s.Require().NoError(err)
	s.Require().True(done)
	info, _ := s.TestClient.GetExecutionState(ahe.Execution.ID)
	s.Require().Equal("SUCCEEDED", info.ExecutionState)
}

func (s *AdHocIntegrationTestSuite) TestAdHocScript() {
	project, _ := s.testCreateProject(false)
	testScript := `#!/bin/bash
echo "hello"
	`
	ahe, err := s.TestClient.RunAdHocScript(project.Name, strings.NewReader(testScript))
	s.Require().NoError(err)

	doneFunc := func() (bool, error) {
		time.Sleep(500 * time.Millisecond)
		info, infoErr := s.TestClient.GetExecutionState(ahe.Execution.ID)
		if infoErr != nil {
			return false, infoErr
		}
		return info.Completed, nil
	}
	done, err := s.TestClient.WaitFor(doneFunc, 5*time.Second)
	s.Require().NoError(err)
	s.Require().True(done)
	info, _ := s.TestClient.GetExecutionState(ahe.Execution.ID)
	s.Require().Equal("SUCCEEDED", info.ExecutionState)
}

func (s *AdHocIntegrationTestSuite) TestAdHocScriptURL() {
	project, _ := s.testCreateProject(false)
	ahe, err := s.TestClient.RunAdHocScriptFromURL(project.Name, testAdHocScriptURL)
	s.Require().NoError(err)
	doneFunc := func() (bool, error) {
		time.Sleep(500 * time.Millisecond)
		info, infoErr := s.TestClient.GetExecutionState(ahe.Execution.ID)
		if infoErr != nil {
			return false, infoErr
		}
		return info.Completed, nil
	}
	done, err := s.TestClient.WaitFor(doneFunc, 5*time.Second)
	s.Require().NoError(err)
	s.Require().True(done)
	info, _ := s.TestClient.GetExecutionState(ahe.Execution.ID)
	s.Require().Equal("SUCCEEDED", info.ExecutionState)
}

func TestIntegrationAdHocSuite(t *testing.T) {
	if testing.Short() || ! testRundeckRunning() {
		t.Skip("skipping integration testing")
	}
	suite.Run(t, &AdHocIntegrationTestSuite{})
}
