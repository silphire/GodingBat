package logic1

import "testing"

func TestTwoAsOne(t *testing.T) {
	if !twoAsOne(1, 2, 3) {
		t.Fatal("expected true but actually false")
	}
	if !twoAsOne(3, 1, 2) {
		t.Fatal("expected true but actually false")
	}
	if twoAsOne(3, 2, 2) {
		t.Fatal("expected false but actually false")
	}
}