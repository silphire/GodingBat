package array2

import "testing"

func TestHas12(t *testing.T) {
	if !has12([]int{1, 3, 2}) {
		t.Fatal("expected true but actual is not")
	}
	if !has12([]int{3, 1, 2}) {
		t.Fatal("expected true but actual is not")
	}
	if !has12([]int{3, 1, 4, 5, 2}) {
		t.Fatal("expected true but actual is not")
	}
	if has12([]int{3, 1, 4, 5, 6}) {
		t.Fatal("expected false but actual is not")
	}
}
