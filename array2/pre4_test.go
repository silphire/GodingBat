package array2

import "testing"

func TestPre4(t *testing.T) {
	expected := []int{1, 2}
	actual := pre4([]int{1, 2, 4, 1})
	if len(actual) != len(expected) {
		t.Fatal("length mismatch")
	}
	for i := 0; i < len(expected); i++ {
		if expected[i] != actual[i] {
			t.Fatalf("expected %d but actual is %d at %d", expected[i], actual[i], i)
		}
	}

	expected = []int{3, 1}
	actual = pre4([]int{3, 1, 4})
	if len(actual) != len(expected) {
		t.Fatal("length mismatch")
	}
	for i := 0; i < len(expected); i++ {
		if expected[i] != actual[i] {
			t.Fatalf("expected %d but actual is %d at %d", expected[i], actual[i], i)
		}
	}

	expected = []int{1}
	actual = pre4([]int{1, 4, 4})
	if len(actual) != len(expected) {
		t.Fatal("length mismatch")
	}
	for i := 0; i < len(expected); i++ {
		if expected[i] != actual[i] {
			t.Fatalf("expected %d but actual is %d at %d", expected[i], actual[i], i)
		}
	}
}
