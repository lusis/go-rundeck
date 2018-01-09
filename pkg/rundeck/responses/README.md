# `responses` package

This is the package that describes all responses from the rundeck API

## Requirements

- All responses must use the following format

```go
type RundeckAPIActionResponse struct {}
```

GetProjectSCMConfigResponse is the response for the GetProjectSCMConfig api call

```go
type GetProjectSCMConfigResponse struct {
    SCMResponse
    Config      map[string]string `json:"config"`
    Enabled     bool              `json:"enabled"`
    Integration string            `json:"integration"`
    Project     string            `json:"project"`
    Type        string            `json:"type"`
}
```

- All responses must satisfy the `VersionedResponse` interface

This is critical to being able to easily manage responses across differing versions of the API

This can be done either through embedding if it makes sense (as with the `SCMResponse` above) or per-response if the response has different version constraints than the top-level construct.

- All responses should not use nested structs unless it makes sense

If there is a nested struct, chances are it's reused across other responses of the same type so to avoid duplication and copy/paste errors, break those out into their own struct

Nested structs aren't forbidden but use some good sense especially if they're complicated

- All responses should provide constants for test data in the form of `StructNameResponseTestData`

```go
// ListKeysResourceResponseTestFile is the test data for a KeyMetaResponse
const ListKeysResourceResponseTestFile = "key_metadata.json"
```

The content should be gathered from an actual live running rundeck server (you can use `rundeck http get`), and saved in the `testdata` directory. The value of the constant should be the name of the saved file.

After saving the file, `make bindata` should be called from the top-level of the repo to ensure the assets are available.

- All responses should provide be tested via strict mapstructure decoding

Example:

```go
func TestHistoryResponse(t *testing.T) {
    obj := &HistoryResponse{}
    data, dataErr := testdata.GetBytes(HistoryResponseTestFile)
    if dataErr != nil {
        t.Fatalf(dataErr.Error())
    }
    // unmarshal into a map[string]interface
    placeholder := make(map[string]interface{})
    _ = json.Unmarshal(data, &placeholder)
    // create a new decoder config and set the result to our instance of type
    config := newMSDecoderConfig()
    config.Result = obj
    decoder, newErr := mapstructure.NewDecoder(config)
    assert.NoError(t, newErr)
    // attempt to decode our unmarshalled data
    dErr := decoder.Decode(placeholder)
    assert.NoError(t, dErr)
}
```

If you get any failures from decoding, your struct definition is wrong for handling the json:

```text
--- FAIL: TestHistoryResponse (0.00s)
    assertions.go:237:

    Error Trace:    history_test.go:26

    Error:        Received unexpected error 6 error(s) decoding:

            * 'events[0]' has invalid keys: statusString
            * 'events[1]' has invalid keys: statusString
            * 'events[2]' has invalid keys: statusString
            * 'events[3]' has invalid keys: job, statusString
            * 'events[4]' has invalid keys: job, statusString
            * 'events[5]' has invalid keys: job, statusString
```

In this case we had some undocumented fields but using output from a live server and attempting to decode it caught that.

Note that we don't need to test the values of individual fields.
**Response tests are ONLY concerned with struct definitions matching real world data**