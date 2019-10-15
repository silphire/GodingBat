package logic2

import "testing"

func TestLuckySum(t *testing.T) {
	if luckySum(1, 2, 3) != 6 {
		t.Fatal("expected 6 but actually not")
	}
	if luckySum(1, 2, 13) != 3 {
		t.Fatal("expected 3 but actually not")
	}
	if luckySum(1, 13, 3) != 1 {
		t.Fatal("expected 1 but actually not")
	}
}
