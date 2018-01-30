# Testing

We have two types of tests for the repo. One is automated, one is not

## Unit Tests

We have the standard unit tests built using stdlib testing and [assert](https://github.com/stretchr/testify/assert).

These are run by Travis on every commit/PR and can be called locally with `script/test`

## Integration Tests

These are not as full featured but they are being developed.
These require an instance of Rundeck and are configured to use a local rundeck server that can be built in this repo.

These are not run automatically right now.

### Setting up the integration environment

```text
make build-test-container
make run-test-container
<wait for rundeck to fully start>
```

### Running the integration tests

`script/test-integration`

These leverage build tags and all the integration test cases are stored in `<name>_integration_test.go` files in the `rundeck` package.

Any integration test should have the following at the top:

```go
// +build integration

package rundeck

/*
standard import statements and test functions
*/
```

Inside your test functions, you can call:

```go
client := testNewTokenAuthClient()
```

and operate as normal on the instance of `Client` returned

### Cleaning up

Because you're testing against an actual rundeck server you MUST clean up after yourself. These tests won't be able to run in parallel in most cases.

There's no need to create the same level of individual tests (i.e. a test for `CreateFoo` and a test for `DeleteFoo`).

Think of the tests as groups of thing. Take the following example where we we test all of the `*Token` functionality:

```go
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
```

In this test we're:

- creating a token for the current user
- getting all tokens for the current user
- ensuring our token is in the list
- cleaning up after ourselves by defering a `DeleteToken` call
