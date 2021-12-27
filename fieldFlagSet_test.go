package structproto

import (
	"reflect"
	"testing"
)

func TestFieldFlagSet_WithEmptySet(t *testing.T) {
	var set FieldFlagSet

	// test IsEmpty()
	{
		if set.IsEmpty() != true {
			t.Errorf("assert 'IsEmpty()':: expected '%+v', got '%+v'", true, set.IsEmpty())
		}
	}
	// test Len
	{
		if set.Len() != 0 {
			t.Errorf("assert 'Len()':: expected '%+v', got '%+v'", 0, set.Len())
		}
	}
	// test Get
	{
		v, found := set.Get(0)
		if found != false {
			t.Errorf("assert found of 'Get(0)':: expected '%+v', got '%+v'", false, found)
		}
		if v != "" {
			t.Errorf("assert value of 'Get(0)':: expected '%+v', got '%+v'", "", v)
		}
	}
	// test IndexOf
	{
		index := set.IndexOf("unknown")
		if index != -1 {
			t.Errorf("assert 'IndexOf()':: expected '%+v', got '%+v'", -1, index)
		}
	}
	// test Has
	{
		existed := set.Has("unknown")
		if existed != false {
			t.Errorf("assert 'Has()':: expected '%+v', got '%+v'", false, existed)
		}
	}
	// test Clone
	{
		Cloned := set.Clone()
		if !reflect.DeepEqual(*Cloned, set) {
			t.Errorf("assert 'Clone()':: expected '%#v', got '%#v'", set, *Cloned)
		}
	}
	// test Append
	{
		set.Append("bob", "alice")
		expected := []string{"alice", "bob"}
		if !reflect.DeepEqual(expected, []string(set)) {
			t.Errorf("assert 'Append()':: expected '%#v', got '%#v'", expected, []string(set))
		}
	}
}

func TestFieldFlagSet_WithSampleSet(t *testing.T) {
	var set FieldFlagSet = []string{"bob", "georgy"}

	// test IsEmpty()
	{
		if set.IsEmpty() != false {
			t.Errorf("assert 'IsEmpty()':: expected '%+v', got '%+v'", false, set.IsEmpty())
		}
	}
	// test count
	{
		if set.Len() != 2 {
			t.Errorf("assert 'Len()':: expected '%+v', got '%+v'", 2, set.Len())
		}
	}
	// test Get
	{
		v, found := set.Get(0)
		if found != true {
			t.Errorf("assert found of 'Get(0)':: expected '%+v', got '%+v'", true, found)
		}
		if v != "bob" {
			t.Errorf("assert value of 'Get(0)':: expected '%+v', got '%+v'", "bob", v)
		}
	}
	// test indexOf
	{
		index := set.IndexOf("georgy")
		if index != 1 {
			t.Errorf("assert 'IndexOf()':: expected '%+v', got '%+v'", 1, index)
		}
	}
	// test contains
	{
		existed := set.Has("bob")
		if existed != true {
			t.Errorf("assert 'Has()':: expected '%+v', got '%+v'", true, existed)
		}
	}
	// test Clone
	{
		Cloned := set.Clone()
		if !reflect.DeepEqual(*Cloned, set) {
			t.Errorf("assert 'Clone()':: expected '%#v', got '%#v'", set, *Cloned)
		}
	}
	// test Append
	{
		set.Append("bob", "alice")
		expected := []string{"alice", "bob", "georgy"}
		if !reflect.DeepEqual(expected, []string(set)) {
			t.Errorf("assert 'Append()':: expected '%#v', got '%#v'", expected, []string(set))
		}
	}
	// test RemoveIndex
	{
		found, value := set.RemoveIndex(8)
		if found != false {
			t.Errorf("assert found of 'RemoveIndex()':: expected '%#v', got '%#v'", false, found)
		}
		if value != "" {
			t.Errorf("assert value of 'RemoveIndex()':: expected '%#v', got '%#v'", "", value)
		}
		expected := []string{"alice", "bob", "georgy"}
		if !reflect.DeepEqual(expected, []string(set)) {
			t.Errorf("assert 'stringSortedSet':: expected '%#v', got '%#v'", expected, []string(set))
		}
	}
	// test RemoveIndex
	{
		found, value := set.RemoveIndex(2)
		if found != true {
			t.Errorf("assert found of 'RemoveIndex()':: expected '%#v', got '%#v'", true, found)
		}
		if value != "georgy" {
			t.Errorf("assert value of 'RemoveIndex()':: expected '%#v', got '%#v'", "georgy", value)
		}
		expected := []string{"alice", "bob"}
		if !reflect.DeepEqual(expected, []string(set)) {
			t.Errorf("assert 'stringSortedSet':: expected '%#v', got '%#v'", expected, []string(set))
		}
	}
}
