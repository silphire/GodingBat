package logic1

import "testing"

func TestGreenTicket(t *testing.T) {
	if greenTicket(1, 2, 3) != 0 {
		t.Fatal("expected 0 but actually not")
	}
	if greenTicket(2, 2, 2) != 20 {
		t.Fatal("expected 20 but actually not")
	}
	if greenTicket(1, 1, 2) != 10 {
		t.Fatal("expected 10 but actually not")
	}
}
