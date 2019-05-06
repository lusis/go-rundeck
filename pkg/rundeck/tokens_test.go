package rundeck

import (
	"errors"
	"testing"

	"github.com/lusis/go-rundeck/pkg/rundeck/responses"

	"github.com/stretchr/testify/require"
)

func testFailedTokenOption() TokenOption {
	return func(t *TokenRequest) error {
		return errors.New("option setting failed")
	}
}

func TestGetTokens(t *testing.T) {
	jsonfile, err := responses.GetTestData(responses.ListTokensResponseTestFile)
	if err != nil {
		t.Fatal(err.Error())
	}
	client, server, cErr := newTestRundeckClient(jsonfile, "application/json", 200)
	defer server.Close()
	if cErr != nil {
		t.Fatalf(cErr.Error())
	}
	s, err := client.ListTokens()
	require.NoError(t, err)
	require.Len(t, s, 4)
	require.Equal(t, "user3", s[0].User)
	require.Equal(t, "ece75ac8-2791-442e-b179-a9907d83fd05", s[0].ID)
	require.Equal(t, "user3", s[0].Creator)
	roles := s[0].Roles
	require.Len(t, roles, 2)
	require.Equal(t, "DEV_99", roles[0])
	require.False(t, s[0].Expired)
	require.NotEmpty(t, s[0].Expiration)

}

func TestGetTokensDecodeError(t *testing.T) {
	client, server, cErr := newTestRundeckClient([]byte(""), "application/json", 200)
	defer server.Close()
	if cErr != nil {
		t.Fatalf(cErr.Error())
	}
	s, err := client.ListTokens()
	require.IsType(t, &UnmarshalError{}, err)
	require.Error(t, err)
	require.Nil(t, s)
}

func TestGetTokensMissing(t *testing.T) {
	client, server, cErr := newTestRundeckClient([]byte(""), "application/json", 404)
	defer server.Close()
	if cErr != nil {
		t.Fatalf(cErr.Error())
	}
	s, err := client.ListTokens()
	require.Error(t, err)
	require.Nil(t, s)
}

func TestGetToken(t *testing.T) {
	jsonfile, err := responses.GetTestData(responses.TokenResponseTestFile)
	if err != nil {
		t.Fatal(err.Error())
	}
	client, server, cErr := newTestRundeckClient(jsonfile, "application/json", 200)
	defer server.Close()
	if cErr != nil {
		t.Fatalf(cErr.Error())
	}
	s, sErr := client.GetToken("XXXXXX")
	require.NoError(t, sErr)
	require.Len(t, s.Roles, 1)
	require.Equal(t, "user3", s.User)
	require.Equal(t, "VjkbX2zUAwnXjDIbRYFp824tF5X2N7W1", s.Token)
	require.Equal(t, "user3", s.Creator)
	require.NotEmpty(t, s.Expiration)
	require.True(t, s.Expired)
	require.Equal(t, "c13de457-c429-4476-9acd-e1c89e3c2928", s.ID)
}

func TestGetTokenMissing(t *testing.T) {
	jsonfile, err := responses.GetTestData(responses.TokenResponseTestFile)
	if err != nil {
		t.Fatal(err.Error())
	}
	client, server, cErr := newTestRundeckClient(jsonfile, "application/json", 404)
	defer server.Close()
	if cErr != nil {
		t.Fatalf(cErr.Error())
	}
	s, sErr := client.GetToken("XXXXXX")
	require.Error(t, sErr)
	require.Nil(t, s)
}

func TestGetTokenInvalidJSON(t *testing.T) {
	client, server, cErr := newTestRundeckClient([]byte(""), "application/json", 200)
	defer server.Close()
	if cErr != nil {
		t.Fatalf(cErr.Error())
	}
	s, sErr := client.GetToken("XXXXXX")
	require.Error(t, sErr)
	require.Nil(t, s)
}

func TestGetUserTokens(t *testing.T) {
	jsonfile, err := responses.GetTestData(responses.ListTokensResponseTestFile)
	if err != nil {
		t.Fatal(err.Error())
	}
	client, server, cErr := newTestRundeckClient(jsonfile, "application/json", 200)
	defer server.Close()
	if cErr != nil {
		t.Fatalf(cErr.Error())
	}
	s, err := client.ListTokensForUser("user3")
	require.NoError(t, err)
	require.Len(t, s, 4)
	require.Equal(t, "user3", s[0].User)
	require.Equal(t, "ece75ac8-2791-442e-b179-a9907d83fd05", s[0].ID)
	require.Equal(t, "user3", s[0].Creator)
	require.Len(t, s[0].Roles, 2)
	require.Equal(t, "DEV_99", s[0].Roles[0])
	require.False(t, s[0].Expired)
	require.NotEmpty(t, s[0].Expiration)

}

func TestGetUserTokensError(t *testing.T) {
	client, server, cErr := newTestRundeckClient([]byte(""), "application/json", 200)
	defer server.Close()
	if cErr != nil {
		t.Fatalf(cErr.Error())
	}
	s, err := client.ListTokensForUser("user3")
	require.Error(t, err)
	require.Nil(t, s)
}

func TestGetUserTokensMissing(t *testing.T) {
	client, server, cErr := newTestRundeckClient([]byte(""), "application/json", 404)
	defer server.Close()
	if cErr != nil {
		t.Fatalf(cErr.Error())
	}
	s, err := client.ListTokensForUser("user3")
	require.Error(t, err)
	require.Nil(t, s)
}

func TestDeleteToken(t *testing.T) {
	client, server, cErr := newTestRundeckClient([]byte(""), "application/json", 204)
	defer server.Close()
	if cErr != nil {
		t.Fatalf(cErr.Error())
	}
	err := client.DeleteToken("abc123")
	require.NoError(t, err)
}

func TestDeleteTokenNotFound(t *testing.T) {
	client, server, cErr := newTestRundeckClient([]byte(""), "application/json", 404)
	defer server.Close()
	if cErr != nil {
		t.Fatalf(cErr.Error())
	}
	err := client.DeleteToken("abc123")
	require.EqualError(t, ErrMissingResource, err.Error())
}

func TestCreateToken(t *testing.T) {
	jsonfile, err := responses.GetTestData(responses.TokenResponseTestFile)
	if err != nil {
		t.Fatal(err.Error())
	}
	client, server, cErr := newTestRundeckClient(jsonfile, "application/json", 201)
	defer server.Close()
	if cErr != nil {
		t.Fatalf(cErr.Error())
	}
	s, sErr := client.CreateToken("auser")
	require.NoError(t, sErr)
	require.Len(t, s.Roles, 1)
	require.Equal(t, "user3", s.User)
	require.Equal(t, "VjkbX2zUAwnXjDIbRYFp824tF5X2N7W1", s.Token)
	require.Equal(t, "user3", s.Creator)
	require.NotEmpty(t, s.Expiration)
	require.True(t, s.Expired)
	require.Equal(t, "c13de457-c429-4476-9acd-e1c89e3c2928", s.ID)
}

func TestCreateTokenJSONError(t *testing.T) {
	client, server, cErr := newTestRundeckClient([]byte(""), "application/json", 201)
	defer server.Close()
	if cErr != nil {
		t.Fatalf(cErr.Error())
	}
	s, sErr := client.CreateToken("auser")
	require.Error(t, sErr)
	require.Nil(t, s)
}

func TestCreateTokenInvalidStatus(t *testing.T) {
	jsonfile, err := responses.GetTestData(responses.TokenResponseTestFile)
	if err != nil {
		t.Fatal(err.Error())
	}
	client, server, cErr := newTestRundeckClient(jsonfile, "application/json", 500)
	defer server.Close()
	if cErr != nil {
		t.Fatalf(cErr.Error())
	}
	s, sErr := client.CreateToken("auser")
	require.Error(t, sErr)
	require.Nil(t, s)
}

func TestCreateTokenWithOpts(t *testing.T) {
	jsonfile, err := responses.GetTestData(responses.TokenResponseTestFile)
	if err != nil {
		t.Fatal(err.Error())
	}
	client, server, cErr := newTestRundeckClient(jsonfile, "application/json", 201)
	defer server.Close()
	if cErr != nil {
		t.Fatalf(cErr.Error())
	}
	s, sErr := client.CreateToken("auser", TokenDuration("120d"))
	require.NoError(t, sErr)
	require.Len(t, s.Roles, 1)
	require.Equal(t, "user3", s.User)
	require.Equal(t, "VjkbX2zUAwnXjDIbRYFp824tF5X2N7W1", s.Token)

	require.Equal(t, "user3", s.Creator)
	require.NotEmpty(t, s.Expiration)
	require.True(t, s.Expired)
	require.Equal(t, "c13de457-c429-4476-9acd-e1c89e3c2928", s.ID)
}

func TestCreateTokenWithOptsError(t *testing.T) {
	jsonfile, err := responses.GetTestData(responses.TokenResponseTestFile)
	if err != nil {
		t.Fatal(err.Error())
	}
	client, server, cErr := newTestRundeckClient(jsonfile, "application/json", 201)
	defer server.Close()
	if cErr != nil {
		t.Fatalf(cErr.Error())
	}
	s, sErr := client.CreateToken("auser", testFailedTokenOption())
	require.Error(t, sErr)
	require.Nil(t, s)
}

func TestTokenOption(t *testing.T) {
	token := &TokenRequest{}
	opts := []TokenOption{
		TokenRoles("admin", "user"),
		TokenDuration("120d"),
	}
	for _, opt := range opts {
		if err := opt(token); err != nil {
			require.NoError(t, err)
		}
	}
	require.Equal(t, "120d", token.Duration)
	require.Contains(t, token.Roles, "admin")
	require.Contains(t, token.Roles, "user")
}
