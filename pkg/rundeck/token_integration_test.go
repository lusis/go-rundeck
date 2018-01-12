// +build integration

package rundeck

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntegrationToken(t *testing.T) {
	client := testNewTokenAuthClient()
	createToken, createErr := client.CreateToken("admin")
	if createErr != nil {
		t.Fatalf("Unable to create token. Cannot continue: %s", createErr.Error())
	}
	getToken, getErr := client.GetToken(createToken.Token)
	assert.NoError(t, getErr)
	assert.ObjectsAreEqualValues(createToken, getToken)
	defer func() {
		deleteErr := client.DeleteToken(createToken.Token)
		if deleteErr != nil {
			t.Errorf("error cleaning up token: %s", deleteErr.Error())
		}
	}()
}

func TestIntegrationTokens(t *testing.T) {
	client := testNewTokenAuthClient()
	createToken, createErr := client.CreateToken("admin")
	if createErr != nil {
		t.Fatalf("Unable to create token. Cannot continue: %s", createErr.Error())
	}
	allTokens, allErr := client.ListTokens()
	assert.NoError(t, allErr)
	assert.ObjectsAreEqualValues(createToken, allTokens[0])
	defer func() {
		deleteErr := client.DeleteToken(createToken.Token)
		if deleteErr != nil {
			t.Errorf("error cleaning up token: %s", deleteErr.Error())
		}
	}()
}

func TestIntegrationTokensForUser(t *testing.T) {
	client := testNewTokenAuthClient()
	createToken, createErr := client.CreateToken("admin")
	if createErr != nil {
		t.Fatalf("Unable to create token. Cannot continue: %s", createErr.Error())
	}
	allTokens, allErr := client.ListTokensForUser("admin")
	assert.NoError(t, allErr)
	assert.ObjectsAreEqualValues(createToken, allTokens[0])
	defer func() {
		deleteErr := client.DeleteToken(createToken.Token)
		if deleteErr != nil {
			t.Logf("error cleaning up token: %s", deleteErr.Error())
		}
	}()
}
