package map2

import (
	"reflect"
	"testing"
)

func TestFirstChar(t *testing.T) {
	expected := firstChar([]string{"salt", "tea", "soda", "toast"})
	actual := map[string]string{"s": "saltsoda", "t": "teatoast"}
	if !reflect.DeepEqual(expected, actual) {
		t.Fatal("not equal")
	}

	expected = firstChar([]string{"aa", "bb", "cc", "aAA", "cCC", "d"})
	actual = map[string]string{"a": "aaaAA", "b": "bb", "c": "cccCC", "d": "d"}
	if !reflect.DeepEqual(expected, actual) {
		t.Fatal("not equal")
	}

	expected = firstChar([]string{})
	actual = map[string]string{}
	if !reflect.DeepEqual(expected, actual) {
		t.Fatal("not equal")
	}
}
