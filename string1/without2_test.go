package string1

import "strings"
import "testing"

func TestWithout21(t *testing.T) {
	expected := "lloHe"
	actual := without2("HelloHe")
	if strings.Compare(expected, actual) != 0 {
		t.Fatalf("expected \"%s\" but actual is \"%s\"", expected, actual)
	}
}

func TestWithout22(t *testing.T) {
	expected := "Hello"
	actual := without2("Hello")
	if strings.Compare(expected, actual) != 0 {
		t.Fatalf("expected \"%s\" but actual is \"%s\"", expected, actual)
	}
}