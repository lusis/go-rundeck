package rundeck

import (
	"testing"

	"github.com/lusis/go-rundeck/pkg/rundeck.v21/responses"
	"github.com/lusis/go-rundeck/pkg/rundeck.v21/responses/testdata"

	"github.com/stretchr/testify/assert"
)

func TestTokens(t *testing.T) {
	jsonfile, err := testdata.GetBytes(responses.TokensResponseTestFile)
	if err != nil {
		t.Fatal(err.Error())
	}
	client, server, cErr := newTestRundeckClient(jsonfile, "application/json", 200)
	defer server.Close()
	if cErr != nil {
		t.Fatalf(cErr.Error())
	}
	s, err := client.GetTokens()
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

func TestUserToken(t *testing.T) {
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
