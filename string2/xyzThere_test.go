package string2

import "testing"

func TestXyzThere(t *testing.T) {
	if !xyzThere("abcxyz") {
		t.Fatal("expected true but actually not ")
	}
	if xyzThere("abc.xyz") {
		t.Fatal("expected false but actually not ")
	}
	if !xyzThere("xyz.abc") {
		t.Fatal("expected true but actually not ")
	}
}
