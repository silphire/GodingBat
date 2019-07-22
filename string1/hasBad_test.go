package string1

import "testing"

func TestHasBad(t *testing.T) {
	if !hasBad("bad") {
		t.Fatal("expected true but actually false")
	}
}

func TestDontHaveBad(t *testing.T) {
	if hasBad("good") {
		t.Fatal("expected false but actually true")
	}
}

func TestHasBadMiddle(t *testing.T) {
	if !hasBad("xxbadx") {
		t.Fatal("expected true but actually false")
	}
}
