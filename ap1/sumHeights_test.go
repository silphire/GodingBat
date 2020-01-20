package ap1

import "testing"

func TestSumHeights(t *testing.T) {
	if sumHeights([]int{5, 3, 6, 7, 2}, 2, 4) != 6 {
		t.Fatal("expected 6 but actual is not")
	}
	if sumHeights([]int{5, 3, 6, 7, 2}, 0, 1) != 2 {
		t.Fatal("expected 2 but actual is not")
	}
	if sumHeights([]int{5, 3, 6, 7, 2}, 0, 4) != 11 {
		t.Fatal("expected 11 but actual is not")
	}
}
