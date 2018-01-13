package rundeck

import (
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
)

type JobIntegrationTestSuite struct {
	suite.Suite
	TestProject *Project
	TestClient  *Client
	TestJobID   string
}

func (s *JobIntegrationTestSuite) SetupSuite() {
	client := testNewTokenAuthClient()
	projectName := testGenerateRandomName("testproject")
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

func (s *JobIntegrationTestSuite) TearDownSuite() {
	e := s.TestClient.DeleteProject(s.TestProject.Name)
	if e != nil {
		s.T().Errorf("unable to clean up test project: %s", e.Error())
	}
}

var testJobDefinition = `
- description: this is a test job
  executionEnabled: true
  group: test/jobs
  id: 8c3176bf-e553-4086-b7b7-38e19974cd89
  loglevel: INFO
  name: testjob
  nodeFilterEditable: false
  nodefilters:
    dispatch:
      excludePrecedence: true
      keepgoing: false
      rankOrder: ascending
      successOnEmptyNodeFilter: false
      threadcount: 1
    filter: .*
  nodesSelectedByDefault: true
  scheduleEnabled: true
  sequence:
    commands:
    - description: ps output
      exec: ps -ef
    keepgoing: false
    strategy: node-first
  uuid: 8c3176bf-e553-4086-b7b7-38e19974cd89
`

func (s *JobIntegrationTestSuite) TestImportJob() {
	importJob, importErr := s.TestClient.ImportJob(s.TestProject.Name,
		strings.NewReader(testJobDefinition),
		ImportFormat("yaml"),
		ImportUUID("remove"))
	if importErr != nil {
		s.T().Fatalf("job did not import. cannot continue: %s", importErr.Error())
	}
	s.NotNil(importJob)
	s.Len(importJob.Failed, 0)
	s.Len(importJob.Skipped, 0)
	s.Len(importJob.Succeeded, 1)
	s.TestJobID = importJob.Succeeded[0].ID
}

func (s *JobIntegrationTestSuite) TestRunJob() {
	runJob, runErr := s.TestClient.RunJob(s.TestJobID)
	if runErr != nil {
		s.T().Fatalf("job did not run. cannot continue: %s", runErr.Error())
	}
	runID := fmt.Sprintf("%d", runJob.ID)
	doneFunc := func() (bool, error) {
		time.Sleep(500 * time.Millisecond)
		info, infoErr := s.TestClient.GetExecutionState(runID)
		if infoErr != nil {
			return false, infoErr
		}
		return info.Completed, nil
	}
	done, doneErr := s.TestClient.WaitFor(doneFunc, 5*time.Second)
	s.NoError(doneErr)
	s.True(done)
	info, _ := s.TestClient.GetExecutionState(runID)
	s.Equal("SUCCEEDED", info.ExecutionState)
}

func TestIntegrationJobSuite(t *testing.T) {
	if testRundeckRunning() {
		suite.Run(t, &JobIntegrationTestSuite{})
	} else {
		t.Skip("rundeck isn't running for integration testing")
	}
}
