package logic2

import "testing"

func TestBlackJack(t *testing.T) {
	if blackjack(19, 21) != 21 {
		t.Fatal("expected 21 but actually not")
	}
	if blackjack(21, 19) != 21 {
		t.Fatal("expected 21 but actually not")
	}
	if blackjack(22, 19) != 19 {
		t.Fatal("expected 19 but actually not")
	}
}
