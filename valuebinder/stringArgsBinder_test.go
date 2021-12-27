package valuebinder

import (
	"reflect"
	"testing"
	"time"

	"github.com/structproto/types"
)

func TestStringArgValueBinder_WithBool(t *testing.T) {
	var target bool = false
	var input = "true"

	rv := reflect.ValueOf(&target).Elem()
	binder := StringArgsBinder(rv)
	err := binder.Bind(input)
	if err != nil {
		t.Error(err)
	}
	if target != true {
		t.Errorf("assert 'target':: expected '%#v', got '%#v'", true, target)
	}
}

func TestStringArgValueBinder_WithInt(t *testing.T) {
	var target int = 0
	var input = "1"

	rv := reflect.ValueOf(&target).Elem()
	binder := StringArgsBinder(rv)
	err := binder.Bind(input)
	if err != nil {
		t.Error(err)
	}
	if target != 1 {
		t.Errorf("assert 'target':: expected '%#v', got '%#v'", 1, target)
	}
}

func TestStringArgValueBinder_WithString(t *testing.T) {
	var target string = ""
	var input = "foo"

	rv := reflect.ValueOf(&target).Elem()
	binder := StringArgsBinder(rv)
	binder.Bind(input)
	if target != "foo" {
		t.Errorf("assert 'target':: expected '%#v', got '%#v'", "foo", target)
	}
}

func TestStringArgValueBinder_WithDuration(t *testing.T) {
	var target time.Duration
	var input = "1m2s"

	rv := reflect.ValueOf(&target).Elem()
	binder := StringArgsBinder(rv)
	err := binder.Bind(input)
	if err != nil {
		t.Error(err)
	}

	expected, _ := time.ParseDuration("1m2s")
	if target != expected {
		t.Errorf("assert 'target':: expected '%#v', got '%#v'", expected, target)
	}
}

func TestStringArgValueBinder_WithStringArray(t *testing.T) {
	var target []string
	var input = "alice,bob,carlo,david,frank,george"

	rv := reflect.ValueOf(&target).Elem()
	binder := StringArgsBinder(rv)
	err := binder.Bind(input)
	if err != nil {
		t.Error(err)
	}

	expected := []string{"alice", "bob", "carlo", "david", "frank", "george"}
	if !reflect.DeepEqual(target, expected) {
		t.Errorf("assert 'target':: expected '%#v', got '%#v'", expected, target)
	}
}

func TestStringArgValueBinder_WithIntArray(t *testing.T) {
	var target []int
	var input = "1,1,2,3,5,8,13"

	rv := reflect.ValueOf(&target).Elem()
	binder := StringArgsBinder(rv)
	err := binder.Bind(input)
	if err != nil {
		t.Error(err)
	}

	expected := []int{1, 1, 2, 3, 5, 8, 13}
	if !reflect.DeepEqual(target, expected) {
		t.Errorf("assert 'target':: expected '%#v', got '%#v'", expected, target)
	}
}

func TestStringArgValueBinder_WithRawContent(t *testing.T) {
	var target types.RawContent
	var input = "binary content"

	rv := reflect.ValueOf(&target).Elem()
	binder := StringArgsBinder(rv)
	err := binder.Bind(input)
	if err != nil {
		t.Error(err)
	}

	expected := types.RawContent("binary content")
	if !reflect.DeepEqual(target, expected) {
		t.Errorf("assert 'target':: expected '%#v', got '%#v'", expected, target)
	}
}
