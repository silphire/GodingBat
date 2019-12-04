package array2

import "testing"

func TestOnly14(t *testing.T) {
	if !only14([]int{1, 4, 1, 4}) {
		t.Fatal("expected true but actual is not")
	}

	if only14([]int{1, 4, 2, 4}) {
		t.Fatal("expected false but actual is not")
	}

	if !only14([]int{1, 1}) {
		t.Fatal("expected true but actual is not")
	}
}
