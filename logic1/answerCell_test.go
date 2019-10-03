package logic1

import "testing"

func TestAnswerCell(t *testing.T) {
	if !answerCell(false, false, false) {
		t.Fatal("expected true but actually not")
	}
	if answerCell(false, false, true) {
		t.Fatal("expected false but actually not")
	}
	if answerCell(true, false, false) {
		t.Fatal("expected false but actually not")
	}
}