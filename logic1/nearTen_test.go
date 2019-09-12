package logic1

import "testing"

func testNearTen(t *testing.T) {
	if !nearTen(12) {
		t.Fatal("expected true but actually not")
	}
	if nearTen(17) {
		t.Fatal("expected false but actually not")
	}
	if !nearTen(19) {
		t.Fatal("expected true but actually not")
	}
}