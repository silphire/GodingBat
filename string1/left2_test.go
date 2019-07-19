package string1

import "testing"
import "strings"

func TestLeft21(t *testing.T) {
	expected := "lloHe"
	actual := left2("Hello")
	if strings.Compare(expected, actual) != 0 {
		t.Fatalf("expected \"%s\", but actual is \"%s\"", expected, actual)
	}
}

func TestLeft22(t *testing.T) {
	expected := "Hi"
	actual := left2("Hi")
	if strings.Compare(expected, actual) != 0 {
		t.Fatalf("expected \"%s\", but actual is \"%s\"", expected, actual)
	}
}