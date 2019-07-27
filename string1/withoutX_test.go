package string1

import "strings"
import "testing"

func TestWithoutX1(t *testing.T) {
	expected := "hello"
	actual := withoutX("xhellox")
	if strings.Compare(expected, actual) != 0 {
		t.Fatalf("expected \"%s\" but actual is \"%s\"", expected, actual)
	}
}

func TestWithoutX2(t *testing.T) {
	expected := "Hello"
	actual := withoutX("Hello")
	if strings.Compare(expected, actual) != 0 {
		t.Fatalf("expected \"%s\" but actual is \"%s\"", expected, actual)
	}
}

func TestWithoutX3(t *testing.T) {
	expected := "hello"
	actual := withoutX("xhello")
	if strings.Compare(expected, actual) != 0 {
		t.Fatalf("expected \"%s\" but actual is \"%s\"", expected, actual)
	}
}

func TestWithoutX4(t *testing.T) {
	expected := "hello"
	actual := withoutX("hellox")
	if strings.Compare(expected, actual) != 0 {
		t.Fatalf("expected \"%s\" but actual is \"%s\"", expected, actual)
	}
}

