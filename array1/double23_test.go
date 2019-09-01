package array1

import "testing"

func TestDouble23(t *testing.T) {
	if !double23([]int{2, 2}) {
		t.Fatal("expected true but actually not")
	}
	if !double23([]int{3, 3}) {
		t.Fatal("expected true but actually not")
	}
	if double23([]int{2, 3}) {
		t.Fatal("expected true but actually not")
	}
}