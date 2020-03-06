package map1

import (
	"reflect"
	"testing"
)

func TestTopping3(t *testing.T) {
	actual := topping3(map[string]string{"potato": "ketchup"})
	expected := map[string]string{"potato": "ketchup", "fries": "ketchup"}
	if !reflect.DeepEqual(expected, actual) {
		t.Fatal("not equal")
	}

	actual = topping3(map[string]string{"potato": "butter"})
	expected = map[string]string{"potato": "butter", "fries": "butter"}
	if !reflect.DeepEqual(expected, actual) {
		t.Fatal("not equal")
	}

	actual = topping3(map[string]string{"salad": "oil", "potato": "ketchup"})
	expected = map[string]string{"spinach": "oil", "salad": "oil", "potato": "ketchup", "fries": "ketchup"}
	if !reflect.DeepEqual(expected, actual) {
		t.Fatal("not equal")
	}
}
