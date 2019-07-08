package string1

import (
	"strings"
	"testing"
)

func TestTheEnd(t *testing.T) {
	expected := "H"
	actual := theEnd("Hello", true)
	if strings.Compare(expected, actual) != 0 {
		t.Fatalf("expected \"%s\", but actual is \"%s\"", expected, actual)
	}
}

func TestTheFront(t *testing.T) {
	expected := "o"
	actual := theEnd("Hello", false)
	if strings.Compare(expected, actual) != 0 {
		t.Fatalf("expected \"%s\", but actual is \"%s\"", expected, actual)
	}
}