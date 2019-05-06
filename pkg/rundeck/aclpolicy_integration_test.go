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
	project, err := s.TestClient.CreateProject(projectName, props)
	s.Require().NoError(err)
	s.Lock()
	s.CreatedProjects = append(s.CreatedProjects, *project)
	s.Unlock()
	importJob, err := s.TestClient.ImportJob(project.Name,
		strings.NewReader(testJobDefinition),
		rundeck.ImportFormat("yaml"),
		rundeck.ImportUUID("remove"))
	s.Require().NoError(err)
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
	err := s.TestClient.CreateSystemACLPolicy(aclName, strings.NewReader(testSystemACLPolicy))
	s.Require().NoError(err)

	list, err := s.TestClient.ListSystemACLPolicies()
	s.Require().NoError(err)
	found := false
	for _, entry := range list.Resources {
		if entry.Name == aclName+".aclpolicy" {
			found = true
			break
		}
	}
	s.Require().True(found)
	acl, err := s.TestClient.GetSystemACLPolicy(aclName)
	s.Require().NoError(err)
	s.Require().Equal(testSystemACLPolicy, string(acl))
	err = s.TestClient.UpdateSystemACLPolicy(aclName, strings.NewReader(testSystemACLPolicySecondary))
	s.Require().NoError(err)
	uacl, err := s.TestClient.GetSystemACLPolicy(aclName)
	s.Require().NoError(err)
	s.Require().Equal(testSystemACLPolicySecondary, string(uacl))
	err = s.TestClient.DeleteSystemACLPolicy(aclName)
	s.Require().NoError(err)
}

func (s *ACLPolicyIntegrationTestSuite) TestProjectACLPolicyLifecycle() {
	project, _ := s.testCreateProject(false)
	aclName := s.T().Name()
	err := s.TestClient.CreateProjectACLPolicy(project.Name, aclName, strings.NewReader(testProjectACLPolicy))
	s.Require().NoError(err)

	list, err := s.TestClient.ListProjectACLPolicies(project.Name)
	s.Require().NoError(err)
	found := false
	for _, entry := range list.Resources {
		if entry.Name == aclName+".aclpolicy" {
			found = true
			break
		}
	}
	s.Require().True(found)
	acl, err := s.TestClient.GetProjectACLPolicy(project.Name, aclName)
	s.Require().NoError(err)
	s.Require().Equal(testProjectACLPolicy, string(acl))
	err = s.TestClient.UpdateProjectACLPolicy(project.Name, aclName, strings.NewReader(testProjectACLPolicySecondary))
	s.Require().NoError(err)

	uacl, err := s.TestClient.GetProjectACLPolicy(project.Name, aclName)
	s.Require().NoError(err)
	s.Require().Equal(testProjectACLPolicySecondary, string(uacl))
	err = s.TestClient.DeleteProjectACLPolicy(project.Name, aclName)
	s.Require().NoError(err)
}

func TestIntegrationACLPolicySuite(t *testing.T) {
	if testing.Short() || testRundeckRunning() == false {
		t.Skip("skipping integration testing")
	}

	suite.Run(t, &ACLPolicyIntegrationTestSuite{})
}
