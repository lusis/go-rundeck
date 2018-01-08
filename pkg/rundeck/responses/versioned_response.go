package responses

// VersionedResponse is an interface for a Rundeck Response that supports versioning information
type VersionedResponse interface {
	minVersion() int
	maxVersion() int
	deprecated() bool
}

// AbsoluteMinimumVersion is the absolute minimum version this library will support
// We set this to `14` as that was the first version of the rundeck API to support JSON
const AbsoluteMinimumVersion = 14

// CurrentVersion is the current version of the API that this library is tested against
const CurrentVersion = 21

// GetMinVersionFor gets the minimum api version required for a response
func GetMinVersionFor(a VersionedResponse) int { return a.minVersion() }

// GetMaxVersionFor gets the maximum api version required for a response
func GetMaxVersionFor(a VersionedResponse) int { return a.maxVersion() }

// IsDeprecated indicates if a response is deprecated or not
func IsDeprecated(a VersionedResponse) bool { return a.deprecated() }
