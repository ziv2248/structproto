package structproto

import (
	"reflect"

	"github.com/structproto/internal"
)

const (
	RequiredFlag = internal.RequiredFlag
)

type (
	ValueBindProvider = internal.ValueBindProvider
	ValueBinder       = internal.ValueBinder
	TagResolver       = internal.TagResolver
	Tag               = internal.Tag

	FieldValueEntry struct {
		Field string
		Value interface{}
	}

	FieldValueCollectionIterator interface {
		Iterate() <-chan FieldValueEntry
	}

	FieldInfo interface {
		Name() string
		Desc() string
		Index() int
		Flags() []string
		HasFlag(v string) bool
		Tag() reflect.StructTag
	}

	StructBinder interface {
		Init(context *StructProtoContext) error
		Bind(field FieldInfo, rv reflect.Value) error
		Deinit(context *StructProtoContext) error
	}

	StructProtoResolveOption struct {
		TagName     string
		TagResolver TagResolver
	}
)
