package string3

import "testing"

func TestMaxBlock(t *testing.T) {
	if maxBlock("hoopla") != 2 {
		t.Fatal("expected 2 but actual is not")
	}

	if maxBlock("abbCCCddBBBxx") != 3 {
		t.Fatal("expected 3 but actual is not")
	}

	if maxBlock("") != 0 {
		t.Fatal("expected 0 but actual is not")
	}
}
