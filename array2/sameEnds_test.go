package array2

import "testing"

func TestSameEnds(t *testing.T) {
	if sameEnds([]int{5, 6, 45, 99, 13, 5, 6}, 1) {
		t.Fatal("expected false but actual is not")
	}
	if !sameEnds([]int{5, 6, 45, 99, 13, 5, 6}, 2) {
		t.Fatal("expected true but actual is not")
	}
	if sameEnds([]int{5, 6, 45, 99, 13, 5, 6}, 3) {
		t.Fatal("expected false but actual is not")
	}
}
