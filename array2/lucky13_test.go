package array2

import "testing"

func TestLucky13(t *testing.T) {
	if !lucky13([]int{0, 2, 4}) {
		t.Fatal("expected true but actually not")
	}
	if lucky13([]int{1, 2, 3}) {
		t.Fatal("expected false but actually not")
	}
	if lucky13([]int{1, 2, 4}) {
		t.Fatal("expected false but actually not")
	}
}
