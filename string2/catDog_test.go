package string2

import "testing"

func TestCatDog(t *testing.T) {
	if !catDog("catdog") {
		t.Fatal("expected true but actually not")
	}
	if catDog("catcat") {
		t.Fatal("expected false but actually not")
	}
	if !catDog("1cat1cadodog") {
		t.Fatal("expected true but actually not")
	}
}
