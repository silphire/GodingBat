package warmup2

import "testing"

func TestHas123(t *testing.T) {
	if !array123([]int{1, 1, 2, 3, 4, 5}) {
		t.Fatal("expected true, but not")
	}
}

func TestDontHave123(t *testing.T) {
	if array123([]int{1, 1, 2, 6, 4, 5}) {
		t.Fatal("expected false, but not")
	}
}
