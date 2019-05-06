package rundeck_test

import (
	"fmt"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/lusis/go-rundeck/pkg/rundeck"
	"github.com/stretchr/testify/suite"
)

type ExecutionIntegrationTestSuite struct {
	suite.Suite
	TestClient      *rundeck.Client
	CreatedProjects []rundeck.Project
	sync.Mutex
}

func (s *ExecutionIntegrationTestSuite) testCreateProject(slow bool) (rundeck.Project, rundeck.JobImportResult) {
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
	importJob, err := s.TestClient.ImportJob(project.Name,
		strings.NewReader(testJobDefinition),
		rundeck.ImportFormat("yaml"),
		rundeck.ImportUUID("remove"))
	s.Require().NoError(err)
	return *project, *importJob
}

func (s *ExecutionIntegrationTestSuite) SetupSuite() {
	client := testNewTokenAuthClient()
	s.TestClient = client
	s.CreatedProjects = []rundeck.Project{}
}

func (s *ExecutionIntegrationTestSuite) TearDownSuite() {

	for _, p := range s.CreatedProjects {
		e := s.TestClient.DeleteProject(p.Name)
		if e != nil {
			s.T().Logf("unable to clean up test project: %s", e.Error())
		}
	}
}

func (s *ExecutionIntegrationTestSuite) TestAbortExecution() {
	_, job := s.testCreateProject(true)
	jobID := job.Succeeded[0].ID
	runTime := time.Now().Add(1 * time.Hour)
	ahe, err := s.TestClient.RunJob(jobID, rundeck.RunJobRunAt(runTime))
	s.Require().NoError(err)

	exec, err := s.TestClient.GetExecutionInfo(ahe.ID)
	s.Require().NoError(err)
	s.Require().Equal("scheduled", exec.Status)

	info, err := s.TestClient.AbortExecution(ahe.ID)
	s.Require().NoError(err)
	s.Require().Equal(fmt.Sprintf("%d", ahe.ID), info.Execution.ID)
}

func (s *ExecutionIntegrationTestSuite) TestAbortExecutionAsUser() {
	_, job := s.testCreateProject(true)
	jobID := job.Succeeded[0].ID
	runtime := time.Now().Add(1 * time.Hour)
	ahe, err := s.TestClient.RunJob(jobID, rundeck.RunJobRunAt(runtime))
	s.Require().NoError(err)

	check, err := s.TestClient.GetExecutionInfo(ahe.ID)
	s.Require().NoError(err)
	s.Require().Equal("scheduled", check.Status)

	info, err := s.TestClient.AbortExecution(ahe.ID, rundeck.AbortExecutionAsUser("auser"))
	s.Require().NoError(err)
	s.Require().Equal(fmt.Sprintf("%d", ahe.ID), info.Execution.ID)

}

func (s *ExecutionIntegrationTestSuite) TestGetExecutionInfo() {
	project, _ := s.testCreateProject(true)
	ahe, err := s.TestClient.RunAdHocCommand(project.Name, "ps -ef", rundeck.CmdThreadCount(1))
	s.Require().NoError(err)

	info, err := s.TestClient.GetExecutionInfo(ahe.Execution.ID)
	s.Require().NoError(err)
	s.Require().NotNil(info)
}

func (s *ExecutionIntegrationTestSuite) TestBulkDeleteExecutions() {
	_, job := s.testCreateProject(false)
	jobID := job.Succeeded[0].ID
	var runningIDs []int
	// start 5 jobs
	for i := 1; i <= 5; i++ {
		ahe, err := s.TestClient.RunJob(jobID)
		s.Require().NoError(err)
		runningIDs = append(runningIDs, ahe.ID)
		doneFunc := func() (bool, error) {
			time.Sleep(500 * time.Millisecond)
			info, infoErr := s.TestClient.GetExecutionState(ahe.ID)
			if infoErr != nil {
				return false, infoErr
			}
			return info.Completed, nil
		}
		_, err = s.TestClient.WaitFor(doneFunc, 10*time.Second)
		s.Require().NoError(err)
	}

	bd, err := s.TestClient.BulkDeleteExecutions(runningIDs...)
	s.Require().NoError(err)
	s.Require().Equal(bd.SuccessCount, len(runningIDs))
}

func (s *ExecutionIntegrationTestSuite) TestListRunningExecutions() {
	project, job := s.testCreateProject(false)
	jobID := job.Succeeded[0].ID

	ahe, err := s.TestClient.RunJob(jobID)
	s.Require().NoError(err)

	bd, err := s.TestClient.ListRunningExecutions(project.Name)
	s.Require().NoError(err)
	s.Require().Len(bd.Executions, 1)
	s.Require().Equal(jobID, bd.Executions[0].Job.ID)
	doneFunc := func() (bool, error) {
		time.Sleep(500 * time.Millisecond)
		info, infoErr := s.TestClient.GetExecutionState(ahe.ID)
		if infoErr != nil {
			return false, infoErr
		}
		return info.Completed, nil
	}
	// wait for the job to be done before we exit
	done, err := s.TestClient.WaitFor(doneFunc, 10*time.Second)
	s.Require().NoError(err)
	s.Require().True(done)
}

func TestIntegrationExecutionSuite(t *testing.T) {
	if testing.Short() || testRundeckRunning() == false {
		t.Skip("skipping integration testing")
	}

	suite.Run(t, &ExecutionIntegrationTestSuite{})
}
