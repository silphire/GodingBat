package array2

import "testing"

func TestFizzArray(t *testing.T) {
	expected := []int{0, 1, 2, 3}
	actual := fizzArray(4)
	if len(expected) != len(actual) {
		t.Fatalf("length mismatch: expected %d but actual is %d", len(expected), len(actual))
	}
	for i := 0; i < len(expected); i++ {
		if expected[i] != actual[i] {
			t.Fatalf("expected %d but actual is %d at %d", expected[i], actual[i], i)
		}
	}

	expected = []int{0}
	actual = fizzArray(1)
	if len(expected) != len(actual) {
		t.Fatalf("length mismatch: expected %d but actual is %d", len(expected), len(actual))
	}
	for i := 0; i < len(expected); i++ {
		if expected[i] != actual[i] {
			t.Fatalf("expected %d but actual is %d at %d", expected[i], actual[i], i)
		}
	}

	expected = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	actual = fizzArray(10)
	if len(expected) != len(actual) {
		t.Fatalf("length mismatch: expected %d but actual is %d", len(expected), len(actual))
	}
	for i := 0; i < len(expected); i++ {
		if expected[i] != actual[i] {
			t.Fatalf("expected %d but actual is %d at %d", expected[i], actual[i], i)
		}
	}
}
