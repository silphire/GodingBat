package array1

import "testing"

func TestHas23(t *testing.T) {
	if !has23([]int{2, 5}) {
		t.Fatal("expected true but actually not")
	}
	if !has23([]int{4, 3}) {
		t.Fatal("expected true but actually not")
	}
	if has23([]int{4, 5}) {
		t.Fatal("expected true but actually not")
	}
}