package logic1

import "testing"

func TestShareDigit(t *testing.T) {
	if !shareDigit(12, 23) {
		t.Fatal("expected true but actually not")
	}
	if shareDigit(12, 43) {
		t.Fatal("expected false but actually not")
	}
	if shareDigit(12, 44) {
		t.Fatal("expected false but actually not")
	}
}
