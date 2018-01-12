// +build integration

package rundeck

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func testGenerateRandomName(resourceType string) string {
	tstamp := fmt.Sprintf("%d", time.Now().UnixNano())
	return fmt.Sprintf("%s-%s", resourceType, tstamp)
}

func TestIntegrationProjects(t *testing.T) {
	client := testNewTokenAuthClient()
	projectName := testGenerateRandomName("testproject")
	createProject, createErr := client.CreateProject(projectName, map[string]string{"fooprop": "fooval"})
	defer func() {
		e := client.DeleteProject(projectName)
		if e != nil {
			t.Errorf("unable to clean up after myself: %s", e.Error())
		}
	}()
	if createErr != nil {
		t.Fatalf("unable to create a test project. cannot continue: %s", createErr.Error())
	}

	listProjects, listErr := client.ListProjects()
	assert.NoError(t, listErr)
	assert.ObjectsAreEqualValues(createProject, listProjects[0])

}

func TestIntegrationProject(t *testing.T) {
	client := testNewTokenAuthClient()
	projectName := testGenerateRandomName("testproject")
	createProject, createErr := client.CreateProject(projectName, map[string]string{"fooprop": "fooval"})
	defer func() {
		e := client.DeleteProject(projectName)
		if e != nil {
			t.Errorf("unable to clean up after myself: %s", e.Error())
		}
	}()
	if createErr != nil {
		t.Fatalf("unable to create a test project. cannot continue: %s", createErr.Error())
	}

	getProject, getErr := client.GetProjectInfo(projectName)
	assert.NoError(t, getErr)
	assert.ObjectsAreEqualValues(createProject, getProject)

}

func TestIntegrationGetProjectConfiguration(t *testing.T) {
	client := testNewTokenAuthClient()
	projectName := testGenerateRandomName("testproject")
	createProject, createErr := client.CreateProject(projectName, map[string]string{"fooprop": "fooval"})
	defer func() {
		e := client.DeleteProject(projectName)
		if e != nil {
			t.Errorf("unable to clean up after myself: %s", e.Error())
		}
	}()
	if createErr != nil {
		t.Fatalf("unable to create a test project. cannot continue: %s", createErr.Error())
	}
	pc, pcerr := client.GetProjectConfiguration(projectName)
	assert.NoError(t, pcerr)
	assert.ObjectsAreEqualValues(createProject.Properties, pc)

}

func TestIntegrationGetProjectArchiveExportImport(t *testing.T) {
	client := testNewTokenAuthClient()
	// Create a project
	projectName := testGenerateRandomName("exportproject")
	/*
		project.description=this is a test project
		project.disable.executions=false
		project.disable.schedule=false
		project.jobs.gui.groupExpandLevel=1
		project.name=testproject
		project.nodeCache.delay=30
		project.nodeCache.enabled=true
		project.ssh-authentication=privateKey
		project.ssh-keypath=/var/lib/rundeck/.ssh/id_rsa
		resources.source.1.config.file=/var/rundeck/projects/testproject/etc/resources.xml
		resources.source.1.config.generateFileAutomatically=true
		resources.source.1.config.includeServerNode=true
		resources.source.1.config.requireFileExists=false
		resources.source.1.config.writeable=false
		resources.source.1.type=file
		service.FileCopier.default.provider=stub
		service.NodeExecutor.default.provider=stub
	*/
	props := map[string]string{
		"project.description":                   projectName,
		"service.NodeExecutor.default.provider": "stub",
		"service.FileCopier.default.provider":   "stub",
	}
	_, createErr := client.CreateProject(projectName, props)
	defer func() {
		e := client.DeleteProject(projectName)
		if e != nil {
			t.Errorf("unable to clean up after myself: %s", e.Error())
		}
	}()
	if createErr != nil {
		t.Fatalf("unable to create a test project. cannot continue: %s", createErr.Error())
	}
	tmpDir := os.TempDir()

	output := filepath.Join(tmpDir, projectName+".zip")
	f, fErr := os.Create(output)
	//defer os.Remove(output) // nolint: errcheck
	if fErr != nil {
		t.Fatalf("unable to create output file. cannot continue: %s", fErr)
	}
	opts := []ProjectExportOption{
		ProjectExportAll(true),
		ProjectExportConfigs(true),
		ProjectExportAcls(true),
		ProjectExportJobs(true),
		ProjectExportReadmes(true),
	}
	// Export created project
	perr := client.GetProjectArchiveExport(projectName, f, opts...)
	assert.NoError(t, perr)

	// import the file into a new project
	destProjectName := testGenerateRandomName("destproject")
	_, destErr := client.CreateProject(destProjectName, nil)
	defer func() {
		e := client.DeleteProject(destProjectName)
		if e != nil {
			t.Errorf("unable to clean up after myself: %s", e.Error())
		}
	}()
	if destErr != nil {
		t.Fatalf("unable to create a test project. cannot continue: %s", destErr.Error())
	}
	data, dataErr := os.Open("/tmp/exportproject-1515766381056320139.zip")
	defer data.Close()
	if dataErr != nil {
		log.Fatalf(dataErr.Error())
	}
	imported, impErr := client.ProjectArchiveImport(destProjectName, data, ProjectImportConfigs(true))
	if impErr != nil {
		t.Fatalf("could not import project. cannot continue: %s", impErr.Error())
	}
	assert.Equal(t, "successful", imported.ImportStatus)
	checkImport, checkErr := client.GetProjectInfo(destProjectName)
	if checkErr != nil {
		t.Fatalf("could not get the newly imported project. cannot continue: %s", checkErr.Error())
	}

	assert.Equal(t, "stub", checkImport.Properties["service.NodeExecutor.default.provider"])
}

func TestIntegrationGetProjectArchiveExportImportNoConfigs(t *testing.T) {
	client := testNewTokenAuthClient()
	// Create a project
	projectName := testGenerateRandomName("exportproject")

	props := map[string]string{
		"project.description":                   projectName,
		"service.NodeExecutor.default.provider": "stub",
		"service.FileCopier.default.provider":   "stub",
	}
	_, createErr := client.CreateProject(projectName, props)
	defer func() {
		e := client.DeleteProject(projectName)
		if e != nil {
			t.Errorf("unable to clean up after myself: %s", e.Error())
		}
	}()
	if createErr != nil {
		t.Fatalf("unable to create a test project. cannot continue: %s", createErr.Error())
	}
	tmpDir := os.TempDir()

	output := filepath.Join(tmpDir, projectName+".zip")
	f, fErr := os.Create(output)
	defer os.Remove(output) // nolint: errcheck
	if fErr != nil {
		t.Fatalf("unable to create output file. cannot continue: %s", fErr)
	}
	opts := []ProjectExportOption{
		ProjectExportAll(true),
		ProjectExportConfigs(true),
		ProjectExportAcls(true),
		ProjectExportJobs(true),
		ProjectExportReadmes(true),
	}
	// Export created project
	perr := client.GetProjectArchiveExport(projectName, f, opts...)
	assert.NoError(t, perr)

	// import the file into a new project
	destProjectName := testGenerateRandomName("destproject")
	_, destErr := client.CreateProject(destProjectName, nil)
	defer func() {
		e := client.DeleteProject(destProjectName)
		if e != nil {
			t.Errorf("unable to clean up after myself: %s", e.Error())
		}
	}()
	if destErr != nil {
		t.Fatalf("unable to create a test project. cannot continue: %s", destErr.Error())
	}
	data, dataErr := os.Open("/tmp/exportproject-1515766381056320139.zip")
	defer data.Close()
	if dataErr != nil {
		log.Fatalf(dataErr.Error())
	}
	imported, impErr := client.ProjectArchiveImport(destProjectName, data)
	if impErr != nil {
		t.Fatalf("could not import project. cannot continue: %s", impErr.Error())
	}
	assert.Equal(t, "successful", imported.ImportStatus)
	checkImport, checkErr := client.GetProjectInfo(destProjectName)
	if checkErr != nil {
		t.Fatalf("could not get the newly imported project. cannot continue: %s", checkErr.Error())
	}

	assert.NotEqual(t, "stub", checkImport.Properties["service.NodeExecutor.default.provider"])
}
