package map2

import (
	"reflect"
	"testing"
)

func TestPairs(t *testing.T) {
	actual := pairs([]string{"code", "bug"})
	expected := map[string]string{"b": "g", "c": "e"}
	if !reflect.DeepEqual(expected, actual) {
		t.Fatal("not equal")
	}

	actual = pairs([]string{"man", "moon", "main"})
	expected = map[string]string{"m": "n"}
	if !reflect.DeepEqual(expected, actual) {
		t.Fatal("not equal")
	}

	actual = pairs([]string{"man", "moon", "good", "night"})
	expected = map[string]string{"g": "d", "m": "n", "n": "t"}
	if !reflect.DeepEqual(expected, actual) {
		t.Fatal("not equal")
	}
}
