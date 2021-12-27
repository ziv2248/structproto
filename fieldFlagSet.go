package structproto

import (
	"reflect"
	"sort"
)

var (
	emptyFieldFlagSet = FieldFlagSet(nil)
)

type FieldFlagSet []string

func (s *FieldFlagSet) Append(values ...string) {
	set := *s
	if len(values) > 0 {
		for _, v := range values {
			i := sort.SearchStrings(set, v)
			var existed = false
			if i < len(set) {
				existed = set[i] == v
			}
			if !existed {
				container := set
				if i < len(container) {
					container = append(container, "")
					copy(container[i+1:], container[i:])
					container[i] = v
				} else {
					container = append(container, v)
				}
				set = container
				*s = set
			}
		}
	}
}

func (s *FieldFlagSet) Clone() *FieldFlagSet {
	set := *s
	if !reflect.ValueOf(set).IsZero() {
		var container = make([]string, len(set))
		copy(container, set)
		cloned := FieldFlagSet(container)
		return &cloned
	}
	return &emptyFieldFlagSet
}

func (s *FieldFlagSet) Get(index int) (string, bool) {
	if !s.IsEmpty() {
		set := *s
		if index >= 0 && index < len(set) {
			return set[index], true
		}
	}
	return "", false
}

func (s *FieldFlagSet) Has(v string) bool {
	if s.IsEmpty() {
		return false
	}

	return -1 != s.IndexOf(v)
}

func (s *FieldFlagSet) IndexOf(v string) int {
	if !s.IsEmpty() {
		set := *s
		if len(set) > 0 {
			i := sort.SearchStrings(set, v)
			if i < len(set) {
				if set[i] == v {
					return i
				}
			}
		}
	}
	return -1
}

func (s *FieldFlagSet) IsEmpty() bool {
	return len(*s) == 0
}

func (s *FieldFlagSet) Len() int {
	if s.IsEmpty() {
		return 0
	}

	return len(*s)
}

func (s FieldFlagSet) Iterate() <-chan string {
	if !s.IsEmpty() {
		c := make(chan string, 1)
		go func() {
			for _, v := range s {
				c <- v
			}
			close(c)
		}()
		return c
	}
	return nil
}

func (s *FieldFlagSet) Remove(v string) bool {
	if !s.IsEmpty() {
		index := s.IndexOf(v)
		deleted, _ := s.RemoveIndex(index)
		return deleted
	}
	return false
}

func (s *FieldFlagSet) RemoveIndex(index int) (bool, string) {
	if !s.IsEmpty() {
		set := *s
		if index >= 0 && index < len(set) {
			value := set[index]
			copy(set[index:], set[index+1:])
			set = set[:len(set)-1]
			*s = set

			return true, value
		}
	}
	return false, ""
}
