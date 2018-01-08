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

- All responses should provide a `FromBytes` and `FromReader` function that can be used in testing.

This is not neccessary for "fractional" responses that were created as part of avoiding nested structs mentioned above

```go
// FromReader returns a GetProjectSCMConfigResponse from an io.Reader
func (a *GetProjectSCMConfigResponse) FromReader(i io.Reader) error {
    b, err := ioutil.ReadAll(i)
    if err != nil {
        return err
    }
    return json.Unmarshal(b, a)
}

// FromBytes returns a GetProjectSCMConfigResponse from a byte slice
func (a *GetProjectSCMConfigResponse) FromBytes(f []byte) error {
    file := bytes.NewReader(f)
    return a.FromReader(file)
}
```