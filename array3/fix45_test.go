package array3

import "testing"

func TestFix45(t *testing.T) {
	expected := []int{9, 4, 5, 4, 5, 9}
	actual := fix45([]int{5, 4, 9, 4, 9, 5})
	if len(expected) != len(actual) {
		t.Fatal("length mismatch")
	}
	for i := 0; i < len(expected); i++ {
		if expected[i] != actual[i] {
			t.Fatalf("expected %d but actual is %d at %d", expected[i], actual[i], i)
		}
	}

	expected = []int{1, 4, 5, 1}
	actual = fix45([]int{1, 4, 1, 5})
	if len(expected) != len(actual) {
		t.Fatal("length mismatch")
	}
	for i := 0; i < len(expected); i++ {
		if expected[i] != actual[i] {
			t.Fatalf("expected %d but actual is %d at %d", expected[i], actual[i], i)
		}
	}

	expected = []int{1, 4, 5, 1, 1, 4, 5}
	actual = fix45([]int{1, 4, 1, 5, 5, 4, 1})
	if len(expected) != len(actual) {
		t.Fatal("length mismatch")
	}
	for i := 0; i < len(expected); i++ {
		if expected[i] != actual[i] {
			t.Fatalf("expected %d but actual is %d at %d", expected[i], actual[i], i)
		}
	}
}