package logic1

import "testing"

func TestCigarParty(t *testing.T) {
	if cigarParty(30, false) {
		t.Fatal("expected false but actually true")
	}
	if !cigarParty(50, false) {
		t.Fatal("expected true but actually false")
	}
	if !cigarParty(70, true) {
		t.Fatal("expected true but actually false")
	}
}