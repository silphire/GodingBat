package string3

import (
	"strings"
	"testing"
)

func TestWithoutString(t *testing.T) {
	expected := "He there"
	actual := withoutString("Hello there", "llo")
	if strings.Compare(expected, actual) != 0 {
		t.Fatalf("expected \"%s\", but actual is \"%s\"", expected, actual)
	}

	expected = "Hllo thr"
	actual = withoutString("Hello there", "e")
	if strings.Compare(expected, actual) != 0 {
		t.Fatalf("expected \"%s\", but actual is \"%s\"", expected, actual)
	}

	expected = "Hello there"
	actual = withoutString("Hello there", "x")
	if strings.Compare(expected, actual) != 0 {
		t.Fatalf("expected \"%s\", but actual is \"%s\"", expected, actual)
	}

	expected = "x"
	actual = withoutString("xxx", "xx")
	if strings.Compare(expected, actual) != 0 {
		t.Fatalf("expected \"%s\", but actual is \"%s\"", expected, actual)
	}
}
