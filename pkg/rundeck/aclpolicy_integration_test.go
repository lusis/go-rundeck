package rundeck_test

import (
	"strings"
	"sync"
	"testing"

	"github.com/lusis/go-rundeck/pkg/rundeck"
	"github.com/stretchr/testify/suite"
)

type ACLPolicyIntegrationTestSuite struct {
	suite.Suite
	TestClient      *rundeck.Client
	CreatedProjects []rundeck.Project
	sync.Mutex
}

func (s *ACLPolicyIntegrationTestSuite) testCreateProject(slow bool) (rundeck.Project, rundeck.JobImportResult) {
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
	importJob, importErr := s.TestClient.ImportJob(project.Name,
		strings.NewReader(testJobDefinition),
		rundeck.ImportFormat("yaml"),
		rundeck.ImportUUID("remove"))
	if importErr != nil {
		s.T().Fatalf("job did not import. cannot continue: %s", importErr.Error())
	}
	return *project, *importJob
}

func (s *ACLPolicyIntegrationTestSuite) SetupSuite() {
	client := testNewTokenAuthClient()
	s.TestClient = client
	s.CreatedProjects = []rundeck.Project{}
}

func (s *ACLPolicyIntegrationTestSuite) TearDownSuite() {

	for _, p := range s.CreatedProjects {
		e := s.TestClient.DeleteProject(p.Name)
		if e != nil {
			s.T().Logf("unable to clean up test project: %s", e.Error())
		}
	}
}

func (s *ACLPolicyIntegrationTestSuite) TestSystemACLPolicyLifecycle() {
	aclName := s.T().Name()
	aclErr := s.TestClient.CreateSystemACLPolicy(aclName, strings.NewReader(testSystemACLPolicy))
	if aclErr != nil {
		s.T().Fatalf("cannot create a system acl policy. cannot continue: %s", aclErr.Error())
	}

	list, listErr := s.TestClient.ListSystemACLPolicies()
	s.NoError(listErr)
	found := false
	for _, entry := range list.Resources {
		if entry.Name == aclName+".aclpolicy" {
			found = true
			break
		}
	}
	s.True(found)
	acl, getErr := s.TestClient.GetSystemACLPolicy(aclName)
	if getErr != nil {
		s.T().Fatalf("cannot get created policy. cannot continue: %s", getErr.Error())
	}
	s.Equal(testSystemACLPolicy, string(acl))
	uAclErr := s.TestClient.UpdateSystemACLPolicy(aclName, strings.NewReader(testSystemACLPolicySecondary))
	if uAclErr != nil {
		s.T().Fatalf("cannot update policy. cannot continue: %s", uAclErr.Error())
	}
	uacl, ugetErr := s.TestClient.GetSystemACLPolicy(aclName)
	if ugetErr != nil {
		s.T().Fatalf("cannot get created policy. cannot continue: %s", ugetErr.Error())
	}
	s.Equal(testSystemACLPolicySecondary, string(uacl))
	delErr := s.TestClient.DeleteSystemACLPolicy(aclName)
	s.NoError(delErr)
}

func (s *ACLPolicyIntegrationTestSuite) TestProjectACLPolicyLifecycle() {
	project, _ := s.testCreateProject(false)
	aclName := s.T().Name()
	aclErr := s.TestClient.CreateProjectACLPolicy(project.Name, aclName, strings.NewReader(testProjectACLPolicy))
	if aclErr != nil {
		s.T().Fatalf("cannot create a project acl policy. cannot continue: %s", aclErr.Error())
	}

	list, listErr := s.TestClient.ListProjectACLPolicies(project.Name)
	s.NoError(listErr)
	found := false
	for _, entry := range list.Resources {
		if entry.Name == aclName+".aclpolicy" {
			found = true
			break
		}
	}
	s.True(found)
	acl, getErr := s.TestClient.GetProjectACLPolicy(project.Name, aclName)
	if getErr != nil {
		s.T().Fatalf("cannot get created policy. cannot continue: %s", getErr.Error())
	}
	s.Equal(testProjectACLPolicy, string(acl))
	uAclErr := s.TestClient.UpdateProjectACLPolicy(project.Name, aclName, strings.NewReader(testProjectACLPolicySecondary))
	if uAclErr != nil {
		s.T().Fatalf("cannot update policy. cannot continue: %s", uAclErr.Error())
	}
	uacl, ugetErr := s.TestClient.GetProjectACLPolicy(project.Name, aclName)
	if ugetErr != nil {
		s.T().Fatalf("cannot get created policy. cannot continue: %s", ugetErr.Error())
	}
	s.Equal(testProjectACLPolicySecondary, string(uacl))
	delErr := s.TestClient.DeleteProjectACLPolicy(project.Name, aclName)
	s.NoError(delErr)

}

func TestIntegrationACLPolicySuite(t *testing.T) {
	if testRundeckRunning() {
		suite.Run(t, &ACLPolicyIntegrationTestSuite{})
	} else {
		t.Skip("rundeck isn't running for integration testing")
	}
}
