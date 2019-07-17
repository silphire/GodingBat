package string1

import "testing"
import "strings"

func TestWithoutEnd1(t *testing.T) {
	expected := "ell"
	actual := withoutEnd("Hello")
	if strings.Compare(expected, actual) != 0 {
		t.Fatalf("expected \"%s\", but actual is \"%s\"", expected, actual)
	}
}

func TestWithoutEnd2(t *testing.T) {
	expected := ""
	actual := withoutEnd("Hi")
	if strings.Compare(expected, actual) != 0 {
		t.Fatalf("expected \"%s\", but actual is \"%s\"", expected, actual)
	}
}