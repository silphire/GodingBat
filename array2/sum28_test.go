package array2

import "testing"

func TestSum28(t *testing.T) {
	if !sum28([]int{2, 3, 2, 2, 4, 2}) {
		t.Fatal("expected true but actual is not")
	}
	if sum28([]int{2, 3, 2, 2, 4, 2, 2}) {
		t.Fatal("expected false but actual is not")
	}
	if sum28([]int{1, 2, 3, 4}) {
		t.Fatal("expected false but actual is not")
	}
}
