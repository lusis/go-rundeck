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
	importJob, importErr := s.TestClient.ImportJob(project.Name,
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
}

func (s *JobIntegrationTestSuite) TestRunJob() {
	_, job := s.testCreateProject(false)
	runJob, runErr := s.TestClient.RunJob(job.ID)
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

func (s *JobIntegrationTestSuite) TestFindJobByName() {
	project, job := s.testCreateProject(false)
	res, err := s.TestClient.FindJobByName(job.Name)
	if err != nil {
		s.T().Fatalf("Cannot find job. Cannot continue: %s", err.Error())
	}
	s.Len(res, 1)
	s.Equal(project.Name, res[0].Project)

}
func TestIntegrationJobSuite(t *testing.T) {
	if testRundeckRunning() {
		suite.Run(t, &JobIntegrationTestSuite{})
	} else {
		t.Skip("rundeck isn't running for integration testing")
	}
}
