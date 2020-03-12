package map1

import (
	"reflect"
	"testing"
)

func TestMapAB3(t *testing.T) {
	actual := mapAB3(map[string]string{"a": "aaa", "c": "cake"})
	expected := map[string]string{"a": "aaa", "b": "aaa", "c": "cake"}
	if !reflect.DeepEqual(expected, actual) {
		t.Fatal("not equal")
	}

	actual = mapAB3(map[string]string{"b": "bbb", "c": "cake"})
	expected = map[string]string{"a": "bbb", "b": "bbb", "c": "cake"}
	if !reflect.DeepEqual(expected, actual) {
		t.Fatal("not equal")
	}

	actual = mapAB3(map[string]string{"a": "aaa", "b": "bbb", "c": "cake"})
	expected = map[string]string{"a": "aaa", "b": "bbb", "c": "cake"}
	if !reflect.DeepEqual(expected, actual) {
		t.Fatal("not equal")
	}
}
