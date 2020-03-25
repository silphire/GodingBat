package map2

import (
	"reflect"
	"testing"
)

func TestWordCount(t *testing.T) {
	actual := wordCount([]string{"a", "b", "a", "c", "b"})
	expected := map[string]int{"a": 2, "b": 2, "c": 1}
	if !reflect.DeepEqual(expected, actual) {
		t.Fatal("not equal")
	}

	actual = wordCount([]string{"c", "b", "a"})
	expected = map[string]int{"a": 1, "b": 1, "c": 1}
	if !reflect.DeepEqual(expected, actual) {
		t.Fatal("not equal")
	}

	actual = wordCount([]string{"c", "c", "c", "c"})
	expected = map[string]int{"c": 4}
	if !reflect.DeepEqual(expected, actual) {
		t.Fatal("not equal")
	}
}
