package logic1

import "testing"

func TestLastDigit(t *testing.T) {
	if !lastDigit(23, 19, 13) {
		t.Fatal("expected true but actually not")
	}
	if lastDigit(23, 19, 12) {
		t.Fatal("expected false but actually not")
	}
	if !lastDigit(23, 19, 3) {
		t.Fatal("expected true but actually not")
	}
}