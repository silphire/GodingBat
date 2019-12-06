package array2

import "testing"

func TestHas77(t *testing.T) {
	if !has77([]int{1, 7, 7}) {
		t.Fatal("expected true but actual is not")
	}
	if !has77([]int{1, 7, 1, 7}) {
		t.Fatal("expected true but actual is not")
	}
	if has77([]int{1, 7, 1, 1, 7}) {
		t.Fatal("expected false but actual is not")
	}
}
