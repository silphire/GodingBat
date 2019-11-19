package array2

import "testing"

func TestSum13(t *testing.T) {
	actual := sum13([]int{1, 2, 2, 1})
	if actual != 6 {
		t.Fatalf("Expected 6 but actual is %d", actual)
	}

	actual = sum13([]int{1, 1})
	if actual != 2 {
		t.Fatalf("Expected 2 but actual is %d", actual)
	}

	actual = sum13([]int{1, 2, 2, 1, 13})
	if actual != 6 {
		t.Fatalf("Expected 6 but actual is %d", actual)
	}
}
