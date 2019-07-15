package string1

import "strings"
import "testing"

func TestStartWord1(t *testing.T) {
	expected := "he"
	actual := startWord("hello", "he")
	if strings.Compare(expected, actual) != 0 {
		t.Fatalf("expected \"%s\" but actual is \"%s\"", expected, actual)
	}
}

func TestStartWord2(t *testing.T) {
	expected := "he"
	actual := startWord("hello", "xe")
	if strings.Compare(expected, actual) != 0 {
		t.Fatalf("expected \"%s\" but actual is \"%s\"", expected, actual)
	}
}

func TestStartWord3(t *testing.T) {
	expected := "h"
	actual := startWord("hello", "x")
	if strings.Compare(expected, actual) != 0 {
		t.Fatalf("expected \"%s\" but actual is \"%s\"", expected, actual)
	}
}