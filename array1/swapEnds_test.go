package array1

import "testing"

func TestSwapEnds(t *testing.T) {
	expected := []int{4, 2, 3, 1}
	actual := swapEnds([]int{1, 2, 3, 4})
	if len(expected) != len(actual) {
		t.Fatal("length mismatch")
	}
	for i := 0; i < len(expected); i++ {
		if expected[i] != actual[i] {
			t.Fatal("content mismatch")
		}
	}
}