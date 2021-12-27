package internal

import "reflect"

const (
	RequiredFlag = "required"
)

type (
	Tag struct {
		Name  string
		Flags []string
		Desc  string
	}

	ValueBinder interface {
		Bind(v interface{}) error
	}

	TagResolver       func(fieldname, token string) (*Tag, error)
	ValueBindProvider func(rv reflect.Value) ValueBinder
)
