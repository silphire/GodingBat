package array1

import "testing"

func TestMakeLast(t *testing.T) {
	expected := []int{0, 0, 0, 0, 0, 6}
	actual := makeLast([]int{4, 5, 6})
	if len(expected) != len(actual) {
		t.Fatal("length mismatch")
	}
	for i := 0; i < len(expected); i++ {
		if expected[i] != actual[i] {
			t.Fatal("value mismatch")
		}
	}
}