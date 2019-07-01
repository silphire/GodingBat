package warmup2

import "testing"

func TestNineFound(t *testing.T) {
	if !arrayFront9([]int{1, 2, 9, 4, 5}) {
		t.Fatal("expected true but not")
	}
}

func TestShortArray(t *testing.T) {
	if arrayFront9([]int{1, 2, 3}) {
		t.Fatal("expected false but not")
	}
}

func TestNineFarFromFront(t *testing.T) {
	if arrayFront9([]int{1, 2, 3, 5, 4, 9, 3}) {
		t.Fatal("expected false but not")
	}
}