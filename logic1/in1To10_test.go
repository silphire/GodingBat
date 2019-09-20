package logic1

import "testing"

func TestIn1To10(t *testing.T) {
	if !in1To10(5, false) {
		t.Fatal("expected true but actually not")
	}
	if in1To10(11, false) {
		t.Fatal("expected false but actually not")
	}
	if !in1To10(11, true) {
		t.Fatal("expected true but actually not")
	}
}