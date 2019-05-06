package rundeck_test

import (
	"testing"

	"github.com/lusis/go-rundeck/pkg/rundeck"
	"github.com/stretchr/testify/suite"
)

type UserIntegrationTestSuite struct {
	suite.Suite
	TestClient     *rundeck.Client
	UserTestClient *rundeck.Client
	AdminProfile   rundeck.User
	UserProfile    rundeck.User
}

func (s *UserIntegrationTestSuite) SetupSuite() {
	client := testNewTokenAuthClient()
	userClient, err := rundeck.NewTokenAuthClient(testIntegrationUserToken, testIntegrationURL)
	s.Require().NoError(err)
	s.TestClient = client
	s.UserTestClient = userClient
	adminProfile := rundeck.User{}
	adminProfile.Login = "admin"
	adminProfile.FirstName = "Mr. Admin"
	adminProfile.LastName = "McAdmin"
	adminProfile.Email = "admin@admin.com"
	s.AdminProfile = adminProfile
	userProfile := rundeck.User{}
	userProfile.Login = "auser"
	userProfile.FirstName = "Alpha"
	userProfile.LastName = "User"
	userProfile.Email = "alpha@user.com"
	s.UserProfile = userProfile
	_, err = s.TestClient.ModifyUserProfile(&s.AdminProfile)
	s.Require().NoError(err)
}

func (s *UserIntegrationTestSuite) TearDownSuite() {

}

func (s *UserIntegrationTestSuite) TestGetCurrentUserProfile() {
	up, err := s.TestClient.GetCurrentUserProfile()
	s.Require().NoError(err)
	s.Require().Equal(s.AdminProfile, *up)
}

func (s *UserIntegrationTestSuite) TestGetUserProfile() {
	up, err := s.TestClient.GetUserProfile("admin")
	s.Require().NoError(err)
	s.Require().Equal(s.AdminProfile, *up)
}

func (s *UserIntegrationTestSuite) TestListUsers() {
	up, err := s.TestClient.ListUsers()
	s.Require().NoError(err)
	s.Require().Len(up, 1)
}

func (s *UserIntegrationTestSuite) TestModifyUserProfile() {
	mup, err := s.TestClient.ModifyUserProfile(&s.AdminProfile)
	s.Require().NoError(err)
	s.Require().NotNil(mup)
	up, err := s.UserTestClient.ModifyUserProfile(&s.UserProfile)
	s.Require().Error(err)
	s.Require().Nil(up)
}

func (s *UserIntegrationTestSuite) TestModifyOtherUserProfile() {

}

func TestIntegrationUserSuite(t *testing.T) {
	if testing.Short() || testRundeckRunning() == false {
		t.Skip("skipping integration testing")
	}

	suite.Run(t, &UserIntegrationTestSuite{})
}
