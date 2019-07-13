package string1

import (
	"testing"
	"strings"
)

func TestRed(t *testing.T) {
	expected := "red"
	actual := seeColor("redxx")
	if strings.Compare(expected, actual) != 0 {
		t.Fatalf("expected \"%s\" but actual is \"%s\"", expected, actual)
	}
}

func TestNotRed(t *testing.T) {
	actual := seeColor("aaredxx")
	if len(actual) != 0 {
		t.Fatalf("expected empty string but actual is \"%s\"", actual)
	}
}

func TestBlue(t *testing.T) {
	expected := "blue"
	actual := seeColor("bluexx")
	if strings.Compare(expected, actual) != 0 {
		t.Fatalf("expected \"%s\" but actual is \"%s\"", expected, actual)
	}
}

func TestNotBlue(t *testing.T) {
	actual := seeColor("aabluexx")
	if len(actual) != 0 {
		t.Fatalf("expected empty string but actual is \"%s\"", actual)
	}
}