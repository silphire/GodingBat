package array2

import "testing"

func TestTenRun(t *testing.T) {
	expected := []int{2, 10, 10, 10, 20, 20}
	actual := tenRun([]int{2, 10, 3, 4, 20, 5})
	if len(expected) != len(actual) {
		t.Fatal("length mismatch")
	}
	for i := 0; i < len(expected); i++ {
		if expected[i] != actual[i] {
			t.Fatalf("expected %d but actual is %d at %d", expected[i], actual[i], i)
		}
	}

	expected = []int{10, 10, 20, 20}
	actual = tenRun([]int{10, 1, 20, 2})
	if len(expected) != len(actual) {
		t.Fatal("length mismatch")
	}
	for i := 0; i < len(expected); i++ {
		if expected[i] != actual[i] {
			t.Fatalf("expected %d but actual is %d at %d", expected[i], actual[i], i)
		}
	}

	expected = []int{10, 10, 10, 20}
	actual = tenRun([]int{10, 1, 9, 20})
	if len(expected) != len(actual) {
		t.Fatal("length mismatch")
	}
	for i := 0; i < len(expected); i++ {
		if expected[i] != actual[i] {
			t.Fatalf("expected %d but actual is %d at %d", expected[i], actual[i], i)
		}
	}
}
