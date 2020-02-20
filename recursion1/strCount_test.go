package recursion1

import "testing"

func TestStrCount(t *testing.T) {
	if strCount("catcowcat", "cat") != 2 {
		t.Fatal("expected 2 but actual is not")
	}
	if strCount("catcowcat", "cow") != 1 {
		t.Fatal("expected 1 but actual is not")
	}
	if strCount("catcowcat", "dog") != 0 {
		t.Fatal("expected 0 but actual is not")
	}
}
