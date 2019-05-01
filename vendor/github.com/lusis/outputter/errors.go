package outputter

import (
	"errors"
)

// ErrorHeadersAlreadyAdded is an error for already having headers set
var ErrorHeadersAlreadyAdded = errors.New("Headers already set. Cannot overwrite")

// ErrorOutputAddRowNoHeaders is an error for not having yet called SetHeaders()
var ErrorOutputAddRowNoHeaders = errors.New("Cannot AddRow with before calling SetHeaders")

// ErrorOutputAddRowTooFewHeaders is an error for having fewer keys than values
var ErrorOutputAddRowTooFewHeaders = errors.New("Cannot AddRow with more values than headers")

// ErrorUnknownOutputter is an error for specifying an unknown ErrorUnknownFormatter
var ErrorUnknownOutputter = errors.New("Unknown formatter specified")

// ErrorInvalidOutputter is an error for specifying that a registered output isn't an Ouputter
var ErrorInvalidOutputter = errors.New("Specified outputter is invalid")

// ErrorCannotChangeWriter is an error for specifying that an output does not allow changing the writer after data is populated
var ErrorCannotChangeWriter = errors.New("Must call SetWriter before any data is added")
