package ap1

import "testing"

func TestCopyEvens(t *testing.T) {
	expected := []int{2, 4}
	actual := copyEvens([]int{3, 2, 4, 5, 8}, 2)
	if len(expected) != len(actual) {
		t.Fatal("size mismatch")
	}
	for i := 0; i < len(expected); i++ {
		if expected[i] != actual[i] {
			t.Fatalf("expected %d but actual is %d at %d", expected[i], actual[i], i)
		}
	}

	expected = []int{2, 4, 8}
	actual = copyEvens([]int{3, 2, 4, 5, 8}, 3)
	if len(expected) != len(actual) {
		t.Fatal("size mismatch")
	}
	for i := 0; i < len(expected); i++ {
		if expected[i] != actual[i] {
			t.Fatalf("expected %d but actual is %d at %d", expected[i], actual[i], i)
		}
	}

	expected = []int{6, 2, 4}
	actual = copyEvens([]int{6, 1, 2, 4, 5, 8}, 3)
	if len(expected) != len(actual) {
		t.Fatal("size mismatch")
	}
	for i := 0; i < len(expected); i++ {
		if expected[i] != actual[i] {
			t.Fatalf("expected %d but actual is %d at %d", expected[i], actual[i], i)
		}
	}
}
