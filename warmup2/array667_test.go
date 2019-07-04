package warmup2

import "testing"

func TestArray66(t *testing.T) {
	if array667([]int{6, 6, 2}) != 1 {
		t.Fatal("expected 1 but actually not")
	}
}

func TestArray66Long(t *testing.T) {
	if array667([]int{6, 6, 2, 6}) != 1 {
		t.Fatal("expected 1 but actually not")
	}
}

func TestArray67(t *testing.T) {
	if array667([]int{6, 7, 2, 6}) != 1 {
		t.Fatal("expected 1 but actually not")
	}
}
