package logic1

import "testing"

func TestMore20(t *testing.T) {
	if more20(20) {
		t.Fatal("expected false but actually not")
	}
	if !more20(21) {
		t.Fatal("expected true but actually not")
	}
	if !more20(22) {
		t.Fatal("expected true but actually not")
	}
}