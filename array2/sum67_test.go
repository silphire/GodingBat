package array2

import "testing"

func TestSum67(t *testing.T) {
	if sum67([]int{1, 2, 2}) != 5 {
		t.Fatal("expected 5 but actual is not")
	}
	if sum67([]int{1, 2, 2, 6, 99, 99, 7}) != 5 {
		t.Fatal("expected 5 but actual is not")
	}
	if sum67([]int{1, 1, 6, 7, 2}) != 4 {
		t.Fatal("expected 5 but actual is not")
	}
}
