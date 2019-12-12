package array2

import "testing"

func TestCenteredAverage(t *testing.T) {
	if centeredAverage([]int{1, 2, 3, 4, 100}) != 3 {
		t.Fatal("expected 3 but actual is not")
	}
	if centeredAverage([]int{1, 1, 5, 5, 10, 8, 7}) != 5 {
		t.Fatal("expected 5 but actual is not")
	}
	if centeredAverage([]int{-10, -4, -2, -4, -2, 0}) != -3 {
		t.Fatal("expected -3 but actual is not")
	}
}
