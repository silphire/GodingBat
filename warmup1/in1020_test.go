package warmup1

import "testing"

func TestOneIsIn(t *testing.T) {
	if !in1020(10, 99) {
		t.Fatal("expected true but actually not")
	}
}

func TestTwoIsIn(t *testing.T) {
	if !in1020(3, 20) {
		t.Fatal("expected true but actually not")
	}
}

func TestNoOneIsIn(t *testing.T) {
	if in1020(3, 33) {
		t.Fatal("expected false but actually not")
	}
}
