package array1

import "testing"

func TestCommonEnd(t *testing.T) {
	if !commonEnd([]int{1, 2, 3}, []int {1, 3}) {
		t.Fatal("expected true but not")
	}
	if commonEnd([]int{0, 2, 3}, []int {1, 3}) {
		t.Fatal("expected false but not")
	}
	if commonEnd([]int{1, 2, 3}, []int {1, 4}) {
		t.Fatal("expected false but not")
	}
}