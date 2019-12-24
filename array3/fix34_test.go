package array3

import "testing"

func TestFix34(t *testing.T) {
	expected := []int{1, 3, 4, 1}
	actual := fix34([]int{1, 3, 1, 4})
	if len(expected) != len(actual) {
		t.Fatal("length mismatch")
	}
	for i := 0; i < len(expected); i++ {
		if expected[i] != actual[i] {
			t.Fatalf("expected %d but actual is %d at %d", expected[i], actual[i], i)
		}
	}
	expected = []int{1, 3, 4, 1}
	actual = fix34([]int{1, 3, 1, 4})
	if len(expected) != len(actual) {
		t.Fatal("length mismatch")
	}
	for i := 0; i < len(expected); i++ {
		if expected[i] != actual[i] {
			t.Fatalf("expected %d but actual is %d at %d", expected[i], actual[i], i)
		}
	}

	expected = []int{1, 3, 4, 1, 1, 3, 4}
	actual = fix34([]int{1, 3, 1, 4, 4, 3, 1})
	if len(expected) != len(actual) {
		t.Fatal("length mismatch")
	}
	for i := 0; i < len(expected); i++ {
		if expected[i] != actual[i] {
			t.Fatalf("expected %d but actual is %d at %d", expected[i], actual[i], i)
		}
	}

	expected = []int{3, 4, 2, 2}
	actual = fix34([]int{3, 2, 2, 4})
	if len(expected) != len(actual) {
		t.Fatal("length mismatch")
	}
	for i := 0; i < len(expected); i++ {
		if expected[i] != actual[i] {
			t.Fatalf("expected %d but actual is %d at %d", expected[i], actual[i], i)
		}
	}
}
