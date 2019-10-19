package string2

import "testing"

func TestBobThere(t *testing.T) {
	if !bobThere("abcbob") {
		t.Fatal("expected true but actually not")
	}
	if !bobThere("b9b") {
		t.Fatal("expected true but actually not")
	}
	if bobThere("bac") {
		t.Fatal("expected false but actually not")
	}
}
