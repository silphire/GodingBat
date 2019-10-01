package logic1

import "testing"

func TestSpecialEleven(t *testing.T) {
	if !specialEleven(22) {
		t.Fatal("expected true but actually not")
	}
	if !specialEleven(23) {
		t.Fatal("expected true but actually not")
	}
	if specialEleven(24) {
		t.Fatal("expected false but actually not")
	}
}