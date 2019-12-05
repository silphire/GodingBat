package array2

import "testing"

func TestIsEverywhere(t *testing.T) {
	if !isEverywhere([]int{1, 2, 1, 3}, 1) {
		t.Fatal("expected true but actual is not")
	}
	if isEverywhere([]int{1, 2, 1, 3}, 2) {
		t.Fatal("expected false but actual is not")
	}
	if isEverywhere([]int{1, 2, 1, 3, 4}, 1) {
		t.Fatal("expected false but actual is not")
	}
}
