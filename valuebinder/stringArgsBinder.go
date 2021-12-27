package valuebinder

import (
	"fmt"
	"net/url"
	"reflect"
	"strconv"
	"strings"

	"github.com/cstockton/go-conv"
	"github.com/structproto/internal"
	"github.com/structproto/types"
	"github.com/structproto/util/reflectutil"
)

var (
	_ internal.ValueBindProvider = BuildStringArgsBinder
	_ internal.ValueBinder       = new(StringArgsBinder)
)

type StringArgsBinder reflect.Value

func (binder StringArgsBinder) Bind(input interface{}) error {
	v, ok := input.(string)
	if !ok {
		return fmt.Errorf("cannot bind type %T from input", input)
	}
	rv := reflect.Value(binder)
	return binder.bindStringValueImpl(rv, v)
}

func (binder StringArgsBinder) bindStringValueImpl(rv reflect.Value, v string) error {
	rv = reflect.Indirect(reflectutil.AssignZero(rv))
	var err error

	switch rv.Kind() {
	case reflect.Struct:
		switch rv.Type() {
		case typeOfUrl:
			url, err := url.Parse(v)
			if err != nil {
				return &ValueBindingError{v, rv.Type().String(), err}
			}
			rv.Set(reflect.ValueOf(*url))
		case typeOfTime:
			time, err := conv.Time(v)
			if err != nil {
				return &ValueBindingError{v, rv.Type().String(), err}
			}
			rv.Set(reflect.ValueOf(time))
		default:
			return &ValueBindingError{v, rv.Type().String(), err}
		}
	case reflect.Array, reflect.Slice:
		switch rv.Type() {
		case typeOfRawContent:
			buf := []byte(v)
			rv.Set(reflect.ValueOf(types.RawContent(buf)))
		default:
			if len(v) > 0 {
				array := strings.Split(v, ",")
				size := len(array)
				container := reflect.MakeSlice(rv.Type(), size, size)
				for i, elem := range array {
					err := binder.bindStringValueImpl(container.Index(i), elem)
					if err != nil {
						return err
					}
				}
				rv.Set(container)
			}
		}
	case reflect.Bool:
		bool, err := strconv.ParseBool(v)
		if err != nil {
			return &ValueBindingError{v, rv.Kind().String(), err}
		}
		rv.SetBool(bool)
	case reflect.String:
		rv.SetString(v)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		switch rv.Type() {
		case typeOfDuration:
			duration, err := conv.Duration(v)
			if err != nil {
				return &ValueBindingError{v, rv.Type().String(), err}
			}
			rv.Set(reflect.ValueOf(duration))
		default:
			int, err := strconv.ParseInt(v, 10, 64)
			if err != nil {
				return &ValueBindingError{v, rv.Kind().String(), err}
			}
			rv.SetInt(int)
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		uint, err := strconv.ParseUint(v, 10, 64)
		if err != nil {
			return &ValueBindingError{v, rv.Kind().String(), err}
		}
		rv.SetUint(uint)
	case reflect.Float32, reflect.Float64:
		float, err := strconv.ParseFloat(v, 64)
		if err != nil {
			return &ValueBindingError{v, rv.Kind().String(), err}
		}
		rv.SetFloat(float)
	default:
		return &ValueBindingError{v, rv.Kind().String(), err}
	}
	return err
}

func BuildStringArgsBinder(rv reflect.Value) internal.ValueBinder {
	return StringArgsBinder(rv)
}
