package rundeck

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type UserIntegrationTestSuite struct {
	suite.Suite
	TestClient     *Client
	UserTestClient *Client
	AdminProfile   User
	UserProfile    User
}

func (s *UserIntegrationTestSuite) SetupSuite() {
	client := testNewTokenAuthClient()
	userClient, _ := NewTokenAuthClient(testIntegrationUserToken, testIntegrationURL)
	s.TestClient = client
	s.UserTestClient = userClient
	adminProfile := User{
		Login:     "admin",
		FirstName: "Mr. Admin",
		LastName:  "McAdmin",
		Email:     "admin@admin.com",
	}
	s.AdminProfile = adminProfile
	userProfile := User{
		Login:     "auser",
		FirstName: "Alpha",
		LastName:  "User",
		Email:     "alpha@user.com",
	}
	s.UserProfile = userProfile
	_, muperr := s.TestClient.ModifyUserProfile(&s.AdminProfile)
	if muperr != nil {
		s.T().Fatalf("can't populate inititial profile for admin user. cannot continue: %s", muperr.Error())
	}
}

func (s *UserIntegrationTestSuite) TearDownSuite() {

}

func (s *UserIntegrationTestSuite) TestGetCurrentUserProfile() {
	up, uperr := s.TestClient.GetCurrentUserProfile()
	s.NoError(uperr)
	s.Equal(s.AdminProfile, *up)
}

func (s *UserIntegrationTestSuite) TestGetUserProfile() {
	up, uperr := s.TestClient.GetUserProfile("admin")
	s.NoError(uperr)
	s.Equal(s.AdminProfile, *up)
}

func (s *UserIntegrationTestSuite) TestListUsers() {
	up, uperr := s.TestClient.ListUsers()
	s.NoError(uperr)
	s.Len(*up, 1)
}

// Need to open some rundeck bugs on this.
// Can't seem to update a profile for any user that hasn't logged in via the UI the first time except for admin
/*
func (s *UserIntegrationTestSuite) TestModifyUserProfile() {
	mup, muperr := s.TestClient.ModifyUserProfile(s.AdminProfile)
	s.NoError(muperr)
	s.NotNil(mup)
	up, uperr := s.UserTestClient.ModifyUserProfile(s.UserProfile)
	s.Error(uperr)
	s.Nil(up)
}

func (s *UserIntegrationTestSuite) TestModifyOtherUserProfile() {

}
*/
func TestIntegrationUserSuite(t *testing.T) {
	if testRundeckRunning() {
		suite.Run(t, &UserIntegrationTestSuite{})
	} else {
		t.Skip("rundeck isn't running for integration testing")
	}
}
