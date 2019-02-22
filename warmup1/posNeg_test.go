package warmup1

import "testing"

func TestOneIsNegative(t *testing.T) {
	result := posNeg(2, -1, false)
	if !result {
		t.Fatal("expected true but false actually")
	}
}

func TestBothAreNegative(t *testing.T) {
	result := posNeg(-3, -5, false)
	if result {
		t.Fatal("expected false but true actually")
	}
}

func TestBothAreNegativeAndFlagOn(t *testing.T) {
	result := posNeg(-2, -9, true)
	if !result {
		t.Fatal("expected true but false actually")
	}
}
