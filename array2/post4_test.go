package array2

import "testing"

func TestPost4(t *testing.T) {
	expected := []int{1, 2}
	actual := post4([]int{2, 4, 1, 2})
	if len(expected) != len(actual) {
		t.Fatalf("size mismatch, expected:%d actual:%d", len(expected), len(actual))
	}
	for i := 0; i < len(expected); i++ {
		if expected[i] != actual[i] {
			t.Fatalf("expected %d but actual is %d at %d", expected[i], actual[i], i)
		}
	}

	expected = []int{2}
	actual = post4([]int{4, 1, 4, 2})
	if len(expected) != len(actual) {
		t.Fatalf("size mismatch, expected:%d actual:%d", len(expected), len(actual))
	}
	for i := 0; i < len(expected); i++ {
		if expected[i] != actual[i] {
			t.Fatalf("expected %d but actual is %d at %d", expected[i], actual[i], i)
		}
	}

	expected = []int{1, 2, 3}
	actual = post4([]int{4, 4, 1, 2, 3})
	if len(expected) != len(actual) {
		t.Fatalf("size mismatch, expected:%d actual:%d", len(expected), len(actual))
	}
	for i := 0; i < len(expected); i++ {
		if expected[i] != actual[i] {
			t.Fatalf("expected %d but actual is %d at %d", expected[i], actual[i], i)
		}
	}
}
