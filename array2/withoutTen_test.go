package array2

import "testing"

func TestWithoutTen(t *testing.T) {
	expected := []int{1, 2, 0, 0}
	actual := withoutTen([]int{1, 10, 10, 2})
	if len(expected) != len(actual) {
		t.Fatalf("size mismatch, expected:%d actual:%d", len(expected), len(actual))
	}
	for i := 0; i < len(expected); i++ {
		if expected[i] != actual[i] {
			t.Fatalf("expected %d but actual is %d at %d", expected[i], actual[i], i)
		}
	}

	expected = []int{2, 0, 0}
	actual = withoutTen([]int{10, 2, 10})
	if len(expected) != len(actual) {
		t.Fatalf("size mismatch, expected:%d actual:%d", len(expected), len(actual))
	}
	for i := 0; i < len(expected); i++ {
		if expected[i] != actual[i] {
			t.Fatalf("expected %d but actual is %d at %d", expected[i], actual[i], i)
		}
	}

	expected = []int{1, 99, 0}
	actual = withoutTen([]int{1, 99, 10})
	if len(expected) != len(actual) {
		t.Fatalf("size mismatch, expected:%d actual:%d", len(expected), len(actual))
	}
	for i := 0; i < len(expected); i++ {
		if expected[i] != actual[i] {
			t.Fatalf("expected %d but actual is %d at %d", expected[i], actual[i], i)
		}
	}
}
