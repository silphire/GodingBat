package array2

import "testing"

func TestModThree(t *testing.T) {
	if !modThree([]int{2, 1, 3, 5}) {
		t.Fatal("expected true but actual is not")
	}
	if modThree([]int{2, 1, 2, 5}) {
		t.Fatal("expected false but actual is not")
	}
	if !modThree([]int{2, 4, 2, 5}) {
		t.Fatal("expected true but actual is not")
	}
}
