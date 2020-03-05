package map1

import (
	"reflect"
	"testing"
)

func TestTopping2(t *testing.T) {
	actual := topping2(map[string]string{"ice cream": "cherry"})
	expected := map[string]string{"yogurt": "cherry", "ice cream": "cherry"}
	if !reflect.DeepEqual(expected, actual) {
		t.Fatal("not equal")
	}

	actual = topping2(map[string]string{"spinach": "dirt", "ice cream": "cherry"})
	expected = map[string]string{"yogurt": "cherry", "spinach": "nuts", "ice cream": "cherry"}
	if !reflect.DeepEqual(expected, actual) {
		t.Fatal("not equal")
	}

	actual = topping2(map[string]string{"yogurt": "salt"})
	expected = map[string]string{"yogurt": "salt"}
	if !reflect.DeepEqual(expected, actual) {
		t.Fatal("not equal")
	}
}
