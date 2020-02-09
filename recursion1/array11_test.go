package recursion1

import "testing"

func TestArray11(t *testing.T) {
	if array11([]int{1, 2, 11}, 0) != 1 {
		t.Fatal("expected 1 but actual is not")
	}

	if array11([]int{11, 11}, 0) != 2 {
		t.Fatal("expected 2 but actual is not")
	}

	if array11([]int{1, 2, 3, 4}, 0) != 0 {
		t.Fatal("expected 0 but actual is not")
	}
}
