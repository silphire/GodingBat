package logic1

import "testing"

func TestInOrder(t *testing.T) {
	if !inOrder(1, 2, 4, false) {
		t.Fatal("expected true but actually not")
	}
	if inOrder(1, 2, 1, false) {
		t.Fatal("expected false but actually not")
	}
	if !inOrder(1, 1, 2, true) {
		t.Fatal("expected true but actually not")
	}
}
