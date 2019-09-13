package logic1

import "testing"

func TestTeaParty(t *testing.T) {
	if teaParty(6, 8) != 1 {
		t.Fatal("expected 1 but actual is not")
	}
	if teaParty(3, 8) != 0 {
		t.Fatal("expected 0 but actual is not")
	}
	if teaParty(20, 6) != 2 {
		t.Fatal("expected 2 but actual is not")
	}
}