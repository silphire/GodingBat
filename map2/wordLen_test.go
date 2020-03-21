package map2

import (
	"reflect"
	"testing"
)

func TestWordLen(t *testing.T) {
	actual := wordLen([]string{"a", "bb", "a", "bb"})
	expected := map[string]int{"bb": 2, "a": 1}
	if !reflect.DeepEqual(expected, actual) {
		t.Fatal("not equal")
	}

	actual = wordLen([]string{"this", "and", "that", "and"})
	expected = map[string]int{"that": 4, "and": 3, "this": 4}
	if !reflect.DeepEqual(expected, actual) {
		t.Fatal("not equal")
	}

	actual = wordLen([]string{"code", "code", "code", "bug"})
	expected = map[string]int{"code": 4, "bug": 3}
	if !reflect.DeepEqual(expected, actual) {
		t.Fatal("not equal")
	}
}
