package map1

import (
	"reflect"
	"testing"
)

func TestMapAB4(t *testing.T) {
	actual := mapAB4(map[string]string{"a": "aaa", "b": "bb", "c": "cake"})
	expected := map[string]string{"a": "aaa", "b": "bb", "c": "aaa"}
	if !reflect.DeepEqual(expected, actual) {
		t.Fatal("not equal")
	}

	actual = mapAB4(map[string]string{"a": "aa", "b": "bbb", "c": "cake"})
	expected = map[string]string{"a": "aa", "b": "bbb", "c": "bbb"}
	if !reflect.DeepEqual(expected, actual) {
		t.Fatal("not equal")
	}

	actual = mapAB4(map[string]string{"a": "aa", "b": "bbb"})
	expected = map[string]string{"a": "aa", "b": "bbb", "c": "bbb"}
	if !reflect.DeepEqual(expected, actual) {
		t.Fatal("not equal")
	}
}
