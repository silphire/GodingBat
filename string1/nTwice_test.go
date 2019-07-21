package string1

import "strings"
import "testing"

func TestNTwice1(t *testing.T) {
	expected := "Helo"
	actual := nTwice("Hello", 2)
	if strings.Compare(expected, actual) != 0 {
		t.Fatalf("expected \"%s\", but actual is \"%s\"", expected, actual)
	}
}

func TestNTwice2(t *testing.T) {
	expected := "Co"
	actual := nTwice("Choco", 1)
	if strings.Compare(expected, actual) != 0 {
		t.Fatalf("expected \"%s\", but actual is \"%s\"", expected, actual)
	}
}