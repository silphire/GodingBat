package ap1

import "testing"

func TestBigHeights(t *testing.T) {
	if bigHeights([]int{5, 3, 6, 7, 2}, 2, 4) != 1 {
		t.Fatal("expected 1 but actual is not")
	}
	if bigHeights([]int{5, 3, 6, 7, 2}, 0, 1) != 0 {
		t.Fatal("expected 0 but actual is not")
	}
	if bigHeights([]int{5, 3, 6, 7, 2}, 0, 4) != 1 {
		t.Fatal("expected 1 but actual is not")
	}
}
