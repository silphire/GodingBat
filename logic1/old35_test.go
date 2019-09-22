package logic1

import "testing"

func TestOld35(t *testing.T) {
	if !old35(3) {
		t.Fatal("expected true but actually not")
	}
	if !old35(10) {
		t.Fatal("expected true but actually not")
	}
	if old35(15) {
		t.Fatal("expected false but actually not")
	}
	if old35(4) {
		t.Fatal("expected false but actually not")
	}
}
