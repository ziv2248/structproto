package structproto

import (
	"reflect"
	"testing"
	"time"

	"github.com/structproto/valuebinder"
)

type mockCharacter struct {
	Name       string    `demo:"*NAME"`
	Age        *int      `demo:"*AGE"`
	Alias      []string  `demo:"ALIAS"`
	DatOfBirth time.Time `demo:"DATE_OF_BIRTH;the character's birth of date"`
	Remark     string    `demo:"REMARK,flag1,flag2,,;note the character's personal favor"`
	Numbers    []int     `demo:"NUMBERS"`
}

func TestResolveCharacterStruct(t *testing.T) {
	c := mockCharacter{}
	prototype, err := Prototypify(&c, &StructProtoResolveOption{
		TagName: "demo",
	})
	if err != nil {
		t.Error(err)
	}
	expectedFields := map[string]Field{
		"NAME": {
			name:  "NAME",
			desc:  "",
			index: 0,
			flags: []string{"required"},
			tag:   `demo:"*NAME"`,
		},
		"AGE": {
			name:  "AGE",
			desc:  "",
			index: 1,
			flags: []string{"required"},
			tag:   `demo:"*AGE"`,
		},
		"ALIAS": {
			name:  "ALIAS",
			desc:  "",
			index: 2,
			flags: []string(nil),
			tag:   `demo:"ALIAS"`,
		},
		"DATE_OF_BIRTH": {
			name:  "DATE_OF_BIRTH",
			desc:  "the character's birth of date",
			index: 3,
			flags: []string(nil),
			tag:   `demo:"DATE_OF_BIRTH;the character's birth of date"`,
		},
		"REMARK": {
			name:  "REMARK",
			desc:  "note the character's personal favor",
			index: 4,
			flags: []string{"flag1", "flag2"},
			tag:   `demo:"REMARK,flag1,flag2,,;note the character's personal favor"`,
		},
		"NUMBERS": {
			name:  "NUMBERS",
			desc:  "",
			index: 5,
			flags: []string(nil),
			tag:   `demo:"NUMBERS"`,
		},
	}

	if len(expectedFields) != len(prototype.fields) {
		t.Errorf("assert 'structproto.fields' length :: expected '%v', got '%v'", len(expectedFields), len(prototype.fields))
	}
	for k, v := range expectedFields {
		if f, ok := prototype.fields[k]; !ok {
			t.Errorf("assert 'structproto.fields' key '%s' not found", k)
		} else {
			if (f.Name() != v.name) ||
				(f.Index() != v.index) ||
				(f.Desc() != v.desc) ||
				(f.Tag() != v.tag) ||
				(!reflect.DeepEqual([]string(f.Flags()), []string(v.flags))) {
				t.Errorf("assert 'structproto.fields' key '%s' :: expected '%#v', got '%#v'", k, v, f)
			}
		}
	}
	expectedRequiredFields := FieldFlagSet([]string{"AGE", "NAME"})
	if !reflect.DeepEqual(expectedRequiredFields, prototype.requiredFields) {
		t.Errorf("assert 'mockCharacter.requiredFields':: expected '%#v', got '%#v'", expectedRequiredFields, prototype.requiredFields)
	}
}

func TestBindFields_MissingRequiredField(t *testing.T) {
	c := mockCharacter{}
	prototype, err := Prototypify(&c, &StructProtoResolveOption{
		TagName: "demo",
	})
	if err != nil {
		t.Error(err)
	}
	err = prototype.BindFields(map[string]interface{}{
		"NAME":          "luffy",
		"ALIAS":         "lucy",
		"DATE_OF_BIRTH": "2020-05-05T00:00:00Z",
	}, valuebinder.BuildStringArgsBinder)
	if err == nil {
		t.Errorf("the 'Mapper.Map()' should throw '%s' error", "with missing symbol 'AGE'")
	}
}

func TestBindValues_Well(t *testing.T) {
	c := mockCharacter{}
	prototype, err := Prototypify(&c, &StructProtoResolveOption{
		TagName: "demo",
	})
	if err != nil {
		t.Error(err)
	}
	err = prototype.BindValues(FieldValueMap{
		"NAME":          "luffy",
		"AGE":           "19",
		"ALIAS":         "lucy",
		"DATE_OF_BIRTH": "2020-05-05T00:00:00Z",
		"NUMBERS":       "5,12",
	}, valuebinder.BuildStringArgsBinder)
	if err != nil {
		t.Error(err)
	}

	if c.Name != "luffy" {
		t.Errorf("assert 'mockCharacter.Name':: expected '%v', got '%v'", "luffy", c.Name)
	}
	if *c.Age != 19 {
		t.Errorf("assert 'mockCharacter.Age':: expected '%v', got '%v'", 19, c.Age)
	}
	expectedAlias := []string{"lucy"}
	if !reflect.DeepEqual(c.Alias, expectedAlias) {
		t.Errorf("assert 'mockCharacter.Alias':: expected '%#v', got '%#v'", expectedAlias, c.Alias)
	}
	expectedDateOfBirth := time.Date(2020, 5, 5, 0, 0, 0, 0, time.UTC)
	if c.DatOfBirth != expectedDateOfBirth {
		t.Errorf("assert 'mockCharacter.DatOfBirth':: expected '%v', got '%v'", expectedDateOfBirth, c.DatOfBirth)
	}
	expectedNumbers := []int{5, 12}
	if !reflect.DeepEqual(c.Numbers, expectedNumbers) {
		t.Errorf("assert 'mockCharacter.Numbers':: expected '%#v', got '%#v'", expectedNumbers, c.Numbers)
	}
}

func TestBind_MissingRequiredField(t *testing.T) {
	input := map[string]string{
		"NAME": "luffy",
		// "AGE":           "19",    -- we won't set the required field
		"ALIAS":         "lucy",
		"DATE_OF_BIRTH": "2020-05-05T00:00:00Z",
	}
	binder := &mockMapStructBinder{
		values: input,
	}

	c := mockCharacter{}
	prototype, err := Prototypify(&c, &StructProtoResolveOption{
		TagName: "demo",
	})

	err = prototype.Bind(binder)
	if err == nil {
		t.Errorf("the 'Process()' should throw '%s' error", "missing required symbol 'AGE'")
	} else {
		missingRequiredFieldError, ok := err.(*MissingRequiredFieldError)
		if !ok {
			t.Errorf("the error expected '%T', got '%T'", &MissingRequiredFieldError{}, err)
		}
		if missingRequiredFieldError.Field != "AGE" {
			t.Errorf("assert 'MissingRequiredFieldError.Field':: expected '%v', got '%v'", "AGE", missingRequiredFieldError.Field)
		}
	}
}

func TestBind_Well(t *testing.T) {
	input := map[string]string{
		"NAME":          "luffy",
		"AGE":           "19",
		"ALIAS":         "lucy",
		"DATE_OF_BIRTH": "2020-05-05T00:00:00Z",
	}
	binder := &mockMapStructBinder{
		values: input,
	}

	c := mockCharacter{}
	prototype, err := Prototypify(&c, &StructProtoResolveOption{
		TagName: "demo",
	})
	err = prototype.Bind(binder)
	if err != nil {
		t.Error(err)
	}

	if c.Name != "luffy" {
		t.Errorf("assert 'mockCharacter.Name':: expected '%v', got '%v'", "luffy", c.Name)
	}
	if *c.Age != 19 {
		t.Errorf("assert 'mockCharacter.Age':: expected '%v', got '%v'", 19, c.Age)
	}
	expectedAlias := []string{"lucy"}
	if !reflect.DeepEqual(c.Alias, expectedAlias) {
		t.Errorf("assert 'mockCharacter.Alias':: expected '%#v', got '%#v'", expectedAlias, c.Alias)
	}
	expectedDateOfBirth := time.Date(2020, 5, 5, 0, 0, 0, 0, time.UTC)
	if c.DatOfBirth != expectedDateOfBirth {
		t.Errorf("assert 'mockCharacter.DatOfBirth':: expected '%v', got '%v'", expectedDateOfBirth, c.DatOfBirth)
	}
}

type mockMapStructBinder struct {
	values map[string]string
}

func (p *mockMapStructBinder) Init(context *StructProtoContext) error {
	return nil
}

func (p *mockMapStructBinder) Bind(field FieldInfo, rv reflect.Value) error {
	name := field.Name()
	if v, ok := p.values[name]; ok {
		return valuebinder.StringArgsBinder(rv).Bind(v)
	}
	return nil
}

func (p *mockMapStructBinder) Deinit(context *StructProtoContext) error {
	return context.CheckIfMissingRequiredFields(func() <-chan string {
		c := make(chan string)
		go func() {
			for k := range p.values {
				c <- k
			}
			close(c)
		}()
		return c
	})
}

func TestUrlTagResolver(t *testing.T) {
	v := mockUrlPathManager{}
	prototype, err := Prototypify(&v,
		&StructProtoResolveOption{
			TagName:     "url",
			TagResolver: resolveUrlTag,
		})
	if err != nil {
		t.Error(err)
	}

	err = prototype.BindValues(FieldValueMap{
		"/":     "root",
		"/Echo": "echo",
	}, valuebinder.BuildStringArgsBinder)
	if err != nil {
		t.Error(err)
	}

	if v.Root != "root" {
		t.Errorf("assert 'mockUrlPathManager.Root':: expected '%v', got '%v'", "root", v.Root)
	}
	if v.Echo != "echo" {
		t.Errorf("assert 'mockUrlPathManager.Echo':: expected '%v', got '%v'", "echo", v.Echo)
	}
}

type mockUrlPathManager struct {
	Root string `url:"/"`
	Echo string `url:"/Echo"`
}

func resolveUrlTag(fieldname, token string) (*Tag, error) {
	var tag *Tag
	if len(token) > 0 {
		if token != "-" {
			tag = &Tag{
				Name: token,
			}
		}
	}
	return tag, nil
}
