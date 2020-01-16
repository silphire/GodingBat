package ap1

import (
	"testing"
)

func TestScoreUp(t *testing.T) {
	expected := 6
	actual := scoreUp([]string{"a", "a", "b", "b"}, []string{"a", "c", "b", "c"})
	if expected != actual {
		t.Fatalf("expected %d but actual is %d", expected, actual)
	}

	expected = 11
	actual = scoreUp([]string{"a", "a", "b", "b"}, []string{"a", "a", "b", "c"})
	if expected != actual {
		t.Fatalf("expected %d but actual is %d", expected, actual)
	}

	expected = 16
	actual = scoreUp([]string{"a", "a", "b", "b"}, []string{"a", "a", "b", "b"})
	if expected != actual {
		t.Fatalf("expected %d but actual is %d", expected, actual)
	}
}
