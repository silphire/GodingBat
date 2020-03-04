package map1

import (
	"reflect"
	"testing"
)

func TestTopping1(t *testing.T) {
	actual := topping1(map[string]string{"ice cream": "peanuts"})
	expected := map[string]string{"bread": "butter", "ice cream": "cherry"}
	if !reflect.DeepEqual(expected, actual) {
		t.Fatal("not equal")
	}

	actual = topping1(map[string]string{})
	expected = map[string]string{"bread": "butter"}
	if !reflect.DeepEqual(expected, actual) {
		t.Fatal("not equal")
	}

	actual = topping1(map[string]string{"pancake": "syrup"})
	expected = map[string]string{"bread": "butter", "pancake": "syrup"}
	if !reflect.DeepEqual(expected, actual) {
		t.Fatal("not equal")
	}
}
