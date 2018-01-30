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
	project, createErr := s.TestClient.CreateProject(projectName, props)
	if createErr != nil {
		s.T().Fatalf("Unable to create test project: %s", createErr.Error())
	}
	s.Lock()
	s.CreatedProjects = append(s.CreatedProjects, *project)
	s.Unlock()
	importJob, importErr := s.TestClient.ImportJob(project.Name,
		strings.NewReader(testJobDefinition),
		rundeck.ImportFormat("yaml"),
		rundeck.ImportUUID("remove"))
	if importErr != nil {
		s.T().Fatalf("job did not import. cannot continue: %s", importErr.Error())
	}
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
	ahe, aheErr := s.TestClient.RunJob(jobID, rundeck.RunJobRunAt(runTime))
	if aheErr != nil {
		s.T().Fatalf("unable to start job. cannot continue: %s", aheErr.Error())
	}
	exec, execErr := s.TestClient.GetExecutionInfo(ahe.ID)
	if execErr != nil {
		s.T().Fatalf("unable to get execution. cannot continue: %s", execErr.Error())
	}
	if exec.Status != "scheduled" {
		s.T().Fatal("cannot schedule job. cannot continue")
	}
	info, infoErr := s.TestClient.AbortExecution(ahe.ID)
	if infoErr != nil {
		s.T().Fatalf("unable to abort execution. cannot continue: %s", infoErr.Error())
	}
	s.Equal(fmt.Sprintf("%d", ahe.ID), info.Execution.ID)
}

func (s *ExecutionIntegrationTestSuite) TestAbortExecutionAsUser() {
	_, job := s.testCreateProject(true)
	jobID := job.Succeeded[0].ID
	runtime := time.Now().Add(1 * time.Hour)
	ahe, aheErr := s.TestClient.RunJob(jobID, rundeck.RunJobRunAt(runtime))
	if aheErr != nil {
		s.T().Fatalf("unable to run adhoc command. cannot continue: %s", aheErr.Error())
	}

	check, checkErr := s.TestClient.GetExecutionInfo(ahe.ID)
	if checkErr != nil {
		s.T().Fatalf("unable to get execution info. cannot continue: %s", checkErr)
	}
	if check.Status != "scheduled" {
		s.T().Fatal("cannot schedule job. cannot continue")
	}

	info, infoErr := s.TestClient.AbortExecution(ahe.ID, rundeck.AbortExecutionAsUser("auser"))
	if infoErr != nil {
		s.T().Fatalf("unabled to abort execution. cannot continue: %s", infoErr.Error())
	}
	s.Equal(fmt.Sprintf("%d", ahe.ID), info.Execution.ID)

}

func (s *ExecutionIntegrationTestSuite) TestGetExecutionInfo() {
	project, _ := s.testCreateProject(true)
	ahe, aheErr := s.TestClient.RunAdHocCommand(project.Name, "ps -ef", rundeck.CmdThreadCount(1))
	if aheErr != nil {
		s.T().Fatalf("unable to run adhoc command. cannot continue: %s", aheErr.Error())
	}

	info, infoErr := s.TestClient.GetExecutionInfo(ahe.Execution.ID)
	s.NoError(infoErr)
	s.NotNil(info)
}

func (s *ExecutionIntegrationTestSuite) TestBulkDeleteExecutions() {
	_, job := s.testCreateProject(false)
	jobID := job.Succeeded[0].ID
	var runningIDs []int
	// start 5 jobs
	for i := 1; i <= 5; i++ {
		ahe, aheErr := s.TestClient.RunJob(jobID)
		if aheErr != nil {
			s.T().Fatalf("unable to run job. cannot continue: %s", aheErr.Error())
		}
		runningIDs = append(runningIDs, ahe.ID)
		doneFunc := func() (bool, error) {
			time.Sleep(500 * time.Millisecond)
			info, infoErr := s.TestClient.GetExecutionState(ahe.ID)
			if infoErr != nil {
				return false, infoErr
			}
			return info.Completed, nil
		}
		_, doneErr := s.TestClient.WaitFor(doneFunc, 10*time.Second)
		if doneErr != nil {
			s.T().Fatalf("job could not be checked. cannot continue: %s", doneErr.Error())
		}
	}

	bd, bderr := s.TestClient.BulkDeleteExecutions(runningIDs...)
	s.NoError(bderr)
	s.Equal(bd.SuccessCount, len(runningIDs))
}

func (s *ExecutionIntegrationTestSuite) TestListRunningExecutions() {
	project, job := s.testCreateProject(false)
	jobID := job.Succeeded[0].ID

	ahe, aheErr := s.TestClient.RunJob(jobID)
	if aheErr != nil {
		s.T().Fatalf("unable to run job. cannot continue: %s", aheErr.Error())
	}

	bd, bderr := s.TestClient.ListRunningExecutions(project.Name)
	s.NoError(bderr)
	s.Len(bd.Executions, 1)
	s.Equal(jobID, bd.Executions[0].Job.ID)
	doneFunc := func() (bool, error) {
		time.Sleep(500 * time.Millisecond)
		info, infoErr := s.TestClient.GetExecutionState(ahe.ID)
		if infoErr != nil {
			return false, infoErr
		}
		return info.Completed, nil
	}
	// wait for the job to be done before we exit
	done, doneErr := s.TestClient.WaitFor(doneFunc, 10*time.Second)
	s.NoError(doneErr)
	s.True(done)
}

func TestIntegrationExecutionSuite(t *testing.T) {
	if testRundeckRunning() {
		suite.Run(t, &ExecutionIntegrationTestSuite{})
	} else {
		t.Skip("rundeck isn't running for integration testing")
	}
}
