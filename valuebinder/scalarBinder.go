package valuebinder

import (
	"reflect"

	"github.com/cstockton/go-conv"
	"github.com/structproto/internal"
	"github.com/structproto/util/reflectutil"
	"github.com/structproto/valuebinder/converter"
)

var (
	_ internal.ValueBindProvider = BuildScalarBinder
	_ internal.ValueBinder       = new(ScalarBinder)
)

type ScalarBinder reflect.Value

func (binder ScalarBinder) Bind(v interface{}) error {
	rv := reflect.Value(binder)
	rv = reflect.Indirect(reflectutil.AssignZero(rv))
	var err error

	switch rv.Kind() {
	case reflect.Struct:
		switch rv.Type() {
		case typeOfUrl:
			url, err := converter.Url(v)
			if err != nil {
				return &ValueBindingError{v, rv.Type().String(), err}
			}
			rv.Set(reflect.ValueOf(url))
		case typeOfTime:
			time, err := conv.Time(v)
			if err != nil {
				return &ValueBindingError{v, rv.Type().String(), err}
			}
			rv.Set(reflect.ValueOf(time))
		default:
			return &ValueBindingError{v, rv.Type().String(), err}
		}
	case reflect.Bool:
		bool, err := conv.Bool(v)
		if err != nil {
			return &ValueBindingError{v, rv.Kind().String(), err}
		}
		rv.SetBool(bool)
	case reflect.String:
		string, err := conv.String(v)
		if err != nil {
			return &ValueBindingError{v, rv.Kind().String(), err}
		}
		rv.SetString(string)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		switch rv.Type() {
		case typeOfDuration:
			duration, err := conv.Duration(v)
			if err != nil {
				return &ValueBindingError{v, rv.Type().String(), err}
			}
			rv.Set(reflect.ValueOf(duration))
		default:
			int, err := conv.Int64(v)
			if err != nil {
				return &ValueBindingError{v, rv.Kind().String(), err}
			}
			rv.SetInt(int)
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		uint, err := conv.Uint64(v)
		if err != nil {
			return &ValueBindingError{v, rv.Kind().String(), err}
		}
		rv.SetUint(uint)
	case reflect.Float32, reflect.Float64:
		float, err := conv.Float64(v)
		if err != nil {
			return &ValueBindingError{v, rv.Kind().String(), err}
		}
		rv.SetFloat(float)
	default:
		return &ValueBindingError{v, rv.Kind().String(), err}
	}
	return err
}

func BuildScalarBinder(rv reflect.Value) internal.ValueBinder {
	return ScalarBinder(rv)
}
