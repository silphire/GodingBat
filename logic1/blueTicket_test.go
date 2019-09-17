package logic1

import "testing"

func TestBlueTicket(t *testing.T) {
	if blueTicket(9, 1, 0) != 10 {
		t.Fatal("expected 10 but actually not")
	}
	if blueTicket(9, 2, 0) != 0 {
		t.Fatal("expected 10 but actually not")
	}
	if blueTicket(15, 0, 5) != 5 {
		t.Fatal("expected 5 but actually not")
	}
}