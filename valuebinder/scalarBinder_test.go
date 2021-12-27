package valuebinder

import (
	"reflect"
	"testing"
	"time"
)

func TestScalarValueBinder_WithBool(t *testing.T) {
	var target bool = false
	var input = []byte("true")

	rv := reflect.ValueOf(&target).Elem()
	binder := ScalarBinder(rv)
	binder.Bind(input)
	if target != true {
		t.Errorf("assert 'target':: expected '%#v', got '%#v'", true, target)
	}
}

func TestScalarValueBinder_WithInt(t *testing.T) {
	var target int = 0
	var input = []byte("1")

	rv := reflect.ValueOf(&target).Elem()
	binder := ScalarBinder(rv)
	binder.Bind(input)
	if target != 1 {
		t.Errorf("assert 'target':: expected '%#v', got '%#v'", 1, target)
	}
}

func TestScalarValueBinder_WithString(t *testing.T) {
	var target string = ""
	var input = []byte("foo")

	rv := reflect.ValueOf(&target).Elem()
	binder := ScalarBinder(rv)
	binder.Bind(input)
	if target != "foo" {
		t.Errorf("assert 'target':: expected '%#v', got '%#v'", "foo", target)
	}
}

func TestScalarValueBinder_WithDuration(t *testing.T) {
	var target time.Duration
	var input = "1m2s"

	rv := reflect.ValueOf(&target).Elem()
	binder := ScalarBinder(rv)
	err := binder.Bind(input)
	if err != nil {
		t.Error(err)
	}

	expected, _ := time.ParseDuration("1m2s")
	if target != expected {
		t.Errorf("assert 'target':: expected '%#v', got '%#v'", expected, target)
	}
}
