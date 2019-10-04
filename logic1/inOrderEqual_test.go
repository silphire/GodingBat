package logic1

import "testing"

func TestInOrderEqual(t *testing.T) {
	if !inOrderEqual(2, 5, 11, false) {
		t.Fatal("expected true but actually not")
	}
	if inOrderEqual(5, 7, 6, false) {
		t.Fatal("expected false but actually not")
	}
	if !inOrderEqual(5, 5, 7, true) {
		t.Fatal("expected true but actually not")
	}
}
