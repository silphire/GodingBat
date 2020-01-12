package ap1

import "testing"

func TestCopyEndy(t *testing.T) {
	expected := []int{9, 90}
	actual := copyEndy([]int{9, 11, 90, 22, 6}, 2)
	if len(expected) != len(actual) {
		t.Fatal("size mismatch")
	}
	for i := 0; i < len(expected); i++ {
		if expected[i] != actual[i] {
			t.Fatalf("expected %d but actual is %d at %d", expected[i], actual[i], i)
		}
	}

	expected = []int{9, 90, 6}
	actual = copyEndy([]int{9, 11, 90, 22, 6}, 3)
	if len(expected) != len(actual) {
		t.Fatal("size mismatch")
	}
	for i := 0; i < len(expected); i++ {
		if expected[i] != actual[i] {
			t.Fatalf("expected %d but actual is %d at %d", expected[i], actual[i], i)
		}
	}

	expected = []int{1, 1}
	actual = copyEndy([]int{12, 1, 1, 13, 0, 20}, 2)
	if len(expected) != len(actual) {
		t.Fatal("size mismatch")
	}
	for i := 0; i < len(expected); i++ {
		if expected[i] != actual[i] {
			t.Fatalf("expected %d but actual is %d at %d", expected[i], actual[i], i)
		}
	}
}
