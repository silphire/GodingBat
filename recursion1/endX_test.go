package recursion1

import (
	"strings"
	"testing"
)

func TestEndX(t *testing.T) {
	expected := "rexx"
	actual := endX("xxre")
	if strings.Compare(expected, actual) != 0 {
		t.Fatalf("expected \"%s\" but actual is \"%s\"", expected, actual)
	}

	expected = "hixxxx"
	actual = endX("xxhixx")
	if strings.Compare(expected, actual) != 0 {
		t.Fatalf("expected \"%s\" but actual is \"%s\"", expected, actual)
	}

	expected = "hihixxx"
	actual = endX("xhixhix")
	if strings.Compare(expected, actual) != 0 {
		t.Fatalf("expected \"%s\" but actual is \"%s\"", expected, actual)
	}
}
