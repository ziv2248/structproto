package structproto

import "reflect"

var _ FieldInfo = new(Field)

type Field struct {
	name  string
	desc  string
	index int
	flags FieldFlagSet
	tag   reflect.StructTag
}

func (f *Field) Name() string {
	return f.name
}

func (f *Field) Desc() string {
	return f.desc
}

func (f *Field) Index() int {
	return f.index
}

func (f *Field) Flags() []string {
	return f.flags
}

func (f *Field) HasFlag(v string) bool {
	return f.flags.Has(v)
}

func (f *Field) Tag() reflect.StructTag {
	return f.tag
}

func (f *Field) appendFlags(values ...string) {
	if len(values) > 0 {
		for _, v := range values {
			if len(v) == 0 {
				continue
			}
			f.flags.Append(v)
		}
	}
}
