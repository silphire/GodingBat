package ap1

import "testing"

func TestSumHeights2(t *testing.T) {
	if sumHeights2([]int{5, 3, 6, 7, 2}, 2, 4) != 7 {
		t.Fatal("expected 7 but actual is not")
	}
	if sumHeights2([]int{5, 3, 6, 7, 2}, 0, 1) != 2 {
		t.Fatal("expected 2 but actual is not")
	}
	if sumHeights2([]int{5, 3, 6, 7, 2}, 0, 4) != 15 {
		t.Fatal("expected 15 but actual is not")
	}
}
