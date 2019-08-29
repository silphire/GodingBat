package array1

import "testing"

func TestRotateLeft31(t *testing.T) {
	expected := []int{2, 3, 1}
	actual := rotateLeft3([]int{1, 2, 3})

	if len(expected) != len(actual) {
		t.Fatal("length mismatch")
	}

	for i := 0; i < len(expected); i++ {
		if expected[i] != actual[i] {
			t.Fatal("content mismatch")
		}
	}
}