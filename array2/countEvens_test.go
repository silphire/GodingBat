package array2

import "testing"

func TestCountEvents(t *testing.T) {
	if countEvents([]int{2, 1, 2, 3, 4}) != 3 {
		t.Fatal("expected 3 but actual is not")
	}
	if countEvents([]int{2, 2, 0}) != 3 {
		t.Fatal("expected 3 but actual is not")
	}
	if countEvents([]int{1, 3, 5}) != 0 {
		t.Fatal("expected 0 but actual is not")
	}
}
