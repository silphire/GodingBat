package array3

import "testing"

func TestCanBalance(t *testing.T) {
	if !canBalance([]int{1, 1, 1, 2, 1}) {
		t.Fatal("expected true but actual is not")
	}
	if canBalance([]int{2, 1, 1, 2, 1}) {
		t.Fatal("expected false but actual is not")
	}
	if !canBalance([]int{10, 10}) {
		t.Fatal("expected true but actual is not")
	}
}
