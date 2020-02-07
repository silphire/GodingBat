package recursion1

import (
	"strings"
	"testing"
)

func TestNoX(t *testing.T) {
	expected := "ab"
	actual := noX("xaxb")
	if strings.Compare(expected, actual) != 0 {
		t.Fatalf("expected \"%s\" but actual is \"%s\"", expected, actual)
	}

	expected = "abc"
	actual = noX("abc")
	if strings.Compare(expected, actual) != 0 {
		t.Fatalf("expected \"%s\" but actual is \"%s\"", expected, actual)
	}

	expected = ""
	actual = noX("xx")
	if strings.Compare(expected, actual) != 0 {
		t.Fatalf("expected \"%s\" but actual is \"%s\"", expected, actual)
	}
}
