package rundeck_test

import (
	"bytes"
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/lusis/go-rundeck/pkg/rundeck"
	"github.com/stretchr/testify/suite"
)

type SCMIntegrationTestSuite struct {
	suite.Suite
	CreatedProjects []rundeck.Project
	TestClient      *rundeck.Client
	sync.Mutex
}

func (s *SCMIntegrationTestSuite) testCreateProject(slow bool) (rundeck.Project, rundeck.JobMetaData) {
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

func (s *SCMIntegrationTestSuite) SetupSuite() {
	client := testNewTokenAuthClient()
	s.TestClient = client
	s.CreatedProjects = []rundeck.Project{}
}

func (s *SCMIntegrationTestSuite) TearDownSuite() {
	for _, p := range s.CreatedProjects {
		e := s.TestClient.DeleteProject(p.Name)
		if e != nil {
			s.T().Logf("unable to clean up test project: %s", e.Error())
		}
	}
}

func TestIntegrationSCMSuite(t *testing.T) {
	if testing.Short() || testRundeckRunning() == false {
		t.Skip("skipping integration testing")
	}

	suite.Run(t, &SCMIntegrationTestSuite{})
}

func (s *SCMIntegrationTestSuite) TestListSCMPlugins() {
	project, _ := s.testCreateProject(false)
	res, err := s.TestClient.ListSCMPlugins(project.Name)
	s.Require().NoError(err)
	s.Require().Len(res.Import, 1)
	s.Require().Len(res.Export, 1)
	s.Require().Equal("git-export", res.Export[0].Type)
	s.Require().Equal("git-import", res.Import[0].Type)
}

func (s *SCMIntegrationTestSuite) TestGetSCMPluginInputFields() {
	project, _ := s.testCreateProject(false)
	res, err := s.TestClient.ListSCMPlugins(project.Name)
	s.Require().NoError(err)

	// We're going to spot check a few fields
	for _, p := range res.Export {
		fields, fieldsErr := s.TestClient.GetProjectSCMPluginInputFields(project.Name, "export", p.Type)
		s.Require().NoError(fieldsErr)
		s.Require().NotEmpty(fields.Integration)
		s.Require().NotEmpty(fields.Type)
		for _, f := range fields.Fields {
			s.Require().NotEmpty(f.Name)
			s.Require().NotEmpty(f.Description)
		}
	}
	for _, p := range res.Import {
		fields, fieldsErr := s.TestClient.GetProjectSCMPluginInputFields(project.Name, "import", p.Type)
		s.Require().NoError(fieldsErr)
		s.Require().NotEmpty(fields.Integration)
		s.Require().NotEmpty(fields.Type)
		for _, f := range fields.Fields {
			s.Require().NotEmpty(f.Name)
			s.Require().NotEmpty(f.Description)
		}
	}
}

// For now we name these with a
func (s *SCMIntegrationTestSuite) TestSetupSCMPluginExport() {
	project, _ := s.testCreateProject(false)
	params := map[string]string{
		"committerName":         "John E. Vincent",
		"committerEmail":        "lusis.org+github.com@gmail.com",
		"url":                   "/home/rundeck-export.git/",
		"format":                "yaml",
		"dir":                   fmt.Sprintf("/var/rundeck/projects/%s/scm", project.Name),
		"pathTemplate":          "${job.group}${job.name}-${job.id}.${config.format}",
		"branch":                "master",
		"strictHostKeyChecking": "no",
	}
	res, err := s.TestClient.SetupSCMPluginForProject(project.Name, "export", "git-export", params)
	s.Require().NoError(err)
	s.Require().True(res.Success)
	s.Require().Nil(res.ValidationErrors)
}

func (s *SCMIntegrationTestSuite) TestSetupSCMPluginImport() {
	project, _ := s.testCreateProject(false)
	params := map[string]string{
		"url":                   "/home/rundeck-import.git/",
		"format":                "yaml",
		"dir":                   fmt.Sprintf("/var/rundeck/projects/%s/scm", project.Name),
		"pathTemplate":          "${job.group}${job.name}-${job.id}.${config.format}",
		"branch":                "master",
		"strictHostKeyChecking": "no",
		"useFilePattern":        "true",
		"filePattern":           ".*yaml",
	}
	res, err := s.TestClient.SetupSCMPluginForProject(project.Name, "import", "git-import", params)
	s.Require().NoError(err)
	s.Require().True(res.Success)
	s.Require().Nil(res.ValidationErrors)
}

// TestPerformProjectSCMActionExport
// this is a multi-step test that does the following:
// - import a job definition to our test project
// - waits until project scm status has pending work
// - gets the pending action
// - performs the pending action
func (s *SCMIntegrationTestSuite) TestSCMActionProjectExport() {
	project, _ := s.testCreateProject(false)
	scmparams := map[string]string{
		"committerName":         "John E. Vincent",
		"committerEmail":        "lusis.org+github.com@gmail.com",
		"url":                   "/home/rundeck-export.git/",
		"format":                "yaml",
		"dir":                   fmt.Sprintf("/var/rundeck/projects/%s/scm", project.Name),
		"pathTemplate":          "${job.group}${job.name}-${job.id}.${config.format}",
		"branch":                "master",
		"strictHostKeyChecking": "no",
	}
	_, err := s.TestClient.SetupSCMPluginForProject(project.Name, "export", "git-export", scmparams)
	s.Require().NoError(err)
	params := map[string]string{
		"message": fmt.Sprintf("commit for integration test: %s", project.Name),
		"push":    "true",
	}

	plugins, err := s.TestClient.ListSCMPlugins(project.Name)
	s.Require().NoError(err)
	s.Require().True(plugins.Export[0].Enabled)


	jobbytes, err := testJobFromTemplate(project.Name+"-job-2", "job for "+project.Name)
	s.Require().NoError(err)
	_, err = s.TestClient.ImportJob(project.Name,
		bytes.NewReader(jobbytes),
		rundeck.ImportFormat("yaml"),
		rundeck.ImportUUID("remove"))
	s.Require().NoError(err)
	doneFunc := func() (bool, error) {
		time.Sleep(1 * time.Second)
		info, infoErr := s.TestClient.GetProjectSCMStatus(project.Name, "export")
		if infoErr != nil && infoErr.Error() == "Rundeck could not find the resource you requested" {
			// just not ready yet
			return false, nil
		}
		if infoErr != nil {
			return false, infoErr
		}
		if info != nil {
			if len(info.Actions) != 0 {
				// we have pending actions, let's return
				return true, nil
			}
		}
		return false, nil
	}
	_, err = s.TestClient.WaitFor(doneFunc, 15*time.Second)
	s.Require().NoError(err)

	aif, err := s.TestClient.GetProjectSCMActionInputFields(project.Name, "export", "project-commit")
	s.Require().NoError(err)

	exportItems := aif.ExportItems[0]

	action, err := s.TestClient.PerformProjectSCMAction(project.Name, aif.Integration, aif.ActionID,
		rundeck.SCMActionInput(params),
		rundeck.SCMActionJobs(exportItems.Job.JobID))
	s.Require().NoError(err)
	s.Require().True(action.Success)
	s.Require().Len(action.ValidationErrors, 0)
}

func (s *SCMIntegrationTestSuite) TestSCMActionProjectFailure() {
	project, _ := s.testCreateProject(false)
	scmparams := map[string]string{
		"committerName":  "John E. Vincent",
		"committerEmail": "lusis.org+github.com@gmail.com",
	}
	res, err := s.TestClient.SetupSCMPluginForProject(project.Name, "export", "git-export", scmparams)
	s.Require().Nil(res)
	s.Require().Error(err)
}

func (s *SCMIntegrationTestSuite) TestSCMActionProjectDisableEnable() {
	project, _ := s.testCreateProject(false)
	scmparams := map[string]string{
		"committerName":         "John E. Vincent",
		"committerEmail":        "lusis.org+github.com@gmail.com",
		"url":                   "/home/rundeck-export.git/",
		"format":                "yaml",
		"dir":                   fmt.Sprintf("/var/rundeck/projects/%s/scm", project.Name),
		"pathTemplate":          "${job.group}${job.name}-${job.id}.${config.format}",
		"branch":                "master",
		"strictHostKeyChecking": "no",
	}
	_, err := s.TestClient.SetupSCMPluginForProject(project.Name, "export", "git-export", scmparams)
	s.Require().NoError(err)
	err = s.TestClient.DisableSCMPluginForProject(project.Name, "export", "git-export")
	s.Require().NoError(err)
	err = s.TestClient.EnableSCMPluginForProject(project.Name, "export", "git-export")
	s.Require().NoError(err)
}

func (s *SCMIntegrationTestSuite) TestSCMActionJobDiff() {
	project, _ := s.testCreateProject(false)
	scmparams := map[string]string{
		"committerName":         "John E. Vincent",
		"committerEmail":        "lusis.org+github.com@gmail.com",
		"url":                   "/home/rundeck-export.git/",
		"format":                "yaml",
		"dir":                   fmt.Sprintf("/var/rundeck/projects/%s/scm", project.Name),
		"pathTemplate":          "${job.group}${job.name}-${job.id}.${config.format}",
		"branch":                "master",
		"strictHostKeyChecking": "no",
	}
	_, err := s.TestClient.SetupSCMPluginForProject(project.Name, "export", "git-export", scmparams)
	s.Require().NoError(err)

	plugins, err := s.TestClient.ListSCMPlugins(project.Name)
	s.Require().NoError(err)
	s.Require().True(plugins.Export[0].Enabled)

	jobbytes, err := testJobFromTemplate(project.Name+"-job-2", "job for "+project.Name)
	s.Require().NoError(err)
	importedJob, err := s.TestClient.ImportJob(project.Name,
		bytes.NewReader(jobbytes),
		rundeck.ImportFormat("yaml"),
		rundeck.ImportUUID("remove"))
	s.Require().NoError(err)
	doneFunc := func() (bool, error) {
		time.Sleep(1 * time.Second)
		info, infoErr := s.TestClient.GetProjectSCMStatus(project.Name, "export")
		if infoErr != nil && infoErr.Error() == "Rundeck could not find the resource you requested" {
			// just not ready yet
			return false, nil
		}
		if infoErr != nil {
			return false, infoErr
		}
		if info != nil {
			if len(info.Actions) != 0 {
				// we have pending actions, let's return
				return true, nil
			}
		}
		return false, nil
	}
	_, err = s.TestClient.WaitFor(doneFunc, 15*time.Second)
	s.Require().NoError(err)

	jobid := importedJob.Succeeded[0].ID
	diff, err := s.TestClient.GetJobSCMDiff(jobid, "export")
	s.Require().NoError(err)
	s.Require().NotEmpty(diff.ID)
	s.Require().NotEmpty(diff.Project)
}
