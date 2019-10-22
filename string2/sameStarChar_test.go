package string2

import "testing"

func TestSameStarChar(t *testing.T) {
	if !sameStarChar("xy*yzz") {
		t.Fatal("expected true but actually not")
	}
	if sameStarChar("xy*zzz") {
		t.Fatal("expected false but actually not")
	}
	if !sameStarChar("*xa*az") {
		t.Fatal("expected true but actually not")
	}
}
