package ap1

import "testing"

func TestScoresIncreasing(t *testing.T) {
	if !scoresIncreasing([]int{1, 3, 4}) {
		t.Fatal("expected true but actual is not")
	}
	if scoresIncreasing([]int{1, 3, 2}) {
		t.Fatal("expected false but actual is not")
	}
	if !scoresIncreasing([]int{1, 1, 4}) {
		t.Fatal("expected true but actual is not")
	}
}
