package structproto

import "fmt"

// A MissingRequiredFieldError represents an error when the required fields
// cannot be binded or missing.
type MissingRequiredFieldError struct {
	Field string
	Err   error
}

func (e *MissingRequiredFieldError) Error() string {
	return fmt.Sprintf("missing required symbol '%s'", e.Field)
}

// Unwrap returns the underlying error.
func (e *MissingRequiredFieldError) Unwrap() error {
	return e.Err
}
