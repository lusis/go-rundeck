package rundeck_test

import (
	"bytes"
	"fmt"
	"net/http"
	"testing"
	"text/template"
	"time"

	"github.com/lusis/go-rundeck/pkg/rundeck"
	"github.com/stretchr/testify/assert"
)

// These are the preconfigured settings with the docker image built in this repo
var testIntegrationURL = "http://localhost:4440"
var testIntegrationUsername = "admin"
var testIntegrationPassword = "admin"
var testIntegrationToken = "yays72hw87aK2AfxWifTSdcMcY81GL1p"
var testIntegrationUserToken = "jHpBIeJRkfVHfWlmiPRxXH2GSk2DF3wy"
var testAdHocScriptURL = "https://gist.github.com/lusis/c230f2d8323e0d440a29d25a8b3bb7af/raw/ccfa844799d375293b1028b1c8f85c2c786be0d1/test.py"

var testJobDefinition = `
- description: this is a test job
  executionEnabled: true
  group: test/jobs
  id: 8c3176bf-e553-4086-b7b7-38e19974cd89
  loglevel: INFO
  name: testjob
  multipleExecutions: true
  nodeFilterEditable: false
  nodefilters:
    dispatch:
      excludePrecedence: true
      keepgoing: false
      rankOrder: ascending
      successOnEmptyNodeFilter: false
      threadcount: 1
    filter: .*
  nodesSelectedByDefault: true
  scheduleEnabled: true
  sequence:
    commands:
    - description: ps output
      exec: ps -ef
    keepgoing: false
    strategy: node-first
  uuid: 8c3176bf-e553-4086-b7b7-38e19974cd89
`
var testJobDefinitionTemplate = `
- description: {{.Description}}
  executionEnabled: true
  group: {{.Group}}
  id: 8c3176bf-e553-4086-b7b7-38e19974cd89
  loglevel: INFO
  name: {{.Name}}
  multipleExecutions: true
  nodeFilterEditable: false
  nodefilters:
    dispatch:
      excludePrecedence: true
      keepgoing: false
      rankOrder: ascending
      successOnEmptyNodeFilter: false
      threadcount: 1
    filter: .*
  nodesSelectedByDefault: true
  scheduleEnabled: true
  sequence:
    commands:
    - description: ps output
      exec: ps -ef
    keepgoing: false
    strategy: node-first
  uuid: 8c3176bf-e553-4086-b7b7-38e19974cd89
`

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

var testSystemACLPolicy = `
description: Admin, all access.
context:
  project: '.*' # all projects
for:
  resource:
    - allow: '*' # allow read/create all kinds
  adhoc:
    - allow: '*' # allow read/running/killing adhoc jobs
  job:
    - allow: '*' # allow read/write/delete/run/kill of all jobs
  node:
    - allow: '*' # allow read/run for all nodes
by:
  group: admin

---

description: Admin, all access.
context:
  application: 'rundeck'
for:
  resource:
    - allow: '*' # allow create of projects
  project:
    - allow: '*' # allow view/admin of all projects
  project_acl:
    - allow: '*' # allow admin of all project-level ACL policies
  storage:
    - allow: '*' # allow read/create/update/delete for all /keys/* storage content
  system_acl:
    - allow: '*' # allow full access to the system
by:
  group: admin
`

var testSystemACLPolicySecondary = `
description: Admin, all access.
context:
  project: '.*' # all projects
for:
  resource:
    - allow: '*' # allow read/create all kinds
  adhoc:
    - allow: '*' # allow read/running/killing adhoc jobs
  job:
    - allow: '*' # allow read/write/delete/run/kill of all jobs
  node:
    - allow: '*' # allow read/run for all nodes
by:
  group: foogroup

---

description: Admin, all access.
context:
  application: 'rundeck'
for:
  resource:
    - allow: '*' # allow create of projects
  project:
    - allow: '*' # allow view/admin of all projects
  project_acl:
    - allow: '*' # allow admin of all project-level ACL policies
  storage:
    - allow: '*' # allow read/create/update/delete for all /keys/* storage content
  system_acl:
    - allow: '*' # allow full access to the system
by:
  group: foogroup
`

var testProjectACLPolicy = `
description: auser access
for:
  resource:
    - allow: '*' # allow read/create all kinds
  adhoc:
    - allow: '*' # allow read/running/killing adhoc jobs
  job:
    - allow: '*' # allow read/write/delete/run/kill of all jobs
  node:
    - allow: '*' # allow read/run for all nodes
by:
  username: 'auser'
`

var testProjectACLPolicySecondary = `
description: auser access
for:
  resource:
    - allow: '*' # allow read/create all kinds
  adhoc:
    - allow: '*' # allow read/running/killing adhoc jobs
  job:
    - allow: '*' # allow read/write/delete/run/kill of all jobs
  node:
    - allow: '*' # allow read/run for all nodes
by:
  username: 'bob'
`

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

func testJobFromTemplate(name, description string) ([]byte, error) {
	t, tErr := template.New("tmp_template").Parse(testJobDefinitionTemplate)
	if tErr != nil {
		return nil, tErr
	}
	var b bytes.Buffer
	funcMap := template.FuncMap{"Description": name, "Name": description + "-job"}
	err := t.Execute(&b, &funcMap)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}

func testNewBasicAuthClient() *rundeck.Client {
	client, _ := rundeck.NewBasicAuthClient(testIntegrationUsername, testIntegrationPassword, testIntegrationURL)
	return client
}

func testNewTokenAuthClient() *rundeck.Client {
	client, _ := rundeck.NewTokenAuthClient(testIntegrationToken, testIntegrationURL)
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
	client, err := rundeck.NewBasicAuthClient(testIntegrationUsername, testIntegrationPassword, testIntegrationURL)
	assert.NoError(t, err)
	info, infoErr := client.GetSystemInfo()
	assert.NoError(t, infoErr)
	assert.NotNil(t, info)
}

func TestIntegrationTokenAuth(t *testing.T) {
	if !testRundeckRunning() {
		t.Skip("rundeck not running for integration tests")
	}
	client, err := rundeck.NewTokenAuthClient(testIntegrationToken, testIntegrationURL)
	assert.NoError(t, err)
	info, infoErr := client.GetSystemInfo()
	assert.NoError(t, infoErr)
	assert.NotNil(t, info)
}

func TestIntegrationInvalidBasicAuth(t *testing.T) {
	if !testRundeckRunning() {
		t.Skip("rundeck not running for integration tests")
	}
	client, _ := rundeck.NewBasicAuthClient("bob", "bob", testIntegrationURL)
	info, infoErr := client.GetSystemInfo()
	assert.Error(t, infoErr)
	assert.Nil(t, info)
}
