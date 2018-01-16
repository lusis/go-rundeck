package rundeck_test

import (
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/lusis/go-rundeck/pkg/rundeck"
	"github.com/stretchr/testify/suite"
)

type SCMIntegrationTestSuite struct {
	suite.Suite
	TestProjectName string
	TestProject     *rundeck.Project
	TestClient      *rundeck.Client
}

func (s *SCMIntegrationTestSuite) SetupSuite() {
	client := testNewTokenAuthClient()
	s.TestProjectName = testGenerateRandomName("scmproject")
	s.TestClient = client
	props := map[string]string{
		"project.description": s.TestProjectName,
	}
	for k, v := range testDefaultProjectProperties {
		props[k] = v
	}
	project, createErr := s.TestClient.CreateProject(s.TestProjectName, props)
	if createErr != nil {
		s.T().Fatalf("Unable to create test project: %s", createErr.Error())
	}
	s.TestProject = project
}

func (s *SCMIntegrationTestSuite) TearDownSuite() {
	e := s.TestClient.DeleteProject(s.TestProject.Name)
	if e != nil {
		s.T().Errorf("unable to clean up test project: %s", e.Error())
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
	res, resErr := s.TestClient.ListSCMPlugins(s.TestProjectName)
	if resErr != nil {
		s.T().Fatalf("cannot list plugins. cannot continue: %s", resErr.Error())
	}
	s.Len(res.Import, 1)
	s.Len(res.Export, 1)
	s.Equal("git-export", res.Export[0].Type)
	s.Equal("git-import", res.Import[0].Type)
}

func (s *SCMIntegrationTestSuite) TestGetSCMPluginInputFields() {
	res, resErr := s.TestClient.ListSCMPlugins(s.TestProjectName)
	if resErr != nil {
		s.T().Fatalf("cannot list plugins. cannot continue: %s", resErr.Error())
	}
	// We're going to spot check a few fields
	for _, p := range res.Export {
		fields, fieldsErr := s.TestClient.GetSCMPluginInputFields(s.TestProjectName, "export", p.Type)
		s.NoError(fieldsErr)
		s.NotEmpty(fields.Integration)
		s.NotEmpty(fields.Type)
		for _, f := range fields.Fields {
			s.NotEmpty(f.Name)
			s.NotEmpty(f.Description)
		}
	}
	for _, p := range res.Import {
		fields, fieldsErr := s.TestClient.GetSCMPluginInputFields(s.TestProjectName, "import", p.Type)
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
func (s *SCMIntegrationTestSuite) Test1SetupSCMPluginExport() {
	params := map[string]string{
		"committerName":         "John E. Vincent",
		"committerEmail":        "lusis.org+github.com@gmail.com",
		"url":                   "/home/rundeck-export.git/",
		"format":                "yaml",
		"dir":                   fmt.Sprintf("/var/rundeck/projects/%s/scm", s.TestProjectName),
		"pathTemplate":          "${job.group}${job.name}-${job.id}.${config.format}",
		"branch":                "master",
		"strictHostKeyChecking": "no",
	}
	res, resErr := s.TestClient.SetupSCMPluginForProject(s.TestProjectName, "export", "git-export", params)
	if resErr != nil {
		s.T().Fatalf("could not setup export plugin. cannot continue: %s", resErr)
	}
	s.True(res.Success)
	s.Nil(res.ValidationErrors)
}

func (s *SCMIntegrationTestSuite) Test1SetupSCMPluginImport() {
	params := map[string]string{
		"url":                   "/home/rundeck-import.git/",
		"format":                "yaml",
		"dir":                   fmt.Sprintf("/var/rundeck/projects/%s/scm", s.TestProjectName),
		"pathTemplate":          "${job.group}${job.name}-${job.id}.${config.format}",
		"branch":                "master",
		"strictHostKeyChecking": "no",
	}
	res, resErr := s.TestClient.SetupSCMPluginForProject(s.TestProjectName, "import", "git-import", params)
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
	params := map[string]string{
		"message": fmt.Sprintf("commit for integration test: %s", s.TestProjectName),
		"push":    "true",
	}
	importJob, importErr := s.TestClient.ImportJob(s.TestProject.Name,
		strings.NewReader(testJobDefinition),
		rundeck.ImportFormat("yaml"),
		rundeck.ImportUUID("remove"))
	if importErr != nil {
		s.T().Fatalf("job did not import. cannot continue: %s", importErr.Error())
	}

	plugins, pluginsErr := s.TestClient.ListSCMPlugins(s.TestProjectName)
	if pluginsErr != nil {
		s.T().Fatalf("cannot list plugins. cannot continue: %s", pluginsErr.Error())
	}
	if !plugins.Export[0].Enabled {
		s.T().Fatalf("export plugin is not enabled. cannot continue: %s", plugins.Export[0].Description)
	}
	_, jobsErr := s.TestClient.GetJobInfo(importJob.Succeeded[0].ID)
	if jobsErr != nil {
		s.T().Fatalf("could not get job info. cannot continue: %s", jobsErr.Error())
	}
	doneFunc := func() (bool, error) {
		time.Sleep(1 * time.Second)
		info, infoErr := s.TestClient.GetProjectSCMStatus(s.TestProjectName, "export")
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
	_, doneErr := s.TestClient.WaitFor(doneFunc, 5*time.Second)
	if doneErr != nil {
		s.T().Fatalf("never saw a status message. cannot continue: %s", doneErr)
	}
	aif, aifErr := s.TestClient.GetProjectSCMActionInputFields(s.TestProjectName, "export", "project-commit")
	if aifErr != nil {
		s.T().Fatalf("could not get action input fields. cannot continue: %s", aifErr.Error())
	}

	exportItems := aif.ExportItems[0]

	res, resErr := s.TestClient.PerformProjectSCMAction(s.TestProjectName, aif.Integration, aif.ActionID,
		rundeck.SCMActionInput(params),
		rundeck.SCMActionJobs(exportItems.Job.JobID))
	if resErr != nil {
		s.T().Fatalf("could not perform action. cannot continue: %s", resErr)
	}
	s.True(res.Success)
	s.Len(res.ValidationErrors, 0)
}
