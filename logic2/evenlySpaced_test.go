package logic2

import "testing"

func TestEvenlySpaced(t *testing.T) {
	if !evenlySpaced(2, 4, 6) {
		t.Fatal("expected true but actually not")
	}
	if !evenlySpaced(4, 6, 2) {
		t.Fatal("expected true but actually not")
	}
	if evenlySpaced(4, 6, 3) {
		t.Fatal("expected false but actually not")
	}
}
