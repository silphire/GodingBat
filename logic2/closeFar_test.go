package logic2

import "testing"

func TestCloseFar(t *testing.T) {
	if !closeFar(1, 2, 10) {
		t.Fatal("expected true but actually not")
	}
	if closeFar(1, 2, 3) {
		t.Fatal("expected false but actually not")
	}
	if !closeFar(4, 1, 3) {
		t.Fatal("expected true but actually not")
	}
}
