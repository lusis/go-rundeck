package rundeck_test

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/lusis/go-rundeck/pkg/rundeck"
	"github.com/stretchr/testify/suite"
)

type ProjectIntegrationTestSuite struct {
	suite.Suite
	TestProjectName string
	TestProject     *rundeck.Project
	TestClient      *rundeck.Client
}

func (s *ProjectIntegrationTestSuite) SetupSuite() {
	client := testNewTokenAuthClient()
	s.TestProjectName = testGenerateRandomName("testproject")
	s.TestClient = client
}

func (s *ProjectIntegrationTestSuite) TearDownSuite() {
	e := s.TestClient.DeleteProject(s.TestProject.Name)
	if e != nil {
		s.T().Errorf("unable to clean up test project: %s", e.Error())
	}
}

func (s *ProjectIntegrationTestSuite) TestIntegrationCreateProject() {
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

	listProjects, listErr := s.TestClient.ListProjects()
	s.NoError(listErr)
	var foundProject bool

	for _, p := range listProjects {
		if p.Name == s.TestProjectName {
			foundProject = true
		}
	}
	s.True(foundProject)
}

func (s *ProjectIntegrationTestSuite) TestIntegrationGetProjectInfo() {
	getProject, getErr := s.TestClient.GetProjectInfo(s.TestProjectName)
	s.NoError(getErr)
	for k, v := range s.TestProject.Properties {
		s.Equal(v, getProject.Properties[k])
	}
}

func (s *ProjectIntegrationTestSuite) TestIntegrationGetProjectResources() {
	resources, getErr := s.TestClient.ListResourcesForProject(s.TestProjectName)
	s.NoError(getErr)
	s.Len(*resources, 5)
}

func (s *ProjectIntegrationTestSuite) TestIntegrationGetProjectReadme() {
	var readme = `# project readme\nthis is the project readme`
	putErr := s.TestClient.PutProjectReadme(s.TestProjectName, strings.NewReader(readme))
	if putErr != nil {
		s.Fail("cannot upload readme. cannot continue. %s", putErr)
	}
	defer s.TestClient.DeleteProjectReadme(s.TestProjectName) // nolint: errcheck
	get, getErr := s.TestClient.GetProjectReadme(s.TestProjectName)
	s.NoError(getErr)
	s.Equal(readme, get)

}

func (s *ProjectIntegrationTestSuite) TestIntegrationGetProjectMotd() {
	var motd = `# project motd\n*stuff is broken*`
	putErr := s.TestClient.PutProjectMotd(s.TestProjectName, strings.NewReader(motd))
	if putErr != nil {
		s.Fail("cannot upload motd. cannot continue. %s", putErr)
	}
	defer s.TestClient.DeleteProjectMotd(s.TestProjectName) // nolint: errcheck
	get, getErr := s.TestClient.GetProjectMotd(s.TestProjectName)
	s.NoError(getErr)
	s.Equal(motd, get)

}

func (s *ProjectIntegrationTestSuite) TestIntegrationGetProjectResource() {
	resource, err := s.TestClient.GetResourceInfo(s.TestProjectName, "node-1-stub")
	s.NoError(err)
	s.NotEmpty(resource.FileCopier)
	s.NotEmpty(resource.NodeExectutor)
	s.NotEmpty(resource.NodeName)
	s.NotEmpty(resource.HostName)
	s.NotEmpty(resource.UserName)
	s.Equal(resource.Tags, "stub")
}

func (s *ProjectIntegrationTestSuite) TestIntegrationGetProjectConfiguration() {
	pc, pcerr := s.TestClient.GetProjectConfiguration(s.TestProjectName)
	s.NoError(pcerr)
	for k, v := range s.TestProject.Properties {
		s.Equal(v, pc[k])
	}
}

func (s *ProjectIntegrationTestSuite) TestIntegrationProjectImport() {
	tmpDir := os.TempDir()

	output := filepath.Join(tmpDir, s.TestProjectName+".zip")
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
	perr := s.TestClient.GetProjectArchiveExport(s.TestProjectName, f, opts...)
	if perr != nil {
		s.T().Fatalf("cannot export project: %s", perr.Error())
	}

	// import the file into a new project
	destProjectName := testGenerateRandomName("destproject")
	_, destErr := s.TestClient.CreateProject(destProjectName, nil)
	defer func() {
		e := s.TestClient.DeleteProject(destProjectName)
		if e != nil {
			s.T().Errorf("unable to clean up after myself: %s", e.Error())
		}
	}()
	if destErr != nil {
		s.T().Fatalf("unable to create a test project. cannot continue: %s", destErr.Error())
	}
	data, dataErr := os.Open(output)
	defer data.Close() // nolint: errcheck
	if dataErr != nil {
		s.T().Fatalf("cannot open import file: %s", dataErr.Error())
	}
	imported, impErr := s.TestClient.ProjectArchiveImport(destProjectName, data, rundeck.ProjectImportConfigs(true))
	if impErr != nil {
		s.T().Fatalf("could not import project. cannot continue: %s", impErr.Error())
	}
	s.Equal("successful", imported.ImportStatus)
	checkImport, checkErr := s.TestClient.GetProjectInfo(destProjectName)
	if checkErr != nil {
		s.T().Fatalf("could not get the newly imported project. cannot continue: %s", checkErr.Error())
	}
	for k, v := range s.TestProject.Properties {
		s.Equal(v, checkImport.Properties[k])
	}
}

func (s *ProjectIntegrationTestSuite) TestIntegrationProjectAsyncExportImport() {
	tmpDir := os.TempDir()

	output := filepath.Join(tmpDir, s.TestProjectName+"-async.zip")
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
	token, perr := s.TestClient.GetProjectArchiveExportAsync(s.TestProjectName, opts...)
	if perr != nil {
		s.T().Fatalf("cannot export project: %s", perr.Error())
	}

	s.NotEmpty(token)

	// this is our doneFunc.
	// we poll with the token until it's done
	doneFunc := func() (bool, error) {
		time.Sleep(500 * time.Millisecond)
		info, infoErr := s.TestClient.GetProjectArchiveExportAsyncStatus(s.TestProjectName, token)
		if infoErr != nil {
			return false, infoErr
		}
		return info.Ready, nil
	}

	done, doneErr := s.TestClient.WaitFor(doneFunc, 5*time.Second)
	s.NoError(doneErr)
	s.True(done)

	downloadErr := s.TestClient.GetProjectArchiveExportAsyncDownload(s.TestProjectName, token, f)
	if downloadErr != nil {
		s.T().Fatalf("could not download project archive. cannot continue: %s", downloadErr.Error())
	}
	// import the file into a new project
	destProjectName := testGenerateRandomName("destproject")
	_, destErr := s.TestClient.CreateProject(destProjectName, nil)
	defer func() {
		e := s.TestClient.DeleteProject(destProjectName)
		if e != nil {
			s.T().Errorf("unable to clean up after myself: %s", e.Error())
		}
	}()
	if destErr != nil {
		s.T().Fatalf("unable to create a test project. cannot continue: %s", destErr.Error())
	}
	data, dataErr := os.Open(output)
	defer data.Close() // nolint: errcheck
	if dataErr != nil {
		s.T().Fatalf("cannot open import file: %s", dataErr.Error())
	}
	imported, impErr := s.TestClient.ProjectArchiveImport(destProjectName, data, rundeck.ProjectImportConfigs(true))
	if impErr != nil {
		s.T().Fatalf("could not import project. cannot continue: %s", impErr.Error())
	}
	s.Equal("successful", imported.ImportStatus)
	checkImport, checkErr := s.TestClient.GetProjectInfo(destProjectName)
	if checkErr != nil {
		s.T().Fatalf("could not get the newly imported project. cannot continue: %s", checkErr.Error())
	}
	for k, v := range s.TestProject.Properties {
		s.Equal(v, checkImport.Properties[k])
	}
}

func (s *ProjectIntegrationTestSuite) TestIntegrationGetProjectArchiveExportImportNoConfigs() {
	tmpDir := os.TempDir()

	output := filepath.Join(tmpDir, s.TestProjectName+".zip")
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
	perr := s.TestClient.GetProjectArchiveExport(s.TestProjectName, f, opts...)
	if perr != nil {
		s.T().Fatalf("cannot export project: %s", perr.Error())
	}

	// import the file into a new project
	destProjectName := testGenerateRandomName("destproject")
	_, destErr := s.TestClient.CreateProject(destProjectName, nil)
	defer func() {
		e := s.TestClient.DeleteProject(destProjectName)
		if e != nil {
			s.T().Errorf("unable to clean up import project after myself: %s", e.Error())
		}
	}()
	if destErr != nil {
		s.T().Fatalf("unable to create a destination test project. cannot continue: %s", destErr.Error())
	}
	data, dataErr := os.Open(output)
	defer data.Close() // nolint: errcheck
	if dataErr != nil {
		s.T().Fatalf(dataErr.Error())
	}
	imported, impErr := s.TestClient.ProjectArchiveImport(destProjectName, data)
	if impErr != nil {
		s.T().Fatalf("could not import project. cannot continue: %s", impErr.Error())
	}
	s.Equal("successful", imported.ImportStatus)
	checkImport, checkErr := s.TestClient.GetProjectInfo(destProjectName)
	if checkErr != nil {
		s.T().Fatalf("could not get the newly imported project. cannot continue: %s", checkErr.Error())
	}
	s.NotEqual("stub", checkImport.Properties["service.NodeExecutor.default.provider"])
}

func (s *ProjectIntegrationTestSuite) TestIntegrationProjectConfigurationKeys() {
	destProjectName := testGenerateRandomName("keyproject")
	props := map[string]string{
		"project.description": s.TestProjectName,
	}
	for k, v := range testDefaultProjectProperties {
		props[k] = v
	}
	_, destErr := s.TestClient.CreateProject(destProjectName, props)
	defer func() {
		e := s.TestClient.DeleteProject(destProjectName)
		if e != nil {
			s.T().Errorf("unable to clean up import project after myself: %s", e.Error())
		}
	}()
	if destErr != nil {
		s.T().Fatalf("unable to create a destination test project. cannot continue: %s", destErr.Error())
	}

	_, checkErr := s.TestClient.GetProjectInfo(destProjectName)
	if checkErr != nil {
		s.T().Fatalf("could not get the newly imported project. cannot continue: %s", checkErr.Error())
	}
	nodeSource, nodeSourceErr := s.TestClient.GetProjectConfigurationKey(destProjectName, "resources.source.1.type")
	s.NoError(nodeSourceErr)
	s.Equal("stub", nodeSource)
	putKeyErr := s.TestClient.PutProjectConfigurationKey(destProjectName, "foo", "bar")
	s.NoError(putKeyErr)
	checkKey, checkKeyErr := s.TestClient.GetProjectConfigurationKey(destProjectName, "foo")
	s.NoError(checkKeyErr)
	s.Equal(checkKey, "bar")
	replaceConf, replaceConfErr := s.TestClient.PutProjectConfiguration(destProjectName, props)
	s.NoError(replaceConfErr)
	s.NotContains(replaceConf, "foo")
	deleteErr := s.TestClient.DeleteProjectConfigurationKey(destProjectName, "project.description")
	s.NoError(deleteErr)
	_, checkAgainKeyErr := s.TestClient.GetProjectConfigurationKey(destProjectName, "project.description")
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
