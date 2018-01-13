package rundeck

import (
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// These are the preconfigured settings with the docker image built in this repo
var testIntegrationURL = "http://localhost:4440"
var testIntegrationUsername = "admin"
var testIntegrationPassword = "admin"
var testIntegrationToken = "yays72hw87aK2AfxWifTSdcMcY81GL1p"
var testIntegrationUserToken = "jHpBIeJRkfVHfWlmiPRxXH2GSk2DF3wy"
var testAdHocScriptURL = "https://gist.github.com/lusis/c230f2d8323e0d440a29d25a8b3bb7af/raw/ccfa844799d375293b1028b1c8f85c2c786be0d1/test.py"

// create a project with 5 stub nodes
var testDefaultProjectProperties = map[string]string{
	"service.NodeExecutor.default.provider": "stub",
	"service.FileCopier.default.provider":   "stub",
	"resources.source.1.config.count":       "5",
	"resources.source.1.config.delay":       "0",
	"resources.source.1.config.prefix":      "node",
	"resources.source.1.config.suffix":      "-stub",
	"resources.source.1.config.tags":        "stub",
	"resources.source.1.type":               "stub",
	"project.nodeCache.delay":               "0",
	"project.nodeCache.enabled":             "false",
}

func testRundeckRunning() bool {
	res, err := http.Get(testIntegrationURL)
	if err != nil {
		return false
	}
	if res.StatusCode > 400 {
		return false
	}
	return true
}

func testNewBasicAuthClient() *Client {
	client, _ := NewBasicAuthClient(testIntegrationUsername, testIntegrationPassword, testIntegrationURL)
	return client
}

func testNewTokenAuthClient() *Client {
	client, _ := NewTokenAuthClient(testIntegrationToken, testIntegrationURL)
	return client
}

func testGenerateRandomName(resourceType string) string {
	tstamp := fmt.Sprintf("%d", time.Now().UnixNano())
	return fmt.Sprintf("%s-%s", resourceType, tstamp)
}

func TestIntegrationBasicAuth(t *testing.T) {
	if !testRundeckRunning() {
		t.Skip("rundeck not running for integration tests")
	}
	client, err := NewBasicAuthClient(testIntegrationUsername, testIntegrationPassword, testIntegrationURL)
	assert.NoError(t, err)
	info, infoErr := client.GetSystemInfo()
	assert.NoError(t, infoErr)
	assert.NotNil(t, info)
}

func TestIntegrationTokenAuth(t *testing.T) {
	if !testRundeckRunning() {
		t.Skip("rundeck not running for integration tests")
	}
	client, err := NewTokenAuthClient(testIntegrationToken, testIntegrationURL)
	assert.NoError(t, err)
	info, infoErr := client.GetSystemInfo()
	assert.NoError(t, infoErr)
	assert.NotNil(t, info)
}

func TestIntegrationInvalidBasicAuth(t *testing.T) {
	if !testRundeckRunning() {
		t.Skip("rundeck not running for integration tests")
	}
	client, _ := NewBasicAuthClient("bob", "bob", testIntegrationURL)
	info, infoErr := client.GetSystemInfo()
	assert.Error(t, infoErr)
	assert.Nil(t, info)
}
