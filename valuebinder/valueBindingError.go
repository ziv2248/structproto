package valuebinder

import "fmt"

const (
	errStringValueLength = 24
)

// A ValueBindingError represents an error when value binding failure.
type ValueBindingError struct {
	Value interface{}
	Kind  string
	Err   error
}

func (e *ValueBindingError) Error() string {
	if v, ok := e.Value.(string); ok {
		if len(v) > errStringValueLength {
			return fmt.Sprintf("cannot bind type %s with value (type %[1]T) '%v'", e.Kind, v[:errStringValueLength])
		}
	}
	return fmt.Sprintf("cannot bind type %s with value (type %[1]T) '%v'", e.Kind, e.Value)
}

// Unwrap returns the underlying error.
func (e *ValueBindingError) Unwrap() error {
	return e.Err
}
