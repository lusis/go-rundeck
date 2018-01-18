package rundeck_test

import (
	"strings"
	"testing"
	"time"

	"github.com/lusis/go-rundeck/pkg/rundeck"
	"github.com/stretchr/testify/suite"
)

type AdHocIntegrationTestSuite struct {
	suite.Suite
	TestProject *rundeck.Project
	TestClient  *rundeck.Client
}

func (s *AdHocIntegrationTestSuite) SetupSuite() {
	client := testNewTokenAuthClient()
	projectName := testGenerateRandomName("adhoc")
	props := map[string]string{
		"project.description": projectName,
	}
	for k, v := range testDefaultProjectProperties {
		props[k] = v
	}
	project, createErr := client.CreateProject(projectName, props)
	if createErr != nil {
		s.T().Fatalf("Unable to create test project: %s", createErr.Error())
	}
	s.TestProject = project
	s.TestClient = client
}

func (s *AdHocIntegrationTestSuite) TearDownSuite() {
	e := s.TestClient.DeleteProject(s.TestProject.Name)
	if e != nil {
		s.T().Errorf("unable to clean up test project: %s", e.Error())
	}
}

func (s *AdHocIntegrationTestSuite) TestAdHocCommand() {
	ahe, aheErr := s.TestClient.RunAdHocCommand(s.TestProject.Name, "ps -ef", rundeck.CmdThreadCount(3))
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
	testScript := `#!/bin/bash
echo "hello"
	`
	ahe, aheErr := s.TestClient.RunAdHocScript(s.TestProject.Name, strings.NewReader(testScript))
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
	ahe, aheErr := s.TestClient.RunAdHocScriptFromURL(s.TestProject.Name, testAdHocScriptURL)
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
