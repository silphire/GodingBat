package array3

import "testing"

func TestLinearIn(t *testing.T) {
	if !linearIn([]int{1, 2, 4, 6}, []int{2, 4}) {
		t.Fatal("expected true but actual is not")
	}
	if linearIn([]int{1, 2, 4, 6}, []int{2, 3, 4}) {
		t.Fatal("expected false but actual is not")
	}
	if !linearIn([]int{1, 2, 4, 4, 6}, []int{2, 4}) {
		t.Fatal("expected true but actual is not")
	}
}
