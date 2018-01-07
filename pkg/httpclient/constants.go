package httpclient

import "errors"

const (
	// ContentTypeJSON is the mimetype for json
	ContentTypeJSON = "application/json"
	// ContentTypeXML is the mimetype for xml
	ContentTypeXML = "application/xml"
	// DefaultAccept is the default Accept mimetype for requests
	DefaultAccept = "*/*"
)

var (
	// ErrInvalidStatusCode is the error type returned when the user sets expected
	// status code with `ExpectStatus`, but it does not match
	ErrInvalidStatusCode = errors.New("response had an invalid status code")
)
