package array1

import "testing"

func TestUnlicky1(t *testing.T) {
	if !unlucky1([]int{1, 3, 4, 5}) {
		t.Fatal("expected true but not")
	}
	if !unlucky1([]int{4, 2, 1, 3}) {
		t.Fatal("expected true but not")
	}
	if unlucky1([]int{1, 2, 4, 5}) {
		t.Fatal("expected false but not")
	}
}