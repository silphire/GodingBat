package map1

import (
	"reflect"
	"testing"
)

func TestMapAB(t *testing.T) {
	expected := map[string]string{"a": "Hi", "ab": "HiThere", "b": "There"}
	actual := mapAB(map[string]string{"a": "Hi", "b": "There"})
	if !reflect.DeepEqual(expected, actual) {
		t.Fatal("not equal")
	}

	expected = map[string]string{"a": "Hi"}
	actual = mapAB(map[string]string{"a": "Hi"})
	if !reflect.DeepEqual(expected, actual) {
		t.Fatal("not equal")
	}

	expected = map[string]string{"b": "There"}
	actual = mapAB(map[string]string{"b": "There"})
	if !reflect.DeepEqual(expected, actual) {
		t.Fatal("not equal")
	}
}
