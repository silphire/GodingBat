package ap1

import "testing"

func TestScoresSpecial(t *testing.T) {
	if scoresSpecial([]int{12, 10, 4}, []int{2, 20, 30}) != 40 {
		t.Fatal("expected 40 but actual is not")
	}
	if scoresSpecial([]int{20, 10, 4}, []int{2, 20, 10}) != 40 {
		t.Fatal("expected 40 but actual is not")
	}
	if scoresSpecial([]int{12, 11, 4}, []int{2, 20, 31}) != 20 {
		t.Fatal("expected 20 but actual is not")
	}
}
