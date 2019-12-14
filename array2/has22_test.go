package array2

import "testing"

func TestHas22(t *testing.T) {
	if !has22([]int{1, 2, 2}) {
		t.Fatal("expected true but actual is not")
	}
	if has22([]int{1, 2, 1, 2}) {
		t.Fatal("expected false but actual is not")
	}
	if has22([]int{2, 1, 2}) {
		t.Fatal("expected false but actual is not")
	}
}
