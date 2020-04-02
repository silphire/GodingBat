package map2

import (
	"reflect"
	"testing"
)

func TestWordMultiple(t *testing.T) {
	actual := wordMultiple([]string{"a", "b", "a", "c", "b"})
	expected := map[string]bool{"a": true, "b": true, "c": false}
	if !reflect.DeepEqual(expected, actual) {
		t.Fatal("not equal")
	}

	actual = wordMultiple([]string{"c", "b", "a"})
	expected = map[string]bool{"a": false, "b": false, "c": false}
	if !reflect.DeepEqual(expected, actual) {
		t.Fatal("not equal")
	}

	actual = wordMultiple([]string{"c", "c", "c", "c"})
	expected = map[string]bool{"c": true}
	if !reflect.DeepEqual(expected, actual) {
		t.Fatal("not equal")
	}
}
