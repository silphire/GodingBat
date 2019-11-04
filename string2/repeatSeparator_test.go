package string2

import (
	"strings"
	"testing"
)

func TestRepeatSeparator(t *testing.T) {
	expected := "WordXWordXWord"
	actual := repeatSeparator("Word", "X", 3)
	if strings.Compare(expected, actual) != 0 {
		t.Fatalf("expected \"%s\" but actual is \"%s\"", expected, actual)
	}

	expected = "ThisAndThis"
	actual = repeatSeparator("This", "And", 2)
	if strings.Compare(expected, actual) != 0 {
		t.Fatalf("expected \"%s\" but actual is \"%s\"", expected, actual)
	}

	expected = "This"
	actual = repeatSeparator("This", "And", 1)
	if strings.Compare(expected, actual) != 0 {
		t.Fatalf("expected \"%s\" but actual is \"%s\"", expected, actual)
	}
}
