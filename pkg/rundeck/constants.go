package rundeck

import "errors"

// defaultNodeFilter is the default filter to use when things require a node filter string
const defaultNodeFilter = "name: .*"

// MaxRundeckVersion is the maximum version of the api this library supports
// can be overridden
// TODO: make this a min/max option and validate
const MaxRundeckVersion = "21"

// minimum version of rundeck api version that supports json
const minJSONSupportedAPIVersion = 14

const (
	basicAuthType = "basic"
	tokenAuthType = "token"

	/*
		scmStateClean   = "CLEAN"          // nolint: deadcode
		scmStateUnknown = "UNKNOWN"        // nolint: deadcode
		scmStateRefresh = "REFRESH_NEEDED" // nolint: deadcode
		scmStateImport  = "IMPORT_NEEDED"  // nolint: deadcode
		scmStateDelete  = "DELETE_NEEDED"  // nolint: deadcode
		scmStateExport  = "EXPORT_NEEDED"  // nolint: deadcode
		scmStateCreate  = "CREATE_NEEDED"  // nolint: deadcode
	*/
)

var (
	errInvalidUsernamePassword = errors.New("Invalid username or password returned by rundeck")

	// ErrInvalidRundeckURL is the error for an invalid rundeck url
	ErrInvalidRundeckURL = errors.New("Invalid Rundeck URL")

	// ErrAuthFailed is the error for a auth failure in an api call
	// this is slightly different the ErrInvalidUsernamePassword
	// as this means auth succeeded with basic auth but a 401 could be returned farther down
	ErrAuthFailed = errors.New("API call failed due to authentication failure")

	// ErrMissingResource is the error type for 404 not found
	ErrMissingResource = errors.New("Rundeck could not find the resource you requested")

	// ErrResourceConflict is the error type for 409 responses
	ErrResourceConflict = errors.New("resource already exists on the rundeck server")

	errDecoding   = errors.New("Could not parse response from the Rundeck server")
	errEncoding   = errors.New("could not encode payload for rundeck server")
	errOption     = errors.New("Passed option returned an error")
	errValidation = errors.New("Error validating input")
)
