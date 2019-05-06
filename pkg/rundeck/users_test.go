package rundeck

import (
	"testing"

	"github.com/lusis/go-rundeck/pkg/rundeck/responses"

	"github.com/stretchr/testify/require"
)

func TestGetUsers(t *testing.T) {
	jsonfile, err := responses.GetTestData(responses.ListUsersResponseTestFile)
	require.NoError(t, err)
	client, server, cErr := newTestRundeckClient(jsonfile, "application/json", 200)
	defer server.Close()
	require.NoError(t, cErr)
	s, err := client.ListUsers()
	require.NoError(t, err)
	require.NotNil(t, s)
	require.Len(t, s, 2)
}

func TestGetUsersJSONError(t *testing.T) {
	client, server, cErr := newTestRundeckClient([]byte(""), "application/json", 200)
	defer server.Close()
	require.NoError(t, cErr)
	s, err := client.ListUsers()
	require.Error(t, err)
	require.Nil(t, s)
}

func TestGetUsersInvalidStatus(t *testing.T) {
	client, server, cErr := newTestRundeckClient([]byte(""), "application/json", 500)
	defer server.Close()
	require.NoError(t, cErr)
	s, err := client.ListUsers()
	require.Error(t, err)
	require.Nil(t, s)
}

func TestGetUserInfo(t *testing.T) {
	jsonfile, err := responses.GetTestData(responses.UserProfileResponseTestFile)
	require.NoError(t, err)
	client, server, cErr := newTestRundeckClient(jsonfile, "application/json", 200)
	defer server.Close()
	require.NoError(t, cErr)

	s, err := client.GetUserProfile("admin")
	require.NoError(t, err)
	require.NotNil(t, s)
	require.Equal(t, "admin", s.Login)
	require.Equal(t, "Admin", s.FirstName)
	require.Equal(t, "McAdmin", s.LastName)
	require.Equal(t, "admin@server.com", s.Email)
}

func TestGetUserInfoJSONError(t *testing.T) {
	client, server, cErr := newTestRundeckClient([]byte(""), "application/json", 200)
	defer server.Close()
	require.NoError(t, cErr)
	s, err := client.GetUserProfile("admin")
	require.Error(t, err)
	require.Nil(t, s)
}

func TestGetUserInfoInvalidStatus(t *testing.T) {
	client, server, cErr := newTestRundeckClient([]byte(""), "application/json", 500)
	defer server.Close()
	require.NoError(t, cErr)
	s, err := client.GetUserProfile("admin")
	require.Error(t, err)
	require.Nil(t, s)
}

func TestGetCurrentUserInfo(t *testing.T) {
	jsonfile, err := responses.GetTestData(responses.UserProfileResponseTestFile)
	require.NoError(t ,err)
	client, server, cErr := newTestRundeckClient(jsonfile, "application/json", 200)
	defer server.Close()
	require.NoError(t, cErr)
	s, err := client.GetCurrentUserProfile()
	require.NoError(t, err)
	require.NotNil(t, s)
	require.Equal(t, "admin", s.Login)
	require.Equal(t, "Admin", s.FirstName)
	require.Equal(t, "McAdmin", s.LastName)
	require.Equal(t, "admin@server.com", s.Email)
}

func TestGetCurrentUserInfoInvalidStatus(t *testing.T) {
	client, server, cErr := newTestRundeckClient([]byte(""), "application/json", 500)
	defer server.Close()
	require.NoError(t, cErr)

	s, err := client.GetCurrentUserProfile()
	require.Error(t, err)
	require.Nil(t, s)
}

func TestGetCurrentUserInfoJSONError(t *testing.T) {
	client, server, cErr := newTestRundeckClient([]byte(""), "application/json", 200)
	defer server.Close()
	require.NoError(t, cErr)

	s, err := client.GetCurrentUserProfile()
	require.Error(t, err)
	require.Nil(t, s)
}
