package functional1

import "testing"

func TestDoubling(t *testing.T) {
	actual := doubling([]int{1, 2, 3})
	expected := []int{2, 4, 6}
	if len(actual) != len(expected) {
		t.Fatal("size mismatch")
	}
	for i := 0; i < len(expected); i++ {
		if expected[i] != actual[i] {
			t.Fatalf("expected %d but actual is %d at %d", expected[i], actual[i], i)
		}
	}

	actual = doubling([]int{6, 8, 6, 8, -1})
	expected = []int{12, 16, 12, 16, -2}
	if len(actual) != len(expected) {
		t.Fatal("size mismatch")
	}
	for i := 0; i < len(expected); i++ {
		if expected[i] != actual[i] {
			t.Fatalf("expected %d but actual is %d at %d", expected[i], actual[i], i)
		}
	}

	actual = doubling([]int{})
	expected = []int{}
	if len(actual) != len(expected) {
		t.Fatal("size mismatch")
	}
	for i := 0; i < len(expected); i++ {
		if expected[i] != actual[i] {
			t.Fatalf("expected %d but actual is %d at %d", expected[i], actual[i], i)
		}
	}
}
