package rundeck

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/suite"
)

type ProjectIntegrationTestSuite struct {
	suite.Suite
	TestProjectName string
	TestProject     *Project
	TestClient      *Client
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

func (s *ProjectIntegrationTestSuite) TestIntegrationGetProjectConfiguration() {
	pc, pcerr := s.TestClient.GetProjectConfiguration(s.TestProjectName)
	s.NoError(pcerr)
	for k, v := range s.TestProject.Properties {
		s.Equal(v, (*pc)[k])
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
	opts := []ProjectExportOption{
		ProjectExportAll(true),
		ProjectExportConfigs(true),
		ProjectExportAcls(true),
		ProjectExportJobs(true),
		ProjectExportReadmes(true),
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
	imported, impErr := s.TestClient.ProjectArchiveImport(destProjectName, data, ProjectImportConfigs(true))
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
	opts := []ProjectExportOption{
		ProjectExportAll(true),
		ProjectExportConfigs(true),
		ProjectExportAcls(true),
		ProjectExportJobs(true),
		ProjectExportReadmes(true),
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

func TestIntegrationProjectSuite(t *testing.T) {
	if testRundeckRunning() {
		suite.Run(t, new(ProjectIntegrationTestSuite))
	} else {
		t.Skip("rundeck isn't running for integration testing")
	}
}
