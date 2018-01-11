// +build integration

package rundeck

// These are the preconfigured settings with the docker image built in this repo
var testIntegrationURL = "http://localhost:4440"
var testIntegrationUsername = "admin"
var testIntegrationPassword = "admin"
var testIntegrationToken = "yays72hw87aK2AfxWifTSdcMcY81GL1p"

func testNewBasicAuthClient() *Client {
	client, _ := NewBasicAuthClient(testIntegrationUsername, testIntegrationPassword, testIntegrationURL)
	return client
}

func testNewTokenAuthClient() *Client {
	client, _ := NewTokenAuthClient(testIntegrationToken, testIntegrationURL)
	return client
}
