package array3

import "testing"

func TestMaxSpan(t *testing.T) {
	if maxSpan([]int{1, 2, 1, 1, 3}) != 4 {
		t.Fatal("expected 4 but actual is not")
	}
	if maxSpan([]int{1, 4, 2, 1, 4, 1, 4}) != 6 {
		t.Fatal("expected 6 but actual is not")
	}
	if maxSpan([]int{1, 4, 2, 1, 4, 4, 4}) != 6 {
		t.Fatal("expected 6 but actual is not")
	}
}
