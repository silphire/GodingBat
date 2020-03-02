package map1

import (
	"reflect"
	"testing"
)

func TestMapShare(t *testing.T) {
	actual := mapShare(map[string]string{"a": "aaa", "b": "bbb", "c": "ccc"})
	expected := map[string]string{"a": "aaa", "b": "aaa"}
	if !reflect.DeepEqual(expected, actual) {
		t.Fatal("not equal")
	}

	actual = mapShare(map[string]string{"b": "xyz", "c": "ccc"})
	expected = map[string]string{"b": "xyz"}
	if !reflect.DeepEqual(expected, actual) {
		t.Fatal("not equal")
	}

	actual = mapShare(map[string]string{"a": "aaa", "c": "meh", "d": "hi"})
	expected = map[string]string{"a": "aaa", "b": "aaa", "d": "hi"}
	if !reflect.DeepEqual(expected, actual) {
		t.Fatal("not equal")
	}
}
