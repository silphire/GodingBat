package string1

import "testing"
import "strings"

func TestDeFront1(t *testing.T) {
	expected := "llo"
	actual := deFront("Hello")

	if strings.Compare(expected, actual) != 0 {
		t.Fatalf("expected \"%s\", but actual is \"%s\"", expected, actual)
	}
}

func TestDeFront2(t *testing.T) {
	expected := "anual"
	actual := deFront("annual")

	if strings.Compare(expected, actual) != 0 {
		t.Fatalf("expected \"%s\", but actual is \"%s\"", expected, actual)
	}
}

func TestDeFront3(t *testing.T) {
	expected := "brk"
	actual := deFront("sbrk")

	if strings.Compare(expected, actual) != 0 {
		t.Fatalf("expected \"%s\", but actual is \"%s\"", expected, actual)
	}
}

func TestDeFront4(t *testing.T) {
	expected := "abroad"
	actual := deFront("abroad")

	if strings.Compare(expected, actual) != 0 {
		t.Fatalf("expected \"%s\", but actual is \"%s\"", expected, actual)
	}
}