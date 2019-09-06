package array1

import "testing"

func TestFront111(t *testing.T) {
	expected := []int{1, 7}
	actual := front11([]int{1, 2, 3}, []int{7, 9, 8})

	if len(expected) != len(actual) {
		t.Fatal("length mismatch")
	}
	for i := 0; i < len(expected); i++ {
		if expected[i] != actual[i] {
			t.Fatal("content mismatch")
		}
	}
}

func TestFront112(t *testing.T) {
	expected := []int{1}
	actual := front11([]int{1, 2, 3}, []int{})

	if len(expected) != len(actual) {
		t.Fatal("length mismatch")
	}
	for i := 0; i < len(expected); i++ {
		if expected[i] != actual[i] {
			t.Fatal("content mismatch")
		}
	}
}