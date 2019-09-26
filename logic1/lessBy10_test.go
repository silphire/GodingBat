package logic1

import "testing"

func TestLessBy10(t *testing.T) {
	if !lessBy10(1, 7, 11) {
		t.Fatal("expected true but actually false")
	}
	if lessBy10(1, 7, 10) {
		t.Fatal("expected false but actually false")
	}
	if !lessBy10(11, 7, 1) {
		t.Fatal("expected true but actually false")
	}
}
