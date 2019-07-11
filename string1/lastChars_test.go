package string1

import (
	"testing"
	"strings"
)

func TestLastChars1(t *testing.T) {
	expected := "ab"
	actual := lastChars("and", "bitwise")
	if strings.Compare(expected, actual) != 0 {
		t.Fatalf("expected \"%s\", actual is \"%s\"", expected, actual)
	}
}