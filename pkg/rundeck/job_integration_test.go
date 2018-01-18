package rundeck_test

import (
	"strings"
	"testing"
	"time"

	"github.com/lusis/go-rundeck/pkg/rundeck"
	"github.com/stretchr/testify/suite"
)

type JobIntegrationTestSuite struct {
	suite.Suite
	TestProject *rundeck.Project
	TestClient  *rundeck.Client
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

func (s *JobIntegrationTestSuite) TestImportJob() {
	importJob, importErr := s.TestClient.ImportJob(s.TestProject.Name,
		strings.NewReader(testJobDefinition),
		rundeck.ImportFormat("yaml"),
		rundeck.ImportUUID("remove"))
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
	doneFunc := func() (bool, error) {
		time.Sleep(500 * time.Millisecond)
		info, infoErr := s.TestClient.GetExecutionState(runJob.ID)
		if infoErr != nil {
			return false, infoErr
		}
		return info.Completed, nil
	}
	done, doneErr := s.TestClient.WaitFor(doneFunc, 5*time.Second)
	s.NoError(doneErr)
	s.True(done)
	info, _ := s.TestClient.GetExecutionState(runJob.ID)
	s.Equal("SUCCEEDED", info.ExecutionState)
}

func TestIntegrationJobSuite(t *testing.T) {
	if testRundeckRunning() {
		suite.Run(t, &JobIntegrationTestSuite{})
	} else {
		t.Skip("rundeck isn't running for integration testing")
	}
}
