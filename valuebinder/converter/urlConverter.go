package converter

import (
	"net/url"
	"reflect"

	reflectutil "github.com/structproto/util/reflectutil"
)

var (
	emptyUrl  = url.URL{}
	typeOfUrl = reflect.TypeOf(emptyUrl)
)

func Url(from interface{}) (url.URL, error) {

	if T, ok := from.(url.URL); ok {
		return T, nil
	} else if T, ok := from.(*url.URL); ok {
		return *T, nil
	}

	rv := reflect.ValueOf(reflectutil.Indirect(from))
	switch rv.Kind() {
	case reflect.String:
		return convStringToUrl(rv.String())
	case reflect.Struct:
		if rv.Type().ConvertibleTo(typeOfUrl) {
			valueConv := rv.Convert(typeOfUrl)
			if valueConv.CanInterface() {
				return valueConv.Interface().(url.URL), nil
			}
		}
	}
	return emptyUrl, newConvErr(from, "url.URL")
}

func convStringToUrl(value string) (url.URL, error) {
	T, err := url.Parse(value)
	return *T, err
}
