package structproto

import "fmt"

const (
	errStringValueLength = 24
)

type FieldBindingError struct {
	Field string
	Value interface{}
	Err   error
}

func (e *FieldBindingError) Error() string {
	if v, ok := e.Value.(string); ok && len(v) > errStringValueLength {
		return fmt.Sprintf("cannot bind field tag '%s' with value '%s...'", e.Field, v[:errStringValueLength])
	}
	return fmt.Sprintf("cannot bind field tag '%s' with value '%v'", e.Field, e.Value)
}

// Unwrap returns the underlying error.
func (e *FieldBindingError) Unwrap() error {
	return e.Err
}
