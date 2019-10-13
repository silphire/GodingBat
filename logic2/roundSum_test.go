package logic2

import "testing"

func TestRoundSum(t *testing.T) {
	if roundSum(16, 17, 18) != 60 {
		t.Fatal("expected 60 but actualy not")
	}
	if roundSum(12, 13, 14) != 30 {
		t.Fatal("expected 30 but actualy not")
	}
	if roundSum(6, 4, 5) != 20 {
		t.Fatal("expected 20 but actualy not")
	}
}