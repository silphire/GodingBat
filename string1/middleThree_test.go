package string1

import (
	"testing"
	"strings"
)

func TestMiddleThree1(t *testing.T) {
	expected := "and"
	actual := middleThree("Candy")
	if strings.Compare(expected, actual) != 0 {
		t.Fatalf("expected \"%s\", but actual is \"%s\"", expected, actual)
	}
}

func TestMiddleThree2(t *testing.T) {
	expected := "owd"
	actual := middleThree("Howdy")
	if strings.Compare(expected, actual) != 0 {
		t.Fatalf("expected \"%s\", but actual is \"%s\"", expected, actual)
	}
}

func TestMiddleThree3(t *testing.T) {
	expected := "and"
	actual := middleThree("and")
	if strings.Compare(expected, actual) != 0 {
		t.Fatalf("expected \"%s\", but actual is \"%s\"", expected, actual)
	}
}
