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

	project, createErr := s.TestClient.CreateProject(name, props)
	if createErr != nil {
		s.T().Fatalf("Unable to create test project: %s", createErr.Error())
	}
	s.CreatedProjects = append(s.CreatedProjects, *project)

	listProjects, listErr := s.TestClient.ListProjects()
	s.NoError(listErr)
	var foundProject bool

	for _, p := range listProjects {
		if p.Name == project.Name {
			foundProject = true
		}
	}
	s.True(foundProject)
}

func (s *ProjectIntegrationTestSuite) TestIntegrationGetProjectInfo() {
	project, _ := s.testCreateProject(false)
	getProject, getErr := s.TestClient.GetProjectInfo(project.Name)
	s.NoError(getErr)
	for k, v := range project.Properties {
		s.Equal(v, getProject.Properties[k])
	}
}

func (s *ProjectIntegrationTestSuite) TestIntegrationGetProjectResources() {
	project, _ := s.testCreateProject(true)
	resources, getErr := s.TestClient.ListResourcesForProject(project.Name)
	s.NoError(getErr)
	s.Len(*resources, 100)
}

func (s *ProjectIntegrationTestSuite) TestIntegrationGetProjectReadme() {
	project, _ := s.testCreateProject(true)
	var readme = `# project readme\nthis is the project readme`
	putErr := s.TestClient.PutProjectReadme(project.Name, strings.NewReader(readme))
	if putErr != nil {
		s.Fail("cannot upload readme. cannot continue. %s", putErr)
	}
	defer s.TestClient.DeleteProjectReadme(project.Name) // nolint: errcheck
	get, getErr := s.TestClient.GetProjectReadme(project.Name)
	s.NoError(getErr)
	s.Equal(readme, get)

}

func (s *ProjectIntegrationTestSuite) TestIntegrationGetProjectMotd() {
	project, _ := s.testCreateProject(true)
	var motd = `# project motd\n*stuff is broken*`
	putErr := s.TestClient.PutProjectMotd(project.Name, strings.NewReader(motd))
	if putErr != nil {
		s.Fail("cannot upload motd. cannot continue. %s", putErr)
	}
	defer s.TestClient.DeleteProjectMotd(project.Name) // nolint: errcheck
	get, getErr := s.TestClient.GetProjectMotd(project.Name)
	s.NoError(getErr)
	s.Equal(motd, get)

}

func (s *ProjectIntegrationTestSuite) TestIntegrationGetProjectResource() {
	project, _ := s.testCreateProject(true)
	resource, err := s.TestClient.GetResourceInfo(project.Name, "node-1-stub")
	s.NoError(err)
	s.NotEmpty(resource.FileCopier)
	s.NotEmpty(resource.NodeExectutor)
	s.NotEmpty(resource.NodeName)
	s.NotEmpty(resource.HostName)
	s.NotEmpty(resource.UserName)
	s.Equal(resource.Tags, "stub")
}

func (s *ProjectIntegrationTestSuite) TestIntegrationGetProjectConfiguration() {
	project, _ := s.testCreateProject(true)
	pc, pcerr := s.TestClient.GetProjectConfiguration(project.Name)
	s.NoError(pcerr)
	for k, v := range project.Properties {
		s.Equal(v, pc[k])
	}
}

func (s *ProjectIntegrationTestSuite) TestIntegrationProjectImport() {
	project, _ := s.testCreateProject(true)
	tmpDir := os.TempDir()

	output := filepath.Join(tmpDir, project.Name+".zip")
	f, fErr := os.Create(output)
	defer os.Remove(output) // nolint: errcheck
	if fErr != nil {
		s.T().Fatalf("unable to create output file. cannot continue: %s", fErr)
	}
	opts := []rundeck.ProjectExportOption{
		rundeck.ProjectExportAll(true),
		rundeck.ProjectExportConfigs(true),
		rundeck.ProjectExportAcls(true),
		rundeck.ProjectExportJobs(true),
		rundeck.ProjectExportReadmes(true),
	}
	// Export created project
	perr := s.TestClient.GetProjectArchiveExport(project.Name, f, opts...)
	if perr != nil {
		s.T().Fatalf("cannot export project: %s", perr.Error())
	}

	// import the file into a new project
	destProject, _ := s.testCreateProject(true)

	data, dataErr := os.Open(output)
	defer data.Close() // nolint: errcheck
	if dataErr != nil {
		s.T().Fatalf("cannot open import file: %s", dataErr.Error())
	}
	imported, impErr := s.TestClient.ProjectArchiveImport(destProject.Name, data, rundeck.ProjectImportConfigs(true), rundeck.ProjectImportJobUUIDs("remove"))
	if impErr != nil {
		s.T().Fatalf("could not import project. cannot continue: %s", impErr.Error())
	}
	s.Equal("successful", imported.ImportStatus)
	checkImport, checkErr := s.TestClient.GetProjectInfo(destProject.Name)
	if checkErr != nil {
		s.T().Fatalf("could not get the newly imported project. cannot continue: %s", checkErr.Error())
	}
	for k, v := range project.Properties {
		s.Equal(v, checkImport.Properties[k])
	}
}

func (s *ProjectIntegrationTestSuite) TestIntegrationProjectAsyncExportImport() {
	project, _ := s.testCreateProject(true)
	tmpDir := os.TempDir()

	output := filepath.Join(tmpDir, project.Name+"-async.zip")
	f, fErr := os.Create(output)
	defer os.Remove(output) // nolint: errcheck
	if fErr != nil {
		s.T().Fatalf("unable to create output file. cannot continue: %s", fErr)
	}
	opts := []rundeck.ProjectExportOption{
		rundeck.ProjectExportAll(true),
		rundeck.ProjectExportConfigs(true),
		rundeck.ProjectExportAcls(true),
		rundeck.ProjectExportJobs(true),
		rundeck.ProjectExportReadmes(true),
	}
	// Export created project
	token, perr := s.TestClient.GetProjectArchiveExportAsync(project.Name, opts...)
	if perr != nil {
		s.T().Fatalf("cannot export project: %s", perr.Error())
	}

	s.NotEmpty(token)

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

	done, doneErr := s.TestClient.WaitFor(doneFunc, 5*time.Second)
	s.NoError(doneErr)
	s.True(done)

	downloadErr := s.TestClient.GetProjectArchiveExportAsyncDownload(project.Name, token, f)
	if downloadErr != nil {
		s.T().Fatalf("could not download project archive. cannot continue: %s", downloadErr.Error())
	}
	// import the file into a new project
	destProject, _ := s.testCreateProject(true)

	data, dataErr := os.Open(output)
	defer data.Close() // nolint: errcheck
	if dataErr != nil {
		s.T().Fatalf("cannot open import file: %s", dataErr.Error())
	}
	imported, impErr := s.TestClient.ProjectArchiveImport(destProject.Name, data, rundeck.ProjectImportConfigs(true), rundeck.ProjectImportJobUUIDs("remove"))
	if impErr != nil {
		s.T().Fatalf("could not import project. cannot continue: %s", impErr.Error())
	}
	s.Equal("successful", imported.ImportStatus)
	checkImport, checkErr := s.TestClient.GetProjectInfo(destProject.Name)
	if checkErr != nil {
		s.T().Fatalf("could not get the newly imported project. cannot continue: %s", checkErr.Error())
	}
	for k, v := range project.Properties {
		s.Equal(v, checkImport.Properties[k])
	}
}

func (s *ProjectIntegrationTestSuite) TestIntegrationGetProjectArchiveExportImportNoConfigs() {
	project, _ := s.testCreateProject(true)
	tmpDir := os.TempDir()

	output := filepath.Join(tmpDir, project.Name+".zip")
	f, fErr := os.Create(output)
	defer os.Remove(output) // nolint: errcheck
	if fErr != nil {
		s.T().Fatalf("unable to create output file. cannot continue: %s", fErr)
	}
	opts := []rundeck.ProjectExportOption{
		rundeck.ProjectExportAll(true),
		rundeck.ProjectExportConfigs(true),
		rundeck.ProjectExportAcls(true),
		rundeck.ProjectExportJobs(true),
		rundeck.ProjectExportReadmes(true),
	}
	// Export created project
	perr := s.TestClient.GetProjectArchiveExport(project.Name, f, opts...)
	if perr != nil {
		s.T().Fatalf("cannot export project: %s", perr.Error())
	}

	// import the file into a new project
	destProject, _ := s.testCreateProject(true)
	data, dataErr := os.Open(output)
	defer data.Close() // nolint: errcheck
	if dataErr != nil {
		s.T().Fatalf(dataErr.Error())
	}
	imported, impErr := s.TestClient.ProjectArchiveImport(destProject.Name, data, rundeck.ProjectImportJobUUIDs("remove"))
	if impErr != nil {
		s.T().Fatalf("could not import project. cannot continue: %s", impErr.Error())
	}
	s.Empty(imported.Errors)
	s.Equal("successful", imported.ImportStatus)
	checkImport, checkErr := s.TestClient.GetProjectInfo(destProject.Name)
	if checkErr != nil {
		s.T().Fatalf("could not get the newly imported project. cannot continue: %s", checkErr.Error())
	}
	s.NotEqual(project.Description, checkImport.Description)

}

func (s *ProjectIntegrationTestSuite) TestIntegrationProjectConfigurationKeys() {
	project, _ := s.testCreateProject(true)

	_, checkErr := s.TestClient.GetProjectInfo(project.Name)
	if checkErr != nil {
		s.T().Fatalf("could not get the newly imported project. cannot continue: %s", checkErr.Error())
	}
	nodeSource, nodeSourceErr := s.TestClient.GetProjectConfigurationKey(project.Name, "resources.source.1.type")
	s.NoError(nodeSourceErr)
	s.Equal("stub", nodeSource)
	putKeyErr := s.TestClient.PutProjectConfigurationKey(project.Name, "foo", "bar")
	s.NoError(putKeyErr)
	checkKey, checkKeyErr := s.TestClient.GetProjectConfigurationKey(project.Name, "foo")
	s.NoError(checkKeyErr)
	s.Equal(checkKey, "bar")
	replaceConf, replaceConfErr := s.TestClient.PutProjectConfiguration(project.Name, project.Properties)
	s.NoError(replaceConfErr)
	s.NotContains(replaceConf, "foo")
	deleteErr := s.TestClient.DeleteProjectConfigurationKey(project.Name, "project.description")
	s.NoError(deleteErr)
	_, checkAgainKeyErr := s.TestClient.GetProjectConfigurationKey(project.Name, "project.description")
	// key should be missing
	s.Error(checkAgainKeyErr)
}

func TestIntegrationProjectSuite(t *testing.T) {
	if testRundeckRunning() {
		suite.Run(t, new(ProjectIntegrationTestSuite))
	} else {
		t.Skip("rundeck isn't running for integration testing")
	}
}
