package string2

import (
	"strings"
	"testing"
)

func TestWordEnds(t *testing.T) {
	expected := "c13i"
	actual := wordEnds("abcXY123XYijk", "XY")
	if strings.Compare(expected, actual) != 0 {
		t.Fatalf("expected \"%s\" but actual is \"%s\"", expected, actual)
	}

	expected = "13"
	actual = wordEnds("XY123XY", "XY")
	if strings.Compare(expected, actual) != 0 {
		t.Fatalf("expected \"%s\" but actual is \"%s\"", expected, actual)
	}

	expected = "11"
	actual = wordEnds("XY1XY", "XY")
	if strings.Compare(expected, actual) != 0 {
		t.Fatalf("expected \"%s\" but actual is \"%s\"", expected, actual)
	}
}
