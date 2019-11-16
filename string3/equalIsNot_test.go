package string3

import "testing"

func TestEqualsNot(t *testing.T) {
	if equalIsNot("This is not") {
		t.Fatal("expected false but actually not")
	}

	if !equalIsNot("This is notnot") {
		t.Fatal("expected true but actually not")
	}

	if !equalIsNot("noisxxnotyynotxisi") {
		t.Fatal("expected true but actually not")
	}
}
