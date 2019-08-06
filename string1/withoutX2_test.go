package string1

import "testing"
import "strings"

func TestWithoutX21(t *testing.T) {
	expected := "Hi"
	actual := withoutX2("xHi")
	if strings.Compare(expected, actual) != 0 {
		t.Fatalf("expected \"%s\" but actual is \"%s\"", expected, actual)
	}
}

func TestWithoutX22(t *testing.T) {
	expected := "Hi"
	actual := withoutX2("Hxi")
	if strings.Compare(expected, actual) != 0 {
		t.Fatalf("expected \"%s\" but actual is \"%s\"", expected, actual)
	}
}

func TestWithoutX23(t *testing.T) {
	expected := "Hi"
	actual := withoutX2("Hi")
	if strings.Compare(expected, actual) != 0 {
		t.Fatalf("expected \"%s\" but actual is \"%s\"", expected, actual)
	}
}