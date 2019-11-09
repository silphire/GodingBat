package string3

import (
	"strings"
	"testing"
)

func TestSameEnds(t *testing.T) {
	expected := "ab"
	actual := sameEnds("abXYab")
	if strings.Compare(expected, actual) != 0 {
		t.Fatalf("expected \"%s\" but actual is \"%s\"", expected, actual)
	}

	expected = "x"
	actual = sameEnds("xx")
	if strings.Compare(expected, actual) != 0 {
		t.Fatalf("expected \"%s\" but actual is \"%s\"", expected, actual)
	}

	expected = "x"
	actual = sameEnds("xxx")
	if strings.Compare(expected, actual) != 0 {
		t.Fatalf("expected \"%s\" but actual is \"%s\"", expected, actual)
	}
}
