package map2

import (
	"reflect"
	"testing"
)

func TestWord0(t *testing.T) {
	expected := map[string]int{"a": 0, "b": 0}
	actual := word0([]string{"a", "b", "a", "b"})
	if !reflect.DeepEqual(expected, actual) {
		t.Fatal("not equal")
	}

	actual = word0([]string{"a", "b", "a", "c", "b"})
	expected = map[string]int{"a": 0, "b": 0, "c": 0}
	if !reflect.DeepEqual(expected, actual) {
		t.Fatal("not equal")
	}

	actual = word0([]string{"c", "b", "a"})
	expected = map[string]int{"a": 0, "b": 0, "c": 0}
	if !reflect.DeepEqual(expected, actual) {
		t.Fatal("not equal")
	}
}
