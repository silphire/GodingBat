package array2

import "testing"

func TestShiftLeft(t *testing.T) {
	expected := []int{2, 5, 3, 6}
	actual := shiftLeft([]int{6, 2, 5, 3})
	if len(expected) != len(actual) {
		t.Fatal("size mismatch")
	}
	for i := 0; i < len(expected); i++ {
		if expected[i] != actual[i] {
			t.Fatalf("expected %d but actual is %d at %d", expected[i], actual[i], i)
		}
	}

	expected = []int{2, 1}
	actual = shiftLeft([]int{1, 2})
	if len(expected) != len(actual) {
		t.Fatal("size mismatch")
	}
	for i := 0; i < len(expected); i++ {
		if expected[i] != actual[i] {
			t.Fatalf("expected %d but actual is %d at %d", expected[i], actual[i], i)
		}
	}

	expected = []int{1}
	actual = shiftLeft([]int{1})
	if len(expected) != len(actual) {
		t.Fatal("size mismatch")
	}
	for i := 0; i < len(expected); i++ {
		if expected[i] != actual[i] {
			t.Fatalf("expected %d but actual is %d at %d", expected[i], actual[i], i)
		}
	}
}
