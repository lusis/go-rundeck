package rundeck

import (
	"errors"
	"testing"

	"github.com/lusis/go-rundeck/pkg/rundeck/responses"
	"github.com/lusis/go-rundeck/pkg/rundeck/responses/testdata"

	"github.com/stretchr/testify/assert"
)

func testFailedTokenOption() TokenOption {
	return func(t *Token) error {
		return errors.New("option setting failed")
	}
}

func TestGetTokens(t *testing.T) {
	jsonfile, err := testdata.GetBytes(responses.ListTokensResponseTestFile)
	if err != nil {
		t.Fatal(err.Error())
	}
	client, server, cErr := newTestRundeckClient(jsonfile, "application/json", 200)
	defer server.Close()
	if cErr != nil {
		t.Fatalf(cErr.Error())
	}
	s, err := client.ListTokens()
	assert.NoError(t, err)
	assert.Len(t, (*s), 4)
	assert.Equal(t, "user3", (*s)[0].User)
	assert.Equal(t, "ece75ac8-2791-442e-b179-a9907d83fd05", (*s)[0].ID)
	assert.Equal(t, "user3", (*s)[0].Creator)
	roles := (*s)[0].Roles
	assert.Len(t, roles, 2)
	assert.Equal(t, "DEV_99", roles[0])
	assert.False(t, (*s)[0].Expired)
	assert.NotEmpty(t, (*s)[0].Expiration)

}

func TestGetTokensDecodeError(t *testing.T) {
	client, server, cErr := newTestRundeckClient([]byte(""), "application/json", 200)
	defer server.Close()
	if cErr != nil {
		t.Fatalf(cErr.Error())
	}
	s, err := client.ListTokens()
	assert.IsType(t, &UnmarshalError{}, err)
	assert.Error(t, err)
	assert.Nil(t, s)
}

func TestGetTokensMissing(t *testing.T) {
	client, server, cErr := newTestRundeckClient([]byte(""), "application/json", 404)
	defer server.Close()
	if cErr != nil {
		t.Fatalf(cErr.Error())
	}
	s, err := client.ListTokens()
	assert.Error(t, err)
	assert.Nil(t, s)
}

func TestGetToken(t *testing.T) {
	jsonfile, err := testdata.GetBytes(responses.TokenResponseTestFile)
	if err != nil {
		t.Fatal(err.Error())
	}
	client, server, cErr := newTestRundeckClient(jsonfile, "application/json", 200)
	defer server.Close()
	if cErr != nil {
		t.Fatalf(cErr.Error())
	}
	s, sErr := client.GetToken("XXXXXX")
	assert.NoError(t, sErr)
	assert.Len(t, s.Roles, 1)
	assert.Equal(t, "user3", s.User)
	assert.Equal(t, "VjkbX2zUAwnXjDIbRYFp824tF5X2N7W1", s.Token)
	assert.Equal(t, "user3", s.Creator)
	assert.NotEmpty(t, s.Expiration)
	assert.True(t, s.Expired)
	assert.Equal(t, "c13de457-c429-4476-9acd-e1c89e3c2928", s.ID)
}

func TestGetTokenMissing(t *testing.T) {
	jsonfile, err := testdata.GetBytes(responses.TokenResponseTestFile)
	if err != nil {
		t.Fatal(err.Error())
	}
	client, server, cErr := newTestRundeckClient(jsonfile, "application/json", 404)
	defer server.Close()
	if cErr != nil {
		t.Fatalf(cErr.Error())
	}
	s, sErr := client.GetToken("XXXXXX")
	assert.Error(t, sErr)
	assert.Nil(t, s)
}

func TestGetTokenInvalidJSON(t *testing.T) {
	client, server, cErr := newTestRundeckClient([]byte(""), "application/json", 200)
	defer server.Close()
	if cErr != nil {
		t.Fatalf(cErr.Error())
	}
	s, sErr := client.GetToken("XXXXXX")
	assert.Error(t, sErr)
	assert.Nil(t, s)
}

func TestGetUserTokens(t *testing.T) {
	jsonfile, err := testdata.GetBytes(responses.ListTokensResponseTestFile)
	if err != nil {
		t.Fatal(err.Error())
	}
	client, server, cErr := newTestRundeckClient(jsonfile, "application/json", 200)
	defer server.Close()
	if cErr != nil {
		t.Fatalf(cErr.Error())
	}
	s, err := client.ListTokensForUser("user3")
	assert.NoError(t, err)
	assert.Len(t, s, 4)
	assert.Equal(t, "user3", s[0].User)
	assert.Equal(t, "ece75ac8-2791-442e-b179-a9907d83fd05", s[0].ID)
	assert.Equal(t, "user3", s[0].Creator)
	assert.Len(t, s[0].Roles, 2)
	assert.Equal(t, "DEV_99", s[0].Roles[0])
	assert.False(t, s[0].Expired)
	assert.NotEmpty(t, s[0].Expiration)

}

func TestGetUserTokensError(t *testing.T) {
	client, server, cErr := newTestRundeckClient([]byte(""), "application/json", 200)
	defer server.Close()
	if cErr != nil {
		t.Fatalf(cErr.Error())
	}
	s, err := client.ListTokensForUser("user3")
	assert.Error(t, err)
	assert.Nil(t, s)
}

func TestGetUserTokensMissing(t *testing.T) {
	client, server, cErr := newTestRundeckClient([]byte(""), "application/json", 404)
	defer server.Close()
	if cErr != nil {
		t.Fatalf(cErr.Error())
	}
	s, err := client.ListTokensForUser("user3")
	assert.Error(t, err)
	assert.Nil(t, s)
}

func TestDeleteToken(t *testing.T) {
	client, server, cErr := newTestRundeckClient([]byte(""), "application/json", 204)
	defer server.Close()
	if cErr != nil {
		t.Fatalf(cErr.Error())
	}
	err := client.DeleteToken("abc123")
	assert.NoError(t, err)
}

func TestDeleteTokenNotFound(t *testing.T) {
	client, server, cErr := newTestRundeckClient([]byte(""), "application/json", 404)
	defer server.Close()
	if cErr != nil {
		t.Fatalf(cErr.Error())
	}
	err := client.DeleteToken("abc123")
	assert.EqualError(t, ErrMissingResource, err.Error())
}

func TestCreateToken(t *testing.T) {
	jsonfile, err := testdata.GetBytes(responses.TokenResponseTestFile)
	if err != nil {
		t.Fatal(err.Error())
	}
	client, server, cErr := newTestRundeckClient(jsonfile, "application/json", 201)
	defer server.Close()
	if cErr != nil {
		t.Fatalf(cErr.Error())
	}
	s, sErr := client.CreateToken("auser")
	assert.NoError(t, sErr)
	assert.Len(t, s.Roles, 1)
	assert.Equal(t, "user3", s.User)
	assert.Equal(t, "VjkbX2zUAwnXjDIbRYFp824tF5X2N7W1", s.Token)
	assert.Equal(t, "user3", s.Creator)
	assert.NotEmpty(t, s.Expiration)
	assert.True(t, s.Expired)
	assert.Equal(t, "c13de457-c429-4476-9acd-e1c89e3c2928", s.ID)
}

func TestCreateTokenJSONError(t *testing.T) {
	client, server, cErr := newTestRundeckClient([]byte(""), "application/json", 201)
	defer server.Close()
	if cErr != nil {
		t.Fatalf(cErr.Error())
	}
	s, sErr := client.CreateToken("auser")
	assert.Error(t, sErr)
	assert.Nil(t, s)
}

func TestCreateTokenInvalidStatus(t *testing.T) {
	jsonfile, err := testdata.GetBytes(responses.TokenResponseTestFile)
	if err != nil {
		t.Fatal(err.Error())
	}
	client, server, cErr := newTestRundeckClient(jsonfile, "application/json", 500)
	defer server.Close()
	if cErr != nil {
		t.Fatalf(cErr.Error())
	}
	s, sErr := client.CreateToken("auser")
	assert.Error(t, sErr)
	assert.Nil(t, s)
}

func TestCreateTokenWithOpts(t *testing.T) {
	jsonfile, err := testdata.GetBytes(responses.TokenResponseTestFile)
	if err != nil {
		t.Fatal(err.Error())
	}
	client, server, cErr := newTestRundeckClient(jsonfile, "application/json", 201)
	defer server.Close()
	if cErr != nil {
		t.Fatalf(cErr.Error())
	}
	s, sErr := client.CreateToken("auser", TokenDuration("120d"))
	assert.NoError(t, sErr)
	assert.Len(t, s.Roles, 1)
	assert.Equal(t, "user3", s.User)
	assert.Equal(t, "VjkbX2zUAwnXjDIbRYFp824tF5X2N7W1", s.Token)

	assert.Equal(t, "user3", s.Creator)
	assert.NotEmpty(t, s.Expiration)
	assert.True(t, s.Expired)
	assert.Equal(t, "c13de457-c429-4476-9acd-e1c89e3c2928", s.ID)
}

func TestCreateTokenWithOptsError(t *testing.T) {
	jsonfile, err := testdata.GetBytes(responses.TokenResponseTestFile)
	if err != nil {
		t.Fatal(err.Error())
	}
	client, server, cErr := newTestRundeckClient(jsonfile, "application/json", 201)
	defer server.Close()
	if cErr != nil {
		t.Fatalf(cErr.Error())
	}
	s, sErr := client.CreateToken("auser", testFailedTokenOption())
	assert.Error(t, sErr)
	assert.Nil(t, s)
}

func TestTokenOption(t *testing.T) {
	token := &Token{}
	opts := []TokenOption{
		TokenRoles("admin", "user"),
		TokenDuration("120d"),
	}
	for _, opt := range opts {
		if err := opt(token); err != nil {
			assert.NoError(t, err)
		}
	}
	assert.Equal(t, "120d", token.Duration)
	assert.Contains(t, token.Roles, "admin")
	assert.Contains(t, token.Roles, "user")
}
