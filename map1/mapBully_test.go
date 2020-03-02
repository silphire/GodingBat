package map1

import (
	"reflect"
	"testing"
)

func TestMapBully(t *testing.T) {
	actual := mapBully(map[string]string{"a": "candy", "b": "dirt"})
	expected := map[string]string{"a": "", "b": "candy"}
	if !reflect.DeepEqual(expected, actual) {
		t.Fatal("not equal")
	}

	actual = mapBully(map[string]string{"a": "candy"})
	expected = map[string]string{"a": "", "b": "candy"}
	if !reflect.DeepEqual(expected, actual) {
		t.Fatal("not equal")
	}

	actual = mapBully(map[string]string{"a": "candy", "b": "carrot", "c": "meh"})
	expected = map[string]string{"a": "", "b": "candy", "c": "meh"}
	if !reflect.DeepEqual(expected, actual) {
		t.Fatal("not equal")
	}
}
