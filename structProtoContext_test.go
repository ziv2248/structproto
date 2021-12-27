package structproto

import (
	"reflect"
	"testing"
)

func TestStructProtoContext(t *testing.T) {
	c := mockCharacter{}
	prototype, err := Prototypify(&c, &StructProtoResolveOption{
		TagName: "demo",
	})
	if err != nil {
		t.Error(err)
	}

	context, err := buildStructProtoContext(prototype)
	if err != nil {
		t.Error(err)
	}

	expectedFieldNames := []string{"NAME", "AGE", "ALIAS", "DATE_OF_BIRTH", "REMARK", "NUMBERS"}
	if !reflect.DeepEqual(expectedFieldNames, context.FieldNames()) {
		t.Errorf("assert 'structprotoContext.AllFields()':: expected '%#v', got '%#v'", expectedFieldNames, context.FieldNames())
	}
	expectedRequiredFields := []string{"AGE", "NAME"}
	if !reflect.DeepEqual(expectedRequiredFields, context.RequiredFields()) {
		t.Errorf("assert 'structprotoContext.AllRequiredFields()':: expected '%#v', got '%#v'", expectedRequiredFields, context.RequiredFields())
	}

	{
		field := context.FieldInfo("NAME")
		if field == nil {
			t.Errorf("assert 'structprotoContext.Field(\"NAME\")':: expected not nil, got '%#v'", field)
		}
		if field.Name() != "NAME" {
			t.Errorf("assert 'structprotoField.Name()':: expected '%#v', got '%#v'", "NAME", field.Name())
		}
		if field.Index() != 0 {
			t.Errorf("assert 'structprotoField.Index()':: expected '%#v', got '%#v'", "NAME", field.Name())
		}
		expectedFlags := []string{"required"}
		if !reflect.DeepEqual(expectedFlags, field.Flags()) {
			t.Errorf("assert 'structprotoField.Flags()':: expected '%#v', got '%#v'", expectedFlags, field.Flags())
		}
	}

	if !context.IsRequired("NAME") {
		t.Errorf("assert 'structprotoContext.IsRequiredField(\"NAME\")':: expected '%#v', got '%#v'", expectedRequiredFields, context.IsRequired("NAME"))
	}
	if context.IsRequired("unknown") {
		t.Errorf("assert 'structprotoContext.IsRequiredField(\"unknown\")':: expected '%#v', got '%#v'", expectedRequiredFields, context.IsRequired("unknown"))
	}

	// TODO: test context.ChechIfMissingRequireFields
}
