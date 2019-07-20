package string1

import "testing"
import "strings"

func TestWithoutEnd21(t *testing.T) {
	expected := "ell"
	actual := withoutEnd2("Hello")
	if strings.Compare(expected, actual) != 0 {
		t.Fatalf("expected \"%s\" but actual is \"%s\"", expected, actual)
	}
}

func TestWithoutEnd22(t *testing.T) {
	expected := ""
	actual := withoutEnd2("hi")
	if strings.Compare(expected, actual) != 0 {
		t.Fatalf("expected \"%s\" but actual is \"%s\"", expected, actual)
	}
}


func TestWithoutEnd23(t *testing.T) {
	expected := ""
	actual := withoutEnd2("x")
	if strings.Compare(expected, actual) != 0 {
		t.Fatalf("expected \"%s\" but actual is \"%s\"", expected, actual)
	}
}

func TestWithoutEnd24(t *testing.T) {
	expected := ""
	actual := withoutEnd2("")
	if strings.Compare(expected, actual) != 0 {
		t.Fatalf("expected \"%s\" but actual is \"%s\"", expected, actual)
	}
}