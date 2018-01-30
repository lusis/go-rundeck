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
	if testRundeckRunning() {
		suite.Run(t, &SCMIntegrationTestSuite{})
	} else {
		t.Skip("rundeck isn't running for integration testing")
	}
}

func (s *SCMIntegrationTestSuite) TestListSCMPlugins() {
	project, _ := s.testCreateProject(false)
	res, resErr := s.TestClient.ListSCMPlugins(project.Name)
	if resErr != nil {
		s.T().Fatalf("cannot list plugins. cannot continue: %s", resErr.Error())
	}
	s.Len(res.Import, 1)
	s.Len(res.Export, 1)
	s.Equal("git-export", res.Export[0].Type)
	s.Equal("git-import", res.Import[0].Type)
}

func (s *SCMIntegrationTestSuite) TestGetSCMPluginInputFields() {
	project, _ := s.testCreateProject(false)
	res, resErr := s.TestClient.ListSCMPlugins(project.Name)
	if resErr != nil {
		s.T().Fatalf("cannot list plugins. cannot continue: %s", resErr.Error())
	}
	// We're going to spot check a few fields
	for _, p := range res.Export {
		fields, fieldsErr := s.TestClient.GetSCMPluginInputFields(project.Name, "export", p.Type)
		s.NoError(fieldsErr)
		s.NotEmpty(fields.Integration)
		s.NotEmpty(fields.Type)
		for _, f := range fields.Fields {
			s.NotEmpty(f.Name)
			s.NotEmpty(f.Description)
		}
	}
	for _, p := range res.Import {
		fields, fieldsErr := s.TestClient.GetSCMPluginInputFields(project.Name, "import", p.Type)
		s.NoError(fieldsErr)
		s.NotEmpty(fields.Integration)
		s.NotEmpty(fields.Type)
		for _, f := range fields.Fields {
			s.NotEmpty(f.Name)
			s.NotEmpty(f.Description)
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
	res, resErr := s.TestClient.SetupSCMPluginForProject(project.Name, "export", "git-export", params)
	if resErr != nil {
		s.T().Fatalf("could not setup export plugin. cannot continue: %s", resErr)
	}
	s.True(res.Success)
	s.Nil(res.ValidationErrors)
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
	}
	res, resErr := s.TestClient.SetupSCMPluginForProject(project.Name, "import", "git-import", params)
	if resErr != nil {
		s.T().Fatalf("could not setup import plugin. cannot continue: %s", resErr)
	}
	s.True(res.Success)
	s.Nil(res.ValidationErrors)
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
	_, resErr := s.TestClient.SetupSCMPluginForProject(project.Name, "export", "git-export", scmparams)
	if resErr != nil {
		s.T().Fatalf("could not setup export plugin. cannot continue: %s", resErr)
	}
	params := map[string]string{
		"message": fmt.Sprintf("commit for integration test: %s", project.Name),
		"push":    "true",
	}

	plugins, pluginsErr := s.TestClient.ListSCMPlugins(project.Name)
	if pluginsErr != nil {
		s.T().Fatalf("cannot list plugins. cannot continue: %s", pluginsErr.Error())
	}
	if !plugins.Export[0].Enabled {
		s.T().Fatalf("export plugin is not enabled. cannot continue: %s", plugins.Export[0].Description)
	}

	jobbytes, joberr := testJobFromTemplate(project.Name+"-job-2", "job for "+project.Name)
	if joberr != nil {
		s.T().Fatalf("cannot create a job import from template. cannot continue: %s", joberr.Error())
	}
	_, importErr := s.TestClient.ImportJob(project.Name,
		bytes.NewReader(jobbytes),
		rundeck.ImportFormat("yaml"),
		rundeck.ImportUUID("remove"))
	if importErr != nil {
		s.T().Fatalf("job did not import. cannot continue: %s", importErr.Error())
	}
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
	_, doneErr := s.TestClient.WaitFor(doneFunc, 15*time.Second)
	if doneErr != nil {
		s.T().Fatalf("never saw a status message. cannot continue: %s", doneErr)
	}
	aif, aifErr := s.TestClient.GetProjectSCMActionInputFields(project.Name, "export", "project-commit")
	if aifErr != nil {
		s.T().Fatalf("could not get action input fields. cannot continue: %s", aifErr.Error())
	}

	exportItems := aif.ExportItems[0]

	action, actionErr := s.TestClient.PerformProjectSCMAction(project.Name, aif.Integration, aif.ActionID,
		rundeck.SCMActionInput(params),
		rundeck.SCMActionJobs(exportItems.Job.JobID))
	if actionErr != nil {
		s.T().Fatalf("could not perform action. cannot continue: %s", resErr)
	}
	s.True(action.Success)
	s.Len(action.ValidationErrors, 0)
}
