package map1

import (
	"reflect"
	"testing"
)

func TestMapAB2(t *testing.T) {
	actual := mapAB2(map[string]string{"a": "aaa", "b": "aaa", "c": "cake"})
	expected := map[string]string{"c": "cake"}
	if !reflect.DeepEqual(expected, actual) {
		t.Fatal("not equal")
	}

	actual = mapAB2(map[string]string{"a": "aaa", "b": "bbb"})
	expected = map[string]string{"a": "aaa", "b": "bbb"}
	if !reflect.DeepEqual(expected, actual) {
		t.Fatal("not equal")
	}

	actual = mapAB2(map[string]string{"a": "aaa", "b": "bbb", "c": "aaa"})
	expected = map[string]string{"a": "aaa", "b": "bbb", "c": "aaa"}
	if !reflect.DeepEqual(expected, actual) {
		t.Fatal("not equal")
	}
}
