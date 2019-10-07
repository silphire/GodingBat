package logic1

import "testing"

func TestSumLimit(t *testing.T) {
	if sumLimit(2, 3) != 5 {
		t.Fatal("expected 5 but actually not")
	}
	if sumLimit(8, 3) != 8 {
		t.Fatal("expected 8 but actually not")
	}
	if sumLimit(8, 1) != 9 {
		t.Fatal("expected 9 but actually not")
	}
}