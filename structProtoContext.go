package structproto

import "reflect"

type StructProtoContext Struct

func (ctx *StructProtoContext) Target() reflect.Value {
	return ctx.target
}

func (ctx *StructProtoContext) FieldInfo(name string) FieldInfo {
	if field, ok := ctx.fields[name]; ok {
		return field
	}
	return nil
}

func (ctx *StructProtoContext) Field(name string) (v reflect.Value, ok bool) {
	info := ctx.FieldInfo(name)
	if info != nil {
		return ctx.target.Field(info.Index()), true
	}
	return reflect.Value{}, false
}

func (ctx *StructProtoContext) FieldNames() []string {
	var fields []string = make([]string, len(ctx.fields))
	for _, v := range ctx.fields {
		fields[v.index] = v.name
	}
	return fields
}

func (ctx *StructProtoContext) RequiredFields() []string {
	return ctx.requiredFields
}

func (ctx *StructProtoContext) IsRequired(name string) bool {
	field := ctx.FieldInfo(name)
	if field != nil {
		return field.HasFlag(RequiredFlag)
	}
	return false
}

func (ctx *StructProtoContext) CheckIfMissingRequiredFields(visitFieldProc func() <-chan string) error {
	if ctx.requiredFields.IsEmpty() {
		return nil
	}

	var requiredFields = ctx.requiredFields.Clone()

	for field := range visitFieldProc() {
		index := requiredFields.IndexOf(field)
		if index != -1 {
			requiredFields.RemoveIndex(index)
		}

		// break loop if no more required fields
		if requiredFields.IsEmpty() {
			return nil
		}
	}

	if !requiredFields.IsEmpty() {
		field, _ := requiredFields.Get(0)
		return &MissingRequiredFieldError{field, nil}
	}
	return nil
}

func buildStructProtoContext(prototype *Struct) (*StructProtoContext, error) {
	context := StructProtoContext(*prototype)
	return &context, nil
}
