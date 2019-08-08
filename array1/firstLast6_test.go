package array1

import "testing"

func TestFirstLast6(t *testing.T) {
	if !firstLast6([]int {1, 2, 6}) {
		t.Fatal("expected true but not")
	}
	if !firstLast6([]int {6, 2, 1}) {
		t.Fatal("expected true but not")
	}
	if firstLast6([]int {4, 2, 1}) {
		t.Fatal("expected false but not")
	}
}