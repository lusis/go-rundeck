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

type JobIntegrationTestSuite struct {
	suite.Suite
	TestClient      *rundeck.Client
	CreatedProjects []rundeck.Project
	sync.Mutex
}

func (s *JobIntegrationTestSuite) testCreateProject(slow bool) (rundeck.Project, rundeck.JobMetaData) {
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

func (s *JobIntegrationTestSuite) SetupSuite() {
	client := testNewTokenAuthClient()
	s.CreatedProjects = []rundeck.Project{}
	s.TestClient = client
}

func (s *JobIntegrationTestSuite) TearDownSuite() {
	for _, p := range s.CreatedProjects {
		e := s.TestClient.DeleteProject(p.Name)
		if e != nil {
			s.T().Logf("unable to clean up test project: %s", e.Error())
		}
	}
}

func (s *JobIntegrationTestSuite) TestImportJob() {
	project, _ := s.testCreateProject(false)
	importJob, err := s.TestClient.ImportJob(project.Name,
		strings.NewReader(testJobDefinition),
		rundeck.ImportFormat("yaml"),
		rundeck.ImportUUID("remove"))
	s.Require().NoError(err)
	s.Require().NotNil(importJob)
	s.Require().Len(importJob.Failed, 0)
	s.Require().Len(importJob.Skipped, 0)
	s.Require().Len(importJob.Succeeded, 1)
}

func (s *JobIntegrationTestSuite) TestRunJob() {
	_, job := s.testCreateProject(false)
	runJob, err := s.TestClient.RunJob(job.ID)
	s.Require().NoError(err)
	doneFunc := func() (bool, error) {
		time.Sleep(500 * time.Millisecond)
		info, infoErr := s.TestClient.GetExecutionState(runJob.ID)
		if infoErr != nil {
			return false, infoErr
		}
		return info.Completed, nil
	}
	done, err := s.TestClient.WaitFor(doneFunc, 5*time.Second)
	s.Require().NoError(err)
	s.Require().True(done)
	info, _ := s.TestClient.GetExecutionState(runJob.ID)
	s.Require().Equal("SUCCEEDED", info.ExecutionState)
}

func (s *JobIntegrationTestSuite) TestFindJobByName() {
	project, job := s.testCreateProject(false)
	res, err := s.TestClient.FindJobByName(job.Name)
	s.Require().NoError(err)
	s.Require().Len(res, 1)
	s.Require().Equal(project.Name, res[0].Project)

}
func TestIntegrationJobSuite(t *testing.T) {
	if testing.Short() || testRundeckRunning() == false {
		t.Skip("skipping integration testing")
	}

	suite.Run(t, &JobIntegrationTestSuite{})
}
