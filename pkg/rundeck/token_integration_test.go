// +build integration

package rundeck

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntegrationBasicAuth(t *testing.T) {
	client, err := NewBasicAuthClient(testIntegrationUsername, testIntegrationPassword, testIntegrationURL)
	assert.NoError(t, err)
	info, infoErr := client.GetSystemInfo()
	assert.NoError(t, infoErr)
	assert.NotNil(t, info)
}

func TestIntegrationTokenAuth(t *testing.T) {
	client, err := NewTokenAuthClient(testIntegrationToken, testIntegrationURL)
	assert.NoError(t, err)
	info, infoErr := client.GetSystemInfo()
	assert.NoError(t, infoErr)
	assert.NotNil(t, info)
}

func TestIntegrationToken(t *testing.T) {
	client := testNewTokenAuthClient()
	createToken, createErr := client.CreateToken("admin")
	if createErr != nil {
		t.Fatalf("Unable to create token. Cannot continue: %s", createErr.Error())
	}
	t.Logf("Created token: %s", createToken.Token)
	getToken, getErr := client.GetToken(createToken.Token)
	assert.NoError(t, getErr)
	assert.ObjectsAreEqualValues(createToken, getToken)
	defer func() {
		deleteErr := client.DeleteToken(createToken.Token)
		if deleteErr != nil {
			t.Logf("error cleaning up token: %s", deleteErr.Error())
		}
	}()
}

func TestIntegrationTokens(t *testing.T) {
	client := testNewTokenAuthClient()
	createToken, createErr := client.CreateToken("admin")
	if createErr != nil {
		t.Fatalf("Unable to create token. Cannot continue: %s", createErr.Error())
	}
	t.Logf("Created token: %s", createToken.Token)
	allTokens, allErr := client.ListTokens()
	assert.NoError(t, allErr)
	assert.ObjectsAreEqualValues(createToken, allTokens[0])
	defer func() {
		deleteErr := client.DeleteToken(createToken.Token)
		if deleteErr != nil {
			t.Logf("error cleaning up token: %s", deleteErr.Error())
		}
	}()
}

func TestIntegrationTokensForUser(t *testing.T) {
	client := testNewTokenAuthClient()
	createToken, createErr := client.CreateToken("admin")
	if createErr != nil {
		t.Fatalf("Unable to create token. Cannot continue: %s", createErr.Error())
	}
	t.Logf("Created token: %s", createToken.Token)
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
