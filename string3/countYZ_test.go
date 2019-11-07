package string3

import "testing"

func TestCountYZ(t *testing.T) {
	if countYZ("fez day") != 2 {
		t.Fatal("expected 2 but actually not")
	}
	if countYZ("day fez") != 2 {
		t.Fatal("expected 2 but actually not")
	}
	if countYZ("day fyyyz") != 2 {
		t.Fatal("expected 2 but actually not")
	}
}
