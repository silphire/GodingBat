package string1

import "testing"
import "strings"

func TestRight21(t *testing.T) {
	expected := "loHel"
	actual := right2("Hello")
	if strings.Compare(expected, actual) != 0 {
		t.Fatalf("expected \"%s\" but actual is \"%s\"", expected, actual)
	}
}

func TestRight22(t *testing.T) {
	expected := "Hi"
	actual := right2("Hi")
	if strings.Compare(expected, actual) != 0 {
		t.Fatalf("expected \"%s\" but actual is \"%s\"", expected, actual)
	}
}

func TestRight23(t *testing.T) {
	expected := "X"
	actual := right2("X")
	if strings.Compare(expected, actual) != 0 {
		t.Fatalf("expected \"%s\" but actual is \"%s\"", expected, actual)
	}
}
