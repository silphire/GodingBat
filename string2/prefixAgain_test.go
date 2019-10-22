package string2

import "testing"

func TestPrefixAgain(t *testing.T) {
	if !prefixAgain("abXYabc", 1) {
		t.Fatal("expected true but actually not")
	}
	if !prefixAgain("abXYabc", 2) {
		t.Fatal("expected true but actually not")
	}
	if prefixAgain("abXYabc", 3) {
		t.Fatal("expected false but actually not")
	}
}
