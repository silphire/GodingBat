package logic1

import "testing"

func TestMaxMod5(t *testing.T) {
	if maxMod5(2, 3) != 3 {
		t.Fatal("expected 3 but actually not")
	}
	if maxMod5(6, 2) != 6 {
		t.Fatal("expected 6 but actually not")
	}
	if maxMod5(3, 2) != 3 {
		t.Fatal("expected 3 but actually not")
	}
}