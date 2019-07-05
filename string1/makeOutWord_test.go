package string1

import (
	"strings"
	"testing"
)

func TestMakeOutWord1(t *testing.T) {
	expected := "<<Hoo>>"
	actual := makeOutWord("<<>>", "Hoo")
	if strings.Compare(expected, actual) != 0 {
		t.Fatalf("expected \"%s\", but actual is \"%s\"", expected, actual)
	}
}

func TestMakeOutWord2(t *testing.T) {
	expected := "[[Foo Bar]]"
	actual := makeOutWord("[[]]", "Foo Bar")
	if strings.Compare(expected, actual) != 0 {
		t.Fatalf("expected \"%s\", but actual is \"%s\"", expected, actual)
	}
}
