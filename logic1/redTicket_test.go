package logic1

import "testing"

func TestRedTicket(t *testing.T) {
	if redTicket(2, 2, 2) != 10 {
		t.Fatal("expected 10 but actually noy")
	}
	if redTicket(2, 2, 1) != 0 {
		t.Fatal("expected 0 but actually noy")
	}
	if redTicket(0, 0, 0) != 5 {
		t.Fatal("expected 5 but actually noy")
	}
}
