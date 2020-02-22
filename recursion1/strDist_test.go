package recursion1

import "testing"

func TestStrDist(t *testing.T) {
	if strDist("catcowcat", "cat") != 9 {
		t.Fatal("expected 9 but actual is not")
	}

	if strDist("catcowcat", "cow") != 3 {
		t.Fatal("expected 3 but actual is not")
	}

	if strDist("cccatcowcatxx", "cat") != 9 {
		t.Fatal("expected 9 but actual is not")
	}
}
