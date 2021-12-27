package reflectutil
import (
	"reflect"
)

func AssignZero(rv reflect.Value) reflect.Value {
	if rv.Kind() == reflect.Ptr {
		if rv.IsNil() {
			rv.Set(reflect.New(rv.Type().Elem()))
		}
	}
	return rv
}

func Indirect(value interface{}) interface{} {
	for {
		val := reflect.ValueOf(value)
		if val.Kind() != reflect.Ptr {
			// Value is not a pointer.
			return value
		}

		res := reflect.Indirect(val)
		if !res.IsValid() || !res.CanInterface() {
			// Invalid value or can't be returned as interface{}.
			return value
		}

		// Test for a circular type.
		if res.Kind() == reflect.Ptr && val.Pointer() == res.Pointer() {
			return value
		}

		// Next round.
		value = res.Interface()
	}
}
