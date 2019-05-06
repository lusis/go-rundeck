package rundeck_test

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/lusis/go-rundeck/pkg/rundeck"
	"github.com/stretchr/testify/suite"
)

type ProjectIntegrationTestSuite struct {
	suite.Suite
	TestClient      *rundeck.Client
	CreatedProjects []rundeck.Project
	sync.Mutex
}

func (s *ProjectIntegrationTestSuite) testCreateProject(slow bool) (rundeck.Project, rundeck.JobMetaData) {
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

func (s *ProjectIntegrationTestSuite) SetupSuite() {
	client := testNewTokenAuthClient()
	s.TestClient = client
	s.CreatedProjects = []rundeck.Project{}
}

func (s *ProjectIntegrationTestSuite) TearDownSuite() {
	for _, p := range s.CreatedProjects {
		e := s.TestClient.DeleteProject(p.Name)
		if e != nil {
			s.T().Logf("unable to clean up test project: %s", e.Error())
		}
	}
}

func (s *ProjectIntegrationTestSuite) TestIntegrationCreateProject() {
	name := testGenerateRandomName(s.T().Name())
	props := map[string]string{
		"project.description": name,
	}
	for k, v := range testDefaultProjectProperties {
		props[k] = v
	}

	project, err := s.TestClient.CreateProject(name, props)
	s.Require().NoError(err)
	s.CreatedProjects = append(s.CreatedProjects, *project)

	listProjects, err := s.TestClient.ListProjects()
	s.Require().NoError(err)
	var foundProject bool

	for _, p := range listProjects {
		if p.Name == project.Name {
			foundProject = true
		}
	}
	s.Require().True(foundProject)
}

func (s *ProjectIntegrationTestSuite) TestIntegrationGetProjectInfo() {
	project, _ := s.testCreateProject(false)
	getProject, err := s.TestClient.GetProjectInfo(project.Name)
	s.Require().NoError(err)
	for k, v := range project.Properties {
		s.Require().Equal(v, getProject.Properties[k])
	}
}

func (s *ProjectIntegrationTestSuite) TestIntegrationGetProjectResources() {
	project, _ := s.testCreateProject(true)
	resources, err := s.TestClient.ListResourcesForProject(project.Name)
	s.Require().NoError(err)
	s.Require().Len(*resources, 100)
}

func (s *ProjectIntegrationTestSuite) TestIntegrationGetProjectReadme() {
	project, _ := s.testCreateProject(true)
	var readme = `# project readme\nthis is the project readme`
	err := s.TestClient.PutProjectReadme(project.Name, strings.NewReader(readme))
	s.Require().NoError(err)
	defer s.TestClient.DeleteProjectReadme(project.Name) // nolint: errcheck
	get, err := s.TestClient.GetProjectReadme(project.Name)
	s.Require().NoError(err)
	s.Require().Equal(readme, get)

}

func (s *ProjectIntegrationTestSuite) TestIntegrationGetProjectMotd() {
	project, _ := s.testCreateProject(true)
	var motd = `# project motd\n*stuff is broken*`
	err := s.TestClient.PutProjectMotd(project.Name, strings.NewReader(motd))
	s.Require().NoError(err)
	defer s.TestClient.DeleteProjectMotd(project.Name) // nolint: errcheck
	get, err := s.TestClient.GetProjectMotd(project.Name)
	s.Require().NoError(err)
	s.Require().Equal(motd, get)

}

func (s *ProjectIntegrationTestSuite) TestIntegrationGetProjectResource() {
	project, _ := s.testCreateProject(true)
	resource, err := s.TestClient.GetResourceInfo(project.Name, "node-1-stub")
	s.Require().NoError(err)
	s.Require().NotEmpty(resource.FileCopier)
	s.Require().NotEmpty(resource.NodeExectutor)
	s.Require().NotEmpty(resource.NodeName)
	s.Require().NotEmpty(resource.HostName)
	s.Require().NotEmpty(resource.UserName)
	s.Require().Equal(resource.Tags, "stub")
}

func (s *ProjectIntegrationTestSuite) TestIntegrationGetProjectConfiguration() {
	project, _ := s.testCreateProject(true)
	pc, err := s.TestClient.GetProjectConfiguration(project.Name)
	s.Require().NoError(err)
	for k, v := range project.Properties {
		s.Require().Equal(v, pc[k])
	}
}

func (s *ProjectIntegrationTestSuite) TestIntegrationProjectImport() {
	project, _ := s.testCreateProject(true)
	tmpDir := os.TempDir()

	output := filepath.Join(tmpDir, project.Name+".zip")
	f, err := os.Create(output)
	defer os.Remove(output) // nolint: errcheck
	s.Require().NoError(err)
	opts := []rundeck.ProjectExportOption{
		rundeck.ProjectExportAll(true),
		rundeck.ProjectExportConfigs(true),
		rundeck.ProjectExportAcls(true),
		rundeck.ProjectExportJobs(true),
		rundeck.ProjectExportReadmes(true),
	}
	// Export created project
	err = s.TestClient.GetProjectArchiveExport(project.Name, f, opts...)
	s.Require().NoError(err)

	// import the file into a new project
	destProject, _ := s.testCreateProject(true)

	data, err := os.Open(output)
	defer data.Close() // nolint: errcheck
	s.Require().NoError(err)

	imported, err := s.TestClient.ProjectArchiveImport(destProject.Name, data, rundeck.ProjectImportConfigs(true), rundeck.ProjectImportJobUUIDs("remove"))
	s.Require().NoError(err)

	s.Require().Equal("successful", imported.ImportStatus)
	checkImport, err := s.TestClient.GetProjectInfo(destProject.Name)
	s.Require().NoError(err)

	for k, v := range project.Properties {
		s.Require().Equal(v, checkImport.Properties[k])
	}
}

func (s *ProjectIntegrationTestSuite) TestIntegrationProjectAsyncExportImport() {
	project, _ := s.testCreateProject(true)
	tmpDir := os.TempDir()

	output := filepath.Join(tmpDir, project.Name+"-async.zip")
	f, err := os.Create(output)
	defer os.Remove(output) // nolint: errcheck
	s.Require().NoError(err)

	opts := []rundeck.ProjectExportOption{
		rundeck.ProjectExportAll(true),
		rundeck.ProjectExportConfigs(true),
		rundeck.ProjectExportAcls(true),
		rundeck.ProjectExportJobs(true),
		rundeck.ProjectExportReadmes(true),
	}
	// Export created project
	token, err := s.TestClient.GetProjectArchiveExportAsync(project.Name, opts...)
	s.Require().NoError(err)
	s.Require().NotEmpty(token)

	// this is our doneFunc.
	// we poll with the token until it's done
	doneFunc := func() (bool, error) {
		time.Sleep(500 * time.Millisecond)
		info, infoErr := s.TestClient.GetProjectArchiveExportAsyncStatus(project.Name, token)
		if infoErr != nil {
			return false, infoErr
		}
		return info.Ready, nil
	}

	done, err := s.TestClient.WaitFor(doneFunc, 5*time.Second)
	s.Require().NoError(err)
	s.Require().True(done)

	err = s.TestClient.GetProjectArchiveExportAsyncDownload(project.Name, token, f)
	s.Require().NoError(err)

	// import the file into a new project
	destProject, _ := s.testCreateProject(true)

	data, err := os.Open(output)
	defer data.Close() // nolint: errcheck
	s.Require().NoError(err)

	imported, err := s.TestClient.ProjectArchiveImport(destProject.Name, data, rundeck.ProjectImportConfigs(true), rundeck.ProjectImportJobUUIDs("remove"))
	s.Require().NoError(err)

	s.Require().Equal("successful", imported.ImportStatus)
	checkImport, err := s.TestClient.GetProjectInfo(destProject.Name)
	s.Require().NoError(err)

	for k, v := range project.Properties {
		s.Require().Equal(v, checkImport.Properties[k])
	}
}

func (s *ProjectIntegrationTestSuite) TestIntegrationGetProjectArchiveExportImportNoConfigs() {
	project, _ := s.testCreateProject(true)
	tmpDir := os.TempDir()

	output := filepath.Join(tmpDir, project.Name+".zip")
	f, err := os.Create(output)
	defer os.Remove(output) // nolint: errcheck
	s.Require().NoError(err)

	opts := []rundeck.ProjectExportOption{
		rundeck.ProjectExportAll(true),
		rundeck.ProjectExportConfigs(true),
		rundeck.ProjectExportAcls(true),
		rundeck.ProjectExportJobs(true),
		rundeck.ProjectExportReadmes(true),
	}
	// Export created project
	err = s.TestClient.GetProjectArchiveExport(project.Name, f, opts...)
	s.Require().NoError(err)

	// import the file into a new project
	destProject, _ := s.testCreateProject(true)
	data, err := os.Open(output)
	defer data.Close() // nolint: errcheck
	s.Require().NoError(err)

	imported, err := s.TestClient.ProjectArchiveImport(destProject.Name, data, rundeck.ProjectImportJobUUIDs("remove"))
	s.Require().NoError(err)

	s.Require().Empty(imported.Errors)
	s.Require().Equal("successful", imported.ImportStatus)
	checkImport, err := s.TestClient.GetProjectInfo(destProject.Name)
	s.Require().NoError(err)
	s.Require().NotEqual(project.Description, checkImport.Description)

}

func (s *ProjectIntegrationTestSuite) TestIntegrationProjectConfigurationKeys() {
	project, _ := s.testCreateProject(true)

	_, err := s.TestClient.GetProjectInfo(project.Name)
	s.Require().NoError(err)

	nodeSource, err := s.TestClient.GetProjectConfigurationKey(project.Name, "resources.source.1.type")
	s.Require().NoError(err)
	s.Require().Equal("stub", nodeSource)
	err = s.TestClient.PutProjectConfigurationKey(project.Name, "foo", "bar")
	s.Require().NoError(err)
	checkKey, err := s.TestClient.GetProjectConfigurationKey(project.Name, "foo")
	s.Require().NoError(err)
	s.Require().Equal(checkKey, "bar")
	replaceConf, err := s.TestClient.PutProjectConfiguration(project.Name, project.Properties)
	s.Require().NoError(err)
	s.Require().NotContains(replaceConf, "foo")
	err = s.TestClient.DeleteProjectConfigurationKey(project.Name, "project.description")
	s.Require().NoError(err)
	_, err = s.TestClient.GetProjectConfigurationKey(project.Name, "project.description")
	// key should be missing
	s.Require().Error(err)
}

func TestIntegrationProjectSuite(t *testing.T) {
	if testing.Short() || testRundeckRunning() == false {
		t.Skip("skipping integration testing")
	}

	suite.Run(t, new(ProjectIntegrationTestSuite))
}
