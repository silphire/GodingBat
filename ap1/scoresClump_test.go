package ap1

import "testing"

func TestScoresClump(t *testing.T) {
	if !scoresClump([]int{3, 4, 5}) {
		t.Fatal("expected true but actual is not")
	}
	if scoresClump([]int{3, 4, 6}) {
		t.Fatal("expected false but actual is not")
	}
	if !scoresClump([]int{1, 3, 5, 5}) {
		t.Fatal("expected true but actual is not")
	}
}
