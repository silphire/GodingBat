package logic1

import "testing"

func TestLess20(t *testing.T) {
	if !less20(18) {
		t.Fatal("expected true but actually not")
	}
	if !less20(19) {
		t.Fatal("expected true but actually not")
	}
	if less20(20) {
		t.Fatal("expected false but actually not")
	}
}
