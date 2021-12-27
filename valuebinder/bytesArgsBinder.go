package valuebinder

import (
	"fmt"
	"reflect"

	"github.com/structproto/internal"
)

var (
	_ internal.ValueBindProvider = BuildBytesArgsBinder
	_ internal.ValueBinder       = new(BytesArgsBinder)
)

type BytesArgsBinder reflect.Value

func (binder BytesArgsBinder) Bind(input interface{}) error {
	buf, ok := input.([]byte)
	if !ok {
		return fmt.Errorf("cannot bind type %T from input", input)
	}
	v := string(buf)

	return StringArgsBinder(reflect.Value(binder)).Bind(v)
}

func BuildBytesArgsBinder(rv reflect.Value) internal.ValueBinder {
	return BytesArgsBinder(rv)
}
