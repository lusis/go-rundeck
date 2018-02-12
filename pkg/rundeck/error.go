package rundeck

// UnmarshalError is a custom error type for decoding errors
type UnmarshalError struct {
	msg string
}

// Error returns the error message
func (e *UnmarshalError) Error() string {
	return e.msg
}

// MarshalError is a custom error type for decoding errors
type MarshalError struct {
	msg string
}

// Error returns the error message
func (e *MarshalError) Error() string {
	return e.msg
}

// OptionError is a custom error type for option errors
type OptionError struct {
	msg string
}

// Error returns the error message
func (e *OptionError) Error() string {
	return e.msg
}

// PolicyValidationError is a custom error type for policy validation errors
type PolicyValidationError struct {
	msg string
}

// Error returns the error message
func (e *PolicyValidationError) Error() string {
	return e.msg
}

// AuthError is a custom error type for decoding errors
type AuthError struct {
	msg string
}

// Error returns the error message
func (e *AuthError) Error() string {
	return e.msg
}

// SCMValidationError is a custom error type for scm validation errors
type SCMValidationError struct {
	msg string
}

// Error returns the error message
func (e *SCMValidationError) Error() string {
	return e.msg
}
