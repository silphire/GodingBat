package array1

import "testing"

func TestSameFirstLast(t *testing.T) {
	if sameFirstLast([]int{1, 2, 3}) {
		t.Fatal("expected false but actually not")
	}
	if !sameFirstLast([]int{1, 2, 3, 1}) {
		t.Fatal("expected true but actually not")
	}
	if !sameFirstLast([]int{1, 2, 1}) {
		t.Fatal("expected true but actually not")
	}
	if sameFirstLast([]int{}) {
		t.Fatal("expected false but actually not")
	}
}