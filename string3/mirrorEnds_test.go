package string3

import (
	"strings"
	"testing"
)

func TestMirrorEnds(t *testing.T) {
	expected := "ab"
	actual := mirrorEnds("abXYZba")
	if strings.Compare(expected, actual) != 0 {
		t.Fatalf("expected \"%s\" but actual is \"%s\"", expected, actual)
	}

	expected = "a"
	actual = mirrorEnds("abca")
	if strings.Compare(expected, actual) != 0 {
		t.Fatalf("expected \"%s\" but actual is \"%s\"", expected, actual)
	}

	expected = "aba"
	actual = mirrorEnds("aba")
	if strings.Compare(expected, actual) != 0 {
		t.Fatalf("expected \"%s\" but actual is \"%s\"", expected, actual)
	}
}
