package ap1

import "testing"

func TestScoresAverage(t *testing.T) {
	expected := 4
	actual := scoresAverage([]int{2, 2, 4, 4})
	if actual != expected {
		t.Fatalf("expected %d but actual is %d", expected, actual)
	}

	expected = 4
	actual = scoresAverage([]int{4, 4, 4, 2, 2, 2})
	if actual != expected {
		t.Fatalf("expected %d but actual is %d", expected, actual)
	}

	expected = 4
	actual = scoresAverage([]int{3, 4, 5, 1, 2, 3})
	if actual != expected {
		t.Fatalf("expected %d but actual is %d", expected, actual)
	}
}
