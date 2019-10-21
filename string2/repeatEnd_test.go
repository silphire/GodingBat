package string2

import (
	"strings"
	"testing"
)

func TestRepeatEnd(t *testing.T) {
	expected := "llollollo"
	actual := repeatEnd("Hello", 3)
	if strings.Compare(actual, expected) != 0 {
		t.Fatalf("expected \"%s\", but actual is \"%s\"", expected, actual)
	}

	expected = "lolo"
	actual = repeatEnd("Hello", 2)
	if strings.Compare(actual, expected) != 0 {
		t.Fatalf("expected \"%s\", but actual is \"%s\"", expected, actual)
	}

	expected = "o"
	actual = repeatEnd("Hello", 1)
	if strings.Compare(actual, expected) != 0 {
		t.Fatalf("expected \"%s\", but actual is \"%s\"", expected, actual)
	}
}
