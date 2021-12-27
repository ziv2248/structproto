package valuebinder

import (
	"reflect"
	"testing"
	"time"
)

func TestByteArgsValueBinder_WithBool(t *testing.T) {
	var v bool = false
	var input = []byte("true")

	rv := reflect.ValueOf(&v).Elem()
	binder := BytesArgsBinder(rv)
	binder.Bind(input)
	if v != true {
		t.Errorf("assert 'v':: expected '%#v', got '%#v'", true, v)
	}
}

func TestByteArgsValueBinder_WithInt(t *testing.T) {
	var v int = 0
	var input = []byte("1")

	rv := reflect.ValueOf(&v).Elem()
	binder := BytesArgsBinder(rv)
	binder.Bind(input)
	if v != 1 {
		t.Errorf("assert 'v':: expected '%#v', got '%#v'", 1, v)
	}
}

func TestByteArgsValueBinder_WithString(t *testing.T) {
	var v string = ""
	var input = []byte("foo")

	rv := reflect.ValueOf(&v).Elem()
	binder := BytesArgsBinder(rv)
	binder.Bind(input)
	if v != "foo" {
		t.Errorf("assert 'v':: expected '%#v', got '%#v'", "foo", v)
	}
}

func TestByteArgsValueBinder_WithDuration(t *testing.T) {
	var v time.Duration
	var input = []byte("1m2s")

	rv := reflect.ValueOf(&v).Elem()
	binder := BytesArgsBinder(rv)
	err := binder.Bind(input)
	if err != nil {
		t.Error(err)
	}

	expected, _ := time.ParseDuration("1m2s")
	if v != expected {
		t.Errorf("assert 'v':: expected '%#v', got '%#v'", expected, v)
	}
}

func TestByteArgsValueBinder_WithStringArray(t *testing.T) {
	var v []string
	var input = []byte("alice,bob,carlo,david,frank,george")

	rv := reflect.ValueOf(&v).Elem()
	binder := BytesArgsBinder(rv)
	err := binder.Bind(input)
	if err != nil {
		t.Error(err)
	}

	expected := []string{"alice", "bob", "carlo", "david", "frank", "george"}
	if !reflect.DeepEqual(v, expected) {
		t.Errorf("assert 'v':: expected '%#v', got '%#v'", expected, v)
	}
}

func TestByteArgsValueBinder_WithIntArray(t *testing.T) {
	var v []int
	var input = []byte("1,1,2,3,5,8,13")

	rv := reflect.ValueOf(&v).Elem()
	binder := BytesArgsBinder(rv)
	err := binder.Bind(input)
	if err != nil {
		t.Error(err)
	}

	expected := []int{1, 1, 2, 3, 5, 8, 13}
	if !reflect.DeepEqual(v, expected) {
		t.Errorf("assert 'v':: expected '%#v', got '%#v'", expected, v)
	}
}
