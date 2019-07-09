package string1

import "testing"

func TestEndsLy1(t *testing.T) {
	if !endsLy("fooly") {
		t.Fatal("expected true but not")
	}
}

func TestEndsLy2(t *testing.T) {
	if endsLy("y") {
		t.Fatal("expected false but not")
	}
}

func TestEndsLy3(t *testing.T) {
	if endsLy("truethy") {
		t.Fatal("expected false but not")
	}
}
