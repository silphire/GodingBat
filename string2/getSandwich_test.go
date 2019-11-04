package string2

import (
	"strings"
	"testing"
)

func TestGetSandwich(t *testing.T) {
	expected := "jam"
	actual := getSandwich("breadjambread")
	if strings.Compare(expected, actual) != 0 {
		t.Fatalf("expected \"%s\" but actual is \"%s\"", expected, actual)
	}

	expected = "jam"
	actual = getSandwich("xxbreadjambreadyy")
	if strings.Compare(expected, actual) != 0 {
		t.Fatalf("expected \"%s\" but actual is \"%s\"", expected, actual)
	}

	expected = ""
	actual = getSandwich("xxbreadyy")
	if strings.Compare(expected, actual) != 0 {
		t.Fatalf("expected \"%s\" but actual is \"%s\"", expected, actual)
	}
}
