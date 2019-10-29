package string2

import (
	"strings"
	"testing"
)

func TestOneTwo(t *testing.T) {
	expected := "bca"
	actual := oneTwo("abc")
	if strings.Compare(expected, actual) != 0 {
		t.Fatalf("expected \"%s\" but actual is \"%s\"", expected, actual)
	}

	expected = "cat"
	actual = oneTwo("tca")
	if strings.Compare(expected, actual) != 0 {
		t.Fatalf("expected \"%s\" but actual is \"%s\"", expected, actual)
	}

	expected = "catdog"
	actual = oneTwo("tcagdo")
	if strings.Compare(expected, actual) != 0 {
		t.Fatalf("expected \"%s\" but actual is \"%s\"", expected, actual)
	}
}
